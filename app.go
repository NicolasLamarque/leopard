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

// ✅ MODIFICATION 1 : Ajouter l'initialisation des dossiers
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	db, err := database.New("./app.db")
	if err != nil {
		panic(err)
	}
	a.db = db

	// 👇 AJOUTE CETTE LIGNE
	if err := a.initializeLeopardFolders(); err != nil {
		// On log juste l'erreur mais on ne crash pas l'app
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

	// Vérifier avec bcrypt
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

	// Récupérer le hash actuel
	var currentHash string
	err := a.db.QueryRow(
		"SELECT password FROM users WHERE id = ?",
		a.currentUser.ID,
	).Scan(&currentHash)

	if err != nil {
		return errors.New("erreur base de données")
	}

	// Vérifier l'ancien mot de passe avec bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(currentHash), []byte(req.CurrentPassword))
	if err != nil {
		return errors.New("mot de passe actuel incorrect")
	}

	// Générer nouveau hash avec bcrypt
	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 14)
	if err != nil {
		return errors.New("erreur génération mot de passe")
	}

	// Mettre à jour
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
	// On passe l'ID et le nom de l'utilisateur connecté pour la signature automatique
	return a.db.CreateNote(req, int(a.currentUser.ID), a.currentUser.FullName)
}

func (a *App) GetNoteByID(id int) (*models.Note, error) {
	return a.db.GetNoteByID(id)
}

// Utilise NoteListItem pour une liste légère et rapide à charger
func (a *App) GetClientNotes(clientID int) ([]models.NoteListItem, error) {
	return a.db.GetClientNotes(clientID)
}

// On adapte pour recevoir le struct de filtrage que tu as défini dans models
func (a *App) GetClientNotesFiltered(filter models.NotesFilter) ([]models.NoteListItem, error) {
	return a.db.GetClientNotesFiltered(filter)
}

// On garde UpdateNote pour la forme, même si on bloque la modif dans le repo
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
