package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"leopard/internal/crypto"
	models "leopard/internal/model"
	"os"
	"strings"
)

// GetAllPharmacies récupère toutes les pharmacies
func (db *Database) GetAllPharmacies() ([]models.Pharmacie, error) {
	var pharmacies []models.Pharmacie

	query := `
		SELECT 
			id, NomOrganisation, Banniere, Adresse, Ville, Region, Tel, Fax,
			DimancheOuvert, HeureOuvertureDimanche, HeureFermetureDimanche,
			LundiOuvert, HeureOuvertureLundi, HeureFermetureLundi,
			MardiOuvert, HeureOuvertureMardi, HeureFermetureMardi,
			MercrediOuvert, HeureOuvertureMercredi, HeureFermetureMercredi,
			JeudiOuvert, HeureOuvertureJeudi, HeureFermetureJeudi,
			VendrediOuvert, HeureOuvertureVendredi, HeureFermetureVendredi,
			SamediOuvert, HeureOuvertureSamedi, HeureFermetureSamedi,
			DateMaj, note
		FROM pharmacies
		ORDER BY NomOrganisation ASC
	`

	err := db.Select(&pharmacies, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des pharmacies: %w", err)
	}

	return pharmacies, nil
}

// GetPharmacieByID récupère une pharmacie par son ID
func (db *Database) GetPharmacieByID(id int) (*models.Pharmacie, error) {
	var pharmacie models.Pharmacie

	query := `
		SELECT 
			id, NomOrganisation, Banniere, Adresse, Ville, Region, Tel, Fax,
			DimancheOuvert, HeureOuvertureDimanche, HeureFermetureDimanche,
			LundiOuvert, HeureOuvertureLundi, HeureFermetureLundi,
			MardiOuvert, HeureOuvertureMardi, HeureFermetureMardi,
			MercrediOuvert, HeureOuvertureMercredi, HeureFermetureMercredi,
			JeudiOuvert, HeureOuvertureJeudi, HeureFermetureJeudi,
			VendrediOuvert, HeureOuvertureVendredi, HeureFermetureVendredi,
			SamediOuvert, HeureOuvertureSamedi, HeureFermetureSamedi,
			DateMaj, note
		FROM pharmacies
		WHERE id = ?
	`

	err := db.Get(&pharmacie, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pharmacie avec l'ID %d non trouvée", id)
		}
		return nil, fmt.Errorf("erreur lors de la récupération de la pharmacie: %w", err)
	}

	return &pharmacie, nil
}

// CreatePharmacie crée une nouvelle pharmacie
func (db *Database) CreatePharmacie(p models.Pharmacie) (int64, error) {
	query := `
		INSERT INTO pharmacies (
			NomOrganisation, Banniere, Adresse, Ville, Region, Tel, Fax,
			DimancheOuvert, HeureOuvertureDimanche, HeureFermetureDimanche,
			LundiOuvert, HeureOuvertureLundi, HeureFermetureLundi,
			MardiOuvert, HeureOuvertureMardi, HeureFermetureMardi,
			MercrediOuvert, HeureOuvertureMercredi, HeureFermetureMercredi,
			JeudiOuvert, HeureOuvertureJeudi, HeureFermetureJeudi,
			VendrediOuvert, HeureOuvertureVendredi, HeureFermetureVendredi,
			SamediOuvert, HeureOuvertureSamedi, HeureFermetureSamedi
		) VALUES (
			:NomOrganisation, :Banniere, :Adresse, :Ville, :Region, :Tel, :Fax,
			:DimancheOuvert, :HeureOuvertureDimanche, :HeureFermetureDimanche,
			:LundiOuvert, :HeureOuvertureLundi, :HeureFermetureLundi,
			:MardiOuvert, :HeureOuvertureMardi, :HeureFermetureMardi,
			:MercrediOuvert, :HeureOuvertureMercredi, :HeureFermetureMercredi,
			:JeudiOuvert, :HeureOuvertureJeudi, :HeureFermetureJeudi,
			:VendrediOuvert, :HeureOuvertureVendredi, :HeureFermetureVendredi,
			:SamediOuvert, :HeureOuvertureSamedi, :HeureFermetureSamedi
		)
	`

	result, err := db.NamedExec(query, p)
	if err != nil {
		return 0, fmt.Errorf("erreur création pharmacie: %w", err)
	}

	return result.LastInsertId()
}

