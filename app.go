// app.go
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"leopard/internal/crypto"
	database "leopard/internal/db"
	models "leopard/internal/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type App struct {
	ctx         context.Context
	db          *database.Database
	currentUser *models.User
	cryptoSvc   *crypto.CryptoService
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

	// 🔐 MOT DE PASSE SIMPLE - Personne ne va le voir dans ton .exe
	masterPassword := "MonMotDePasseSecret2025!"

	cryptoSvc, err := crypto.NewCryptoService(masterPassword)
	if err != nil {
		panic(err)
	}
	a.cryptoSvc = cryptoSvc

	if err := a.initializeLeopardFolders(); err != nil {
		println("⚠️ Avertissement:", err.Error())
	}
}

// ========== AUTH ==========

func (a *App) Login(username, password string) (map[string]interface{}, error) {
	var user models.User
	var passwordHash string

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

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return nil, errors.New("identifiants invalides")
	}

	a.currentUser = &user

	return map[string]interface{}{
		"id":       user.ID,
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

func (a *App) CreateNote(req models.CreateNoteRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateNote(req, int(a.currentUser.ID), a.currentUser.FullName)
}

func (a *App) GetNoteByID(id int) (*models.Note, error) {
	return a.db.GetNoteByID(id)
}

func (a *App) GetClientNotes(clientID int) ([]models.NoteListItem, error) {
	return a.db.GetClientNotes(clientID)
}

func (a *App) GetClientNotesFiltered(filter models.NotesFilter) ([]models.NoteListItem, error) {
	return a.db.GetClientNotesFiltered(filter)
}

func (a *App) UpdateNote(req models.Note) error {
	return a.db.UpdateNote(req)
}

func (a *App) DeleteNote(id int) error {
	return a.db.DeleteNote(id)
}

func (a *App) LockNote(id int) error {
	return a.db.LockNote(id)
}

func (a *App) GetNotesStats(clientID int) (map[string]interface{}, error) {
	return a.db.GetNotesStats(clientID)
}

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
