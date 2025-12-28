package database

import (
	"fmt"
	models "leopard/internal/model"
)

// GetAllClients récupère tous les clients (Ultra simplifié avec sqlx)
func (db *Database) GetAllClients() ([]models.Client, error) {
	var listeClients []models.Client

	// Plus besoin de scanClient, sqlx fait le mapping tout seul avec les tags db:""
	query := `SELECT * FROM clients ORDER BY nom, prenom`

	err := db.Select(&listeClients, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des clients: %w", err)
	}

	return listeClients, nil
}

// GetClientByID récupère un client unique par son ID
func (db *Database) GetClientByID(id int) (*models.Client, error) {
	var clientRecupere models.Client

	query := `SELECT * FROM clients WHERE id = ?`

	// db.Get est l'équivalent de QueryRow + Scan pour un seul objet
	err := db.Get(&clientRecupere, query, id)
	if err != nil {
		return nil, fmt.Errorf("client avec l'ID %d non trouvé: %w", id, err)
	}

	return &clientRecupere, nil
}

// CreateClient crée un nouveau client
func (db *Database) CreateClient(req models.CreateClientRequest, createdBy int) (int64, error) {
	query := `
        INSERT INTO clients (
            nom, prenom, date_naissance, telephone, cellulaire, email, adresse,
            code_postal, ville, pays,
            numero_assurance_maladie, numero_securite_sociale, no_hcm, no_chaur, 
            no_dossier_leopard, medecin_famille_No_Licence,
            notaire_id, pivot_id, rpa_id, chsld_id, ri_id,
            note_fixe, Actif, dcd, created_by
        ) VALUES (
            :nom, :prenom, :date_naissance, :telephone, :cellulaire, :email, :adresse,
            :code_postal, :ville, :pays,
            :numero_assurance_maladie, :numero_securite_sociale, :no_hcm, :no_chaur,
            :no_dossier_leopard, :medecin_famille_No_Licence,
            :notaire_id, :pivot_id, :rpa_id, :chsld_id, :ri_id,
            :note_fixe, :Actif, :dcd, :created_by
        )`

	resultat, err := db.NamedExec(query, map[string]interface{}{
		"nom":                        req.Nom,
		"prenom":                     req.Prenom,
		"date_naissance":             req.DateNaissance,
		"telephone":                  req.Telephone,
		"cellulaire":                 req.Cellulaire,
		"email":                      req.Email,
		"adresse":                    req.Adresse,
		"code_postal":                req.CodePostal,
		"ville":                      req.Ville,
		"pays":                       req.Pays,
		"numero_assurance_maladie":   req.NumeroAssuranceMaladie,
		"numero_securite_sociale":    req.NumeroSecuriteSociale,
		"no_hcm":                     req.NoHCM,
		"no_chaur":                   req.NoCHAUR,
		"no_dossier_leopard":         req.NoDossierLeopard,
		"medecin_famille_No_Licence": req.MedecinFamilleNoLicence,
		"notaire_id":                 req.NotaireID,
		"pivot_id":                   req.PivotID,
		"rpa_id":                     req.RPAID,
		"chsld_id":                   req.CHSLDID,
		"ri_id":                      req.RIID,
		"note_fixe":                  req.NoteFixe,
		"Actif":                      req.Actif,
		"dcd":                        req.DCD,
		"created_by":                 createdBy,
	})

	if err != nil {
		return 0, fmt.Errorf("erreur création client: %w", err)
	}

	return resultat.LastInsertId()
}

// UpdateClient met à jour un client existant
func (db *Database) UpdateClient(req models.UpdateClientRequest) error {
	query := `
		UPDATE clients SET
			nom = :nom,
			prenom = :prenom,
			date_naissance = :date_naissance,
			telephone = :telephone,
			cellulaire = :cellulaire,
			email = :email,
			adresse = :adresse,
			code_postal = :code_postal,
			ville = :ville,
			pays = :pays,
			numero_assurance_maladie = :numero_assurance_maladie,
			numero_securite_sociale = :numero_securite_sociale,
			no_hcm = :no_hcm,
			no_chaur = :no_chaur,
			no_dossier_leopard = :no_dossier_leopard,
			medecin_famille_No_Licence = :medecin_famille_No_Licence,
			notaire_id = :notaire_id,
			pivot_id = :pivot_id,
			rpa_id = :rpa_id,
			chsld_id = :chsld_id,
			ri_id = :ri_id,
			note_fixe = :note_fixe,
			Actif = :Actif,
			dcd = :dcd
		WHERE id = :id`

	_, err := db.NamedExec(query, req)
	if err != nil {
		return fmt.Errorf("erreur mise à jour client %d: %w", req.ID, err)
	}

	return nil
}

// ArchiveClient désactive un client (suppression logique)
func (db *Database) ArchiveClient(id int) error {
	query := `
		UPDATE clients
		SET Actif = 0
		WHERE id = ?
	`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur archivage client %d: %w", id, err)
	}
	return nil
}
// supprimer DeleteClient supprime un client de la base de données
// DeleteClient supprime définitivement un client
func (db *Database) DeleteClient(id int) error {
	query := `DELETE FROM clients WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur suppression physique client %d: %w", id, err)
	}
	return nil
}

