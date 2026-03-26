package schema

// TableCim11Codes contient les codes diagnostics CIM-11 (FR + EN)
// Peuplée par le script d'import via l'API OMS (voir cim11_import.go)
var TableCim11Codes = `
CREATE TABLE IF NOT EXISTS cim11_codes (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    code            TEXT    NOT NULL UNIQUE,   -- ex: "8B20"
    titre_fr        TEXT    NOT NULL DEFAULT '',
    titre_en        TEXT    NOT NULL DEFAULT '',
    chapitre_code   TEXT    NOT NULL DEFAULT '',  -- ex: "08"
    chapitre_titre  TEXT    NOT NULL DEFAULT '',  -- ex: "Maladies du système nerveux"
    bloc_code       TEXT    NOT NULL DEFAULT '',
    bloc_titre      TEXT    NOT NULL DEFAULT '',
    parent_code     TEXT    NOT NULL DEFAULT '',
    niveau          INTEGER NOT NULL DEFAULT 0,   -- 0=chapitre, 1=bloc, 2=catégorie, 3=sous-cat
    is_billable     INTEGER NOT NULL DEFAULT 1,   -- 1 = code facturable/utilisable
    actif           INTEGER NOT NULL DEFAULT 1,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);
 
CREATE INDEX IF NOT EXISTS idx_cim11_code     ON cim11_codes(code);
CREATE INDEX IF NOT EXISTS idx_cim11_titre_fr ON cim11_codes(titre_fr);
CREATE INDEX IF NOT EXISTS idx_cim11_titre_en ON cim11_codes(titre_en);
CREATE INDEX IF NOT EXISTS idx_cim11_chapitre ON cim11_codes(chapitre_code);
 
-- Métadonnées de la dernière synchronisation CIM-11
CREATE TABLE IF NOT EXISTS cim11_sync_meta (
    id          INTEGER PRIMARY KEY CHECK (id = 1),  -- Singleton
    derniere_sync   DATETIME,
    version_oms     TEXT NOT NULL DEFAULT '',        -- ex: "2024-01"
    nb_codes_fr     INTEGER NOT NULL DEFAULT 0,
    nb_codes_en     INTEGER NOT NULL DEFAULT 0,
    statut          TEXT NOT NULL DEFAULT 'jamais'   -- 'jamais' | 'ok' | 'erreur'
);
 
INSERT OR IGNORE INTO cim11_sync_meta (id, statut) VALUES (1, 'jamais');
`

// TableSuivisMedicaux : suivi médical d'un client par un médecin avec diagnostic CIM-11
var TableSuivisMedicaux = `
CREATE TABLE IF NOT EXISTS suivis_medicaux (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id           INTEGER NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    medecin_licence     TEXT    REFERENCES medecins(no_licence),   -- médecin responsable du suivi
    cim11_code          TEXT    REFERENCES cim11_codes(code),      -- diagnostic principal
    cim11_code_second   TEXT    REFERENCES cim11_codes(code),      -- diagnostic secondaire (optionnel)
    motif_libre         TEXT    NOT NULL DEFAULT '',               -- champ libre si Dx pas encore connu
    date_debut          DATE,
    frequence_id        INTEGER REFERENCES ref_listes(id),         -- catégorie: 'frequence_suivi'
    recurrence_id       INTEGER REFERENCES ref_listes(id),         -- catégorie: 'recurrence_suivi'
    prochain_rv         DATE,
    suivi_hopital       INTEGER NOT NULL DEFAULT 0,                -- 0=non, 1=oui
    hopital_nom         TEXT    NOT NULL DEFAULT '',               -- nom de l'hôpital si applicable
    notes_ts            TEXT    NOT NULL DEFAULT '',               -- notes du travailleur social
    actif               INTEGER NOT NULL DEFAULT 1,
    created_by          INTEGER REFERENCES users(id),
    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP
);
 
CREATE INDEX IF NOT EXISTS idx_suivis_client   ON suivis_medicaux(client_id);
CREATE INDEX IF NOT EXISTS idx_suivis_medecin  ON suivis_medicaux(medecin_licence);
CREATE INDEX IF NOT EXISTS idx_suivis_cim11    ON suivis_medicaux(cim11_code);
CREATE INDEX IF NOT EXISTS idx_suivis_actif    ON suivis_medicaux(actif);
`
