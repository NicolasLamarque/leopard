package schema

// ═══════════════════════════════════════════════════════════════════════════════
// SCHÉMA — Organisations & Payeurs
//
// Pourquoi ces tables :
//   La table `clients` est conçue pour des personnes physiques (patients).
//   Pour facturer une école, un CIUSSS, une entreprise, ou la mandataire
//   d'un client inapte — il faut un modèle de payeur flexible.
//
// Architecture :
//   organisations  → entités morales (CIUSSS, école, entreprise, organisme)
//   payeurs        → vue unifiée du "qui paie" :
//                    peut pointer vers un client, un contact, ou une organisation
//
// Règle de chiffrement (identique au reste du module) :
//   CHIFFRÉ  → tout ce qui identifie une personne ou organisation nominativement
//   EN CLAIR → montants, dates, statuts, types, flags
// ═══════════════════════════════════════════════════════════════════════════════

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : organisations
// Entités morales : CIUSSS, CLSC, écoles, entreprises, organismes communautaires
// ─────────────────────────────────────────────────────────────────────────────
var TableOrganisations = `
CREATE TABLE IF NOT EXISTS organisations (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    -- EN CLAIR : classification et statut
    type_org        TEXT    NOT NULL DEFAULT 'entreprise',
    -- ciusss | clsc | css | ecole | cpe | organisme | entreprise | autre
    statut          TEXT    NOT NULL DEFAULT 'actif',
    -- actif | inactif | archive
    no_fournisseur  TEXT,   -- EN CLAIR : code interne Leopard (ex: ORG-2025-001)

    -- CHIFFRÉ : identité nominative
    nom             TEXT    NOT NULL,   -- "CIUSSS de l'Estrie-CHUS"
    acronyme        TEXT,               -- "CIUSSS MCQ"
    no_organisme    TEXT,               -- NEQ, no. agrément, etc.

    -- CHIFFRÉ : coordonnées
    adresse         TEXT,
    ville           TEXT,
    code_postal     TEXT,
    province        TEXT,
    pays            TEXT,
    telephone       TEXT,
    telecopieur     TEXT,
    email_general   TEXT,
    site_web        TEXT,

    -- CHIFFRÉ : personne-ressource principale
    contact_nom         TEXT,   -- "Marie Tremblay"
    contact_titre       TEXT,   -- "Coordonnatrice clinique"
    contact_telephone   TEXT,
    contact_email       TEXT,

    -- EN CLAIR : informations contractuelles
    conditions_paiement INTEGER DEFAULT 30,  -- jours
    taux_tps_exempt     INTEGER DEFAULT 0,   -- 1 si org exempte de TPS
    mode_paiement_pref  TEXT    DEFAULT 'cheque',
    -- cheque | virement | interac | mandat

    -- CHIFFRÉ : notes et conditions particulières
    notes               TEXT,
    conditions_speciales TEXT,  -- clauses tarifaires négociées

    actif           INTEGER DEFAULT 1,
    created_by      INTEGER,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(created_by) REFERENCES users(id)
);
`

var IndexOrganisations = `
CREATE INDEX IF NOT EXISTS idx_organisations_type   ON organisations(type_org);
CREATE INDEX IF NOT EXISTS idx_organisations_statut ON organisations(statut);
CREATE INDEX IF NOT EXISTS idx_organisations_actif  ON organisations(actif);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : payeurs
// Vue unifiée "qui paie" — trois sources possibles :
//   1. Le client lui-même (suivi individuel standard)
//   2. Un contact du client (mandataire, parent, conjoint)
//   3. Une organisation (CIUSSS, école, employeur)
//
// Une facture/contrat pointe vers UN payeur.
// Un payeur pointe vers UNE des trois sources.
// ─────────────────────────────────────────────────────────────────────────────
var TablePayeurs = `
CREATE TABLE IF NOT EXISTS payeurs (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    -- EN CLAIR : type de payeur et liens FK
    type_payeur     TEXT    NOT NULL DEFAULT 'client',
    -- client | contact | organisation

    -- Une seule de ces trois FK sera non-nulle
    client_id       INTEGER,        -- si type_payeur = 'client'
    contact_id      INTEGER,        -- si type_payeur = 'contact'
    organisation_id INTEGER,        -- si type_payeur = 'organisation'

    -- EN CLAIR : métadonnées
    actif           INTEGER DEFAULT 1,
    est_defaut      INTEGER DEFAULT 0,  -- payeur par défaut pour ce client

    -- CHIFFRÉ : surcharge optionnelle (si l'adresse de facturation diffère)
    nom_facturation     TEXT,   -- peut différer du nom légal
    adresse_facturation TEXT,
    email_facturation   TEXT,
    notes               TEXT,

    created_by      INTEGER,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(client_id)       REFERENCES clients(id),
    FOREIGN KEY(contact_id)      REFERENCES contacts(id),
    FOREIGN KEY(organisation_id) REFERENCES organisations(id),
    FOREIGN KEY(created_by)      REFERENCES users(id),

    -- Contrainte : exactement une FK non-nulle
    CHECK (
        (client_id IS NOT NULL AND contact_id IS NULL AND organisation_id IS NULL) OR
        (client_id IS NULL AND contact_id IS NOT NULL AND organisation_id IS NULL) OR
        (client_id IS NULL AND contact_id IS NULL AND organisation_id IS NOT NULL)
    )
);
`

var IndexPayeurs = `
CREATE INDEX IF NOT EXISTS idx_payeurs_client  ON payeurs(client_id);
CREATE INDEX IF NOT EXISTS idx_payeurs_contact ON payeurs(contact_id);
CREATE INDEX IF NOT EXISTS idx_payeurs_org     ON payeurs(organisation_id);
CREATE INDEX IF NOT EXISTS idx_payeurs_actif   ON payeurs(actif);
`

// ─────────────────────────────────────────────────────────────────────────────
// MIGRATIONS — Ajouter payeur_id aux tables existantes
// À exécuter dans db.go → InitDB() après création des tables de base
// ─────────────────────────────────────────────────────────────────────────────
var MigrationFacturesPayeur = `
ALTER TABLE factures ADD COLUMN payeur_id INTEGER REFERENCES payeurs(id);
`

var MigrationContratsPayeur = `
ALTER TABLE contrats ADD COLUMN payeur_id INTEGER REFERENCES payeurs(id);
`

var MigrationRevenusPayeur = `
ALTER TABLE revenus ADD COLUMN payeur_id INTEGER REFERENCES payeurs(id);
`

// Index sur les nouvelles colonnes
var IndexPayeursForeignKeys = `
CREATE INDEX IF NOT EXISTS idx_factures_payeur ON factures(payeur_id);
CREATE INDEX IF NOT EXISTS idx_contrats_payeur ON contrats(payeur_id);
CREATE INDEX IF NOT EXISTS idx_revenus_payeur  ON revenus(payeur_id);
`

// ─────────────────────────────────────────────────────────────────────────────
// AllOrganisationTables — à ajouter dans AllFinanceTables dans finance_schema.go
// ─────────────────────────────────────────────────────────────────────────────
var AllOrganisationTables = []string{
	TableOrganisations,
	IndexOrganisations,
	TablePayeurs,
	IndexPayeurs,
}

// Migrations séparées — à exécuter avec ALTER TABLE (idempotent via recover)
var AllOrganisationMigrations = []string{
	MigrationFacturesPayeur,
	MigrationContratsPayeur,
	MigrationRevenusPayeur,
	IndexPayeursForeignKeys,
}
