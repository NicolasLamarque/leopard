package schema

var TableEvalModels = `
-- 1. Les Modèles (ex: Curateur Public v1, Suivi Psychosocial, Annexe A)
CREATE TABLE IF NOT EXISTS evaluation_definitions (
    id TEXT PRIMARY KEY,        -- ex: 'qc_curateur_2024'
    nom TEXT NOT NULL,          -- ex: 'Évaluation du Curateur Public'
    schema_json TEXT NOT NULL,  -- Le "plan" du formulaire (sections, champs, types)
    version INTEGER DEFAULT 1,
    active INTEGER DEFAULT 1
);

-- 2. Les Évaluations (La puissance flexible)
-- On remplace votre table actuelle par celle-ci
CREATE TABLE IF NOT EXISTS evaluations_sociales_v2 (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    model_id TEXT NOT NULL,     -- Référence à evaluation_definitions
    no_leopard TEXT NOT NULL,
    
    -- LE COEUR DU SYSTÈME :
    -- On stocke TOUTES les réponses dans un seul BLOC JSON CRYPTÉ.
    -- Plus besoin de modifier la DB quand une question change !
    payload_crypte TEXT NOT NULL, 
    
    statut TEXT DEFAULT 'brouillon', -- 'brouillon', 'finalise', 'verrouille'
    signature_nom TEXT,
    date_signature DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(client_id) REFERENCES clients(id),
    FOREIGN KEY(model_id) REFERENCES evaluation_definitions(id)
);;`
