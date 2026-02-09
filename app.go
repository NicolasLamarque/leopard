// app.go
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"leopard/internal/crypto"
	database "leopard/internal/db"
	repo "leopard/internal/db"
	"leopard/internal/export"
	models "leopard/internal/model"
	"leopard/internal/services"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type App struct {
	ctx            context.Context
	db             *database.Database
	authSvc        *services.AuthService // <--- IL MANQUE ÇA
	currentUser    *models.User
	cryptoSvc      *crypto.CryptoService
	LogRepo        *repo.LogRepo
	currentUserID  int
	PIService      *services.PIService
	IntervenantSvc *services.IntervenantService
	appPath        string
}

func NewApp() *App {
	return &App{}
}
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	db, err := database.New("./app.db")
	if err != nil {
		panic(err)
	}
	a.db = db
	// Initialisez le service d'auth ici
	a.authSvc = services.NewAuthService(db)
	masterPassword := "MonMotDePasseSecret2025!"
	cryptoSvc, err := crypto.NewCryptoService(masterPassword)
	if err != nil {
		panic(err)
	}
	a.cryptoSvc = cryptoSvc
	a.IntervenantSvc = services.NewIntervenantService(db, cryptoSvc)

	if err := a.initializeLeopardFolders(); err != nil {
		println("⚠️ Avertissement:", err.Error())
	}

	// --- ON AJOUTE ÇA ICI : ---
	templatePath := "data/Modele_Import_Notaires.xlsx"
	// Si le fichier n'existe pas, on demande à la DB de le créer
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		// Pas besoin de Mkdir ici car initializeLeopardFolders s'en occupe sûrement
		a.db.CreateNotaireTemplate(templatePath)
	}
}

// ========== AUTH ==========

func (a *App) Login(username, password string) (map[string]interface{}, error) {
	var user models.User
	var passwordHash string

	// 1. On garde ta requête SQL intacte
	err := a.db.QueryRow(
		"SELECT id, username, password, full_name, role FROM users WHERE username = ?",
		username,
	).Scan(&user.ID, &user.Username, &passwordHash, &user.FullName, &user.Role)

	if err == sql.ErrNoRows {
		return nil, errors.New("identifiants invalides")
	}
	if err != nil {
		return nil, errors.New("erreur base de données")
	}

	// 2. Ta vérification Bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return nil, errors.New("identifiants invalides")
	}

	// 3. ON FIXE L'ID ICI
	// On s'assure que currentUser est bien assigné
	a.currentUser = &user

	// On force la conversion en int si jamais user.ID est un int64 dans le model
	a.currentUserID = int(user.ID)

	// 4. Le retour pour ton FormKit (Vue)
	return map[string]interface{}{
		"id":       int(user.ID), // On cast en int pour être certain
		"username": user.Username,
		"fullName": user.FullName,
		"role":     user.Role,
	}, nil
}
func (a *App) Logout() {
	a.currentUser = nil
}

// ========== SETUP PREMIER UTILISATEUR ==========

// GetUserCount retourne le nombre d'utilisateurs dans la base
func (a *App) GetUserCount() (int, error) {
	var count int
	err := a.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CreateFirstUserRequest structure pour la création du premier user
type CreateFirstUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

// CreateFirstUser crée le premier utilisateur (admin obligatoirement)
func (a *App) CreateFirstUser(req CreateFirstUserRequest) error {
	// Vérifier qu'il n'y a vraiment aucun utilisateur
	count, err := a.GetUserCount()
	if err != nil {
		return fmt.Errorf("erreur vérification: %w", err)
	}

	if count > 0 {
		return errors.New("des utilisateurs existent déjà")
	}

	// Validation
	if req.Username == "" || req.Password == "" || req.FullName == "" {
		return errors.New("tous les champs sont requis")
	}

	if len(req.Password) < 5 {
		return errors.New("le mot de passe doit contenir au moins 5 caractères")
	}

	// Hash du mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return fmt.Errorf("erreur hashage mot de passe: %w", err)
	}

	// Insertion du premier user (toujours admin)
	query := `
		INSERT INTO users (username, password, full_name, role, created_at)
		VALUES (?, ?, ?, 'admin', CURRENT_TIMESTAMP)
	`

	_, err = a.db.Exec(query, req.Username, string(hashedPassword), req.FullName)
	if err != nil {
		return fmt.Errorf("erreur création utilisateur: %w", err)
	}

	println("✅ Premier utilisateur créé:", req.Username, "(admin)")
	return nil
}

