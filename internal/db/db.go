package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite" // Driver SQLite pur Go
)

type Database struct {
	*sqlx.DB
}

func New(path string) (*Database, error) {
	// Connexion simple
	db, err := sqlx.Connect("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("impossible de se connecter à SQLite: %w", err)
	}

	// Activation des clés étrangères
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("impossible d'activer les foreign_keys: %w", err)
	}

	// TON SCHÉMA EXACT (Aucune modification)
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		salt TEXT,
		full_name TEXT NOT NULL,
		role TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS T_Mun (
		IDMun	INTEGER,
		MunCode	INTEGER,
		Désignation	TEXT,
		MunName	TEXT,
		MRC	TEXT,
		RegAdm	TEXT,
		AjoutListeMun	TEXT
	);
	CREATE TABLE IF NOT EXISTS "T_pays" (
		"ID_Pays"	INTEGER,
		"PAYS"	TEXT,
		"Alpha2"	TEXT,
		"Alpha3"	TEXT,
		"numerique"	INTEGER
	);

	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		token TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT NOT NULL,
		prenom TEXT NOT NULL,
		date_naissance DATE,
		sexe TEXT,
		lieu_naissance TEXT,
		statut_marital TEXT,
		occupation TEXT,
		employeur TEXT,
		profession TEXT,
		niveau_scolaire TEXT,
		langue_preferee TEXT,
		origine_ethnique TEXT,
		premiere_nation INTEGER DEFAULT 0,
		identite_genre TEXT,
		orientation_sexuelle TEXT,
		religion TEXT,
		premiere_langue TEXT,
		telephone TEXT,
		cellulaire TEXT,
		email TEXT,
		adresse TEXT,
		appartement TEXT,
		code_postal TEXT,
		ville TEXT,
		pays TEXT,
		province TEXT,
		numero_assurance_maladie TEXT,
		numero_securite_sociale TEXT,
		no_hcm TEXT,
		no_chaur TEXT,
		no_dossier_leopard TEXT,
		medecin_famille_No_Licence TEXT,
		notaire_id INTEGER,
		pivot_id INTEGER,
		rpa_id INTEGER,
		chsld_id INTEGER,
		ri_id INTEGER,
		procuration_bancaire TEXT,
		procuration_notariee TEXT,
		note_fixe TEXT,
		Actif INTEGER DEFAULT 1,
		dcd INTEGER DEFAULT 0,
		date_deces DATE,
		created_by INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(created_by) REFERENCES users(id)
	);

CREATE TABLE IF NOT EXISTS contacts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	Nom TEXT NOT NULL,
	Prenom TEXT NOT NULL,
	Telephone TEXT,
	Cellulaire TEXT,
	Adresse TEXT,
	Code_Postal TEXT,
	Ville TEXT,
	Pays TEXT,
	Email TEXT,
	employeur TEXT,
	Profession TEXT,
	Relation_Client TEXT,
	lien_familial TEXT,
	force_lien INTEGER DEFAULT 0,
	contact_urgence INTEGER DEFAULT 0,
	aidant_naturel INTEGER DEFAULT 0,
	type_de_contact TEXT,
	procuration_bancaire INTEGER DEFAULT 0,
	procuration_notariee INTEGER DEFAULT 0,
	note_fixe TEXT,
	relation_type TEXT NOT NULL,
	client_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(client_id) REFERENCES clients(id) 
);

