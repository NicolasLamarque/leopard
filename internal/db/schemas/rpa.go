package schema

var TableResidence = `
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
	
`
