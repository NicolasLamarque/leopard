package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllClients récupère tous les clients et déchiffre leurs données
func (db *Database) GetAllClients(cryptoSvc *crypto.CryptoService) ([]models.Client, error) {
	var listeClients []models.Client

	query := `SELECT * FROM clients WHERE Actif = 1 ORDER BY nom, prenom`

	err := db.Select(&listeClients, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des clients: %w", err)
	}

	for i := range listeClients {
		if err := decryptClient(&listeClients[i], cryptoSvc); err != nil {
			return nil, fmt.Errorf("erreur déchiffrement client %d: %w", listeClients[i].ID, err)
		}
	}

	return listeClients, nil
}

// GetClientByID récupère un client unique par son ID
func (db *Database) GetClientByID(id int, cryptoSvc *crypto.CryptoService) (*models.Client, error) {
	var clientRecupere models.Client
	query := `SELECT * FROM clients WHERE id = ?`

	err := db.Get(&clientRecupere, query, id)
	if err != nil {
		return nil, fmt.Errorf("client avec l'ID %d non trouvé: %w", id, err)
	}

	if err := decryptClient(&clientRecupere, cryptoSvc); err != nil {
		return nil, fmt.Errorf("erreur déchiffrement client: %w", err)
	}

	return &clientRecupere, nil
}

