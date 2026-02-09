package schema

var TableIntervenants = `
CREATE TABLE IF NOT EXISTS intervenants (id INTEGER PRIMARY KEY AUTOINCREMENT,
    nom_complet TEXT NOT NULL,
    titre_emploi TEXT,
    direction TEXT,
    installation TEXT,
    
    -- Contact
    telephone TEXT,
    poste TEXT,
    cellulaire_pro TEXT,
    cellulaire_perso TEXT,
    
    -- Emails
    courriel_personnel TEXT,
    courriel_professionnel TEXT,
    courrier_interne TEXT,
    
    -- Statuts
    actif INTEGER DEFAULT 1,
    is_intervenant_social INTEGER DEFAULT 0,
    numero_permis TEXT,
	Ordre_professionnel TEXT,
    date_naissance DATE,

	note_fixe TEXT,
    
    personne_ressource_reseau_dir TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`
