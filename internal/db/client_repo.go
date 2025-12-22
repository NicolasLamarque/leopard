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
	// On utilise NamedExec pour ne plus se tromper dans l'ordre des arguments
	query := `
        INSERT INTO clients (
            nom, prenom, date_naissance, telephone, email, adresse,
            numero_assurance_maladie, numero_securite_sociale, no_hcm, 
            no_chaur, no_dossier_leopard, created_by
        ) VALUES (
            :nom, :prenom, :date_naissance, :telephone, :email, :adresse,
            :numero_assurance_maladie, :numero_securite_sociale, :no_hcm, 
            :no_chaur, :no_dossier_leopard, :created_by
        )`

	// On peut préparer une map pour injecter created_by qui n'est pas dans la struct req
	// Ou simplement ajouter le tag db:"created_by" à createdBy si on veut être fancy.
	// Ici, on reste simple :

	resultat, err := db.NamedExec(query, map[string]interface{}{
		"nom":                      req.Nom,
		"prenom":                   req.Prenom,
		"date_naissance":           req.DateNaissance,
		"telephone":                req.Telephone,
		"email":                    req.Email,
		"adresse":                  req.Adresse,
		"numero_assurance_maladie": req.NumeroAssuranceMaladie,
		"numero_securite_sociale":  req.NumeroSecuriteSociale,
		"no_hcm":                   req.NoHCM,
		"no_chaur":                 req.NoCHAUR,
		"no_dossier_leopard":       req.NoDossierLeopard,
		"created_by":               createdBy,
	})

	if err != nil {
		return 0, fmt.Errorf("erreur création client: %w", err)
	}

	return resultat.LastInsertId()
}

// UpdateClient met à jour un client existant
func (db *Database) UpdateClient(req models.UpdateClientRequest) error {
	// NamedExec est magique ici : il prend ta struct req et apparie
	// les noms :nom avec les tags db:"nom" de la struct.
	query := `
		UPDATE clients SET
			nom = :nom, 
			prenom = :prenom, 
			date_naissance = :date_naissance, 
			telephone = :telephone, 
			email = :email, 
			adresse = :adresse, 
			numero_assurance_maladie = :numero_assurance_maladie,
			numero_securite_sociale = :numero_securite_sociale, 
			no_hcm = :no_hcm, 
			no_chaur = :no_chaur,
			no_dossier_leopard = :no_dossier_leopard
		WHERE id = :id`

	_, err := db.NamedExec(query, req)
	if err != nil {
		return fmt.Errorf("erreur mise à jour client %d: %w", req.ID, err)
	}

	return nil
}

// DeleteClient reste simple
func (db *Database) DeleteClient(id int) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = ?", id)
	return err
}