CREATE TABLE IF NOT EXISTS medecins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		licence TEXT NOT NULL UNIQUE,
		civilite TEXT,
		prenom TEXT,
		nom TEXT,
		nomComplet TEXT,
		statut TEXT,
		specialisation TEXT,
		service TEXT,
		departement TEXT,
		installation_principale TEXT,
		rls TEXT,
		telephone TEXT,
		extension TEXT,
		cellulaire TEXT,
		email TEXT,
		adresse TEXT,
		code_postal TEXT,
		ville TEXT,
		pays TEXT DEFAULT 'Canada',
		Note_fixe TEXT,
		actif INTEGER DEFAULT 1,
		created_by INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    date_note DATETIME DEFAULT CURRENT_TIMESTAMP,
    date_intervention DATETIME,
	heure_intervention TEXT,
	duree_intervention TEXT,
    mode_intervention TEXT,
    type_intervention TEXT,
    type_note TEXT,
    titre TEXT,
    contenu TEXT,
    verrouille INTEGER DEFAULT 0,
    signature_nom TEXT,
    signature_date DATETIME,
    note_liee_id INTEGER DEFAULT NULL,       
    type_lien TEXT DEFAULT NULL,             
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(client_id) REFERENCES clients(id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(note_liee_id) REFERENCES notes(id) 
);

CREATE TABLE IF NOT EXISTS user_settings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER UNIQUE NOT NULL,
		theme TEXT DEFAULT 'light',
		language TEXT DEFAULT 'fr',
		notifications_enabled INTEGER DEFAULT 1,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	
CREATE TABLE IF NOT EXISTS "T_CHSLD" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		"Region"	TEXT,
		"TitreCHSLD"	TEXT,
		"Adresse"	TEXT,
		"Municipalite"	TEXT,
		"CodePostal"	TEXT,
		"Telephone"	TEXT,
		telecopieur	TEXT,
		"Poste_Garde_infirmiere" TEXT,
		"Capacite"	INTEGER,	
		"TypeCHSLD"	TEXT,
		"Proprietaire"	TEXT,
		"DateCertification"	TEXT,
		"Statut"	TEXT,
		"SourceURL"	TEXT,
		"InfosCHSLD"	TEXT,
		"Notes"	TEXT,
		"DateAjout"	TEXT,
		"DateMaj"	TEXT
	);

CREATE TABLE IF NOT EXISTS "T_RI" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		"Region"	TEXT,
		"TitreRI"	TEXT,
		"Adresse"	TEXT,
		"Municipalite"	TEXT,
		"CodePostal"	TEXT,
		"Telephone"	TEXT,
		"telecopieur"	TEXT,
		"direction_generale"	TEXT,
		"direction_tel"	TEXT,
		"courriel"	TEXT,
		"Poste_Garde_infirmiere" TEXT,
		"Capacite"	INTEGER,	
		"TypeRI"	TEXT,
		"Proprietaire"	TEXT,
		"DateCertification"	TEXT,
		"Statut"	TEXT,
		"SourceURL"	TEXT,
		"InfosRI"	TEXT,
		"Notes"	TEXT,
		"DateAjout"	TEXT,
		"DateMaj"	TEXT
	);

CREATE TABLE IF NOT EXISTS residences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    region TEXT,
    registre TEXT NOT NULL UNIQUE,
	numero_interne TEXT,
    titre TEXT NOT NULL,
    municipalite TEXT,
    adresse TEXT,
    ville TEXT,
    code_postal TEXT,
    telephone TEXT,
    capacite INTEGER,
    type_resid TEXT,
    proprietaires TEXT,
    services TEXT,
    date_certification TEXT,
    statut TEXT DEFAULT 'actif',
    source_url TEXT,
	source_url_detaillee TEXT,
    notes TEXT,
    date_ajout DATETIME,
    date_maj DATETIME,
    date_fermeture DATETIME,
    derniere_verification DATETIME
);
CREATE TABLE IF NOT EXISTS notaires (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	civilite TEXT NOT NULL,
	prenom TEXT NOT NULL,
	nom TEXT NOT NULL,
	telephone TEXT,
	cellulaire TEXT,
	telecopieur TEXT,
	adresse TEXT,
	code_postal TEXT,
	ville TEXT,
	pays TEXT DEFAULT 'Canada',
	email TEXT,
	secteur_activite TEXT,
	note_fixe TEXT,
	actif INTEGER DEFAULT 1,
	created_by INTEGER,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du schéma: %w", err)
	}

	return &Database{db}, nil
}
