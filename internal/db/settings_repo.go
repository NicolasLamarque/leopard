package database

import (
	"database/sql"
	models "leopard/internal/model"
)

// On utilise (int, models.UpdateSettingsRequest) pour matcher l'appel dans app.go
func (db *Database) UpdateUserSettings(userID int, req models.UpdateSettingsRequest) error {
	_, err := db.Exec(`
        UPDATE user_settings 
        SET theme = ?, language = ?, notifications_enabled = ?, updated_at = CURRENT_TIMESTAMP
        WHERE user_id = ?
    `, req.Theme, req.Language, req.NotificationsEnabled, userID)
	return err
}

func (db *Database) UpdateUserProfile(userID int, req models.UpdateProfileRequest) error {
	_, err := db.Exec(`
		UPDATE users 
		SET full_name = ?
		WHERE id = ?
	`, req.FullName, userID)
	return err
}

func (db *Database) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	// sqlx va mapper id, username, full_name, role vers ta struct User
	err := db.Get(&user, `SELECT id, username, full_name, role FROM users WHERE id = ?`, userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// On garde GetUserSettings pour la récupération initiale
func (db *Database) GetUserSettings(userID int) (*models.UserSettings, error) {
	var s models.UserSettings
	err := db.Get(&s, `SELECT * FROM user_settings WHERE user_id = ?`, userID)

	if err == sql.ErrNoRows {
		return db.CreateDefaultSettings(userID)
	}
	return &s, err
}

func (db *Database) CreateDefaultSettings(userID int) (*models.UserSettings, error) {
	_, err := db.Exec(`
        INSERT INTO user_settings (user_id, theme, language, notifications_enabled)
        VALUES (?, 'light', 'fr', 1)
    `, userID)
	if err != nil {
		return nil, err
	}
	return db.GetUserSettings(userID)
}
