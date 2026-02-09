package schema

// Tables de référence des notaires
var TableNotaires = `
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
