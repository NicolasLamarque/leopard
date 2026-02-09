package schema

var TableRi = `
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
`
