package database

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// GenerateLeopardNumber génère un numéro de dossier unique
// Format: NOM(3) + PRENOM(1) + _ + YYYYMMDD + _ + COUNTER(3)
// Exemple: DUPA_20250105_001
func (db *Database) GenerateLeopardNumber(nom, prenom string) (string, error) {
	// 1. Nettoyer et normaliser
	cleanNom := normalizeString(nom)
	cleanPrenom := normalizeString(prenom)

	if cleanNom == "" || cleanPrenom == "" {
		return "", fmt.Errorf("nom et prénom requis")
	}

	// 2. Extraire les parties
	nomPart := extractChars(cleanNom, 3)
	prenomPart := extractChars(cleanPrenom, 1)

	// 3. Date du jour
	today := time.Now().Format("20060102") // YYYYMMDD

	// 4. Base du numéro (sans compteur)
	baseNumber := fmt.Sprintf("%s%s_%s", nomPart, prenomPart, today)

	// 5. Trouver le prochain compteur disponible
	counter := 1
	var finalNumber string

	for {
		finalNumber = fmt.Sprintf("%s_%03d", baseNumber, counter)

		// Vérifier si ce numéro existe déjà
		var exists int
		err := db.QueryRow(
			"SELECT COUNT(*) FROM clients WHERE no_dossier_leopard = ?",
			finalNumber,
		).Scan(&exists)

		if err != nil {
			return "", fmt.Errorf("erreur vérification: %w", err)
		}

		// Si le numéro n'existe pas, on le prend !
		if exists == 0 {
			break
		}

		// Sinon, incrémenter et réessayer
		counter++

		// Sécurité : max 999 clients avec le même code le même jour
		if counter > 999 {
			return "", fmt.Errorf("limite de collisions atteinte pour %s", baseNumber)
		}
	}

	return finalNumber, nil
}

// normalizeString enlève les accents et caractères spéciaux
func normalizeString(s string) string {
	// Transformer en majuscules
	s = strings.ToUpper(s)

	// Enlever les accents (é → E, à → A, etc.)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)

	// Garder seulement les lettres
	var cleaned strings.Builder
	for _, r := range result {
		if unicode.IsLetter(r) {
			cleaned.WriteRune(r)
		}
	}

	return cleaned.String()
}

// extractChars extrait N caractères d'une chaîne, pad avec X si trop court
func extractChars(s string, n int) string {
	runes := []rune(s)

	if len(runes) >= n {
		return string(runes[:n])
	}

	// Pad avec X si trop court
	result := s
	for len(result) < n {
		result += "X"
	}

	return result
}

// GetLeopardNumberInfo retourne les infos d'un numéro Leopard
type LeopardNumberInfo struct {
	Number       string    `json:"number"`
	NomPart      string    `json:"nom_part"`
	PrenomPart   string    `json:"prenom_part"`
	DateCreation time.Time `json:"date_creation"`
	Counter      int       `json:"counter"`
}

// ParseLeopardNumber décompose un numéro Leopard
// DUPA_20250105_001 → {nom: DUPA, prenom: A, date: 2025-01-05, counter: 1}
func ParseLeopardNumber(leopardNumber string) (*LeopardNumberInfo, error) {
	parts := strings.Split(leopardNumber, "_")

	if len(parts) != 3 {
		return nil, fmt.Errorf("format invalide: attendu NOM_DATE_COUNTER")
	}

	// Extraire nom/prénom
	if len(parts[0]) < 4 {
		return nil, fmt.Errorf("partie nom invalide")
	}
	nomPart := parts[0][:3]
	prenomPart := string(parts[0][3])

	// Parser la date
	dateStr := parts[1]
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return nil, fmt.Errorf("date invalide: %w", err)
	}

	// Parser le compteur
	var counter int
	_, err = fmt.Sscanf(parts[2], "%03d", &counter)
	if err != nil {
		return nil, fmt.Errorf("compteur invalide: %w", err)
	}

	return &LeopardNumberInfo{
		Number:       leopardNumber,
		NomPart:      nomPart,
		PrenomPart:   prenomPart,
		DateCreation: date,
		Counter:      counter,
	}, nil
}
