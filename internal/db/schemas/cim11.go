package schema

// TableCim11 contient les tables CIM-11 et leurs métadonnées de sync OMS.
// À ajouter dans db.go : schema.TableCim11 dans le fullSchema.
// Données publiques OMS — pas de chiffrement.
var TableCim11 = `
CREATE TABLE IF NOT EXISTS cim11_codes (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    code            TEXT    NOT NULL UNIQUE,
    titre_fr        TEXT    NOT NULL DEFAULT '',
    titre_en        TEXT    NOT NULL DEFAULT '',
    chapitre_code   TEXT    NOT NULL DEFAULT '',
    chapitre_titre  TEXT    NOT NULL DEFAULT '',
    bloc_code       TEXT    NOT NULL DEFAULT '',
    bloc_titre      TEXT    NOT NULL DEFAULT '',
    parent_code     TEXT    NOT NULL DEFAULT '',
    niveau          INTEGER NOT NULL DEFAULT 0,
    is_billable     INTEGER NOT NULL DEFAULT 1,
    actif           INTEGER NOT NULL DEFAULT 1,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);
 
CREATE INDEX IF NOT EXISTS idx_cim11_code     ON cim11_codes(code);
CREATE INDEX IF NOT EXISTS idx_cim11_titre_fr ON cim11_codes(titre_fr);
CREATE INDEX IF NOT EXISTS idx_cim11_titre_en ON cim11_codes(titre_en);
CREATE INDEX IF NOT EXISTS idx_cim11_chapitre ON cim11_codes(chapitre_code);
 
CREATE TABLE IF NOT EXISTS cim11_sync_meta (
    id            INTEGER PRIMARY KEY CHECK (id = 1),
    derniere_sync DATETIME,
    version_oms   TEXT    NOT NULL DEFAULT '',
    nb_codes_fr   INTEGER NOT NULL DEFAULT 0,
    nb_codes_en   INTEGER NOT NULL DEFAULT 0,
    statut        TEXT    NOT NULL DEFAULT 'jamais'
);
 
INSERT OR IGNORE INTO cim11_sync_meta (id, statut) VALUES (1, 'jamais');
`
