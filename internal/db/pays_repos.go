package database

import (
	"fmt"
	models "leopard/internal/model"
	"strings"
)

// ImportPaysResult est défini dans le package main (pays_handler.go)
// On le redéclare ici pour éviter l'import circulaire
type ImportPaysResult struct {
	Total   int    `json:"total"`
	Success int    `json:"success"`
	Errors  int    `json:"errors"`
	Message string `json:"message"`
}

// GetAllPays récupère tous les pays triés par nom
func (db *Database) GetAllPays() ([]models.Pays, error) {
	var pays []models.Pays
	query := `SELECT * FROM pays ORDER BY pays`

	err := db.Select(&pays, query)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération pays: %w", err)
	}
	return pays, nil
}

// GetPaysByID récupère un pays par son ID
func (db *Database) GetPaysByID(id int) (*models.Pays, error) {
	var pays models.Pays
	query := `SELECT * FROM pays WHERE id = ?`

	err := db.Get(&pays, query, id)
	if err != nil {
		return nil, fmt.Errorf("pays ID %d non trouvé: %w", id, err)
	}
	return &pays, nil
}

// GetPaysByAlpha2 récupère un pays par son code Alpha-2
func (db *Database) GetPaysByAlpha2(alpha2 string) (*models.Pays, error) {
	var pays models.Pays
	query := `SELECT * FROM pays WHERE alpha2 = ?`

	err := db.Get(&pays, query, strings.ToUpper(alpha2))
	if err != nil {
		return nil, fmt.Errorf("pays Alpha-2 %s non trouvé: %w", alpha2, err)
	}
	return &pays, nil
}

// GetPaysByAlpha3 récupère un pays par son code Alpha-3
func (db *Database) GetPaysByAlpha3(alpha3 string) (*models.Pays, error) {
	var pays models.Pays
	query := `SELECT * FROM pays WHERE alpha3 = ?`

	err := db.Get(&pays, query, strings.ToUpper(alpha3))
	if err != nil {
		return nil, fmt.Errorf("pays Alpha-3 %s non trouvé: %w", alpha3, err)
	}
	return &pays, nil
}

// GetPaysCount retourne le nombre de pays en base
func (db *Database) GetPaysCount() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM pays`

	err := db.Get(&count, query)
	if err != nil {
		// Si la table n'existe pas encore, retourner 0
		return 0, nil
	}
	return count, nil
}

// SearchPays recherche des pays par nom
func (db *Database) SearchPays(searchTerm string) ([]models.Pays, error) {
	var pays []models.Pays
	query := `
		SELECT * FROM pays 
		WHERE pays LIKE ? OR alpha2 LIKE ? OR alpha3 LIKE ?
		ORDER BY pays`

	search := "%" + searchTerm + "%"
	err := db.Select(&pays, query, search, search, search)
	if err != nil {
		return nil, fmt.Errorf("erreur recherche pays: %w", err)
	}
	return pays, nil
}

// GetPaysForSelect retourne la liste des pays formatée pour un select
func (db *Database) GetPaysForSelect() ([]models.PaysListItemForSelect, error) {
	var pays []models.PaysListItemForSelect
	query := `SELECT id, pays, alpha2, alpha3, numerique FROM pays ORDER BY pays`

	err := db.Select(&pays, query)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération pays pour select: %w", err)
	}
	return pays, nil
}
