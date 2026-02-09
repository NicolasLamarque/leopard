package schema

var TableNotes = `
CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    client_NoRAMQ TEXT, 
    client_Nom TEXT,
    client_Prenom TEXT,
    client_Telephone TEXT, 
    client_Cellulaire TEXT, 
    client_NoLeopard TEXT, 
    client_Adresse TEXT, 
	client_appartement TEXT,
	client_code_postal TEXT,
	client_ville TEXT,
	client_pays TEXT,
	client_province TEXT,
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
    note_tardive INTEGER DEFAULT 0,
    note_de_tier INTEGER DEFAULT 0,
    signature_nom TEXT,
    signature_date DATETIME,
    note_liee_id INTEGER DEFAULT NULL,       
    type_lien TEXT DEFAULT NULL,             
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(client_id) REFERENCES clients(id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(note_liee_id) REFERENCES notes(id) 
);

`