// ========== CLIENTS ==========

func (a *App) GetClients() ([]models.Client, error) {
	return a.db.GetAllClients(a.cryptoSvc) // ✅ Ajouté
}

func (a *App) CreateClient(req models.CreateClientRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateClient(req, int(a.currentUser.ID), a.cryptoSvc) // ✅ Ajouté
}

func (a *App) GetClientByID(id int) (*models.Client, error) {
	return a.db.GetClientByID(id, a.cryptoSvc) // ✅ Ajouté
}

func (a *App) UpdateClient(req models.UpdateClientRequest) error {
	return a.db.UpdateClient(req, a.cryptoSvc) // ✅ Ajouté
}

func (a *App) DeleteClient(id int) error {
	return a.db.DeleteClient(id)
}

// ========== SETTINGS ==========

func (a *App) GetSettings() (*models.UserSettings, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetUserSettings(int(a.currentUser.ID))
}

func (a *App) UpdateSettings(req models.UpdateSettingsRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateUserSettings(int(a.currentUser.ID), req)
}

func (a *App) UpdateProfile(req models.UpdateProfileRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	err := a.db.UpdateUserProfile(int(a.currentUser.ID), req)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ChangePassword(req models.ChangePasswordRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	var currentHash string
	err := a.db.QueryRow(
		"SELECT password FROM users WHERE id = ?",
		a.currentUser.ID,
	).Scan(&currentHash)

	if err != nil {
		return errors.New("erreur base de données")
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentHash), []byte(req.CurrentPassword))
	if err != nil {
		return errors.New("mot de passe actuel incorrect")
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 14)
	if err != nil {
		return errors.New("erreur génération mot de passe")
	}

	_, err = a.db.Exec(
		"UPDATE users SET password = ? WHERE id = ?",
		string(newHash), a.currentUser.ID,
	)

	return err
}

func (a *App) GetCurrentUserProfile() (*models.User, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetUserByID(int(a.currentUser.ID))
}

// ========== NOTES ==========
// ========== NOTES CLINIQUES ==========

// GetClientNotes récupère toutes les notes d'un client (pour la sidebar)
func (a *App) GetClientNotes(clientID int) ([]models.NoteListItem, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}

	if clientID == 0 {
		return nil, errors.New("ID client invalide")
	}

	return a.db.GetAllNotesByClientID(clientID, a.cryptoSvc)
}

// GetNoteByID récupère une note complète (pour visualisation/édition)
func (a *App) GetNoteByID(noteID int) (*models.Note, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}

	if noteID == 0 {
		return nil, errors.New("ID note invalide")
	}

	return a.db.GetNoteByID(noteID, a.cryptoSvc)
}

// CreateNote crée une nouvelle note (brouillon)
func (a *App) CreateNote(req models.CreateNoteRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}

	// Validation de base
	if req.ClientID == 0 {
		return 0, errors.New("ID client requis")
	}

	// Assigner l'utilisateur courant
	req.UserID = int(a.currentUser.ID)

	return a.db.CreateNote(req, a.cryptoSvc)
}

// UpdateNoteDraft met à jour un brouillon existant
func (a *App) UpdateNoteDraft(noteID int, req models.UpdateNoteRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if noteID == 0 {
		return errors.New("ID note invalide")
	}

	req.ID = noteID
	return a.db.UpdateNote(req, a.cryptoSvc)
}

// DeleteNote supprime un brouillon
func (a *App) DeleteNote(noteID int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if noteID == 0 {
		return errors.New("ID note invalide")
	}

	return a.db.DeleteNote(noteID)
}

// LockNote verrouille et signe une note (irréversible)
func (a *App) LockNote(noteID int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if noteID == 0 {
		return errors.New("ID note invalide")
	}

	// Utiliser le nom complet de l'utilisateur comme signature
	signatureNom := a.currentUser.FullName
	// DEBUG
	fmt.Printf("🔍 DEBUG LockNote - FullName: '%s'\n", signatureNom)
	return a.db.LockNote(noteID, signatureNom, a.cryptoSvc)
}

