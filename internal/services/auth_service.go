package services

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	database "leopard/internal/db"
	models "leopard/internal/model"
)

type AuthService struct {
	db          *database.Database
	currentUser *models.User
}

func NewAuthService(db *database.Database) *AuthService {
	return &AuthService{db: db}
}

// --- Tes fonctions de Hashage (ON NE TOUCHE PAS) ---

func (s *AuthService) generateSalt() (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func (s *AuthService) hashPassword(password, salt string) string {
	h := sha256.New()
	h.Write([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// --- Inscription ---

func (s *AuthService) Register(username, password, fullName, role string) error {
	if username == "" || password == "" || fullName == "" {
		return errors.New("tous les champs sont requis")
	}

	// 1. Préparation du sel et du hash
	salt, err := s.generateSalt()
	if err != nil {
		return err
	}
	hashedPassword := s.hashPassword(password, salt)

	// 2. Requête avec paramètres nommés
	// On utilise :user, :pass, etc. au lieu des ?
	query := `
		INSERT INTO users (username, password, salt, full_name, role)
		VALUES (:user, :pass, :salt, :name, :role)`

	// 3. On passe une map pour lier les noms aux valeurs
	_, err = s.db.NamedExec(query, map[string]interface{}{
		"user": username,
		"pass": hashedPassword,
		"salt": salt,
		"name": fullName,
		"role": role,
	})

	return err
}

// --- Connexion ---

func (s *AuthService) Login(username, passwordClair string) (map[string]interface{}, error) {
	var user models.User

	// On récupère l'user.
	// Important: Ta struct User doit avoir le tag `db:"password"`
	err := s.db.Get(&user, `SELECT * FROM users WHERE username = ?`, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("identifiants invalides")
		}
		return nil, fmt.Errorf("erreur base de données: %w", err)
	}

	// VERIFICATION DU HASH
	// On reprend le sel de la DB + le mot de passe tapé
	providedHash := s.hashPassword(passwordClair, user.Salt)

	// On compare avec le hash stocké dans user.Password
	if providedHash != user.Password {
		return nil, errors.New("identifiants invalides")
	}

	s.currentUser = &user

	return map[string]interface{}{
		"success": true,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"fullName": user.FullName,
			"role":     user.Role,
		},
	}, nil
}

func (s *AuthService) Logout() error {
	s.currentUser = nil
	return nil
}

func (s *AuthService) GetCurrentUser() *models.User {
	return s.currentUser
}
