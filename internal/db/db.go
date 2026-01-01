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
		telephone TEXT,
		cellulaire TEXT,
		email TEXT,
		adresse TEXT,
		code_postal TEXT,
		ville TEXT,
		pays TEXT,
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
		note_fixe TEXT,
		Actif INTEGER DEFAULT 1,
		dcd INTEGER DEFAULT 0,
		created_by INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(created_by) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS medecins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		licence TEXT NOT NULL,
		nomComplet TEXT,
		specialisation TEXT,
		telephone TEXT,
		extension TEXT,
		cellulaire TEXT,
		email TEXT,
		adresse TEXT,
		code_postal TEXT,
		ville TEXT,
		pays TEXT,
		Note_fixe TEXT,
		actif INTEGER DEFAULT 1,
		created_by INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		date_note DATETIME DEFAULT CURRENT_TIMESTAMP,
		date_intervention DATETIME,
		mode_intervention TEXT,
		type_intervention TEXT,
		type_note TEXT,
		titre TEXT,
		contenu TEXT,
		verrouille INTEGER DEFAULT 0,
		signature_nom TEXT,
		signature_date DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(client_id) REFERENCES clients(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
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
		"Region"	TEXT,
		"TitreCHSLD"	TEXT,
		"Adresse"	TEXT,
		"InfosCHSLD"	TEXT,
		"EtablissementResponsable"	TEXT,
		"TypeInstallation"	TEXT
	);

CREATE TABLE IF NOT EXISTS residences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    region TEXT,
    registre TEXT NOT NULL UNIQUE,
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
    notes TEXT,
    date_ajout DATETIME,
    date_maj DATETIME,
    date_fermeture DATETIME,
    derniere_verification DATETIME
);

	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du schéma: %w", err)
	}

	return &Database{db}, nil
}
