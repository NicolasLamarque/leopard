// app.go
package main

import (
	"context"
	"database/sql"
	"errors"
	database "leopard/internal/db"
	models "leopard/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type App struct {
	ctx         context.Context
	db          *database.Database
	currentUser *models.User
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

	if err := a.initializeLeopardFolders(); err != nil {
		println("⚠️ Avertissement: Impossible d'initialiser les dossiers Leopard:", err.Error())
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

// ========== CLIENTS ==========

func (a *App) GetClients() ([]models.Client, error) {
	return a.db.GetAllClients()
}

func (a *App) CreateClient(req models.CreateClientRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateClient(req, int(a.currentUser.ID))
}

func (a *App) GetClientByID(id int) (*models.Client, error) {
	return a.db.GetClientByID(id)
}

func (a *App) UpdateClient(req models.UpdateClientRequest) error {
	return a.db.UpdateClient(req)
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

func (a *App) GetResidences() ([]models.Residence, error) {
	return a.db.GetAllResidences()
}

// Dans app.go
// GetResidenceForDetails fait le pont entre le frontend et le repo
func (a *App) GetResidenceForDetails(registre string, sync bool) (*models.Residence, error) {
	return a.db.GetResidenceByRegistre(registre, sync)
}

func (a *App) GetResidenceByID(id int) (*models.Residence, error) {
	return a.db.GetResidenceByID(id)
}

func (a *App) InsertResidence(residence *models.Residence) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.InsertResidence(residence)
}

func (a *App) UpdateResidence(residence *models.Residence) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateResidence(residence)
}

func (a *App) DeleteResidence(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteResidence(id)
}

func (a *App) SearchResidences(municipalite, nom, statut string) ([]models.Residence, error) {
	return a.db.SearchResidences(municipalite, nom, statut)
}

// ✅ Synchronisation RPA depuis le site du MSSS
func (a *App) SyncRPA() (map[string]interface{}, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}

	stats, err := a.db.SyncRPAFromMSSS()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_scraped": stats.TotalScraped,
		"nouveaux":      stats.Nouveaux,
		"mis_a_jour":    stats.MisAJour,
		"fermes":        stats.Fermes,
		"erreurs":       stats.Erreurs,
		"date_sync":     stats.DateSync,
	}, nil
}