// CreateClient crée un nouveau client avec tous les nouveaux champs
func (db *Database) CreateClient(req models.CreateClientRequest, createdBy int, cryptoSvc *crypto.CryptoService) (int64, error) {
	// Génération du numéro Leopard (en clair)
	leopardNumber, err := db.GenerateLeopardNumber(req.Nom, req.Prenom)
	if err != nil {
		return 0, fmt.Errorf("erreur génération numéro Leopard: %w", err)
	}
	req.NoDossierLeopard = &leopardNumber

	// Chiffrement des données sensibles (sauf Nom, Prenom, NoDossierLeopard)
	encryptedReq, err := encryptCreateClientRequest(req, cryptoSvc)
	if err != nil {
		return 0, fmt.Errorf("erreur chiffrement des données: %w", err)
	}

	query := `
		INSERT INTO clients (
			nom, prenom, date_naissance, sexe, lieu_naissance, statut_marital,
			occupation, employeur, profession, niveau_scolaire, langue_preferee,
			origine_ethnique, premiere_nation, identite_genre, orientation_sexuelle,
			religion, premiere_langue, telephone, cellulaire, email, adresse,
			appartement, code_postal, ville, pays, province,
			numero_assurance_maladie, numero_securite_sociale, no_hcm, no_chaur, 
			no_dossier_leopard, medecin_famille_No_Licence,
			notaire_id, pivot_id, rpa_id, chsld_id, ri_id, 
			procuration_bancaire, procuration_notariee,
			note_fixe, Actif, dcd, date_deces, created_by
		) VALUES (
			:nom, :prenom, :date_naissance, :sexe, :lieu_naissance, :statut_marital,
			:occupation, :employeur, :profession, :niveau_scolaire, :langue_preferee,
			:origine_ethnique, :premiere_nation, :identite_genre, :orientation_sexuelle,
			:religion, :premiere_langue, :telephone, :cellulaire, :email, :adresse,
			:appartement, :code_postal, :ville, :pays, :province,
			:numero_assurance_maladie, :numero_securite_sociale, :no_hcm, :no_chaur,
			:no_dossier_leopard, :medecin_famille_No_Licence,
			:notaire_id, :pivot_id, :rpa_id, :chsld_id, :ri_id,
			:procuration_bancaire, :procuration_notariee,
			:note_fixe, :Actif, :dcd, :date_deces, :created_by
		)`

	data := map[string]interface{}{
		"nom":                        encryptedReq.Nom,
		"prenom":                     encryptedReq.Prenom,
		"date_naissance":             encryptedReq.DateNaissance,
		"sexe":                       encryptedReq.Sexe,
		"lieu_naissance":             encryptedReq.LieuNaissance,
		"statut_marital":             encryptedReq.StatutMarital,
		"occupation":                 encryptedReq.Occupation,
		"employeur":                  encryptedReq.Employeur,
		"profession":                 encryptedReq.Profession,
		"niveau_scolaire":            encryptedReq.NiveauScolaire,
		"langue_preferee":            encryptedReq.LanguePreferee,
		"origine_ethnique":           encryptedReq.OrigineEthnique,
		"premiere_nation":            encryptedReq.PremiereNation,
		"identite_genre":             encryptedReq.IdentiteGenre,
		"orientation_sexuelle":       encryptedReq.OrientationSexuelle,
		"religion":                   encryptedReq.Religion,
		"premiere_langue":            encryptedReq.PremiereLangue,
		"telephone":                  encryptedReq.Telephone,
		"cellulaire":                 encryptedReq.Cellulaire,
		"email":                      encryptedReq.Email,
		"adresse":                    encryptedReq.Adresse,
		"appartement":                encryptedReq.Appartement,
		"code_postal":                encryptedReq.CodePostal,
		"ville":                      encryptedReq.Ville,
		"pays":                       encryptedReq.Pays,
		"province":                   encryptedReq.Province,
		"numero_assurance_maladie":   encryptedReq.NumeroAssuranceMaladie,
		"numero_securite_sociale":    encryptedReq.NumeroSecuriteSociale,
		"no_hcm":                     encryptedReq.NoHCM,
		"no_chaur":                   encryptedReq.NoCHAUR,
		"no_dossier_leopard":         leopardNumber, // <-- ON UTILISE LA VARIABLE NON CHIFFRÉE
		"medecin_famille_No_Licence": encryptedReq.MedecinFamilleNoLicence,
		"notaire_id":                 encryptedReq.NotaireID,
		"pivot_id":                   encryptedReq.PivotID,
		"rpa_id":                     encryptedReq.RPAID,
		"chsld_id":                   encryptedReq.CHSLDID,
		"ri_id":                      encryptedReq.RIID,
		"procuration_bancaire":       encryptedReq.ProcurationBancaire,
		"procuration_notariee":       encryptedReq.ProcurationNotariee,
		"note_fixe":                  encryptedReq.NoteFixe,
		"Actif":                      encryptedReq.Actif,
		"dcd":                        encryptedReq.DCD,
		"date_deces":                 encryptedReq.DateDeces,
		"created_by":                 createdBy,
	}

	resultat, err := db.NamedExec(query, data)
	if err != nil {
		return 0, fmt.Errorf("erreur création client: %w", err)
	}

	return resultat.LastInsertId()
}

// UpdateClient met à jour un client existant
func (db *Database) UpdateClient(req models.UpdateClientRequest, cryptoSvc *crypto.CryptoService) error {
	encryptedReq, err := encryptUpdateClientRequest(req, cryptoSvc)
	if err != nil {
		return fmt.Errorf("erreur chiffrement des données: %w", err)
	}

	query := `
		UPDATE clients SET
			nom = :nom, prenom = :prenom, date_naissance = :date_naissance, sexe = :sexe,
			lieu_naissance = :lieu_naissance, statut_marital = :statut_marital,
			occupation = :occupation, employeur = :employeur, profession = :profession,
			niveau_scolaire = :niveau_scolaire, langue_preferee = :langue_preferee,
			origine_ethnique = :origine_ethnique, premiere_nation = :premiere_nation,
			identite_genre = :identite_genre, orientation_sexuelle = :orientation_sexuelle,
			religion = :religion, premiere_langue = :premiere_langue, telephone = :telephone,
			cellulaire = :cellulaire, email = :email, adresse = :adresse,
			appartement = :appartement, code_postal = :code_postal, ville = :ville, 
			pays = :pays, province = :province,
			numero_assurance_maladie = :numero_assurance_maladie,
			numero_securite_sociale = :numero_securite_sociale, no_hcm = :no_hcm,
			no_chaur = :no_chaur, no_dossier_leopard = :no_dossier_leopard,
			medecin_famille_No_Licence = :medecin_famille_No_Licence,
			notaire_id = :notaire_id, pivot_id = :pivot_id, rpa_id = :rpa_id,
			chsld_id = :chsld_id, ri_id = :ri_id, 
			procuration_bancaire = :procuration_bancaire, procuration_notariee = :procuration_notariee, 
			note_fixe = :note_fixe,
			Actif = :Actif, dcd = :dcd, date_deces = :date_deces
		WHERE id = :id`

	_, err = db.NamedExec(query, encryptedReq)
	if err != nil {
		return fmt.Errorf("erreur mise à jour client %d: %w", encryptedReq.ID, err)
	}
	return nil
}

