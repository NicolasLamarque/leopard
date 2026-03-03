package models

// Client représente la structure complète de la table en base de données
type Client struct {
	ID                      int     `db:"id" json:"id"`
	Nom                     string  `db:"nom" json:"nom"`
	Prenom                  string  `db:"prenom" json:"prenom"`
	DateNaissance           *string `db:"date_naissance" json:"date_naissance,omitempty"`
	Sexe                    *string `db:"sexe" json:"sexe,omitempty"`
	LieuNaissance           *string `db:"lieu_naissance" json:"lieu_naissance,omitempty"`
	StatutMarital           *string `db:"statut_marital" json:"statut_marital,omitempty"`
	Occupation              *string `db:"occupation" json:"occupation,omitempty"`
	Employeur               *string `db:"employeur" json:"employeur,omitempty"`
	Profession              *string `db:"profession" json:"profession,omitempty"`
	NiveauScolaire          *string `db:"niveau_scolaire" json:"niveau_scolaire,omitempty"`
	LanguePreferee          *string `db:"langue_preferee" json:"langue_preferee,omitempty"`
	OrigineEthnique         *string `db:"origine_ethnique" json:"origine_ethnique,omitempty"`
	PremiereNation          int     `db:"premiere_nation" json:"premiere_nation"`
	IdentiteGenre           *string `db:"identite_genre" json:"identite_genre,omitempty"`
	OrientationSexuelle     *string `db:"orientation_sexuelle" json:"orientation_sexuelle,omitempty"`
	Religion                *string `db:"religion" json:"religion,omitempty"`
	PremiereLangue          *string `db:"premiere_langue" json:"premiere_langue,omitempty"`
	Telephone               *string `db:"telephone" json:"telephone,omitempty"`
	Cellulaire              *string `db:"cellulaire" json:"cellulaire,omitempty"`
	Email                   *string `db:"email" json:"email,omitempty"`
	Adresse                 *string `db:"adresse" json:"adresse,omitempty"`
	Appartement             *string `db:"appartement" json:"appartement,omitempty"`
	CodePostal              *string `db:"code_postal" json:"code_postal,omitempty"`
	Ville                   *string `db:"ville" json:"ville,omitempty"`
	Mcode                   *string `db:"mcode" json:"mcode,omitempty"`
	Arrcod                  *string `db:"arrcod" json:"arrcod,omitempty"`
	Pays                    *string `db:"pays" json:"pays,omitempty"`
	Province                *string `db:"province" json:"province,omitempty"`
	NumeroAssuranceMaladie  *string `db:"numero_assurance_maladie" json:"numero_assurance_maladie,omitempty"`
	NumeroSecuriteSociale   *string `db:"numero_securite_sociale" json:"numero_securite_sociale,omitempty"`
	NoHCM                   *string `db:"no_hcm" json:"no_hcm,omitempty"`
	NoCHAUR                 *string `db:"no_chaur" json:"no_chaur,omitempty"`
	NoDossierLeopard        *string `db:"no_dossier_leopard" json:"no_dossier_leopard,omitempty"`
	MedecinFamilleNoLicence *string `db:"medecin_famille_No_Licence" json:"medecin_famille_No_Licence,omitempty"`
	NotaireID               *int    `db:"notaire_id" json:"notaire_id,omitempty"`
	PharmacieID             *int    `db:"pharmacie_id" json:"pharmacie_id,omitempty"`
	PivotID                 *int    `db:"pivot_id" json:"pivot_id,omitempty"`
	RPAID                   *int    `db:"rpa_id" json:"rpa_id,omitempty"`
	CHSLDID                 *int    `db:"chsld_id" json:"chsld_id,omitempty"`
	RIID                    *int    `db:"ri_id" json:"ri_id,omitempty"`
	ProcurationBancaire     *string `db:"procuration_bancaire" json:"procuration_bancaire,omitempty"`
	ProcurationNotariee     *string `db:"procuration_notariee" json:"procuration_notariee,omitempty"`
	// ── Tutelle ─────────────────────────────────────────────────────────────────
	TutelleActive             int     `db:"tutelle_active" json:"tutelle_active"`
	TutelleType               *string `db:"tutelle_type" json:"tutelle_type,omitempty"`
	TutelleBien               int     `db:"tutelle_bien" json:"tutelle_bien"`
	TutellePersonne           int     `db:"tutelle_personne" json:"tutelle_personne"`
	TutelleDateDebut          *string `db:"tutelle_date_debut" json:"tutelle_date_debut,omitempty"`
	TutelleDateRenouvellement *string `db:"tutelle_date_renouvellement" json:"tutelle_date_renouvellement,omitempty"`
	TuteurNom                 *string `db:"tuteur_nom" json:"tuteur_nom,omitempty"`
	TuteurPrenom              *string `db:"tuteur_prenom" json:"tuteur_prenom,omitempty"`
	TuteurTelephone           *string `db:"tuteur_telephone" json:"tuteur_telephone,omitempty"`
	TuteurCellulaire          *string `db:"tuteur_cellulaire" json:"tuteur_cellulaire,omitempty"`
	TuteurEmail               *string `db:"tuteur_email" json:"tuteur_email,omitempty"`
	TuteurAdresse             *string `db:"tuteur_adresse" json:"tuteur_adresse,omitempty"`
	TuteurCodePostal          *string `db:"tuteur_code_postal" json:"tuteur_code_postal,omitempty"`
	TuteurVille               *string `db:"tuteur_ville" json:"tuteur_ville,omitempty"`
	// ── Métadonnées ──────────────────────────────────────────────────────────────
	NoteFixe  *string `db:"note_fixe" json:"note_fixe,omitempty"`
	Actif     int     `db:"Actif" json:"actif"`
	DCD       int     `db:"dcd" json:"dcd"`
	DateDeces *string `db:"date_deces" json:"date_deces,omitempty"`
	CreatedBy *int    `db:"created_by" json:"created_by,omitempty"`
	CreatedAt string  `db:"created_at" json:"created_at"`
}

