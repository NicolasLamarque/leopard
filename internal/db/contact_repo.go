package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllContactsByClientID récupère tous les contacts d'un client spécifique et déchiffre leurs données
func (db *Database) GetAllContactsByClientID(clientID int, cryptoSvc *crypto.CryptoService) ([]models.Contact, error) {
	var listeContacts []models.Contact

	query := `SELECT * FROM contacts WHERE client_id = ? ORDER BY Nom, Prenom`

	err := db.Select(&listeContacts, query, clientID)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des contacts du client %d: %w", clientID, err)
	}

	// Déchiffrement de chaque contact
	for i := range listeContacts {
		if err := decryptContact(&listeContacts[i], cryptoSvc); err != nil {
			return nil, fmt.Errorf("erreur déchiffrement contact %d: %w", listeContacts[i].ID, err)
		}
	}

	return listeContacts, nil
}

// GetAllContacts récupère tous les contacts (tous clients confondus)
func (db *Database) GetAllContacts(cryptoSvc *crypto.CryptoService) ([]models.Contact, error) {
	var listeContacts []models.Contact

	query := `SELECT * FROM contacts ORDER BY client_id, Nom, Prenom`

	err := db.Select(&listeContacts, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des contacts: %w", err)
	}

	// Déchiffrement de chaque contact
	for i := range listeContacts {
		if err := decryptContact(&listeContacts[i], cryptoSvc); err != nil {
			return nil, fmt.Errorf("erreur déchiffrement contact %d: %w", listeContacts[i].ID, err)
		}
	}

	return listeContacts, nil
}

// GetContactByID récupère un contact unique par son ID et déchiffre ses données
func (db *Database) GetContactByID(id int, cryptoSvc *crypto.CryptoService) (*models.Contact, error) {
	var contactRecupere models.Contact

	query := `SELECT * FROM contacts WHERE id = ?`

	err := db.Get(&contactRecupere, query, id)
	if err != nil {
		return nil, fmt.Errorf("contact avec l'ID %d non trouvé: %w", id, err)
	}

	// Déchiffrement
	if err := decryptContact(&contactRecupere, cryptoSvc); err != nil {
		return nil, fmt.Errorf("erreur déchiffrement contact: %w", err)
	}

	return &contactRecupere, nil
}

// CreateContact crée un nouveau contact avec chiffrement des données sensibles
func (db *Database) CreateContact(req models.CreateContactRequest, cryptoSvc *crypto.CryptoService) (int64, error) {
	// Chiffrement des données sensibles avant insertion
	encryptedReq, err := encryptCreateContactRequest(req, cryptoSvc)
	if err != nil {
		return 0, fmt.Errorf("erreur chiffrement des données: %w", err)
	}

	query := `
        INSERT INTO contacts (
            Nom, Prenom, Telephone, Cellulaire, Adresse, Code_Postal, Ville, Pays,
            Email, employeur, Profession, Relation_Client, lien_familial,
            force_lien, contact_urgence, aidant_naturel, type_de_contact,
            procuration_bancaire, procuration_notariee, note_fixe,
            relation_type, client_id
        ) VALUES (
            :Nom, :Prenom, :Telephone, :Cellulaire, :Adresse, :Code_Postal, :Ville, :Pays,
            :Email, :employeur, :Profession, :Relation_Client, :lien_familial,
            :force_lien, :contact_urgence, :aidant_naturel, :type_de_contact,
            :procuration_bancaire, :procuration_notariee, :note_fixe,
            :relation_type, :client_id
        )`

	resultat, err := db.NamedExec(query, map[string]interface{}{
		"Nom":                  encryptedReq.Nom,
		"Prenom":               encryptedReq.Prenom,
		"Telephone":            encryptedReq.Telephone,
		"Cellulaire":           encryptedReq.Cellulaire,
		"Adresse":              encryptedReq.Adresse,
		"Code_Postal":          encryptedReq.CodePostal,
		"Ville":                encryptedReq.Ville,
		"Pays":                 encryptedReq.Pays,
		"Email":                encryptedReq.Email,
		"employeur":            encryptedReq.Employeur,
		"Profession":           encryptedReq.Profession,
		"Relation_Client":      encryptedReq.RelationClient,
		"lien_familial":        encryptedReq.LienFamilial,
		"force_lien":           encryptedReq.ForceLien,
		"contact_urgence":      encryptedReq.ContactUrgence,
		"aidant_naturel":       encryptedReq.AidantNaturel,
		"type_de_contact":      encryptedReq.TypeDeContact,
		"procuration_bancaire": encryptedReq.ProcurationBancaire,
		"procuration_notariee": encryptedReq.ProcurationNotariee,
		"note_fixe":            encryptedReq.NoteFixe,
		"relation_type":        encryptedReq.RelationType,
		"client_id":            encryptedReq.ClientID,
	})

	if err != nil {
		return 0, fmt.Errorf("erreur création contact: %w", err)
	}

	return resultat.LastInsertId()
}

