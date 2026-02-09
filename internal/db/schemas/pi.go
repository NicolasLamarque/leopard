package schema

var TablePI = `
CREATE TABLE IF NOT EXISTS plans_intervention (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    titre TEXT NOT NULL,
    type_plan TEXT DEFAULT 'psychosocial',
    statut TEXT DEFAULT 'brouillon', -- brouillon, actif, termine, archive
    date_debut DATE,
    date_fin_prevue DATE,
    date_revision_prevue DATE,
    contexte TEXT,
    problematique TEXT,
    forces TEXT,
    objectifs TEXT,
    interventions TEXT,
    resultats TEXT,
    ententes TEXT,
    verrouille INTEGER DEFAULT 0,
    signature_nom TEXT,
    date_signature DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(client_id) REFERENCES clients(id) ON DELETE CASCADE,
    FOREIGN KEY(created_by) REFERENCES users(id)
);
CREATE VIEW IF NOT EXISTS v_pi_details AS
SELECT 
    pi.*,
    c.nom AS client_nom,
    c.prenom AS client_prenom,
    c.no_dossier_leopard AS client_leopard,
    u.full_name AS auteur_nom
FROM plans_intervention pi
JOIN clients c ON pi.client_id = c.id
JOIN users u ON pi.created_by = u.id;
CREATE INDEX IF NOT EXISTS idx_pi_client ON plans_intervention(client_id);
CREATE INDEX IF NOT EXISTS idx_pi_statut ON plans_intervention(statut);

`