// CreateClientRequest pour l'insertion (sans ID et CreatedAt)
type CreateClientRequest struct {
	Nom                     string  `json:"nom"`
	Prenom                  string  `json:"prenom"`
	DateNaissance           *string `json:"date_naissance,omitempty"`
	Sexe                    *string `json:"sexe,omitempty"`
	LieuNaissance           *string `json:"lieu_naissance,omitempty"`
	StatutMarital           *string `json:"statut_marital,omitempty"`
	Occupation              *string `json:"occupation,omitempty"`
	Employeur               *string `json:"employeur,omitempty"`
	Profession              *string `json:"profession,omitempty"`
	NiveauScolaire          *string `json:"niveau_scolaire,omitempty"`
	LanguePreferee          *string `json:"langue_preferee,omitempty"`
	OrigineEthnique         *string `json:"origine_ethnique,omitempty"`
	PremiereNation          int     `json:"premiere_nation"`
	IdentiteGenre           *string `json:"identite_genre,omitempty"`
	OrientationSexuelle     *string `json:"orientation_sexuelle,omitempty"`
	Religion                *string `json:"religion,omitempty"`
	PremiereLangue          *string `json:"premiere_langue,omitempty"`
	Telephone               *string `json:"telephone,omitempty"`
	Cellulaire              *string `json:"cellulaire,omitempty"`
	Email                   *string `json:"email,omitempty"`
	Adresse                 *string `json:"adresse,omitempty"`
	Appartement             *string `json:"appartement,omitempty"`
	CodePostal              *string `json:"code_postal,omitempty"`
	Ville                   *string `json:"ville,omitempty"`
	Mcode                   *string `json:"mcode,omitempty"`
	Arrcod                  *string `json:"arrcod,omitempty"`
	Pays                    *string `json:"pays,omitempty"`
	Province                *string `json:"province,omitempty"`
	NumeroAssuranceMaladie  *string `json:"numero_assurance_maladie,omitempty"`
	NumeroSecuriteSociale   *string `json:"numero_securite_sociale,omitempty"`
	NoHCM                   *string `json:"no_hcm,omitempty"`
	NoCHAUR                 *string `json:"no_chaur,omitempty"`
	NoDossierLeopard        *string `json:"no_dossier_leopard,omitempty"`
	MedecinFamilleNoLicence *string `json:"medecin_famille_No_Licence,omitempty"`
	NotaireID               *int    `json:"notaire_id,omitempty"`
	PharmacieID             *int    `json:"pharmacie_id,omitempty"`
	PivotID                 *int    `json:"pivot_id,omitempty"`
	RPAID                   *int    `json:"rpa_id,omitempty"`
	CHSLDID                 *int    `json:"chsld_id,omitempty"`
	RIID                    *int    `json:"ri_id,omitempty"`
	ProcurationBancaire     *string `json:"procuration_bancaire,omitempty"`
	ProcurationNotariee     *string `json:"procuration_notariee,omitempty"`
	// ── Tutelle ──────────────────────────────────────────────────────────────────
	TutelleActive             int     `json:"tutelle_active"`
	TutelleType               *string `json:"tutelle_type,omitempty"`
	TutelleBien               int     `json:"tutelle_bien"`
	TutellePersonne           int     `json:"tutelle_personne"`
	TutelleDateDebut          *string `json:"tutelle_date_debut,omitempty"`
	TutelleDateRenouvellement *string `json:"tutelle_date_renouvellement,omitempty"`
	TuteurNom                 *string `json:"tuteur_nom,omitempty"`
	TuteurPrenom              *string `json:"tuteur_prenom,omitempty"`
	TuteurTelephone           *string `json:"tuteur_telephone,omitempty"`
	TuteurCellulaire          *string `json:"tuteur_cellulaire,omitempty"`
	TuteurEmail               *string `json:"tuteur_email,omitempty"`
	TuteurAdresse             *string `json:"tuteur_adresse,omitempty"`
	TuteurCodePostal          *string `json:"tuteur_code_postal,omitempty"`
	TuteurVille               *string `json:"tuteur_ville,omitempty"`
	// ── Métadonnées ──────────────────────────────────────────────────────────────
	NoteFixe  *string `json:"note_fixe,omitempty"`
	Actif     int     `json:"actif"`
	DCD       int     `json:"dcd"`
	DateDeces *string `json:"date_deces,omitempty"`
}

