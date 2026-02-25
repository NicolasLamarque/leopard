package schema

// ═══════════════════════════════════════════════════════════════════════════════
// SCHÉMA FINANCIER — Module Finances Leopard
// Profil : TS Privée Autonome · Québec · Wails v2 · SQLite · AES-GCM
//
// ┌─────────────────────────────────────────────────────────────────────────┐
// │  RÈGLE DE CHIFFREMENT — version finale                                  │
// │                                                                         │
// │  EN CLAIR  → Tout sauf les identifiants personnels directs :            │
// │              montants, taxes, km, dates, statuts, catégories,           │
// │              descriptions de séance, noms de services,                  │
// │              clauses de contrat, notes cliniques génériques.            │
// │              → SQL natif : SUM, AVG, GROUP BY, LIKE, ORDER BY           │
// │                                                                         │
// │  CHIFFRÉ   → Uniquement ce qui identifie ou permet de retracer          │
// │              une personne physique (Loi 25 / LPRPDE) :                  │
// │              • Noms & prénoms (clients, praticienne)                    │
// │              • Adresses, téléphones, courriels                          │
// │              • N° membre ordre professionnel, N° TPS/TVQ               │
// │              • Références paiement Interac (liées à un compte)          │
// └─────────────────────────────────────────────────────────────────────────┘
// ═══════════════════════════════════════════════════════════════════════════════

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : parametres_finance
// ─────────────────────────────────────────────────────────────────────────────
var TableParametresFinance = `
CREATE TABLE IF NOT EXISTS parametres_finance (
    id                              INTEGER PRIMARY KEY CHECK (id = 1),

    -- CHIFFRÉ : identifiants personnels de la praticienne
    nom_complet                     TEXT,
    titre_professionnel             TEXT,   -- chiffré : "T.S. — Travailleuse Sociale"
    no_membre_ordre                 TEXT,   -- chiffré : N° OTSTCFQ
    email_professionnel             TEXT,   -- chiffré
    telephone_professionnel         TEXT,   -- chiffré
    adresse_pratique                TEXT,   -- chiffré
    no_tps                          TEXT,   -- chiffré
    no_tvq                          TEXT,   -- chiffré
    interac_email                   TEXT,   -- chiffré : lié à un compte bancaire

    -- EN CLAIR : tarification et configuration
    taux_horaire_defaut             REAL    DEFAULT 0,
    taux_forfait_evaluation         REAL    DEFAULT 0,
    taux_rapport_expertise          REAL    DEFAULT 0,
    taux_km                         REAL    DEFAULT 0.72,
    sous_seuil_tps                  INTEGER DEFAULT 1,
    inscrite_tps                    INTEGER DEFAULT 0,
    inscrite_tvq                    INTEGER DEFAULT 0,
    taux_tps                        REAL    DEFAULT 0.05,
    taux_tvq                        REAL    DEFAULT 0.09975,
    prefixe_facture                 TEXT    DEFAULT 'FAC',
    annee_facture                   INTEGER,
    prochain_no_facture             INTEGER DEFAULT 1,
    jours_paiement_defaut           INTEGER DEFAULT 30,

    -- EN CLAIR : textes génériques affichés sur les PDF
    infos_paiement                  TEXT,   -- "Interac : voir email, libellé : Prénom + date"
    politique_annulation            TEXT,   -- "24h à l'avance"
    mentions_legales_pdf            TEXT,   -- pied de page

    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : services
// Catalogue des prestations — tout en clair (rien de nominatif).
// Conçu pour évoluer : consultations aujourd'hui, tracts/ateliers/capsules demain.
// ─────────────────────────────────────────────────────────────────────────────
var TableServices = `
CREATE TABLE IF NOT EXISTS services (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,

    -- EN CLAIR : tout (catalogue générique, aucun identifiant personnel)
    code                        TEXT    NOT NULL UNIQUE,
    -- Convention : CATEGORIE-TYPE[-DUREE]
    -- Exemples : "CONS-IND-60", "EVAL-PSY", "TRACT-ANX-A4", "CAPS-DEUIL-001"

    categorie                   TEXT    NOT NULL DEFAULT 'consultation',
    -- consultation | evaluation | rapport | forfait |
    -- atelier | groupe | formation | supervision |
    -- tract | capsule | ressource_num | autre

    type_livraison              TEXT    NOT NULL DEFAULT 'presentiel',
    -- presentiel | virtuel | hybride | asynchrone | imprime | numerique

    actif                       INTEGER DEFAULT 1,
    ordre_affichage             INTEGER DEFAULT 0,

    -- Contenu (EN CLAIR : noms génériques de catalogue)
    nom                         TEXT    NOT NULL,
    description_courte          TEXT,   -- 1 ligne → lignes de facture
    description_longue          TEXT,   -- détail → contrats, site
    public_cible                TEXT,   -- "Adultes, troubles anxieux"
    notes_internes              TEXT,

    -- Tarification (EN CLAIR : calculs directs)
    mode_tarification           TEXT    NOT NULL DEFAULT 'horaire',
    -- horaire | acte | forfait | gratuit | prix_libre | abonnement
    duree_minutes               INTEGER,
    taux_horaire                REAL,
    tarif_unitaire              REAL,
    tarif_min                   REAL,
    tarif_max                   REAL,

    -- Fiscal (EN CLAIR)
    exempte_taxes               INTEGER DEFAULT 1,
    tps_applicable              INTEGER DEFAULT 0,
    tvq_applicable              INTEGER DEFAULT 0,

    -- Extensibilité future (EN CLAIR)
    nb_places_max               INTEGER,            -- ateliers / groupes
    nb_seances_forfait          INTEGER,            -- packages
    duree_programme_semaines    INTEGER,
    format_tract                TEXT,               -- "Recto A5", "Trifold A4"
    couleur_tract               TEXT,               -- "N&B", "Couleur"
    duree_video_minutes         INTEGER,
    url_ressource               TEXT,
    tags                        TEXT,               -- JSON : ["anxiété","TCC","adulte"]
    icone                       TEXT,               -- nom lucide-vue icon

    created_by  INTEGER,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(created_by) REFERENCES users(id)
);
`

var IndexServices = `
CREATE INDEX IF NOT EXISTS idx_services_categorie ON services(categorie);
CREATE INDEX IF NOT EXISTS idx_services_actif     ON services(actif);
CREATE INDEX IF NOT EXISTS idx_services_ordre     ON services(ordre_affichage);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : revenus
// Tout en clair sauf la référence de paiement (confirmation Interac = traçable).
// ─────────────────────────────────────────────────────────────────────────────
var TableRevenus = `
CREATE TABLE IF NOT EXISTS revenus (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id                   INTEGER NOT NULL,
    service_id                  INTEGER,
    facture_id                  INTEGER,

    -- EN CLAIR : tout
    date_service                DATE    NOT NULL,
    categorie                   TEXT    NOT NULL,
    mode_facturation            TEXT    NOT NULL DEFAULT 'horaire',
    statut_paiement             TEXT    NOT NULL DEFAULT 'en_attente',
    -- en_attente | paye | partiel | annule
    mode_paiement               TEXT,
    -- interac | especes | cheque
    date_paiement               DATE,
    duree_heures                REAL    DEFAULT 0,
    taux_horaire_applique       REAL    DEFAULT 0,
    montant_brut                REAL    NOT NULL,
    montant_tps                 REAL    DEFAULT 0,
    montant_tvq                 REAL    DEFAULT 0,
    montant_total               REAL    NOT NULL,
    description                 TEXT,   -- "Consultation individuelle", "Suivi sem. 3"
    notes                       TEXT,

    -- CHIFFRÉ : référence Interac (numéro de transaction, lié à un compte bancaire)
    reference_paiement          TEXT,

    created_by  INTEGER,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(client_id)   REFERENCES clients(id),
    FOREIGN KEY(service_id)  REFERENCES services(id),
    FOREIGN KEY(facture_id)  REFERENCES factures(id),
    FOREIGN KEY(created_by)  REFERENCES users(id)
);
`