// UpdateContact met à jour un contact existant avec chiffrement
func (db *Database) UpdateContact(req models.UpdateContactRequest, cryptoSvc *crypto.CryptoService) error {
	// Chiffrement des données sensibles avant mise à jour
	encryptedReq, err := encryptUpdateContactRequest(req, cryptoSvc)
	if err != nil {
		return fmt.Errorf("erreur chiffrement des données: %w", err)
	}

	query := `
		UPDATE contacts SET
			Nom = :Nom,
			Prenom = :Prenom,
			Telephone = :Telephone,
			Cellulaire = :Cellulaire,
			Adresse = :Adresse,
			Code_Postal = :Code_Postal,
			Ville = :Ville,
			Pays = :Pays,
			Email = :Email,
			employeur = :employeur,
			Profession = :Profession,
			Relation_Client = :Relation_Client,
			lien_familial = :lien_familial,
			force_lien = :force_lien,
			contact_urgence = :contact_urgence,
			aidant_naturel = :aidant_naturel,
			type_de_contact = :type_de_contact,
			procuration_bancaire = :procuration_bancaire,
			procuration_notariee = :procuration_notariee,
			note_fixe = :note_fixe,
			relation_type = :relation_type,
			client_id = :client_id
		WHERE id = :id`

	_, err = db.NamedExec(query, map[string]interface{}{
		"id":                   encryptedReq.ID,
		"Nom":                  encryptedReq.Nom,
		"Prenom":               encryptedReq.Prenom,
		"Telephone":            encryptedReq.Telephone,
		"Cellulaire":           encryptedReq.Cellulaire,
		"Adresse":              encryptedReq.Adresse,
		"Code_Postal":          encryptedReq.CodePostal,
		"Ville":                encryptedReq.Ville,
		"Pays":                 encryptedReq.Pays,
		"Email":                encryptedReq.Email,
		"employeur":            encryptedReq.Employeur,
		"Profession":           encryptedReq.Profession,
		"Relation_Client":      encryptedReq.RelationClient,
		"lien_familial":        encryptedReq.LienFamilial,
		"force_lien":           encryptedReq.ForceLien,
		"contact_urgence":      encryptedReq.ContactUrgence,
		"aidant_naturel":       encryptedReq.AidantNaturel,
		"type_de_contact":      encryptedReq.TypeDeContact,
		"procuration_bancaire": encryptedReq.ProcurationBancaire,
		"procuration_notariee": encryptedReq.ProcurationNotariee,
		"note_fixe":            encryptedReq.NoteFixe,
		"relation_type":        encryptedReq.RelationType,
		"client_id":            encryptedReq.ClientID,
	})

	if err != nil {
		return fmt.Errorf("erreur mise à jour contact %d: %w", encryptedReq.ID, err)
	}

	return nil
}

// DeleteContact supprime définitivement un contact
func (db *Database) DeleteContact(id int) error {
	query := `DELETE FROM contacts WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur suppression contact %d: %w", id, err)
	}
	return nil
}

func decryptContact(contact *models.Contact, cryptoSvc *crypto.CryptoService) error {
	// On ne crée pas de variable 'err' globale qui pourrait contaminer la suite

	// 1. CHAMPS OBLIGATOIRES : Si ceux-là pètent, on veut le savoir
	var err error
	contact.Nom, err = cryptoSvc.Decrypt(contact.Nom)
	if err != nil {
		return fmt.Errorf("erreur nom: %w", err)
	}

	contact.Prenom, err = cryptoSvc.Decrypt(contact.Prenom)
	if err != nil {
		return fmt.Errorf("erreur prénom: %w", err)
	}

	// 2. CHAMPS OPTIONNELS : On utilise '_' pour ignorer l'erreur.
	// Si le déchiffrement échoue (donnée vide ou corrompue), on continue quand même.
	contact.Telephone, _ = cryptoSvc.DecryptStringPtr(contact.Telephone)
	contact.Cellulaire, _ = cryptoSvc.DecryptStringPtr(contact.Cellulaire)
	contact.Email, _ = cryptoSvc.DecryptStringPtr(contact.Email)
	contact.Adresse, _ = cryptoSvc.DecryptStringPtr(contact.Adresse)
	contact.CodePostal, _ = cryptoSvc.DecryptStringPtr(contact.CodePostal)
	contact.Ville, _ = cryptoSvc.DecryptStringPtr(contact.Ville)
	contact.Pays, _ = cryptoSvc.DecryptStringPtr(contact.Pays)
	contact.Employeur, _ = cryptoSvc.DecryptStringPtr(contact.Employeur)
	contact.Profession, _ = cryptoSvc.DecryptStringPtr(contact.Profession)
	contact.RelationClient, _ = cryptoSvc.DecryptStringPtr(contact.RelationClient)
	contact.LienFamilial, _ = cryptoSvc.DecryptStringPtr(contact.LienFamilial)
	contact.TypeDeContact, _ = cryptoSvc.DecryptStringPtr(contact.TypeDeContact)
	contact.NoteFixe, _ = cryptoSvc.DecryptStringPtr(contact.NoteFixe)

	// 3. CHAMPS NUMÉRIQUES (IMPORTANT)
	// On n'appelle PAS cryptoSvc sur les INTEGER (force_lien, procurations, etc.)
	// Ils sont déjà mappés correctement par db.Select()

	return nil
}

// encryptCreateContactRequest chiffre les données sensibles lors de la création
func encryptCreateContactRequest(req models.CreateContactRequest, cryptoSvc *crypto.CryptoService) (models.CreateContactRequest, error) {
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

	return encrypted, nil
}

// encryptUpdateContactRequest chiffre les données sensibles lors de la mise à jour
func encryptUpdateContactRequest(req models.UpdateContactRequest, cryptoSvc *crypto.CryptoService) (models.UpdateContactRequest, error) {
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

	return encrypted, nil
}
