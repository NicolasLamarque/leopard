package schema

var TableCHSLD = `
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
`