var IndexRevenus = `
CREATE INDEX IF NOT EXISTS idx_revenus_date      ON revenus(date_service);
CREATE INDEX IF NOT EXISTS idx_revenus_client    ON revenus(client_id);
CREATE INDEX IF NOT EXISTS idx_revenus_statut    ON revenus(statut_paiement);
CREATE INDEX IF NOT EXISTS idx_revenus_categorie ON revenus(categorie);
CREATE INDEX IF NOT EXISTS idx_revenus_service   ON revenus(service_id);
CREATE INDEX IF NOT EXISTS idx_revenus_montant   ON revenus(montant_total);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : depenses
// Entièrement en clair — aucune donnée personnelle ici.
// ─────────────────────────────────────────────────────────────────────────────
var TableDepenses = `
CREATE TABLE IF NOT EXISTS depenses (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,

    -- EN CLAIR : tout
    date_depense                DATE    NOT NULL,
    categorie                   TEXT    NOT NULL,
    -- bureau | formation | assurance | logiciel | materiel |
    -- deplacement | honoraires | marketing | autre
    deductible                  INTEGER DEFAULT 1,
    pct_utilisation_affaires    REAL    DEFAULT 100.0,
    est_kilometrage             INTEGER DEFAULT 0,
    sous_total                  REAL    NOT NULL,
    montant_tps                 REAL    DEFAULT 0,
    montant_tvq                 REAL    DEFAULT 0,
    montant_total               REAL    NOT NULL,
    montant_deductible          REAL    DEFAULT 0,
    km_parcourus                REAL    DEFAULT 0,
    taux_km_utilise             REAL    DEFAULT 0,
    fournisseur                 TEXT,   -- "Videotron", "Jean Coutu" — nom commercial, pas nominatif
    no_recu                     TEXT,
    description                 TEXT    NOT NULL,
    notes                       TEXT,
    chemin_recu                 TEXT,   -- path fichier local

    created_by  INTEGER,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(created_by) REFERENCES users(id)
);
`

var IndexDepenses = `
CREATE INDEX IF NOT EXISTS idx_depenses_date       ON depenses(date_depense);
CREATE INDEX IF NOT EXISTS idx_depenses_categorie  ON depenses(categorie);
CREATE INDEX IF NOT EXISTS idx_depenses_deductible ON depenses(deductible);
CREATE INDEX IF NOT EXISTS idx_depenses_montant    ON depenses(montant_total);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : factures
// ─────────────────────────────────────────────────────────────────────────────
var TableFactures = `
CREATE TABLE IF NOT EXISTS factures (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id                   INTEGER NOT NULL,

    -- EN CLAIR : tout sauf le nom du tiers-payant
    numero                      TEXT    NOT NULL UNIQUE,    -- "FAC-2025-0001"
    date_emission               DATE    NOT NULL,
    date_echeance               DATE    NOT NULL,
    date_paiement               DATE,
    statut                      TEXT    NOT NULL DEFAULT 'brouillon',
    -- brouillon | envoyee | payee | partielle | en_retard | annulee
    exempte_taxes               INTEGER DEFAULT 1,
    avec_tps                    INTEGER DEFAULT 0,
    avec_tvq                    INTEGER DEFAULT 0,
    jours_paiement              INTEGER DEFAULT 30,
    montant_sous_total          REAL    NOT NULL,
    montant_tps                 REAL    DEFAULT 0,
    montant_tvq                 REAL    DEFAULT 0,
    montant_total               REAL    NOT NULL,
    montant_paye                REAL    DEFAULT 0,
    montant_solde               REAL    NOT NULL,
    montant_tiers_payant        REAL    DEFAULT 0,
    montant_du_client           REAL    DEFAULT 0,
    titre_facture               TEXT,
    notes_client                TEXT,
    notes_internes              TEXT,

    -- CHIFFRÉ : nom de l'assureur / organisme (peut identifier indirectement)
    tiers_payant_nom            TEXT,

    created_by  INTEGER,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(client_id)  REFERENCES clients(id),
    FOREIGN KEY(created_by) REFERENCES users(id)
);
`