// ArchiveClient désactive un client
func (db *Database) ArchiveClient(id int) error {
	query := `UPDATE clients SET Actif = 0 WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// DeleteClient supprime définitivement un client
func (db *Database) DeleteClient(id int) error {
	query := `DELETE FROM clients WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// --- Fonctions utilitaires de chiffrement ---

func decryptClient(c *models.Client, cryptoSvc *crypto.CryptoService) error {
	// Nom, Prenom, NoDossierLeopard : RESTENT EN CLAIR
	var err error
	c.DateNaissance, err = cryptoSvc.DecryptStringPtr(c.DateNaissance)
	if err != nil {
		return err
	}
	c.Sexe, _ = cryptoSvc.DecryptStringPtr(c.Sexe)
	c.LieuNaissance, _ = cryptoSvc.DecryptStringPtr(c.LieuNaissance)
	c.StatutMarital, _ = cryptoSvc.DecryptStringPtr(c.StatutMarital)
	c.Occupation, _ = cryptoSvc.DecryptStringPtr(c.Occupation)
	c.Employeur, _ = cryptoSvc.DecryptStringPtr(c.Employeur)
	c.Profession, _ = cryptoSvc.DecryptStringPtr(c.Profession)
	c.NiveauScolaire, _ = cryptoSvc.DecryptStringPtr(c.NiveauScolaire)
	c.LanguePreferee, _ = cryptoSvc.DecryptStringPtr(c.LanguePreferee)
	c.OrigineEthnique, _ = cryptoSvc.DecryptStringPtr(c.OrigineEthnique)
	c.IdentiteGenre, _ = cryptoSvc.DecryptStringPtr(c.IdentiteGenre)
	c.OrientationSexuelle, _ = cryptoSvc.DecryptStringPtr(c.OrientationSexuelle)
	c.Religion, _ = cryptoSvc.DecryptStringPtr(c.Religion)
	c.PremiereLangue, _ = cryptoSvc.DecryptStringPtr(c.PremiereLangue)
	c.Telephone, _ = cryptoSvc.DecryptStringPtr(c.Telephone)
	c.Cellulaire, _ = cryptoSvc.DecryptStringPtr(c.Cellulaire)
	c.Email, _ = cryptoSvc.DecryptStringPtr(c.Email)
	c.Adresse, _ = cryptoSvc.DecryptStringPtr(c.Adresse)
	c.Appartement, _ = cryptoSvc.DecryptStringPtr(c.Appartement)
	c.CodePostal, _ = cryptoSvc.DecryptStringPtr(c.CodePostal)
	c.Ville, _ = cryptoSvc.DecryptStringPtr(c.Ville)
	c.Pays, _ = cryptoSvc.DecryptStringPtr(c.Pays)
	c.Province, _ = cryptoSvc.DecryptStringPtr(c.Province)
	c.NumeroAssuranceMaladie, _ = cryptoSvc.DecryptStringPtr(c.NumeroAssuranceMaladie)
	c.NumeroSecuriteSociale, _ = cryptoSvc.DecryptStringPtr(c.NumeroSecuriteSociale)
	c.NoHCM, _ = cryptoSvc.DecryptStringPtr(c.NoHCM)
	c.NoCHAUR, _ = cryptoSvc.DecryptStringPtr(c.NoCHAUR)
	c.ProcurationBancaire, _ = cryptoSvc.DecryptStringPtr(c.ProcurationBancaire)
	c.ProcurationNotariee, _ = cryptoSvc.DecryptStringPtr(c.ProcurationNotariee)
	c.MedecinFamilleNoLicence, _ = cryptoSvc.DecryptStringPtr(c.MedecinFamilleNoLicence)
	c.NoteFixe, _ = cryptoSvc.DecryptStringPtr(c.NoteFixe)
	c.DateDeces, _ = cryptoSvc.DecryptStringPtr(c.DateDeces)
	return nil
}

func encryptCreateClientRequest(req models.CreateClientRequest, cryptoSvc *crypto.CryptoService) (models.CreateClientRequest, error) {
	enc := req
	// Nom, Prenom, NoDossierLeopard : RESTENT EN CLAIR
	enc.DateNaissance, _ = cryptoSvc.EncryptStringPtr(req.DateNaissance)
	enc.Sexe, _ = cryptoSvc.EncryptStringPtr(req.Sexe)
	enc.LieuNaissance, _ = cryptoSvc.EncryptStringPtr(req.LieuNaissance)
	enc.StatutMarital, _ = cryptoSvc.EncryptStringPtr(req.StatutMarital)
	enc.Occupation, _ = cryptoSvc.EncryptStringPtr(req.Occupation)
	enc.Employeur, _ = cryptoSvc.EncryptStringPtr(req.Employeur)
	enc.Profession, _ = cryptoSvc.EncryptStringPtr(req.Profession)
	enc.NiveauScolaire, _ = cryptoSvc.EncryptStringPtr(req.NiveauScolaire)
	enc.LanguePreferee, _ = cryptoSvc.EncryptStringPtr(req.LanguePreferee)
	enc.OrigineEthnique, _ = cryptoSvc.EncryptStringPtr(req.OrigineEthnique)
	enc.IdentiteGenre, _ = cryptoSvc.EncryptStringPtr(req.IdentiteGenre)
	enc.OrientationSexuelle, _ = cryptoSvc.EncryptStringPtr(req.OrientationSexuelle)
	enc.Religion, _ = cryptoSvc.EncryptStringPtr(req.Religion)
	enc.PremiereLangue, _ = cryptoSvc.EncryptStringPtr(req.PremiereLangue)
	enc.Telephone, _ = cryptoSvc.EncryptStringPtr(req.Telephone)
	enc.Cellulaire, _ = cryptoSvc.EncryptStringPtr(req.Cellulaire)
	enc.Email, _ = cryptoSvc.EncryptStringPtr(req.Email)
	enc.Adresse, _ = cryptoSvc.EncryptStringPtr(req.Adresse)
	enc.Appartement, _ = cryptoSvc.EncryptStringPtr(req.Appartement)
	enc.CodePostal, _ = cryptoSvc.EncryptStringPtr(req.CodePostal)
	enc.Ville, _ = cryptoSvc.EncryptStringPtr(req.Ville)
	enc.Pays, _ = cryptoSvc.EncryptStringPtr(req.Pays)
	enc.Province, _ = cryptoSvc.EncryptStringPtr(req.Province)
	enc.NumeroAssuranceMaladie, _ = cryptoSvc.EncryptStringPtr(req.NumeroAssuranceMaladie)
	enc.NumeroSecuriteSociale, _ = cryptoSvc.EncryptStringPtr(req.NumeroSecuriteSociale)
	enc.NoHCM, _ = cryptoSvc.EncryptStringPtr(req.NoHCM)
	enc.NoCHAUR, _ = cryptoSvc.EncryptStringPtr(req.NoCHAUR)
	enc.MedecinFamilleNoLicence, _ = cryptoSvc.EncryptStringPtr(req.MedecinFamilleNoLicence)
	enc.NoteFixe, _ = cryptoSvc.EncryptStringPtr(req.NoteFixe)
	enc.DateDeces, _ = cryptoSvc.EncryptStringPtr(req.DateDeces)
	return enc, nil
}

func encryptUpdateClientRequest(req models.UpdateClientRequest, cryptoSvc *crypto.CryptoService) (models.UpdateClientRequest, error) {
	enc := req
	// Nom, Prenom, NoDossierLeopard : RESTENT EN CLAIR
	enc.DateNaissance, _ = cryptoSvc.EncryptStringPtr(req.DateNaissance)
	enc.Sexe, _ = cryptoSvc.EncryptStringPtr(req.Sexe)
	enc.LieuNaissance, _ = cryptoSvc.EncryptStringPtr(req.LieuNaissance)
	enc.StatutMarital, _ = cryptoSvc.EncryptStringPtr(req.StatutMarital)
	enc.Occupation, _ = cryptoSvc.EncryptStringPtr(req.Occupation)
	enc.Employeur, _ = cryptoSvc.EncryptStringPtr(req.Employeur)
	enc.Profession, _ = cryptoSvc.EncryptStringPtr(req.Profession)
	enc.NiveauScolaire, _ = cryptoSvc.EncryptStringPtr(req.NiveauScolaire)
	enc.LanguePreferee, _ = cryptoSvc.EncryptStringPtr(req.LanguePreferee)
	enc.OrigineEthnique, _ = cryptoSvc.EncryptStringPtr(req.OrigineEthnique)
	enc.IdentiteGenre, _ = cryptoSvc.EncryptStringPtr(req.IdentiteGenre)
	enc.OrientationSexuelle, _ = cryptoSvc.EncryptStringPtr(req.OrientationSexuelle)
	enc.Religion, _ = cryptoSvc.EncryptStringPtr(req.Religion)
	enc.PremiereLangue, _ = cryptoSvc.EncryptStringPtr(req.PremiereLangue)
	enc.Telephone, _ = cryptoSvc.EncryptStringPtr(req.Telephone)
	enc.Cellulaire, _ = cryptoSvc.EncryptStringPtr(req.Cellulaire)
	enc.Email, _ = cryptoSvc.EncryptStringPtr(req.Email)
	enc.Adresse, _ = cryptoSvc.EncryptStringPtr(req.Adresse)
	enc.Appartement, _ = cryptoSvc.EncryptStringPtr(req.Appartement)
	enc.CodePostal, _ = cryptoSvc.EncryptStringPtr(req.CodePostal)
	enc.Ville, _ = cryptoSvc.EncryptStringPtr(req.Ville)
	enc.Pays, _ = cryptoSvc.EncryptStringPtr(req.Pays)
	enc.Province, _ = cryptoSvc.EncryptStringPtr(req.Province)
	enc.NumeroAssuranceMaladie, _ = cryptoSvc.EncryptStringPtr(req.NumeroAssuranceMaladie)
	enc.NumeroSecuriteSociale, _ = cryptoSvc.EncryptStringPtr(req.NumeroSecuriteSociale)
	enc.NoHCM, _ = cryptoSvc.EncryptStringPtr(req.NoHCM)
	enc.NoCHAUR, _ = cryptoSvc.EncryptStringPtr(req.NoCHAUR)
	enc.MedecinFamilleNoLicence, _ = cryptoSvc.EncryptStringPtr(req.MedecinFamilleNoLicence)
	enc.NoteFixe, _ = cryptoSvc.EncryptStringPtr(req.NoteFixe)
	enc.DateDeces, _ = cryptoSvc.EncryptStringPtr(req.DateDeces)
	return enc, nil
}