// UpdatePharmacie met à jour une pharmacie existante
func (db *Database) UpdatePharmacie(p models.Pharmacie) error {
	query := `
		UPDATE pharmacies SET
			NomOrganisation = :NomOrganisation,
			Banniere = :Banniere,
			Adresse = :Adresse,
			Ville = :Ville,
			Region = :Region,
			Tel = :Tel,
			Fax = :Fax,
			DimancheOuvert = :DimancheOuvert,
			HeureOuvertureDimanche = :HeureOuvertureDimanche,
			HeureFermetureDimanche = :HeureFermetureDimanche,
			LundiOuvert = :LundiOuvert,
			HeureOuvertureLundi = :HeureOuvertureLundi,
			HeureFermetureLundi = :HeureFermetureLundi,
			MardiOuvert = :MardiOuvert,
			HeureOuvertureMardi = :HeureOuvertureMardi,
			HeureFermetureMardi = :HeureFermetureMardi,
			MercrediOuvert = :MercrediOuvert,
			HeureOuvertureMercredi = :HeureOuvertureMercredi,
			HeureFermetureMercredi = :HeureFermetureMercredi,
			JeudiOuvert = :JeudiOuvert,
			HeureOuvertureJeudi = :HeureOuvertureJeudi,
			HeureFermetureJeudi = :HeureFermetureJeudi,
			VendrediOuvert = :VendrediOuvert,
			HeureOuvertureVendredi = :HeureOuvertureVendredi,
			HeureFermetureVendredi = :HeureFermetureVendredi,
			SamediOuvert = :SamediOuvert,
			HeureOuvertureSamedi = :HeureOuvertureSamedi,
			HeureFermetureSamedi = :HeureFermetureSamedi,
			DateMaj = CURRENT_TIMESTAMP
		WHERE id = :ID
	`

	_, err := db.NamedExec(query, p)
	if err != nil {
		return fmt.Errorf("erreur mise à jour pharmacie %d: %w", p.ID, err)
	}

	return nil
}

// DeletePharmacie supprime définitivement une pharmacie
func (db *Database) DeletePharmacie(id int) error {
	_, err := db.Exec("UPDATE clients SET pharmacie_id = NULL WHERE pharmacie_id = ?", id)
	if err != nil {
		return fmt.Errorf("erreur mise à jour clients: %w", err)
	}

	_, err = db.Exec("DELETE FROM pharmacies WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("erreur suppression pharmacie: %w", err)
	}

	return nil
}

// GetPharmacieForClient retourne la pharmacie d'un client
func (db *Database) GetPharmacieForClient(clientID int) (*models.Pharmacie, error) {
	var pharmacie models.Pharmacie

	query := `
		SELECT 
			p.id, p.NomOrganisation, p.Banniere, p.Adresse, p.Ville, p.Region, p.Tel, p.Fax,
			p.DimancheOuvert, p.HeureOuvertureDimanche, p.HeureFermetureDimanche,
			p.LundiOuvert, p.HeureOuvertureLundi, p.HeureFermetureLundi,
			p.MardiOuvert, p.HeureOuvertureMardi, p.HeureFermetureMardi,
			p.MercrediOuvert, p.HeureOuvertureMercredi, p.HeureFermetureMercredi,
			p.JeudiOuvert, p.HeureOuvertureJeudi, p.HeureFermetureJeudi,
			p.VendrediOuvert, p.HeureOuvertureVendredi, p.HeureFermetureVendredi,
			p.SamediOuvert, p.HeureOuvertureSamedi, p.HeureFermetureSamedi,
			p.DateMaj, p.note
		FROM pharmacies p
		INNER JOIN clients c ON c.pharmacie_id = p.id
		WHERE c.id = ?
	`

	err := db.Get(&pharmacie, query, clientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("erreur récupération pharmacie du client: %w", err)
	}

	return &pharmacie, nil
}

