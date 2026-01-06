package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// CryptoService gère le chiffrement/déchiffrement des données sensibles
type CryptoService struct {
	key []byte
}

// NewCryptoService crée un nouveau service de chiffrement
// masterPassword devrait être stocké de manière sécurisée (variable d'environnement, config protégée)
func NewCryptoService(masterPassword string) (*CryptoService, error) {
	if masterPassword == "" {
		return nil, errors.New("le mot de passe maître ne peut pas être vide")
	}

	// Dérivation de clé avec PBKDF2 (plus sécurisé qu'un hash simple)
	// Salt fixe pour cette application (idéalement stocké séparément)
	salt := []byte("leopard-medical-2025-secure-salt")
	key := pbkdf2.Key([]byte(masterPassword), salt, 100000, 32, sha256.New)

	return &CryptoService{key: key}, nil
}

// Encrypt chiffre une chaîne de caractères
func (cs *CryptoService) Encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil // Ne chiffre pas les chaînes vides
	}

	block, err := aes.NewCipher(cs.key)
	if err != nil {
		return "", fmt.Errorf("erreur création cipher: %w", err)
	}

	// GCM (Galois/Counter Mode) - authentifié et sécurisé
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("erreur création GCM: %w", err)
	}

	// Génération d'un nonce aléatoire (crucial pour la sécurité)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("erreur génération nonce: %w", err)
	}

	// Chiffrement (le nonce est préfixé aux données chiffrées)
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encodage en base64 pour stockage en DB
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt déchiffre une chaîne de caractères
func (cs *CryptoService) Decrypt(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", nil // Retourne vide si vide
	}

	// Décodage base64
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("erreur décodage base64: %w", err)
	}

	block, err := aes.NewCipher(cs.key)
	if err != nil {
		return "", fmt.Errorf("erreur création cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("erreur création GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("données chiffrées invalides")
	}

	// Extraction du nonce et des données
	nonce, cipherData := data[:nonceSize], data[nonceSize:]

	// Déchiffrement
	plaintext, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", fmt.Errorf("erreur déchiffrement: %w", err)
	}

	return string(plaintext), nil
}

// EncryptStringPtr chiffre un pointeur de chaîne (utile pour les champs optionnels)
func (cs *CryptoService) EncryptStringPtr(plaintext *string) (*string, error) {
	if plaintext == nil {
		return nil, nil
	}

	encrypted, err := cs.Encrypt(*plaintext)
	if err != nil {
		return nil, err
	}

	return &encrypted, nil
}

// DecryptStringPtr déchiffre un pointeur de chaîne
func (cs *CryptoService) DecryptStringPtr(ciphertext *string) (*string, error) {
	if ciphertext == nil {
		return nil, nil
	}

	decrypted, err := cs.Decrypt(*ciphertext)
	if err != nil {
		return nil, err
	}

	return &decrypted, nil
}