// UpdateClientRequest pour la mise à jour (avec ID)
type UpdateClientRequest struct {
	ID                      int     `db:"id" json:"id"`
	Nom                     string  `db:"nom" json:"nom"`
	Prenom                  string  `db:"prenom" json:"prenom"`
	DateNaissance           *string `db:"date_naissance" json:"date_naissance,omitempty"`
	Sexe                    *string `db:"sexe" json:"sexe,omitempty"`
	LieuNaissance           *string `db:"lieu_naissance" json:"lieu_naissance,omitempty"`
	StatutMarital           *string `db:"statut_marital" json:"statut_marital,omitempty"`
	Occupation              *string `db:"occupation" json:"occupation,omitempty"`
	Employeur               *string `db:"employeur" json:"employeur,omitempty"`
	Profession              *string `db:"profession" json:"profession,omitempty"`
	NiveauScolaire          *string `db:"niveau_scolaire" json:"niveau_scolaire,omitempty"`
	LanguePreferee          *string `db:"langue_preferee" json:"langue_preferee,omitempty"`
	OrigineEthnique         *string `db:"origine_ethnique" json:"origine_ethnique,omitempty"`
	PremiereNation          int     `json:"premiere_nation"`
	IdentiteGenre           *string `db:"identite_genre" json:"identite_genre,omitempty"`
	OrientationSexuelle     *string `db:"orientation_sexuelle" json:"orientation_sexuelle,omitempty"`
	Religion                *string `db:"religion" json:"religion,omitempty"`
	PremiereLangue          *string `db:"premiere_langue" json:"premiere_langue,omitempty"`
	Telephone               *string `db:"telephone" json:"telephone,omitempty"`
	Cellulaire              *string `db:"cellulaire" json:"cellulaire,omitempty"`
	Email                   *string `db:"email" json:"email,omitempty"`
	Adresse                 *string `db:"adresse" json:"adresse,omitempty"`
	Appartement             *string `db:"appartement" json:"appartement,omitempty"`
	CodePostal              *string `db:"code_postal" json:"code_postal,omitempty"`
	Ville                   *string `db:"ville" json:"ville,omitempty"`
	Mcode                   *string `db:"mcode" json:"mcode,omitempty"`
	Arrcod                  *string `db:"arrcod" json:"arrcod,omitempty"`
	Pays                    *string `db:"pays" json:"pays,omitempty"`
	Province                *string `db:"province" json:"province,omitempty"`
	NumeroAssuranceMaladie  *string `db:"numero_assurance_maladie" json:"numero_assurance_maladie,omitempty"`
	NumeroSecuriteSociale   *string `db:"numero_securite_sociale" json:"numero_securite_sociale,omitempty"`
	NoHCM                   *string `db:"no_hcm" json:"no_hcm,omitempty"`
	NoCHAUR                 *string `db:"no_chaur" json:"no_chaur,omitempty"`
	NoDossierLeopard        *string `db:"no_dossier_leopard" json:"no_dossier_leopard,omitempty"`
	MedecinFamilleNoLicence *string `db:"medecin_famille_No_Licence" json:"medecin_famille_No_Licence,omitempty"`
	NotaireID               *int    `db:"notaire_id" json:"notaire_id,omitempty"`
	PharmacieID             *int    `db:"pharmacie_id" json:"pharmacie_id,omitempty"`
	PivotID                 *int    `db:"pivot_id" json:"pivot_id,omitempty"`
	RPAID                   *int    `db:"rpa_id" json:"rpa_id,omitempty"`
	CHSLDID                 *int    `db:"chsld_id" json:"chsld_id,omitempty"`
	RIID                    *int    `db:"ri_id" json:"ri_id,omitempty"`
	ProcurationBancaire     *string `db:"procuration_bancaire" json:"procuration_bancaire,omitempty"`
	ProcurationNotariee     *string `db:"procuration_notariee" json:"procuration_notariee,omitempty"`
	// ── Tutelle ──────────────────────────────────────────────────────────────────
	TutelleActive             int     `db:"tutelle_active" json:"tutelle_active"`
	TutelleType               *string `db:"tutelle_type" json:"tutelle_type,omitempty"`
	TutelleBien               int     `db:"tutelle_bien" json:"tutelle_bien"`
	TutellePersonne           int     `db:"tutelle_personne" json:"tutelle_personne"`
	TutelleDateDebut          *string `db:"tutelle_date_debut" json:"tutelle_date_debut,omitempty"`
	TutelleDateRenouvellement *string `db:"tutelle_date_renouvellement" json:"tutelle_date_renouvellement,omitempty"`
	TuteurNom                 *string `db:"tuteur_nom" json:"tuteur_nom,omitempty"`
	TuteurPrenom              *string `db:"tuteur_prenom" json:"tuteur_prenom,omitempty"`
	TuteurTelephone           *string `db:"tuteur_telephone" json:"tuteur_telephone,omitempty"`
	TuteurCellulaire          *string `db:"tuteur_cellulaire" json:"tuteur_cellulaire,omitempty"`
	TuteurEmail               *string `db:"tuteur_email" json:"tuteur_email,omitempty"`
	TuteurAdresse             *string `db:"tuteur_adresse" json:"tuteur_adresse,omitempty"`
	TuteurCodePostal          *string `db:"tuteur_code_postal" json:"tuteur_code_postal,omitempty"`
	TuteurVille               *string `db:"tuteur_ville" json:"tuteur_ville,omitempty"`
	// ── Métadonnées ──────────────────────────────────────────────────────────────
	NoteFixe  *string `db:"note_fixe" json:"note_fixe,omitempty"`
	Actif     int     `db:"Actif" json:"actif"`
	DCD       int     `db:"dcd" json:"dcd"`
	DateDeces *string `db:"date_deces" json:"date_deces,omitempty"`
}
