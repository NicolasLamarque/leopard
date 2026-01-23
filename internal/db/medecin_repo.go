package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllMedecins récupère tous les médecins actifs
func (db *Database) GetAllMedecins() ([]models.Medecin, error) {
	var medecins []models.Medecin
	query := `SELECT * FROM medecins WHERE actif = 1 ORDER BY nomComplet`

	err := db.Select(&medecins, query)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération médecins: %w", err)
	}
	return medecins, nil
}

// GetMedecinByID récupère un médecin par son ID
func (db *Database) GetMedecinByID(id int) (*models.Medecin, error) {
	var medecin models.Medecin
	query := `SELECT * FROM medecins WHERE id = ?`

	err := db.Get(&medecin, query, id)
	if err != nil {
		return nil, fmt.Errorf("médecin ID %d non trouvé: %w", id, err)
	}
	return &medecin, nil
}

// GetMedecinByLicence récupère un médecin par son numéro de licence
func (db *Database) GetMedecinByLicence(licence string) (*models.Medecin, error) {
	var medecin models.Medecin
	query := `SELECT * FROM medecins WHERE licence = ?`

	err := db.Get(&medecin, query, licence)
	if err != nil {
		return nil, fmt.Errorf("médecin licence %s non trouvé: %w", licence, err)
	}
	return &medecin, nil
}

// CreateMedecin crée un nouveau médecin
func (db *Database) CreateMedecin(req models.CreateMedecinRequest, createdBy int) (int64, error) {
	query := `
        INSERT INTO medecins (
            licence, civilite, prenom, nom, nomComplet, statut,
            specialisation, service, departement, installation_principale, rls,
            telephone, extension, cellulaire, email, 
            adresse, code_postal, ville, pays,
            Note_fixe, actif, created_by
        ) VALUES (
            :licence, :civilite, :prenom, :nom, :nomComplet, :statut,
            :specialisation, :service, :departement, :installation_principale, :rls,
            :telephone, :extension, :cellulaire, :email,
            :adresse, :code_postal, :ville, :pays,
            :Note_fixe, :actif, :created_by
        )`
	result, err := db.NamedExec(query, map[string]interface{}{
		"licence":                 req.Licence,
		"civilite":                req.Civilite,
		"prenom":                  req.Prenom,
		"nom":                     req.Nom,
		"nomComplet":              req.NomComplet,
		"statut":                  req.Statut,
		"specialisation":          req.Specialisation,
		"service":                 req.Service,
		"departement":             req.Departement,
		"installation_principale": req.InstallationPrincipale,
		"rls":                     req.RLS,
		"telephone":               req.Telephone,
		"extension":               req.Extension,
		"cellulaire":              req.Cellulaire,
		"email":                   req.Email,
		"adresse":                 req.Adresse,
		"code_postal":             req.CodePostal,
		"ville":                   req.Ville,
		"pays":                    req.Pays,
		"Note_fixe":               req.NoteFixe,
		"actif":                   req.Actif,
		"created_by":              createdBy,
	})
	if err != nil {
		return 0, fmt.Errorf("erreur création médecin: %w", err)
	}
	return result.LastInsertId()
}

// UpdateMedecin met à jour un médecin
func (db *Database) UpdateMedecin(req models.UpdateMedecinRequest) error {
	query := `
        UPDATE medecins SET
            licence = :licence,
            civilite = :civilite,
            prenom = :prenom,
            nom = :nom,
            nomComplet = :nomComplet,
            statut = :statut,
            specialisation = :specialisation,
            service = :service,
            departement = :departement,
            installation_principale = :installation_principale,
            rls = :rls,
            telephone = :telephone,
            extension = :extension,
            cellulaire = :cellulaire,
            email = :email,
            adresse = :adresse,
            code_postal = :code_postal,
            ville = :ville,
            pays = :pays,
            Note_fixe = :Note_fixe,
            actif = :actif,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = :id`
	_, err := db.NamedExec(query, req)
	if err != nil {
		return fmt.Errorf("erreur mise à jour médecin %d: %w", req.ID, err)
	}
	return nil
}

// ArchiveMedecin désactive un médecin (suppression logique)
func (db *Database) ArchiveMedecin(id int) error {
	query := `UPDATE medecins SET actif = 0 WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur archivage médecin %d: %w", id, err)
	}
	return nil
}

// DeleteMedecin suppression physique (à utiliser avec précaution)
func (db *Database) DeleteMedecin(id int) error {
	query := `DELETE FROM medecins WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur suppression médecin %d: %w", id, err)
	}
	return nil
}

// SearchMedecins recherche par nom ou spécialisation
func (db *Database) SearchMedecins(searchTerm string) ([]models.Medecin, error) {
	var medecins []models.Medecin
	query := `
		SELECT * FROM medecins 
		WHERE actif = 1 
		AND (nomComplet LIKE ? OR specialisation LIKE ? OR licence LIKE ?)
		ORDER BY nomComplet`

	search := "%" + searchTerm + "%"
	err := db.Select(&medecins, query, search, search, search)
	if err != nil {
		return nil, fmt.Errorf("erreur recherche médecins: %w", err)
	}
	return medecins, nil
}

// À ajouter dans internal/db/medecin_repo.go

// GetClientsByMedecinLicence récupère tous les clients actifs d'un médecin
func (db *Database) GetClientsByMedecinLicence(licence string, cryptoSvc *crypto.CryptoService) ([]models.Client, error) {
	var allClients []models.Client
	var matchingClients []models.Client

	// Récupérer TOUS les clients actifs (on va filtrer après déchiffrement)
	query := `SELECT * FROM clients WHERE Actif = 1`

	err := db.Select(&allClients, query)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération clients: %w", err)
	}

	// Déchiffrer et filtrer
	for i := range allClients {
		// Déchiffrer le client
		if err := decryptClient(&allClients[i], cryptoSvc); err != nil {
			continue // Skip si erreur déchiffrement
		}

		// Comparer la licence déchiffrée
		if allClients[i].MedecinFamilleNoLicence != nil &&
			*allClients[i].MedecinFamilleNoLicence == licence {
			matchingClients = append(matchingClients, allClients[i])
		}
	}

	return matchingClients, nil
}

// GetMedecinClientsCount retourne le nombre de clients d'un médecin
func (db *Database) GetMedecinClientsCount(licence string, cryptoSvc *crypto.CryptoService) (int, error) {
	clients, err := db.GetClientsByMedecinLicence(licence, cryptoSvc)
	if err != nil {
		return 0, err
	}
	return len(clients), nil
}
