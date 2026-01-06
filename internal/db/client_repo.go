package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllClients récupère tous les clients et déchiffre leurs données
func (db *Database) GetAllClients(cryptoSvc *crypto.CryptoService) ([]models.Client, error) {
	var listeClients []models.Client

	query := `SELECT * FROM clients ORDER BY nom, prenom`

	err := db.Select(&listeClients, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des clients: %w", err)
	}

	// Déchiffrement de chaque client
	for i := range listeClients {
		if err := decryptClient(&listeClients[i], cryptoSvc); err != nil {
			return nil, fmt.Errorf("erreur déchiffrement client %d: %w", listeClients[i].ID, err)
		}
	}

	return listeClients, nil
}

// GetClientByID récupère un client unique par son ID et déchiffre ses données
func (db *Database) GetClientByID(id int, cryptoSvc *crypto.CryptoService) (*models.Client, error) {
	var clientRecupere models.Client

	query := `SELECT * FROM clients WHERE id = ?`

	err := db.Get(&clientRecupere, query, id)
	if err != nil {
		return nil, fmt.Errorf("client avec l'ID %d non trouvé: %w", id, err)
	}

	// Déchiffrement
	if err := decryptClient(&clientRecupere, cryptoSvc); err != nil {
		return nil, fmt.Errorf("erreur déchiffrement client: %w", err)
	}

	return &clientRecupere, nil
}

// CreateClient crée un nouveau client avec chiffrement des données sensibles
func (db *Database) CreateClient(req models.CreateClientRequest, createdBy int, cryptoSvc *crypto.CryptoService) (int64, error) {
	// 🔢 Génération automatique du numéro Leopard UNIQUE
	leopardNumber, err := db.GenerateLeopardNumber(req.Nom, req.Prenom)
	if err != nil {
		return 0, fmt.Errorf("erreur génération numéro Leopard: %w", err)
	}
	req.NoDossierLeopard = &leopardNumber

	// Chiffrement des données sensibles avant insertion
	encryptedReq, err := encryptCreateClientRequest(req, cryptoSvc)
	if err != nil {
		return 0, fmt.Errorf("erreur chiffrement des données: %w", err)
	}

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
		"nom":                        encryptedReq.Nom,
		"prenom":                     encryptedReq.Prenom,
		"date_naissance":             encryptedReq.DateNaissance,
		"telephone":                  encryptedReq.Telephone,
		"cellulaire":                 encryptedReq.Cellulaire,
		"email":                      encryptedReq.Email,
		"adresse":                    encryptedReq.Adresse,
		"code_postal":                encryptedReq.CodePostal,
		"ville":                      encryptedReq.Ville,
		"pays":                       encryptedReq.Pays,
		"numero_assurance_maladie":   encryptedReq.NumeroAssuranceMaladie,
		"numero_securite_sociale":    encryptedReq.NumeroSecuriteSociale,
		"no_hcm":                     encryptedReq.NoHCM,
		"no_chaur":                   encryptedReq.NoCHAUR,
		"no_dossier_leopard":         encryptedReq.NoDossierLeopard,
		"medecin_famille_No_Licence": encryptedReq.MedecinFamilleNoLicence,
		"notaire_id":                 encryptedReq.NotaireID,
		"pivot_id":                   encryptedReq.PivotID,
		"rpa_id":                     encryptedReq.RPAID,
		"chsld_id":                   encryptedReq.CHSLDID,
		"ri_id":                      encryptedReq.RIID,
		"note_fixe":                  encryptedReq.NoteFixe,
		"Actif":                      encryptedReq.Actif,
		"dcd":                        encryptedReq.DCD,
		"created_by":                 createdBy,
	})

	if err != nil {
		return 0, fmt.Errorf("erreur création client: %w", err)
	}

	return resultat.LastInsertId()
}

// UpdateClient met à jour un client existant avec chiffrement
func (db *Database) UpdateClient(req models.UpdateClientRequest, cryptoSvc *crypto.CryptoService) error {
	// Chiffrement des données sensibles avant mise à jour
	encryptedReq, err := encryptUpdateClientRequest(req, cryptoSvc)
	if err != nil {
		return fmt.Errorf("erreur chiffrement des données: %w", err)
	}

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

	_, err = db.NamedExec(query, encryptedReq)
	if err != nil {
		return fmt.Errorf("erreur mise à jour client %d: %w", encryptedReq.ID, err)
	}

	return nil
}

// ArchiveClient désactive un client (suppression logique)
func (db *Database) ArchiveClient(id int) error {
	query := `UPDATE clients SET Actif = 0 WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur archivage client %d: %w", id, err)
	}
	return nil
}

// DeleteClient supprime définitivement un client
func (db *Database) DeleteClient(id int) error {
	query := `DELETE FROM clients WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur suppression physique client %d: %w", id, err)
	}
	return nil
}

// --- Fonctions utilitaires de chiffrement/déchiffrement ---

