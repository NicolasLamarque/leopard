package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"leopard/internal/crypto"
	models "leopard/internal/model"
	"os"
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
		WHERE id = :id
	`

	_, err := db.NamedExec(query, p)
	if err != nil {
		return fmt.Errorf("erreur mise à jour pharmacie %d: %w", p.ID, err)
	}

	return nil
}

// DeletePharmacie supprime définitivement une pharmacie
func (db *Database) DeletePharmacie(id int) error {
	// Met à NULL le champ pharmacie_id dans les clients
	_, err := db.Exec("UPDATE clients SET pharmacie_id = NULL WHERE pharmacie_id = ?", id)
	if err != nil {
		return fmt.Errorf("erreur mise à jour clients: %w", err)
	}

	// Supprime la pharmacie
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

// GetClientsByPharmacie retourne la liste SIMPLE des clients d'une pharmacie (déchiffrés)
func (db *Database) GetClientsByPharmacie(pharmacieID int, cryptoSvc *crypto.CryptoService) ([]models.ClientPharmacieInfo, error) {
	// 1. Le moule temporaire doit avoir les colonnes dcd et actif
	type clientTemp struct {
		ID               int     `db:"id"`
		Nom              *string `db:"nom"`
		Prenom           *string `db:"prenom"`
		NoDossierLeopard *string `db:"no_dossier_leopard"`
		DCD              int     `db:"dcd"`   // On les récupère ici
		Actif            int     `db:"Actif"` // On les récupère ici
	}

	var tempClients []clientTemp

	// 2. La requête qui ne "brise" rien (on prend tout)
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
		// 3. On initialise le client avec les infos de statut
		client := models.ClientPharmacieInfo{
			ID:    temp.ID,
			DCD:   temp.DCD,
			Actif: temp.Actif,
		}

		// Déchiffrement du nom
		if temp.Nom != nil {
			dec, _ := cryptoSvc.DecryptStringPtr(temp.Nom)
			if dec != nil {
				client.Nom = *dec
			}
		}
		// Déchiffrement du prénom
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
func (r *Database) ImportPharmaciesFromCSV(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.LazyQuotes = true

	// Sauter le header
	_, err = reader.Read()
	if err != nil {
		return 0, err
	}

	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}

	// Utilisation des paramètres nommés (:nom)
	query := `INSERT OR REPLACE INTO pharmacies (
        id, NomOrganisation, Banniere, Adresse, Ville, Region, Tel, Fax,
        DimancheOuvert, HeureOuvertureDimanche, HeuresOuvertDim,
        HeureOuverLundi, HeureOuvertureLundi, HeureFermetureLundi,
        MardiOuvert, HeureOuvertureMardi, HeureFermetureMardi,
        MercrediOuvert, HeureOuvertureMercredi, HeureFermetureMercredi,
        JeudiOuvert, HeureOuvertureJeudi, HeureFermetureJeudi,
        VendrediOuvert, HeureOuvertureVendredi, HeureFermetureVendredi,
        SamediOuvert, HeureOuvertureSamedi, HeureFermetureSamedi,
        DateMaj, note
    ) VALUES (
        :id, :nom, :ban, :adr, :ville, :reg, :tel, :fax,
        :d_o, :d_h_o, :d_h_f,
        :l_o, :l_h_o, :l_h_f,
        :ma_o, :ma_h_o, :ma_h_f,
        :me_o, :me_h_o, :me_h_f,
        :j_o, :j_h_o, :j_h_f,
        :v_o, :v_h_o, :v_h_f,
        :s_o, :s_h_o, :s_h_f,
        :maj, :note
    )`

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if len(record) < 31 {
			continue
		} // Protection contre les lignes tronquées

		_, err = tx.Exec(query,
			sql.Named("id", record[0]),
			sql.Named("nom", record[1]),
			sql.Named("ban", record[2]),
			sql.Named("adr", record[3]),
			sql.Named("ville", record[4]),
			sql.Named("reg", record[5]),
			sql.Named("tel", record[6]),
			sql.Named("fax", record[7]),
			sql.Named("d_o", record[8]),
			sql.Named("d_h_o", record[9]),
			sql.Named("d_h_f", record[10]), // HeuresOuvertDim
			sql.Named("l_o", record[11]),   // HeureOuverLundi
			sql.Named("l_h_o", record[12]),
			sql.Named("l_h_f", record[13]),
			sql.Named("ma_o", record[14]),
			sql.Named("ma_h_o", record[15]),
			sql.Named("ma_h_f", record[16]),
			sql.Named("me_o", record[17]),
			sql.Named("me_h_o", record[18]),
			sql.Named("me_h_f", record[19]),
			sql.Named("j_o", record[20]),
			sql.Named("j_h_o", record[21]),
			sql.Named("j_h_f", record[22]),
			sql.Named("v_o", record[23]),
			sql.Named("v_h_o", record[24]),
			sql.Named("v_h_f", record[25]),
			sql.Named("s_o", record[26]),
			sql.Named("s_h_o", record[27]),
			sql.Named("s_h_f", record[28]),
			sql.Named("maj", record[29]),
			sql.Named("note", record[30]),
		)

		if err == nil {
			count++
		}
	}

	return count, tx.Commit()
}
