package schema

var TableContacts = `
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
`