// decryptClient déchiffre tous les champs sensibles d'un client
func decryptClient(client *models.Client, cryptoSvc *crypto.CryptoService) error {
	var err error

	// Déchiffrement des champs string obligatoires
	client.Nom, err = cryptoSvc.Decrypt(client.Nom)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement nom: %w", err)
	}

	client.Prenom, err = cryptoSvc.Decrypt(client.Prenom)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement prénom: %w", err)
	}

	// Déchiffrement des champs optionnels (*string)
	client.DateNaissance, err = cryptoSvc.DecryptStringPtr(client.DateNaissance)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement date naissance: %w", err)
	}

	client.Telephone, err = cryptoSvc.DecryptStringPtr(client.Telephone)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement téléphone: %w", err)
	}

	client.Cellulaire, err = cryptoSvc.DecryptStringPtr(client.Cellulaire)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement cellulaire: %w", err)
	}

	client.Email, err = cryptoSvc.DecryptStringPtr(client.Email)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement email: %w", err)
	}

	client.Adresse, err = cryptoSvc.DecryptStringPtr(client.Adresse)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement adresse: %w", err)
	}

	client.CodePostal, err = cryptoSvc.DecryptStringPtr(client.CodePostal)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement code postal: %w", err)
	}

	client.NumeroAssuranceMaladie, err = cryptoSvc.DecryptStringPtr(client.NumeroAssuranceMaladie)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement NAM: %w", err)
	}

	client.NumeroSecuriteSociale, err = cryptoSvc.DecryptStringPtr(client.NumeroSecuriteSociale)
	if err != nil {
		return fmt.Errorf("erreur déchiffrement NAS: %w", err)
	}

	return nil
}

// encryptCreateClientRequest chiffre les données sensibles lors de la création
func encryptCreateClientRequest(req models.CreateClientRequest, cryptoSvc *crypto.CryptoService) (models.CreateClientRequest, error) {
	encrypted := req
	var err error

	// Chiffrement des champs obligatoires
	encrypted.Nom, err = cryptoSvc.Encrypt(req.Nom)
	if err != nil {
		return req, err
	}

	encrypted.Prenom, err = cryptoSvc.Encrypt(req.Prenom)
	if err != nil {
		return req, err
	}

	// Chiffrement des champs optionnels
	encrypted.DateNaissance, err = cryptoSvc.EncryptStringPtr(req.DateNaissance)
	if err != nil {
		return req, err
	}

	encrypted.Telephone, err = cryptoSvc.EncryptStringPtr(req.Telephone)
	if err != nil {
		return req, err
	}

	encrypted.Cellulaire, err = cryptoSvc.EncryptStringPtr(req.Cellulaire)
	if err != nil {
		return req, err
	}

	encrypted.Email, err = cryptoSvc.EncryptStringPtr(req.Email)
	if err != nil {
		return req, err
	}

	encrypted.Adresse, err = cryptoSvc.EncryptStringPtr(req.Adresse)
	if err != nil {
		return req, err
	}

	encrypted.CodePostal, err = cryptoSvc.EncryptStringPtr(req.CodePostal)
	if err != nil {
		return req, err
	}

	encrypted.NumeroAssuranceMaladie, err = cryptoSvc.EncryptStringPtr(req.NumeroAssuranceMaladie)
	if err != nil {
		return req, err
	}

	encrypted.NumeroSecuriteSociale, err = cryptoSvc.EncryptStringPtr(req.NumeroSecuriteSociale)
	if err != nil {
		return req, err
	}

	return encrypted, nil
}

// encryptUpdateClientRequest chiffre les données sensibles lors de la mise à jour
func encryptUpdateClientRequest(req models.UpdateClientRequest, cryptoSvc *crypto.CryptoService) (models.UpdateClientRequest, error) {
	encrypted := req
	var err error

	encrypted.Nom, err = cryptoSvc.Encrypt(req.Nom)
	if err != nil {
		return req, err
	}

	encrypted.Prenom, err = cryptoSvc.Encrypt(req.Prenom)
	if err != nil {
		return req, err
	}

	encrypted.DateNaissance, err = cryptoSvc.EncryptStringPtr(req.DateNaissance)
	if err != nil {
		return req, err
	}

	encrypted.Telephone, err = cryptoSvc.EncryptStringPtr(req.Telephone)
	if err != nil {
		return req, err
	}

	encrypted.Cellulaire, err = cryptoSvc.EncryptStringPtr(req.Cellulaire)
	if err != nil {
		return req, err
	}

	encrypted.Email, err = cryptoSvc.EncryptStringPtr(req.Email)
	if err != nil {
		return req, err
	}

	encrypted.Adresse, err = cryptoSvc.EncryptStringPtr(req.Adresse)
	if err != nil {
		return req, err
	}

	encrypted.CodePostal, err = cryptoSvc.EncryptStringPtr(req.CodePostal)
	if err != nil {
		return req, err
	}

	encrypted.NumeroAssuranceMaladie, err = cryptoSvc.EncryptStringPtr(req.NumeroAssuranceMaladie)
	if err != nil {
		return req, err
	}

	encrypted.NumeroSecuriteSociale, err = cryptoSvc.EncryptStringPtr(req.NumeroSecuriteSociale)
	if err != nil {
		return req, err
	}

	return encrypted, nil
}