var IndexFactures = `
CREATE INDEX IF NOT EXISTS idx_factures_client   ON factures(client_id);
CREATE INDEX IF NOT EXISTS idx_factures_statut   ON factures(statut);
CREATE INDEX IF NOT EXISTS idx_factures_echeance ON factures(date_echeance);
CREATE INDEX IF NOT EXISTS idx_factures_emission ON factures(date_emission);
CREATE INDEX IF NOT EXISTS idx_factures_solde    ON factures(montant_solde);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : facture_lignes
// ─────────────────────────────────────────────────────────────────────────────
var TableFactureLignes = `
CREATE TABLE IF NOT EXISTS facture_lignes (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    facture_id      INTEGER NOT NULL,
    service_id      INTEGER,
    revenu_id       INTEGER,

    -- EN CLAIR : tout
    ordre           INTEGER DEFAULT 0,
    description     TEXT    NOT NULL,
    date_service    TEXT,
    quantite        REAL    NOT NULL,
    unite           TEXT,               -- "heure", "séance", "km", "forfait"
    tarif_unitaire  REAL    NOT NULL,
    montant_ligne   REAL    NOT NULL,

    FOREIGN KEY(facture_id) REFERENCES factures(id) ON DELETE CASCADE,
    FOREIGN KEY(service_id) REFERENCES services(id),
    FOREIGN KEY(revenu_id)  REFERENCES revenus(id)
);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : paiements
// ─────────────────────────────────────────────────────────────────────────────
var TablePaiements = `
CREATE TABLE IF NOT EXISTS paiements (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    facture_id      INTEGER NOT NULL,

    -- EN CLAIR : tout sauf la référence de transaction
    date_paiement   DATE    NOT NULL,
    source          TEXT    NOT NULL DEFAULT 'client',
    mode_paiement   TEXT,               -- interac | especes | cheque
    montant         REAL    NOT NULL,
    notes           TEXT,

    -- CHIFFRÉ : numéro de confirmation Interac (lié à un compte bancaire)
    reference       TEXT,

    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(facture_id) REFERENCES factures(id) ON DELETE CASCADE
);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : contrats
// Clauses et descriptions entièrement en clair (modèles standards réutilisés).
// Seule exception : frequence_seances si elle contient des détails personnels
// (ex: "les mardis à 14h" pourrait identifier via agenda) → jugement call.
// Décision : EN CLAIR, cohérent avec le reste.
// ─────────────────────────────────────────────────────────────────────────────
var TableContrats = `
CREATE TABLE IF NOT EXISTS contrats (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id       INTEGER NOT NULL,
    service_id      INTEGER,

    -- EN CLAIR : tout
    type_contrat            TEXT    NOT NULL DEFAULT 'suivi_individuel',
    -- suivi_individuel | evaluation_complete | rapport_expertise | forfait | atelier | autre
    statut                  TEXT    NOT NULL DEFAULT 'actif',
    -- actif | suspendu | termine | expire | resilie
    date_debut              DATE    NOT NULL,
    date_fin                DATE,
    renouvellement_auto     INTEGER DEFAULT 0,
    consentement_signe      INTEGER DEFAULT 0,
    date_signature          DATE,
    mode_facturation        TEXT    NOT NULL DEFAULT 'horaire',
    taux_horaire            REAL    DEFAULT 0,
    taux_forfait            REAL    DEFAULT 0,
    duree_seance_min        INTEGER,
    frequence_seances       TEXT,   -- "1x/semaine"

    -- Clauses (EN CLAIR : modèles standards, rien de nominatif)
    clause_objet                    TEXT,
    clause_objectifs                TEXT,
    clause_services_inclus          TEXT,
    clause_services_exclus          TEXT,
    clause_tarification             TEXT,
    clause_paiement_retard          TEXT,
    clause_annulation               TEXT,
    clause_absence_prolongee        TEXT,
    clause_confidentialite          TEXT,
    clause_limites_confidentialite  TEXT,
    clause_dossier_client           TEXT,
    clause_communication            TEXT,
    clause_urgences                 TEXT,
    clause_resiliation_client       TEXT,
    clause_resiliation_ts           TEXT,
    clauses_specifiques             TEXT,
    notes_internes                  TEXT,

    created_by  INTEGER,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(client_id)  REFERENCES clients(id),
    FOREIGN KEY(service_id) REFERENCES services(id),
    FOREIGN KEY(created_by) REFERENCES users(id)
);
`

var IndexContrats = `
CREATE INDEX IF NOT EXISTS idx_contrats_client ON contrats(client_id);
CREATE INDEX IF NOT EXISTS idx_contrats_statut ON contrats(statut);
`

// ─────────────────────────────────────────────────────────────────────────────
// TABLE : stats_mensuelles
// Cache d'agrégats — mis à jour par triggers SQLite, aucun déchiffrement requis.
// ─────────────────────────────────────────────────────────────────────────────
var TableStatsMensuelles = `
CREATE TABLE IF NOT EXISTS stats_mensuelles (
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    annee   INTEGER NOT NULL,
    mois    INTEGER NOT NULL,

    rev_total           REAL DEFAULT 0,
    rev_consultations   REAL DEFAULT 0,
    rev_evaluations     REAL DEFAULT 0,
    rev_rapports        REAL DEFAULT 0,
    rev_forfaits        REAL DEFAULT 0,
    rev_ateliers        REAL DEFAULT 0,
    rev_autres          REAL DEFAULT 0,
    nb_seances          INTEGER DEFAULT 0,
    nb_clients_uniques  INTEGER DEFAULT 0,

    dep_total           REAL DEFAULT 0,
    dep_deductible      REAL DEFAULT 0,
    dep_bureau          REAL DEFAULT 0,
    dep_formation       REAL DEFAULT 0,
    dep_assurance       REAL DEFAULT 0,
    dep_logiciel        REAL DEFAULT 0,
    dep_materiel        REAL DEFAULT 0,
    dep_deplacement     REAL DEFAULT 0,
    dep_honoraires      REAL DEFAULT 0,
    dep_marketing       REAL DEFAULT 0,
    dep_autres          REAL DEFAULT 0,

    benefice_net           REAL DEFAULT 0,
    cumul_annuel_revenus   REAL DEFAULT 0,
    pct_marge              REAL DEFAULT 0,

    nb_factures_emises     INTEGER DEFAULT 0,
    nb_factures_payees     INTEGER DEFAULT 0,
    nb_factures_en_retard  INTEGER DEFAULT 0,
    montant_a_percevoir    REAL DEFAULT 0,

    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(annee, mois) 
);
`

var IndexStatsMensuelles = `
CREATE INDEX IF NOT EXISTS idx_stats_annee   ON stats_mensuelles(annee);
CREATE INDEX IF NOT EXISTS idx_stats_periode ON stats_mensuelles(annee, mois);
`

// ─────────────────────────────────────────────────────────────────────────────
// TRIGGERS — mise à jour automatique des stats après chaque mutation
// Entièrement en SQL natif, aucun déchiffrement (tout est numérique).
// ─────────────────────────────────────────────────────────────────────────────
var TriggerRevenuInsert = `
CREATE TRIGGER IF NOT EXISTS trg_revenu_insert
AFTER INSERT ON revenus
BEGIN
    INSERT INTO stats_mensuelles (annee, mois)
    VALUES (CAST(strftime('%Y', NEW.date_service) AS INTEGER),
            CAST(strftime('%m', NEW.date_service) AS INTEGER))
    ON CONFLICT(annee, mois) DO NOTHING;

    UPDATE stats_mensuelles SET
        rev_total = (SELECT COALESCE(SUM(montant_total),0) FROM revenus
                     WHERE strftime('%Y-%m', date_service) = strftime('%Y-%m', NEW.date_service)
                       AND statut_paiement != 'annule'),
        nb_seances = (SELECT COUNT(*) FROM revenus
                      WHERE strftime('%Y-%m', date_service) = strftime('%Y-%m', NEW.date_service)
                        AND statut_paiement != 'annule'),
        nb_clients_uniques = (SELECT COUNT(DISTINCT client_id) FROM revenus
                              WHERE strftime('%Y-%m', date_service) = strftime('%Y-%m', NEW.date_service)
                                AND statut_paiement != 'annule'),
        benefice_net = rev_total - dep_total,
        cumul_annuel_revenus = (SELECT COALESCE(SUM(rev_total),0) FROM stats_mensuelles
                                WHERE annee = CAST(strftime('%Y', NEW.date_service) AS INTEGER)),
        updated_at = CURRENT_TIMESTAMP
    WHERE annee = CAST(strftime('%Y', NEW.date_service) AS INTEGER)
      AND mois  = CAST(strftime('%m', NEW.date_service) AS INTEGER);