// ExportNotesToPDF génère un PDF avec les notes sélectionnées
func (a *App) ExportNotesToPDF(leopardNumber string, noteIDs []int) (string, error) {
	if a.currentUser == nil {
		return "", errors.New("non authentifié")
	}

	if leopardNumber == "" {
		return "", errors.New("numéro Léopard requis")
	}

	if len(noteIDs) == 0 {
		return "", errors.New("aucune note sélectionnée")
	}

	// Récupérer les notes
	notes, err := a.db.GetNotesByIDs(noteIDs, a.cryptoSvc)
	if err != nil {
		return "", fmt.Errorf("erreur récupération notes: %w", err)
	}

	if len(notes) == 0 {
		return "", errors.New("aucune note trouvée")
	}

	// Utiliser le système de dossiers existant
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// Trouver le dossier du client
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return "", fmt.Errorf("erreur lecture dossier Clients: %w", err)
	}

	var clientPath string
	targetLength := len(leopardNumber)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				clientPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if clientPath == "" {
		return "", fmt.Errorf("dossier client introuvable pour %s", leopardNumber)
	}

	// Le PDF va dans le sous-dossier "Notes"
	outputDir := filepath.Join(clientPath, "Notes")

	// Créer le dossier Notes s'il n'existe pas
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("erreur création dossier Notes: %w", err)
	}

	// Générer le PDF
	pdfPath, err := export.GenerateNotesPDF(notes, leopardNumber, outputDir)
	if err != nil {
		return "", fmt.Errorf("erreur génération PDF: %w", err)
	}

	return pdfPath, nil
}

// ========== FIN SECTION NOTES ==========

// ========== MÉDECINS ==========

func (a *App) GetMedecins() ([]models.Medecin, error) {
	return a.db.GetAllMedecins()
}

func (a *App) GetMedecinByID(id int) (*models.Medecin, error) {
	return a.db.GetMedecinByID(id)
}

func (a *App) CreateMedecin(req models.CreateMedecinRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateMedecin(req, int(a.currentUser.ID))
}

func (a *App) UpdateMedecin(req models.UpdateMedecinRequest) error {
	return a.db.UpdateMedecin(req)
}

func (a *App) DeleteMedecin(id int) error {
	return a.db.DeleteMedecin(id)
}

func (a *App) SearchMedecins(query string) ([]models.Medecin, error) {
	return a.db.SearchMedecins(query)
}

// À ajouter dans app.go dans la section MÉDECINS

// GetMedecinClients récupère tous les clients d'un médecin
func (a *App) GetMedecinClients(licence string) ([]models.Client, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetClientsByMedecinLicence(licence, a.cryptoSvc)
}

// GetMedecinClientsCount retourne le nombre de clients d'un médecin
func (a *App) GetMedecinClientsCount(licence string) (int, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.GetMedecinClientsCount(licence, a.cryptoSvc)
}

// ========== RÉSIDENCES / RPA ==========

// ✅ CONSERVÉ - GetResidences (utilisé dans RPAPage.vue au démarrage)
func (a *App) GetResidences() ([]models.Residence, error) {
	return a.db.GetAllResidences()
}

// ✅ CONSERVÉ - GetResidenceForDetails (utilisé dans RPADetails.vue)
// Cette fonction fait le pont entre le frontend et le repo
func (a *App) GetResidenceForDetails(registre string, sync bool) (*models.Residence, error) {
	println("📖 [app.go] GetResidenceForDetails appelé")
	println("   - Registre:", registre)
	println("   - Sync demandé:", sync)
	println("   - Utilisateur connecté:", a.currentUser != nil)

	// On passe directement au repo qui gère la logique
	return a.db.GetResidenceByRegistre(registre, sync)
}

// ✅ CONSERVÉ - GetResidenceByID
func (a *App) GetResidenceByID(id int) (*models.Residence, error) {
	return a.db.GetResidenceByID(id)
}

// ✅ CONSERVÉ - InsertResidence
func (a *App) InsertResidence(residence *models.Residence) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.InsertResidence(residence)
}

// ✅ CONSERVÉ - UpdateResidence
func (a *App) UpdateResidence(residence *models.Residence) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateResidence(residence)
}

// ✅ CONSERVÉ - DeleteResidence
func (a *App) DeleteResidence(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteResidence(id)
}

// ✅ CONSERVÉ - SearchResidences (utilisé dans RPAPage.vue avec filtres)
func (a *App) SearchResidences(municipalite, nom, statut string) ([]models.Residence, error) {
	return a.db.SearchResidences(municipalite, nom, statut)
}