// GetClientsByPharmacie retourne la liste des clients d'une pharmacie (déchiffrés)
func (db *Database) GetClientsByPharmacie(pharmacieID int, cryptoSvc *crypto.CryptoService) ([]models.ClientPharmacieInfo, error) {
	type clientTemp struct {
		ID               int     `db:"id"`
		Nom              *string `db:"nom"`
		Prenom           *string `db:"prenom"`
		NoDossierLeopard *string `db:"no_dossier_leopard"`
		DCD              int     `db:"dcd"`
		Actif            int     `db:"Actif"`
	}

	var tempClients []clientTemp

	query := `
        SELECT id, nom, prenom, no_dossier_leopard, dcd, Actif
        FROM clients
        WHERE pharmacie_id = ? 
        ORDER BY dcd ASC, nom ASC
    `

	err := db.Select(&tempClients, query, pharmacieID)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération clients: %w", err)
	}

	var clients []models.ClientPharmacieInfo
	for _, temp := range tempClients {
		client := models.ClientPharmacieInfo{
			ID:    temp.ID,
			DCD:   temp.DCD,
			Actif: temp.Actif,
		}

		if temp.Nom != nil {
			dec, _ := cryptoSvc.DecryptStringPtr(temp.Nom)
			if dec != nil {
				client.Nom = *dec
			}
		}
		if temp.Prenom != nil {
			dec, _ := cryptoSvc.DecryptStringPtr(temp.Prenom)
			if dec != nil {
				client.Prenom = *dec
			}
		}
		if temp.NoDossierLeopard != nil {
			client.NoDossierLeopard = *temp.NoDossierLeopard
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// normaliserOuvert convertit toute valeur CSV en 0 ou 1 propre
// Accepte: "1", "0", "Ouvert", "Fermé", "oui", "non", "true", "false", "O", "F", "Y", "N"
func normaliserOuvert(val string) int {
	switch strings.ToLower(strings.TrimSpace(val)) {
	case "1", "ouvert", "oui", "true", "yes", "y", "o":
		return 1
	default:
		return 0
	}
}

// normaliserHeure nettoie une heure CSV (retire espaces, normalise format HH:MM)
func normaliserHeure(val string) string {
	return strings.TrimSpace(val)
}

// ImportPharmaciesFromCSV importe les pharmacies depuis un fichier CSV.
//
// Format CSV attendu (séparateur ';', 31 colonnes) :
//
//	[0]  id
//	[1]  Nom de lorganisation  → NomOrganisation
//	[2]  Banniere
//	[3]  Adresse
//	[4]  Ville
//	[5]  Region
//	[6]  Tel
//	[7]  Fax
//	[8]  Dimanche1             → DimancheOuvert  (normalisé 0/1)
//	[9]  Heure1ureDimanche     → HeureOuvertureDimanche
//	[10] Heures1Dim            → HeureFermetureDimanche
//	[11] HeureOuverLundi       → LundiOuvert      (normalisé 0/1)
//	[12] Heure1ureLundi        → HeureOuvertureLundi
//	[13] HeureFermetureLundi   → HeureFermetureLundi
//	[14] Mardi1                → MardiOuvert
//	[15] Heure1ureMardi        → HeureOuvertureMardi
//	[16] HeureFermetureMardi   → HeureFermetureMardi
//	[17] Mercredi1             → MercrediOuvert
//	[18] Heure1ureMercredi     → HeureOuvertureMercredi
//	[19] HeureFermetureMercredi
//	[20] Jeudi1                → JeudiOuvert
//	[21] Heure1ureJeudi        → HeureOuvertureJeudi
//	[22] HeureFermetureJeudi
//	[23] Vendredi1             → VendrediOuvert
//	[24] Heure1ureVendredi     → HeureOuvertureVendredi
//	[25] HeureFermetureVendredi
//	[26] Samedi1               → SamediOuvert
//	[27] Heure1ureSamedi       → HeureOuvertureSamedi
//	[28] HeureFermetureSamedi
//	[29] DatMaj                → DateMaj
//	[30] Note                  → note
//
// Stratégie : INSERT OR REPLACE (upsert sur l'id CSV)
// Retourne : (nb importées, nb erreurs, error fatale)
func (db *Database) ImportPharmaciesFromCSV(filePath string) (int, int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, 0, fmt.Errorf("impossible d'ouvrir le fichier: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	// Sauter le header
	if _, err = reader.Read(); err != nil {
		return 0, 0, fmt.Errorf("erreur lecture header: %w", err)
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return 0, 0, fmt.Errorf("erreur démarrage transaction: %w", err)
	}

	query := `
		INSERT OR REPLACE INTO pharmacies (
			id, NomOrganisation, Banniere, Adresse, Ville, Region, Tel, Fax,
			DimancheOuvert, HeureOuvertureDimanche, HeureFermetureDimanche,
			LundiOuvert,    HeureOuvertureLundi,    HeureFermetureLundi,
			MardiOuvert,    HeureOuvertureMardi,    HeureFermetureMardi,
			MercrediOuvert, HeureOuvertureMercredi, HeureFermetureMercredi,
			JeudiOuvert,    HeureOuvertureJeudi,    HeureFermetureJeudi,
			VendrediOuvert, HeureOuvertureVendredi, HeureFermetureVendredi,
			SamediOuvert,   HeureOuvertureSamedi,   HeureFermetureSamedi,
			DateMaj, note
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?,
			?, ?, ?,
			?, ?, ?,
			?, ?, ?,
			?, ?, ?,
			?, ?, ?,
			?, ?, ?,
			?, ?
		)
	`

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("erreur préparation requête: %w", err)
	}
	defer stmt.Close()

	imported := 0
	skipped := 0
	lineNum := 1

	for {
		record, err := reader.Read()
		lineNum++

		if err == io.EOF {
			break
		}
		if err != nil {
			skipped++
			continue
		}
		if len(record) < 31 {
			skipped++
			continue
		}

		// Mapping CSV → DB avec normalisation
		_, execErr := stmt.Exec(
			// Identité
			strings.TrimSpace(record[0]), // id
			strings.TrimSpace(record[1]), // NomOrganisation
			strings.TrimSpace(record[2]), // Banniere
			strings.TrimSpace(record[3]), // Adresse
			strings.TrimSpace(record[4]), // Ville
			strings.TrimSpace(record[5]), // Region
			strings.TrimSpace(record[6]), // Tel
			strings.TrimSpace(record[7]), // Fax
			// Dimanche
			normaliserOuvert(record[8]), // DimancheOuvert
			normaliserHeure(record[9]),  // HeureOuvertureDimanche
			normaliserHeure(record[10]), // HeureFermetureDimanche
			// Lundi
			normaliserOuvert(record[11]), // LundiOuvert
			normaliserHeure(record[12]),  // HeureOuvertureLundi
			normaliserHeure(record[13]),  // HeureFermetureLundi
			// Mardi
			normaliserOuvert(record[14]), // MardiOuvert
			normaliserHeure(record[15]),  // HeureOuvertureMardi
			normaliserHeure(record[16]),  // HeureFermetureMardi
			// Mercredi
			normaliserOuvert(record[17]), // MercrediOuvert
			normaliserHeure(record[18]),  // HeureOuvertureMercredi
			normaliserHeure(record[19]),  // HeureFermetureMercredi
			// Jeudi
			normaliserOuvert(record[20]), // JeudiOuvert
			normaliserHeure(record[21]),  // HeureOuvertureJeudi
			normaliserHeure(record[22]),  // HeureFermetureJeudi
			// Vendredi
			normaliserOuvert(record[23]), // VendrediOuvert
			normaliserHeure(record[24]),  // HeureOuvertureVendredi
			normaliserHeure(record[25]),  // HeureFermetureVendredi
			// Samedi
			normaliserOuvert(record[26]), // SamediOuvert
			normaliserHeure(record[27]),  // HeureOuvertureSamedi
			normaliserHeure(record[28]),  // HeureFermetureSamedi
			// Meta
			strings.TrimSpace(record[29]), // DateMaj
			strings.TrimSpace(record[30]), // note
		)

		if execErr != nil {
			skipped++
		} else {
			imported++
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, 0, fmt.Errorf("erreur commit: %w", err)
	}

	return imported, skipped, nil
}