END;
`

var TriggerRevenuUpdate = `
CREATE TRIGGER IF NOT EXISTS trg_revenu_update
AFTER UPDATE ON revenus
BEGIN
    UPDATE stats_mensuelles SET
        rev_total = (SELECT COALESCE(SUM(montant_total),0) FROM revenus
                     WHERE strftime('%Y-%m', date_service) = strftime('%Y-%m', NEW.date_service)
                       AND statut_paiement != 'annule'),
        nb_seances = (SELECT COUNT(*) FROM revenus
                      WHERE strftime('%Y-%m', date_service) = strftime('%Y-%m', NEW.date_service)
                        AND statut_paiement != 'annule'),
        benefice_net = rev_total - dep_total,
        updated_at = CURRENT_TIMESTAMP
    WHERE annee = CAST(strftime('%Y', NEW.date_service) AS INTEGER)
      AND mois  = CAST(strftime('%m', NEW.date_service) AS INTEGER);
END;
`

var TriggerDepenseInsert = `
CREATE TRIGGER IF NOT EXISTS trg_depense_insert
AFTER INSERT ON depenses
BEGIN
    INSERT INTO stats_mensuelles (annee, mois)
    VALUES (CAST(strftime('%Y', NEW.date_depense) AS INTEGER),
            CAST(strftime('%m', NEW.date_depense) AS INTEGER))
    ON CONFLICT(annee, mois) DO NOTHING;

    UPDATE stats_mensuelles SET
        dep_total = (SELECT COALESCE(SUM(montant_total),0) FROM depenses
                     WHERE strftime('%Y-%m', date_depense) = strftime('%Y-%m', NEW.date_depense)),
        dep_deductible = (SELECT COALESCE(SUM(montant_deductible),0) FROM depenses
                          WHERE strftime('%Y-%m', date_depense) = strftime('%Y-%m', NEW.date_depense)
                            AND deductible = 1),
        benefice_net = rev_total - dep_total,
        updated_at = CURRENT_TIMESTAMP
    WHERE annee = CAST(strftime('%Y', NEW.date_depense) AS INTEGER)
      AND mois  = CAST(strftime('%m', NEW.date_depense) AS INTEGER);
