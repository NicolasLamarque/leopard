package models

type UserSettings struct {
	ID                   int    `db:"id" json:"id"`
	UserID               int    `db:"user_id" json:"user_id"`
	Theme                string `db:"theme" json:"theme"`
	Language             string `db:"language" json:"language"`
	NotificationsEnabled bool   `db:"notifications_enabled" json:"notifications_enabled"`
	EmailNotifications   bool   `db:"email_notifications" json:"email_notifications"`
	CreatedAt            string `db:"created_at" json:"created_at"`
	UpdatedAt            string `db:"updated_at" json:"updated_at"`
}

type UpdateSettingsRequest struct {
	Theme                string `db:"theme" json:"theme"`
	Language             string `db:"language" json:"language"`
	NotificationsEnabled bool   `db:"notifications_enabled" json:"notifications_enabled"`
	EmailNotifications   bool   `db:"email_notifications" json:"email_notifications"`
}

type UpdateProfileRequest struct {
	FullName string `db:"full_name" json:"full_name"`
	Email    string `db:"email" json:"email"`
}

// Pour le changement de mot de passe, on n'a souvent pas besoin de tags 'db'
// car on ne mappe pas directement cette struct à une table (on compare et on update manuellement).
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