// ✅ CONSERVÉ - SyncRPA (utilisé dans RPASyncPanel.vue)
// Synchronisation complète depuis le site du MSSS
func (a *App) SyncRPA() (map[string]interface{}, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}

	println("🚀 [app.go] SyncRPA lancé par:", a.currentUser.FullName)

	stats, err := a.db.SyncRPAFromMSSS()
	if err != nil {
		println("❌ [app.go] Erreur sync:", err.Error())
		return nil, err
	}

	println("✅ [app.go] Sync terminé avec succès")

	return map[string]interface{}{
		"total_scraped": stats.TotalScraped,
		"nouveaux":      stats.Nouveaux,
		"mis_a_jour":    stats.MisAJour,
		"fermes":        stats.Fermes,
		"erreurs":       stats.Erreurs,
		"date_sync":     stats.DateSync,
	}, nil
}

// ========== CHSLD ==========

// GetAllCHSLD retourne tous les CHSLD
func (a *App) GetAllCHSLD() ([]models.CHSLD, error) {
	return a.db.GetAllCHSLD()
}

// GetCHSLDByID retourne un CHSLD par son ID
func (a *App) GetCHSLDByID(id int) (*models.CHSLD, error) {
	return a.db.GetCHSLDByID(id)
}

// SearchCHSLD recherche des CHSLD selon des critères
func (a *App) SearchCHSLD(nom, municipalite, region string) ([]models.CHSLD, error) {
	return a.db.SearchCHSLD(nom, municipalite, region)
}

// CreateCHSLD crée un nouveau CHSLD
func (a *App) CreateCHSLD(chsld models.CHSLD) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	// Validation basique
	if chsld.TitreCHSLD == "" {
		return fmt.Errorf("le nom du CHSLD est requis")
	}

	if chsld.Municipalite == "" {
		return fmt.Errorf("la municipalité est requise")
	}

	// Définir les dates
	now := time.Now().Format(time.RFC3339)
	chsld.DateAjout = now
	chsld.DateMaj = now

	// Créer le CHSLD
	return a.db.CreateCHSLD(&chsld)
}

// UpdateCHSLD met à jour un CHSLD existant
func (a *App) UpdateCHSLD(chsld models.CHSLD) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	// Validation
	if chsld.ID == 0 {
		return fmt.Errorf("ID CHSLD invalide")
	}

	if chsld.TitreCHSLD == "" {
		return fmt.Errorf("le nom du CHSLD est requis")
	}

	// Mettre à jour la date de modification
	chsld.DateMaj = time.Now().Format(time.RFC3339)

	return a.db.UpdateCHSLD(&chsld)
}

// DeleteCHSLD supprime un CHSLD
func (a *App) DeleteCHSLD(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if id == 0 {
		return fmt.Errorf("ID CHSLD invalide")
	}

	return a.db.DeleteCHSLD(id)
}

// GetCHSLDByRegion retourne les CHSLD d'une région
func (a *App) GetCHSLDByRegion(region string) ([]models.CHSLD, error) {
	return a.db.GetCHSLDByRegion(region)
}

// GetCHSLDByStatut retourne les CHSLD par statut
func (a *App) GetCHSLDByStatut(statut string) ([]models.CHSLD, error) {
	return a.db.GetCHSLDByStatut(statut)
}

// GetCHSLDStats retourne des statistiques sur les CHSLD
func (a *App) GetCHSLDStats() (map[string]interface{}, error) {
	return a.db.GetCHSLDStats()
}

// ========== CONTACTS (Correction des arguments) ==========

func (a *App) GetContacts() ([]models.Contact, error) {
	// Le repo "want" (veut) le cryptoSvc en argument
	return a.db.GetAllContacts(a.cryptoSvc)
}

func (a *App) CreateContact(req models.CreateContactRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	// Ajout de a.cryptoSvc ici aussi
	return a.db.CreateContact(req, a.cryptoSvc)
}

func (a *App) UpdateContact(req models.UpdateContactRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	// Et ici également
	return a.db.UpdateContact(req, a.cryptoSvc)
}

// GetAllContactsByClientID fait le pont vers contact_repo.go
func (a *App) GetAllContactsByClientID(clientID int) ([]models.Contact, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	// Ton repo demande le cryptoSvc, on lui passe
	return a.db.GetAllContactsByClientID(clientID, a.cryptoSvc)
}

// DeleteContact fait le pont vers le repo (Vérifie s'il est dans contact_repo.go,
// sinon ajoute-le au repo d'abord)
func (a *App) DeleteContact(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteContact(id)
}

// ========== APPELS SIMPLIFIÉS - À AJOUTER DANS app.go ==========