END;
`

var TriggerDepenseUpdate = `
CREATE TRIGGER IF NOT EXISTS trg_depense_update
AFTER UPDATE ON depenses
BEGIN
    UPDATE stats_mensuelles SET
        dep_total = (SELECT COALESCE(SUM(montant_total),0) FROM depenses
                     WHERE strftime('%Y-%m', date_depense) = strftime('%Y-%m', NEW.date_depense)),
        dep_deductible = (SELECT COALESCE(SUM(montant_deductible),0) FROM depenses
                          WHERE strftime('%Y-%m', date_depense) = strftime('%Y-%m', NEW.date_depense)
                            AND deductible = 1),
        benefice_net = rev_total - dep_total,
        updated_at = CURRENT_TIMESTAMP
    WHERE annee = CAST(strftime('%Y', NEW.date_depense) AS INTEGER)
      AND mois  = CAST(strftime('%m', NEW.date_depense) AS INTEGER);
END;
`

// ─────────────────────────────────────────────────────────────────────────────
// AllFinanceTables — ordre d'exécution dans db.go → InitDB()
// ─────────────────────────────────────────────────────────────────────────────
var AllFinanceTables = []string{
	TableParametresFinance,
	TableServices,
	IndexServices,
	TableRevenus,
	IndexRevenus,
	TableDepenses,
	IndexDepenses,
	TableFactures,
	IndexFactures,
	TableFactureLignes,
	TablePaiements,
	TableContrats,
	IndexContrats,
	TableStatsMensuelles,
	IndexStatsMensuelles,
	TriggerRevenuInsert,
	TriggerRevenuUpdate,
	TriggerDepenseInsert,
	TriggerDepenseUpdate,
}
