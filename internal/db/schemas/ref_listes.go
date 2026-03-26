package schema

// TableRefListes contient le schéma CREATE TABLE et toutes les données
// de départ (seed) pour la table universelle des listes de référence.
// À ajouter dans db.go : schema.TableRefListes dans le fullSchema.
var TableRefListes = `
CREATE TABLE IF NOT EXISTS ref_listes (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    categorie   TEXT    NOT NULL,
    libelle     TEXT    NOT NULL,
    couleur     TEXT    NOT NULL DEFAULT 'slate',
    icone       TEXT    NOT NULL DEFAULT '',
    is_systeme  INTEGER NOT NULL DEFAULT 0,
    ordre       INTEGER NOT NULL DEFAULT 0,
    actif       INTEGER NOT NULL DEFAULT 1,

    UNIQUE(categorie, libelle)   -- ← AJOUTE CETTE LIGNE
);

CREATE INDEX IF NOT EXISTS idx_ref_listes_categorie ON ref_listes(categorie, actif);

-- ============================================================
-- SEED : Données système de base (is_systeme = 1, non supprimables)
-- ============================================================

-- Types d'intervenants
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('type_intervenant', 'Médecin',              'blue',    'stethoscope',  1, 1),
('type_intervenant', 'Travailleur Social',   'emerald', 'heart-handshake', 1, 2),
('type_intervenant', 'Pharmacien',           'purple',  'pill',         1, 3),
('type_intervenant', 'Infirmier(ère)',        'pink',    'activity',     1, 4),
('type_intervenant', 'Notaire',              'amber',   'scale',        1, 5),
('type_intervenant', 'Psychologue',          'cyan',    'brain',        1, 6),
('type_intervenant', 'Ergothérapeute',       'lime',    'hand',         1, 7),
('type_intervenant', 'Physiothérapeute',     'orange',  'person-standing', 1, 8),
('type_intervenant', 'Autre',                'slate',   'user',         1, 9);

-- Statuts de dossier client
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('statut_dossier', 'Actif',        'emerald', 'circle-check',  1, 1),
('statut_dossier', 'En attente',   'amber',   'clock',         1, 2),
('statut_dossier', 'Suspendu',     'orange',  'pause-circle',  1, 3),
('statut_dossier', 'Fermé',        'slate',   'archive',       1, 4),
('statut_dossier', 'Décédé',       'zinc',    'cross',         1, 5),
('statut_dossier', 'Référé',       'blue',    'send',          1, 6);

-- Types de notes
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('type_note', 'Note de suivi',       'blue',   'file-text',    1, 1),
('type_note', 'Note téléphonique',   'teal',   'phone',        1, 2),
('type_note', 'Note de rencontre',   'emerald','users',        1, 3),
('type_note', 'Note administrative', 'slate',  'clipboard',    1, 4),
('type_note', 'Alerte',              'red',    'alert-triangle', 1, 5),
('type_note', 'Observation',         'purple', 'eye',          1, 6);

-- Types de contact (proches du client)
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('type_contact', 'Conjoint(e)',       'pink',    'heart',       1, 1),
('type_contact', 'Enfant',            'blue',    'baby',        1, 2),
('type_contact', 'Parent',            'amber',   'user',        1, 3),
('type_contact', 'Fratrie',           'orange',  'users',       1, 4),
('type_contact', 'Tuteur légal',      'purple',  'scale',       1, 5),
('type_contact', 'Ami(e)',            'teal',    'smile',       1, 6),
('type_contact', 'Voisin(e)',         'lime',    'home',        1, 7),
('type_contact', 'Autre',             'slate',   'user-circle', 1, 8);

-- Types d'appel
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('type_appel', 'Entrant',    'emerald', 'phone-incoming', 1, 1),
('type_appel', 'Sortant',    'blue',    'phone-outgoing', 1, 2),
('type_appel', 'Manqué',     'red',     'phone-missed',   1, 3),
('type_appel', 'Message',    'amber',   'message-square', 1, 4);

-- Langues
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('langue', 'Français',    'blue',   'languages', 1, 1),
('langue', 'Anglais',     'red',    'languages', 1, 2),
('langue', 'Espagnol',    'amber',  'languages', 1, 3),
('langue', 'Arabe',       'emerald','languages', 1, 4),
('langue', 'Créole',      'orange', 'languages', 1, 5),
('langue', 'Mandarin',    'red',    'languages', 1, 6),
('langue', 'Autre',       'slate',  'languages', 1, 7);

-- Sexe / Genre
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('sexe', 'Masculin',       'blue',   'user', 1, 1),
('sexe', 'Féminin',        'pink',   'user', 1, 2),
('sexe', 'Non-binaire',    'purple', 'user', 1, 3),
('sexe', 'Autre',          'slate',  'user', 1, 4),
('sexe', 'Préfère ne pas préciser', 'zinc', 'user', 1, 5);

-- Statut civil
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('statut_civil', 'Célibataire',       'slate',   'user',       1, 1),
('statut_civil', 'Marié(e)',          'emerald', 'heart',      1, 2),
('statut_civil', 'Conjoint de fait',  'teal',    'heart',      1, 3),
('statut_civil', 'Séparé(e)',         'amber',   'user',       1, 4),
('statut_civil', 'Divorcé(e)',        'orange',  'user',       1, 5),
('statut_civil', 'Veuf(ve)',          'zinc',    'user',       1, 6);

-- Liens familiaux
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('lien_familial', 'Conjoint(e)',    'pink',   'heart',  1, 1),
('lien_familial', 'Fils',           'blue',   'user',   1, 2),
('lien_familial', 'Fille',          'pink',   'user',   1, 3),
('lien_familial', 'Père',           'slate',  'user',   1, 4),
('lien_familial', 'Mère',           'slate',  'user',   1, 5),
('lien_familial', 'Frère',          'blue',   'user',   1, 6),
('lien_familial', 'Sœur',           'pink',   'user',   1, 7),
('lien_familial', 'Petit-fils',     'cyan',   'user',   1, 8),
('lien_familial', 'Petite-fille',   'cyan',   'user',   1, 9),
('lien_familial', 'Autre',          'slate',  'user',   1, 10);

-- Types de relation (grande catégorie)
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('type_relation', 'Famille',        'pink',    'heart',       1, 1),
('type_relation', 'Légal',          'purple',  'scale',       1, 2),
('type_relation', 'Professionnel',  'blue',    'briefcase',   1, 3),
('type_relation', 'Social',         'teal',    'users',       1, 4),
('type_relation', 'Institutionnel', 'amber',   'building',    1, 5),
('type_relation', 'Autre',          'slate',   'user-circle', 1, 6);

-- Civilités (universelle - médecins, notaires, contacts, etc.)
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('civilite', 'M.',          'blue',   'user', 1, 1),
('civilite', 'Mme',         'pink',   'user', 1, 2),
('civilite', 'Dr',          'teal',   'stethoscope', 1, 3),
('civilite', 'Dre',         'teal',   'stethoscope', 1, 4),
('civilite', 'Docteur',     'teal',   'stethoscope', 1, 5),
('civilite', 'Docteure',    'teal',   'stethoscope', 1, 6),
('civilite', 'Professeur',  'purple', 'graduation-cap', 1, 7),
('civilite', 'Professeure', 'purple', 'graduation-cap', 1, 8),
('civilite', 'Me',          'amber',  'scale', 1, 9),
('civilite', 'Révérend',    'slate',  'user', 1, 10),
('civilite', 'Autre',       'slate',  'user', 1, 11);

-- Spécialités médicales
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('specialite_medecin', 'Médecin de famille',  'teal',   'stethoscope', 1, 1),
('specialite_medecin', 'Médecine générale',   'teal',   'stethoscope', 1, 2),
('specialite_medecin', 'Médecine interne',    'blue',   'stethoscope', 1, 3),
('specialite_medecin', 'Cardiologie',         'red',    'heart',       1, 4),
('specialite_medecin', 'Neurologie',          'purple', 'brain',       1, 5),
('specialite_medecin', 'Psychiatrie',         'violet', 'brain',       1, 6),
('specialite_medecin', 'Pédiatrie',           'pink',   'user',        1, 7),
('specialite_medecin', 'Gériatrie',           'amber',  'user',        1, 8),
('specialite_medecin', 'Obstétrique',         'pink',   'user',        1, 9),
('specialite_medecin', 'Gynécologie',         'pink',   'user',        1, 10),
('specialite_medecin', 'Chirurgie générale',  'slate',  'user',        1, 11),
('specialite_medecin', 'Orthopédie',          'orange', 'user',        1, 12),
('specialite_medecin', 'Dermatologie',        'yellow', 'user',        1, 13),
('specialite_medecin', 'Ophtalmologie',       'cyan',   'eye',         1, 14),
('specialite_medecin', 'ORL',                 'lime',   'user',        1, 15),
('specialite_medecin', 'Radiologie',          'zinc',   'user',        1, 16),
('specialite_medecin', 'Anesthésiologie',     'slate',  'user',        1, 17),
('specialite_medecin', 'Urgence',             'red',    'alert',       1, 18),
('specialite_medecin', 'Autre',               'slate',  'user',        1, 19);

-- Statuts médecins
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('statut_medecin', 'Actif',     'emerald', 'check',  1, 1),
('statut_medecin', 'Retraité',  'slate',   'user',   1, 2),
('statut_medecin', 'En congé',  'amber',   'clock',  1, 3),
('statut_medecin', 'Autre',     'zinc',    'user',   1, 4);

-- ============================================================
-- Fréquences de suivi médical
-- ============================================================
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('frequence_suivi', 'Au besoin',         'slate',   'calendar',        1, 1),
('frequence_suivi', 'Hebdomadaire',      'blue',    'calendar-days',   1, 2),
('frequence_suivi', 'Aux 2 semaines',    'cyan',    'calendar-days',   1, 3),
('frequence_suivi', 'Mensuel',           'teal',    'calendar',        1, 4),
('frequence_suivi', 'Aux 2 mois',        'emerald', 'calendar',        1, 5),
('frequence_suivi', 'Trimestriel',       'green',   'calendar',        1, 6),
('frequence_suivi', 'Aux 6 mois',        'amber',   'calendar',        1, 7),
('frequence_suivi', 'Annuel',            'orange',  'calendar',        1, 8),
('frequence_suivi', 'Ponctuel',          'zinc',    'calendar-x',      1, 9),
('frequence_suivi', 'Autre',             'slate',   'calendar',        1, 10);
 
-- ============================================================
-- Récurrences (nature du suivi dans le temps)
-- ============================================================
INSERT OR IGNORE INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre) VALUES
('recurrence_suivi', 'Suivi actif continu',   'emerald', 'repeat',          1, 1),
('recurrence_suivi', 'Suivi épisodique',      'blue',    'repeat-1',        1, 2),
('recurrence_suivi', 'Suivi post-crise',      'amber',   'activity',        1, 3),
('recurrence_suivi', 'Suivi préventif',       'teal',    'shield-check',    1, 4),
('recurrence_suivi', 'Suivi post-opératoire', 'purple',  'heart-pulse',     1, 5),
('recurrence_suivi', 'Suivi palliatif',       'zinc',    'heart-handshake', 1, 6),
('recurrence_suivi', 'Consultation unique',   'slate',   'calendar-check',  1, 7),
('recurrence_suivi', 'Suivi hospitalier',     'red',     'building-2',      1, 8),
('recurrence_suivi', 'Autre',                 'slate',   'repeat',          1, 9);



`
