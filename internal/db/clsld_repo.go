package database

import (
	"fmt"
	models "leopard/internal/model"
	"strings"
)

// GetAllCHSLD retourne tous les CHSLD
func (db *Database) GetAllCHSLD() ([]models.CHSLD, error) {
	var chslds []models.CHSLD
	query := `SELECT * FROM T_CHSLD ORDER BY TitreCHSLD`

	err := db.Select(&chslds, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des CHSLD: %w", err)
	}

	return chslds, nil
}

// GetCHSLDByID retourne un CHSLD par son ID
func (db *Database) GetCHSLDByID(id int) (*models.CHSLD, error) {
	var chsld models.CHSLD
	query := `SELECT * FROM T_CHSLD WHERE id = ?`

	err := db.Get(&chsld, query, id)
	if err != nil {
		return nil, fmt.Errorf("CHSLD non trouvé: %w", err)
	}

	return &chsld, nil
}

// SearchCHSLD recherche des CHSLD selon différents critères
func (db *Database) SearchCHSLD(nom, municipalite, region string) ([]models.CHSLD, error) {
	var chslds []models.CHSLD

	// Construction dynamique de la requête
	query := `SELECT * FROM T_CHSLD WHERE 1=1`
	args := []interface{}{}

	if nom != "" {
		query += ` AND TitreCHSLD LIKE ?`
		args = append(args, "%"+nom+"%")
	}

	if municipalite != "" {
		query += ` AND Municipalite LIKE ?`
		args = append(args, "%"+municipalite+"%")
	}

	if region != "" {
		query += ` AND Region = ?`
		args = append(args, region)
	}

	query += ` ORDER BY TitreCHSLD`

	err := db.Select(&chslds, query, args...)
	if err != nil {
		return nil, fmt.Errorf("erreur recherche CHSLD: %w", err)
	}

	return chslds, nil
}

// CreateCHSLD crée un nouveau CHSLD
func (db *Database) CreateCHSLD(chsld *models.CHSLD) error {
	query := `
		INSERT INTO T_CHSLD (
			Region, TitreCHSLD, Adresse, Municipalite, CodePostal,
			Telephone, telecopieur, Poste_Garde_infirmiere, Capacite,
			TypeCHSLD, Proprietaire, DateCertification, Statut,
			SourceURL, InfosCHSLD, Notes, DateAjout, DateMaj
		) VALUES (
			:Region, :TitreCHSLD, :Adresse, :Municipalite, :CodePostal,
			:Telephone, :telecopieur, :Poste_Garde_infirmiere, :Capacite,
			:TypeCHSLD, :Proprietaire, :DateCertification, :Statut,
			:SourceURL, :InfosCHSLD, :Notes, :DateAjout, :DateMaj
		)
	`

	result, err := db.NamedExec(query, chsld)
	if err != nil {
		return fmt.Errorf("erreur création CHSLD: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("erreur récupération ID: %w", err)
	}

	chsld.ID = int(id)
	return nil
}

// UpdateCHSLD met à jour un CHSLD existant
func (db *Database) UpdateCHSLD(chsld *models.CHSLD) error {
	query := `
		UPDATE T_CHSLD SET
			Region = :Region,
			TitreCHSLD = :TitreCHSLD,
			Adresse = :Adresse,
			Municipalite = :Municipalite,
			CodePostal = :CodePostal,
			Telephone = :Telephone,
			telecopieur = :telecopieur,
			Poste_Garde_infirmiere = :Poste_Garde_infirmiere,
			Capacite = :Capacite,
			TypeCHSLD = :TypeCHSLD,
			Proprietaire = :Proprietaire,
			DateCertification = :DateCertification,
			Statut = :Statut,
			SourceURL = :SourceURL,
			InfosCHSLD = :InfosCHSLD,
			Notes = :Notes,
			DateMaj = :DateMaj
		WHERE id = :id
	`

	result, err := db.NamedExec(query, chsld)
	if err != nil {
		return fmt.Errorf("erreur mise à jour CHSLD: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur vérification mise à jour: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("CHSLD non trouvé")
	}

	return nil
}

// DeleteCHSLD supprime un CHSLD
func (db *Database) DeleteCHSLD(id int) error {
	query := `DELETE FROM T_CHSLD WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur suppression CHSLD: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur vérification suppression: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("CHSLD non trouvé")
	}

	return nil
}

// GetCHSLDByRegion retourne tous les CHSLD d'une région
func (db *Database) GetCHSLDByRegion(region string) ([]models.CHSLD, error) {
	var chslds []models.CHSLD
	query := `SELECT * FROM T_CHSLD WHERE Region = ? ORDER BY TitreCHSLD`

	err := db.Select(&chslds, query, region)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération CHSLD par région: %w", err)
	}

	return chslds, nil
}

// GetCHSLDByStatut retourne tous les CHSLD avec un statut donné
func (db *Database) GetCHSLDByStatut(statut string) ([]models.CHSLD, error) {
	var chslds []models.CHSLD
	query := `SELECT * FROM T_CHSLD WHERE Statut = ? ORDER BY TitreCHSLD`

	err := db.Select(&chslds, query, statut)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération CHSLD par statut: %w", err)
	}

	return chslds, nil
}

// GetCHSLDStats retourne des statistiques sur les CHSLD
func (db *Database) GetCHSLDStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total
	var total int
	err := db.Get(&total, `SELECT COUNT(*) FROM T_CHSLD`)
	if err != nil {
		return nil, err
	}
	stats["total"] = total

	// Par statut
	var statutCounts []struct {
		Statut string `db:"Statut"`
		Count  int    `db:"count"`
	}
	err = db.Select(&statutCounts, `SELECT Statut, COUNT(*) as count FROM T_CHSLD GROUP BY Statut`)
	if err != nil {
		return nil, err
	}

	statutMap := make(map[string]int)
	for _, sc := range statutCounts {
		statutMap[strings.ToLower(sc.Statut)] = sc.Count
	}
	stats["par_statut"] = statutMap

	// Capacité totale
	var capaciteTotale int
	err = db.Get(&capaciteTotale, `SELECT COALESCE(SUM(Capacite), 0) FROM T_CHSLD`)
	if err != nil {
		return nil, err
	}
	stats["capacite_totale"] = capaciteTotale

	return stats, nil
}