// GetAllAppels retourne tous les appels (liste)
func (a *App) GetAllAppels() ([]models.AppelListItem, error) {
	repo := database.NewAppelRepository(a.db)
	return repo.GetAll(a.cryptoSvc)
}

// GetAppelByID retourne un appel complet par ID
func (a *App) GetAppelByID(id int) (*models.Appel, error) {
	repo := database.NewAppelRepository(a.db)
	return repo.GetByID(id, a.cryptoSvc)
}

// CreateAppel crée un nouvel appel
func (a *App) CreateAppel(req models.CreateAppelRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	repo := database.NewAppelRepository(a.db)
	_, err := repo.Create(req, int(a.currentUser.ID), a.cryptoSvc)
	return err
}

// UpdateAppel met à jour un appel existant
func (a *App) UpdateAppel(id int, req models.CreateAppelRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	repo := database.NewAppelRepository(a.db)

	// Conversion de int64 vers int ici :
	return repo.Update(id, req, int(a.currentUser.ID), a.cryptoSvc)
}

// DeleteAppel supprime un appel
func (a *App) DeleteAppel(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	repo := database.NewAppelRepository(a.db)
	return repo.Delete(id)
}

// GetStatsAppels retourne les statistiques des appels
func (a *App) GetStatsAppels() (*models.StatsAppels, error) {
	repo := database.NewAppelRepository(a.db)
	return repo.GetStats()
}

// SetCurrentUser définit l'utilisateur courant (à appeler après login)
func (a *App) SetCurrentUser(userID int) {
	a.currentUserID = userID

	// Si vous utilisez le contexte Wails
	if a.ctx != nil {
		a.ctx = context.WithValue(a.ctx, "user_id", userID)
	}
}
func (a *App) GetPlansByClient(clientID int) ([]models.PlanInterventionDetail, error) {
	return a.PIService.GetPlans(clientID)
}

func (a *App) CreatePlan(req models.CreatePlanRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.PIService.Create(req, a.currentUser.ID)
}

func (a *App) LockPlan(planID int, signature string) error {
	return a.PIService.Lock(planID, signature)
}

// app.go
func (a *App) UpdatePlan(planID int, req models.CreatePlanRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	// On appelle le service avec l'ID d'abord, puis la structure req
	// On ne passe PAS userID ici car ton s.db.UpdatePlan ne le demande pas
	return a.PIService.Update(planID, req)
}

// Dans app.go

// GetClientStoragePath garantit l'existence du dossier et sous-dossier
// Usage: path, err := a.GetClientStoragePath("2023-001", "Notes")
func (a *App) GetClientStoragePath(leopardNumber string, subFolder string) (string, error) {
	basePath := a.GetBasePath() // Déjà dans folders.go
	clientsPath := filepath.Join(basePath, "Clients")

	// 1. Trouver le dossier client par le numéro Leopard
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return "", fmt.Errorf("dossier Clients introuvable")
	}

	var targetPath string
	for _, entry := range entries {
		if entry.IsDir() && len(entry.Name()) >= len(leopardNumber) && entry.Name()[:len(leopardNumber)] == leopardNumber {
			targetPath = filepath.Join(clientsPath, entry.Name(), subFolder)
			break
		}
	}

	if targetPath == "" {
		return "", fmt.Errorf("dossier du client %s introuvable", leopardNumber)
	}

	// 2. Créer le sous-dossier (Notes, PI, etc.) s'il n'existe pas
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return "", err
	}

	return targetPath, nil
}

// ========== INTERVENANTS ==========

func (a *App) GetAllIntervenants() ([]models.Intervenant, error) {
	if a.IntervenantSvc == nil {
		return nil, errors.New("service intervenant non initialisé")
	}
	return a.IntervenantSvc.GetList()
}

func (a *App) GetIntervenantByID(id int) (*models.Intervenant, error) {
	if a.IntervenantSvc == nil {
		return nil, errors.New("service intervenant non initialisé")
	}
	return a.IntervenantSvc.GetByID(id)
}

func (a *App) SaveIntervenant(intervenant models.Intervenant) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	if a.IntervenantSvc == nil {
		return errors.New("service intervenant non initialisé")
	}
	return a.IntervenantSvc.Save(intervenant)
}

func (a *App) DeleteIntervenant(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	if a.IntervenantSvc == nil {
		return errors.New("service intervenant non initialisé")
	}
	return a.IntervenantSvc.Delete(id)
}

// AJOUTE AUSSI DANS LA FONCTION startup() APRÈS cryptoSvc :

// Initialiser le service des intervenants
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
