package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

func (db *Database) GetAllClients(cryptoSvc *crypto.CryptoService) ([]models.Client, error) {
	var listeClients []models.Client
	err := db.Select(&listeClients, `SELECT * FROM clients WHERE Actif = 1 ORDER BY nom, prenom`)
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

func (db *Database) GetClientByID(id int, cryptoSvc *crypto.CryptoService) (*models.Client, error) {
	var c models.Client
	err := db.Get(&c, `SELECT * FROM clients WHERE id = ?`, id)
	if err != nil {
		return nil, fmt.Errorf("client avec l'ID %d non trouvé: %w", id, err)
	}
	if err := decryptClient(&c, cryptoSvc); err != nil {
		return nil, fmt.Errorf("erreur déchiffrement client: %w", err)
	}
	return &c, nil
}

func (db *Database) CreateClient(req models.CreateClientRequest, createdBy int, cryptoSvc *crypto.CryptoService) (int64, error) {
	leopardNumber, err := db.GenerateLeopardNumber(req.Nom, req.Prenom)
	if err != nil {
		return 0, fmt.Errorf("erreur génération numéro Leopard: %w", err)
	}
	req.NoDossierLeopard = &leopardNumber

	enc, err := encryptCreateClientRequest(req, cryptoSvc)
	if err != nil {
		return 0, fmt.Errorf("erreur chiffrement des données: %w", err)
	}

	query := `
		INSERT INTO clients (
			nom, prenom, date_naissance, sexe, lieu_naissance, statut_marital,
			occupation, employeur, profession, niveau_scolaire, langue_preferee,
			origine_ethnique, premiere_nation, identite_genre, orientation_sexuelle,
			religion, premiere_langue, telephone, cellulaire, email,
			adresse, appartement, code_postal, ville, mcode, arrcod, pays, province,
			numero_assurance_maladie, numero_securite_sociale, no_hcm, no_chaur,
			no_dossier_leopard, medecin_famille_No_Licence,
			notaire_id, pharmacie_id, pivot_id, rpa_id, chsld_id, ri_id,
			procuration_bancaire, procuration_notariee,
			tutelle_active, tutelle_type, tutelle_bien, tutelle_personne,
			tutelle_date_debut, tutelle_date_renouvellement,
			tuteur_nom, tuteur_prenom, tuteur_telephone, tuteur_cellulaire,
			tuteur_email, tuteur_adresse, tuteur_code_postal, tuteur_ville,
			note_fixe, Actif, dcd, date_deces, created_by
		) VALUES (
			:nom, :prenom, :date_naissance, :sexe, :lieu_naissance, :statut_marital,
			:occupation, :employeur, :profession, :niveau_scolaire, :langue_preferee,
			:origine_ethnique, :premiere_nation, :identite_genre, :orientation_sexuelle,
			:religion, :premiere_langue, :telephone, :cellulaire, :email,
			:adresse, :appartement, :code_postal, :ville, :mcode, :arrcod, :pays, :province,
			:numero_assurance_maladie, :numero_securite_sociale, :no_hcm, :no_chaur,
			:no_dossier_leopard, :medecin_famille_No_Licence,
			:notaire_id, :pharmacie_id, :pivot_id, :rpa_id, :chsld_id, :ri_id,
			:procuration_bancaire, :procuration_notariee,
			:tutelle_active, :tutelle_type, :tutelle_bien, :tutelle_personne,
			:tutelle_date_debut, :tutelle_date_renouvellement,
			:tuteur_nom, :tuteur_prenom, :tuteur_telephone, :tuteur_cellulaire,
			:tuteur_email, :tuteur_adresse, :tuteur_code_postal, :tuteur_ville,
			:note_fixe, :Actif, :dcd, :date_deces, :created_by
		)`

	data := map[string]interface{}{
		"nom":                         enc.Nom,
		"prenom":                      enc.Prenom,
		"date_naissance":              enc.DateNaissance,
		"sexe":                        enc.Sexe,
		"lieu_naissance":              enc.LieuNaissance,
		"statut_marital":              enc.StatutMarital,
		"occupation":                  enc.Occupation,
		"employeur":                   enc.Employeur,
		"profession":                  enc.Profession,
		"niveau_scolaire":             enc.NiveauScolaire,
		"langue_preferee":             enc.LanguePreferee,
		"origine_ethnique":            enc.OrigineEthnique,
		"premiere_nation":             enc.PremiereNation,
		"identite_genre":              enc.IdentiteGenre,
		"orientation_sexuelle":        enc.OrientationSexuelle,
		"religion":                    enc.Religion,
		"premiere_langue":             enc.PremiereLangue,
		"telephone":                   enc.Telephone,
		"cellulaire":                  enc.Cellulaire,
		"email":                       enc.Email,
		"adresse":                     enc.Adresse,
		"appartement":                 enc.Appartement,
		"code_postal":                 enc.CodePostal,
		"ville":                       enc.Ville,
		"mcode":                       enc.Mcode,  // non chiffré
		"arrcod":                      enc.Arrcod, // non chiffré
		"pays":                        enc.Pays,
		"province":                    enc.Province,
		"numero_assurance_maladie":    enc.NumeroAssuranceMaladie,
		"numero_securite_sociale":     enc.NumeroSecuriteSociale,
		"no_hcm":                      enc.NoHCM,
		"no_chaur":                    enc.NoCHAUR,
		"no_dossier_leopard":          leopardNumber, // non chiffré
		"medecin_famille_No_Licence":  enc.MedecinFamilleNoLicence,
		"notaire_id":                  enc.NotaireID,
		"pharmacie_id":                enc.PharmacieID,
		"pivot_id":                    enc.PivotID,
		"rpa_id":                      enc.RPAID,
		"chsld_id":                    enc.CHSLDID,
		"ri_id":                       enc.RIID,
		"procuration_bancaire":        enc.ProcurationBancaire,
		"procuration_notariee":        enc.ProcurationNotariee,
		"tutelle_active":              enc.TutelleActive,   // non chiffré
		"tutelle_type":                enc.TutelleType,     // non chiffré
		"tutelle_bien":                enc.TutelleBien,     // non chiffré
		"tutelle_personne":            enc.TutellePersonne, // non chiffré
		"tutelle_date_debut":          enc.TutelleDateDebut,
		"tutelle_date_renouvellement": enc.TutelleDateRenouvellement,
		"tuteur_nom":                  enc.TuteurNom,
		"tuteur_prenom":               enc.TuteurPrenom,
		"tuteur_telephone":            enc.TuteurTelephone,
		"tuteur_cellulaire":           enc.TuteurCellulaire,
		"tuteur_email":                enc.TuteurEmail,
		"tuteur_adresse":              enc.TuteurAdresse,
		"tuteur_code_postal":          enc.TuteurCodePostal,
		"tuteur_ville":                enc.TuteurVille,
		"note_fixe":                   enc.NoteFixe,
		"Actif":                       enc.Actif,
		"dcd":                         enc.DCD,
		"date_deces":                  enc.DateDeces,
		"created_by":                  createdBy,
	}

	resultat, err := db.NamedExec(query, data)
	if err != nil {
		return 0, fmt.Errorf("erreur création client: %w", err)
	}
	return resultat.LastInsertId()
}

func (db *Database) UpdateClient(req models.UpdateClientRequest, cryptoSvc *crypto.CryptoService) error {
	enc, err := encryptUpdateClientRequest(req, cryptoSvc)
	if err != nil {
		return fmt.Errorf("erreur chiffrement des données: %w", err)
	}

	query := `
		UPDATE clients SET
			nom = :nom, prenom = :prenom,
			date_naissance = :date_naissance, sexe = :sexe,
			lieu_naissance = :lieu_naissance, statut_marital = :statut_marital,
			occupation = :occupation, employeur = :employeur, profession = :profession,
			niveau_scolaire = :niveau_scolaire, langue_preferee = :langue_preferee,
			origine_ethnique = :origine_ethnique, premiere_nation = :premiere_nation,
			identite_genre = :identite_genre, orientation_sexuelle = :orientation_sexuelle,
			religion = :religion, premiere_langue = :premiere_langue,
			telephone = :telephone, cellulaire = :cellulaire, email = :email,
			adresse = :adresse, appartement = :appartement, code_postal = :code_postal,
			ville = :ville, mcode = :mcode, arrcod = :arrcod,
			pays = :pays, province = :province,
			numero_assurance_maladie = :numero_assurance_maladie,
			numero_securite_sociale = :numero_securite_sociale,
			no_hcm = :no_hcm, no_chaur = :no_chaur,
			no_dossier_leopard = :no_dossier_leopard,
			medecin_famille_No_Licence = :medecin_famille_No_Licence,
			notaire_id = :notaire_id, pharmacie_id = :pharmacie_id,
			pivot_id = :pivot_id, rpa_id = :rpa_id,
			chsld_id = :chsld_id, ri_id = :ri_id,
			procuration_bancaire = :procuration_bancaire,
			procuration_notariee = :procuration_notariee,
			tutelle_active = :tutelle_active, tutelle_type = :tutelle_type,
			tutelle_bien = :tutelle_bien, tutelle_personne = :tutelle_personne,
			tutelle_date_debut = :tutelle_date_debut,
			tutelle_date_renouvellement = :tutelle_date_renouvellement,
			tuteur_nom = :tuteur_nom, tuteur_prenom = :tuteur_prenom,
			tuteur_telephone = :tuteur_telephone, tuteur_cellulaire = :tuteur_cellulaire,
			tuteur_email = :tuteur_email, tuteur_adresse = :tuteur_adresse,
			tuteur_code_postal = :tuteur_code_postal, tuteur_ville = :tuteur_ville,
			note_fixe = :note_fixe,
			Actif = :Actif, dcd = :dcd, date_deces = :date_deces
		WHERE id = :id`

	_, err = db.NamedExec(query, enc)
	if err != nil {
		return fmt.Errorf("erreur mise à jour client %d: %w", enc.ID, err)
	}
	return nil
}

func (db *Database) ArchiveClient(id int) error {
	_, err := db.Exec(`UPDATE clients SET Actif = 0 WHERE id = ?`, id)
	return err
}

func (db *Database) DeleteClient(id int) error {
	_, err := db.Exec(`DELETE FROM clients WHERE id = ?`, id)
	return err
}

// ── Chiffrement ───────────────────────────────────────────────────────────────

func decryptClient(c *models.Client, s *crypto.CryptoService) error {
	var err error
	c.DateNaissance, err = s.DecryptStringPtr(c.DateNaissance)
	if err != nil {
		return err
	}
	c.Sexe, _ = s.DecryptStringPtr(c.Sexe)
	c.LieuNaissance, _ = s.DecryptStringPtr(c.LieuNaissance)
	c.StatutMarital, _ = s.DecryptStringPtr(c.StatutMarital)
	c.Occupation, _ = s.DecryptStringPtr(c.Occupation)
	c.Employeur, _ = s.DecryptStringPtr(c.Employeur)
	c.Profession, _ = s.DecryptStringPtr(c.Profession)
	c.NiveauScolaire, _ = s.DecryptStringPtr(c.NiveauScolaire)
	c.LanguePreferee, _ = s.DecryptStringPtr(c.LanguePreferee)
	c.OrigineEthnique, _ = s.DecryptStringPtr(c.OrigineEthnique)
	c.IdentiteGenre, _ = s.DecryptStringPtr(c.IdentiteGenre)
	c.OrientationSexuelle, _ = s.DecryptStringPtr(c.OrientationSexuelle)
	c.Religion, _ = s.DecryptStringPtr(c.Religion)
	c.PremiereLangue, _ = s.DecryptStringPtr(c.PremiereLangue)
	c.Telephone, _ = s.DecryptStringPtr(c.Telephone)
	c.Cellulaire, _ = s.DecryptStringPtr(c.Cellulaire)
	c.Email, _ = s.DecryptStringPtr(c.Email)
	c.Adresse, _ = s.DecryptStringPtr(c.Adresse)
	c.Appartement, _ = s.DecryptStringPtr(c.Appartement)
	c.CodePostal, _ = s.DecryptStringPtr(c.CodePostal)
	c.Ville, _ = s.DecryptStringPtr(c.Ville)
	// mcode et arrcod → non chiffrés, on ne touche pas
	c.Pays, _ = s.DecryptStringPtr(c.Pays)
	c.Province, _ = s.DecryptStringPtr(c.Province)
	c.NumeroAssuranceMaladie, _ = s.DecryptStringPtr(c.NumeroAssuranceMaladie)
	c.NumeroSecuriteSociale, _ = s.DecryptStringPtr(c.NumeroSecuriteSociale)
	c.NoHCM, _ = s.DecryptStringPtr(c.NoHCM)
	c.NoCHAUR, _ = s.DecryptStringPtr(c.NoCHAUR)
	c.MedecinFamilleNoLicence, _ = s.DecryptStringPtr(c.MedecinFamilleNoLicence)
	c.ProcurationBancaire, _ = s.DecryptStringPtr(c.ProcurationBancaire)
	c.ProcurationNotariee, _ = s.DecryptStringPtr(c.ProcurationNotariee)
	// tutelle_active, tutelle_type, tutelle_bien, tutelle_personne → non chiffrés
	c.TutelleDateDebut, _ = s.DecryptStringPtr(c.TutelleDateDebut)
	c.TutelleDateRenouvellement, _ = s.DecryptStringPtr(c.TutelleDateRenouvellement)
	c.TuteurNom, _ = s.DecryptStringPtr(c.TuteurNom)
	c.TuteurPrenom, _ = s.DecryptStringPtr(c.TuteurPrenom)
	c.TuteurTelephone, _ = s.DecryptStringPtr(c.TuteurTelephone)
	c.TuteurCellulaire, _ = s.DecryptStringPtr(c.TuteurCellulaire)
	c.TuteurEmail, _ = s.DecryptStringPtr(c.TuteurEmail)
	c.TuteurAdresse, _ = s.DecryptStringPtr(c.TuteurAdresse)
	c.TuteurCodePostal, _ = s.DecryptStringPtr(c.TuteurCodePostal)
	c.TuteurVille, _ = s.DecryptStringPtr(c.TuteurVille)
	c.NoteFixe, _ = s.DecryptStringPtr(c.NoteFixe)
	c.DateDeces, _ = s.DecryptStringPtr(c.DateDeces)
	return nil
}

func encryptCreateClientRequest(req models.CreateClientRequest, s *crypto.CryptoService) (models.CreateClientRequest, error) {
	enc := req
	enc.DateNaissance, _ = s.EncryptStringPtr(req.DateNaissance)
	enc.Sexe, _ = s.EncryptStringPtr(req.Sexe)
	enc.LieuNaissance, _ = s.EncryptStringPtr(req.LieuNaissance)
	enc.StatutMarital, _ = s.EncryptStringPtr(req.StatutMarital)
	enc.Occupation, _ = s.EncryptStringPtr(req.Occupation)
	enc.Employeur, _ = s.EncryptStringPtr(req.Employeur)
	enc.Profession, _ = s.EncryptStringPtr(req.Profession)
	enc.NiveauScolaire, _ = s.EncryptStringPtr(req.NiveauScolaire)
	enc.LanguePreferee, _ = s.EncryptStringPtr(req.LanguePreferee)
	enc.OrigineEthnique, _ = s.EncryptStringPtr(req.OrigineEthnique)
	enc.IdentiteGenre, _ = s.EncryptStringPtr(req.IdentiteGenre)
	enc.OrientationSexuelle, _ = s.EncryptStringPtr(req.OrientationSexuelle)
	enc.Religion, _ = s.EncryptStringPtr(req.Religion)
	enc.PremiereLangue, _ = s.EncryptStringPtr(req.PremiereLangue)
	enc.Telephone, _ = s.EncryptStringPtr(req.Telephone)
	enc.Cellulaire, _ = s.EncryptStringPtr(req.Cellulaire)
	enc.Email, _ = s.EncryptStringPtr(req.Email)
	enc.Adresse, _ = s.EncryptStringPtr(req.Adresse)
	enc.Appartement, _ = s.EncryptStringPtr(req.Appartement)
	enc.CodePostal, _ = s.EncryptStringPtr(req.CodePostal)
	enc.Ville, _ = s.EncryptStringPtr(req.Ville)
	// mcode et arrcod → non chiffrés
	enc.Pays, _ = s.EncryptStringPtr(req.Pays)
	enc.Province, _ = s.EncryptStringPtr(req.Province)
	enc.NumeroAssuranceMaladie, _ = s.EncryptStringPtr(req.NumeroAssuranceMaladie)
	enc.NumeroSecuriteSociale, _ = s.EncryptStringPtr(req.NumeroSecuriteSociale)
	enc.NoHCM, _ = s.EncryptStringPtr(req.NoHCM)
	enc.NoCHAUR, _ = s.EncryptStringPtr(req.NoCHAUR)
	enc.MedecinFamilleNoLicence, _ = s.EncryptStringPtr(req.MedecinFamilleNoLicence)
	enc.ProcurationBancaire, _ = s.EncryptStringPtr(req.ProcurationBancaire)
	enc.ProcurationNotariee, _ = s.EncryptStringPtr(req.ProcurationNotariee)
	// tutelle statuts → non chiffrés
	enc.TutelleDateDebut, _ = s.EncryptStringPtr(req.TutelleDateDebut)
	enc.TutelleDateRenouvellement, _ = s.EncryptStringPtr(req.TutelleDateRenouvellement)
	enc.TuteurNom, _ = s.EncryptStringPtr(req.TuteurNom)
	enc.TuteurPrenom, _ = s.EncryptStringPtr(req.TuteurPrenom)
	enc.TuteurTelephone, _ = s.EncryptStringPtr(req.TuteurTelephone)
	enc.TuteurCellulaire, _ = s.EncryptStringPtr(req.TuteurCellulaire)
	enc.TuteurEmail, _ = s.EncryptStringPtr(req.TuteurEmail)
	enc.TuteurAdresse, _ = s.EncryptStringPtr(req.TuteurAdresse)
	enc.TuteurCodePostal, _ = s.EncryptStringPtr(req.TuteurCodePostal)
	enc.TuteurVille, _ = s.EncryptStringPtr(req.TuteurVille)
	enc.NoteFixe, _ = s.EncryptStringPtr(req.NoteFixe)
	enc.DateDeces, _ = s.EncryptStringPtr(req.DateDeces)
	return enc, nil
}

func encryptUpdateClientRequest(req models.UpdateClientRequest, s *crypto.CryptoService) (models.UpdateClientRequest, error) {
	enc := req
	enc.DateNaissance, _ = s.EncryptStringPtr(req.DateNaissance)
	enc.Sexe, _ = s.EncryptStringPtr(req.Sexe)
	enc.LieuNaissance, _ = s.EncryptStringPtr(req.LieuNaissance)
	enc.StatutMarital, _ = s.EncryptStringPtr(req.StatutMarital)
	enc.Occupation, _ = s.EncryptStringPtr(req.Occupation)
	enc.Employeur, _ = s.EncryptStringPtr(req.Employeur)
	enc.Profession, _ = s.EncryptStringPtr(req.Profession)
	enc.NiveauScolaire, _ = s.EncryptStringPtr(req.NiveauScolaire)
	enc.LanguePreferee, _ = s.EncryptStringPtr(req.LanguePreferee)
	enc.OrigineEthnique, _ = s.EncryptStringPtr(req.OrigineEthnique)
	enc.IdentiteGenre, _ = s.EncryptStringPtr(req.IdentiteGenre)
	enc.OrientationSexuelle, _ = s.EncryptStringPtr(req.OrientationSexuelle)
	enc.Religion, _ = s.EncryptStringPtr(req.Religion)
	enc.PremiereLangue, _ = s.EncryptStringPtr(req.PremiereLangue)
	enc.Telephone, _ = s.EncryptStringPtr(req.Telephone)
	enc.Cellulaire, _ = s.EncryptStringPtr(req.Cellulaire)
	enc.Email, _ = s.EncryptStringPtr(req.Email)
	enc.Adresse, _ = s.EncryptStringPtr(req.Adresse)
	enc.Appartement, _ = s.EncryptStringPtr(req.Appartement)
	enc.CodePostal, _ = s.EncryptStringPtr(req.CodePostal)
	enc.Ville, _ = s.EncryptStringPtr(req.Ville)
	// mcode et arrcod → non chiffrés
	enc.Pays, _ = s.EncryptStringPtr(req.Pays)
	enc.Province, _ = s.EncryptStringPtr(req.Province)
	enc.NumeroAssuranceMaladie, _ = s.EncryptStringPtr(req.NumeroAssuranceMaladie)
	enc.NumeroSecuriteSociale, _ = s.EncryptStringPtr(req.NumeroSecuriteSociale)
	enc.NoHCM, _ = s.EncryptStringPtr(req.NoHCM)
	enc.NoCHAUR, _ = s.EncryptStringPtr(req.NoCHAUR)
	enc.MedecinFamilleNoLicence, _ = s.EncryptStringPtr(req.MedecinFamilleNoLicence)
	enc.ProcurationBancaire, _ = s.EncryptStringPtr(req.ProcurationBancaire)
	enc.ProcurationNotariee, _ = s.EncryptStringPtr(req.ProcurationNotariee)
	// tutelle statuts → non chiffrés
	enc.TutelleDateDebut, _ = s.EncryptStringPtr(req.TutelleDateDebut)
	enc.TutelleDateRenouvellement, _ = s.EncryptStringPtr(req.TutelleDateRenouvellement)
	enc.TuteurNom, _ = s.EncryptStringPtr(req.TuteurNom)
	enc.TuteurPrenom, _ = s.EncryptStringPtr(req.TuteurPrenom)
	enc.TuteurTelephone, _ = s.EncryptStringPtr(req.TuteurTelephone)
	enc.TuteurCellulaire, _ = s.EncryptStringPtr(req.TuteurCellulaire)
	enc.TuteurEmail, _ = s.EncryptStringPtr(req.TuteurEmail)
	enc.TuteurAdresse, _ = s.EncryptStringPtr(req.TuteurAdresse)
	enc.TuteurCodePostal, _ = s.EncryptStringPtr(req.TuteurCodePostal)
	enc.TuteurVille, _ = s.EncryptStringPtr(req.TuteurVille)
	enc.NoteFixe, _ = s.EncryptStringPtr(req.NoteFixe)
	enc.DateDeces, _ = s.EncryptStringPtr(req.DateDeces)
	return enc, nil
}
