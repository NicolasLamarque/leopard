package schema

var TableAppels = `
CREATE TABLE IF NOT EXISTS appels (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date_appel DATETIME DEFAULT CURRENT_TIMESTAMP,
    heure_appel TEXT,
    
    -- Appelant (qui téléphone) - CRYPTÉ
    appelant_nom TEXT,
    appelant_prenom TEXT,
    appelant_telephone TEXT,
    appelant_relation TEXT, -- Ex: "Conjoint", "Fils", "Aidant naturel", "Lui-même"
    
    -- Client concerné (optionnel - NULL si nouveau prospect)
    client_id INTEGER,
    
    -- Si pas de client_id, infos minimales du prospect - CRYPTÉ
    prospect_nom TEXT,
    prospect_prenom TEXT,
    prospect_telephone TEXT,
    
    -- Détails de l'appel
    type_demande TEXT, -- 'evaluation', 'mandat_protection', 'consultation', 'information', 'autre'
    motif_appel TEXT, -- Description libre - CRYPTÉ
    priorite TEXT DEFAULT 'normale', -- 'urgente', 'haute', 'normale', 'basse'
    
    -- Suivi
    statut TEXT DEFAULT 'nouveau', -- 'nouveau', 'en_attente', 'rdv_planifie', 'complete', 'annule'
    notes_internes TEXT, -- CRYPTÉ
    
    -- RDV
    rdv_date DATE,
    rdv_heure TEXT,
    rdv_lieu TEXT,
    
    -- Gestion
    recu_par INTEGER NOT NULL, -- user_id qui a pris l'appel
    assigne_a INTEGER, -- user_id assigné pour le suivi
    
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY(client_id) REFERENCES clients(id) ON DELETE SET NULL,
    FOREIGN KEY(recu_par) REFERENCES users(id),
    FOREIGN KEY(assigne_a) REFERENCES users(id)
);

-- Index pour performance
CREATE INDEX IF NOT EXISTS idx_appels_date ON appels(date_appel);
CREATE INDEX IF NOT EXISTS idx_appels_statut ON appels(statut);
CREATE INDEX IF NOT EXISTS idx_appels_client ON appels(client_id);
CREATE INDEX IF NOT EXISTS idx_appels_priorite ON appels(priorite);


`
