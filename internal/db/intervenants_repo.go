package database

import (
	"fmt"
	models "leopard/internal/model"

	"github.com/xuri/excelize/v2"
)

// GetAllIntervenants récupère la liste brute de la DB
func (db *Database) GetAllIntervenants() ([]models.Intervenant, error) {
	var list []models.Intervenant
	query := `SELECT * FROM intervenants ORDER BY nom_complet ASC`

	err := db.Select(&list, query)
	if err != nil {
		return nil, fmt.Errorf("erreur repo GetAllIntervenants: %w", err)
	}
	return list, nil
}

// GetIntervenantByID récupère un seul enregistrement
func (db *Database) GetIntervenantByID(id int) (*models.Intervenant, error) {
	var i models.Intervenant
	query := `SELECT * FROM intervenants WHERE id = ?`

	err := db.Get(&i, query, id)
	if err != nil {
		return nil, fmt.Errorf("intervenant ID %d non trouvé en DB: %w", id, err)
	}
	return &i, nil
}

// CreateIntervenant insère les données (déjà chiffrées par le service)
func (db *Database) CreateIntervenant(i models.Intervenant) (int64, error) {
	query := `
        INSERT INTO intervenants (
            nom_complet, titre_emploi, direction, installation,
            telephone, poste, cellulaire_pro, cellulaire_perso,
            courriel_personnel, courriel_professionnel, courrier_interne,
            actif, is_intervenant_social, numero_permis, ordre_professionnel,
            date_naissance, note_fixe, personne_ressource_reseau_dir
        ) VALUES (
            :nom_complet, :titre_emploi, :direction, :installation,
            :telephone, :poste, :cellulaire_pro, :cellulaire_perso,
            :courriel_personnel, :courriel_professionnel, :courrier_interne,
            :actif, :is_intervenant_social, :numero_permis, :ordre_professionnel,
            :date_naissance, :note_fixe, :personne_ressource_reseau_dir
        )`

	result, err := db.NamedExec(query, i)
	if err != nil {
		return 0, fmt.Errorf("erreur repo CreateIntervenant: %w", err)
	}
	return result.LastInsertId()
}

// UpdateIntervenant met à jour l'enregistrement
func (db *Database) UpdateIntervenant(i models.Intervenant) error {
	query := `
        UPDATE intervenants SET
            nom_complet = :nom_complet,
            titre_emploi = :titre_emploi,
            direction = :direction,
            installation = :installation,
            telephone = :telephone,
            poste = :poste,
            cellulaire_pro = :cellulaire_pro,
            cellulaire_perso = :cellulaire_perso,
            courriel_personnel = :courriel_personnel,
            courriel_professionnel = :courriel_professionnel,
            courrier_interne = :courrier_interne,
            actif = :actif,
            is_intervenant_social = :is_intervenant_social,
            numero_permis = :numero_permis,
            ordre_professionnel = :ordre_professionnel,
            date_naissance = :date_naissance,
            note_fixe = :note_fixe,
            personne_ressource_reseau_dir = :personne_ressource_reseau_dir,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = :id`

	_, err := db.NamedExec(query, i)
	if err != nil {
		return fmt.Errorf("erreur repo UpdateIntervenant: %w", err)
	}
	return nil
}

// DeleteIntervenant suppression physique
func (db *Database) DeleteIntervenant(id int) error {
	_, err := db.Exec(`DELETE FROM intervenants WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("erreur repo DeleteIntervenant: %w", err)
	}
	return nil
}

// SearchIntervenants recherche textuelle simple
func (db *Database) SearchIntervenants(term string) ([]models.Intervenant, error) {
	var list []models.Intervenant
	query := `SELECT * FROM intervenants 
              WHERE nom_complet LIKE ? 
              OR direction LIKE ? 
              OR installation LIKE ? 
              ORDER BY nom_complet ASC`

	pattern := "%" + term + "%"
	err := db.Select(&list, query, pattern, pattern, pattern)
	return list, err
}

// ImportIntervenantsFromExcel lit le fichier et écrit en DB
func (db *Database) ImportIntervenantsFromExcel(filePath string) (int, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	rows, _ := f.GetRows(f.GetSheetName(0))
	count := 0

	for i, row := range rows {
		if i == 0 || len(row) < 1 {
			continue
		}

		intervenant := models.Intervenant{
			NomComplet:            getString(row, 0),
			TitreEmploi:           getString(row, 1),
			Direction:             getString(row, 2),
			Installation:          getString(row, 3),
			Telephone:             getString(row, 4),
			Poste:                 getString(row, 5),
			CellulairePro:         getString(row, 6),
			CellulairePerso:       getString(row, 7),
			CourrielPersonnel:     getString(row, 8),
			CourrielProfessionnel: getString(row, 9),
			// Continue pour les autres colonnes...
		}

		_, err := db.CreateIntervenant(intervenant)
		if err == nil {
			count++
		}
	}
	return count, nil
}

// Petite fonction d'aide à mettre aussi dans intervenants_repo.go
func getString(row []string, index int) string {
	if index < len(row) {
		return row[index]
	}
	return ""
}
