package schema

// 1. La table des définitions (Le "Cerveau")
var TableEvalDefinitions = `
CREATE TABLE IF NOT EXISTS evaluation_definitions (
    id TEXT PRIMARY KEY,        
    nom TEXT NOT NULL,          
    icone TEXT DEFAULT 'FileText', 
    couleur TEXT DEFAULT 'teal',   
    schema_json TEXT NOT NULL,  
    version INTEGER DEFAULT 1,
    active INTEGER DEFAULT 1
);`

// 2. La table des résultats (Le "Coffre-fort")
var TableEvaluations = `
CREATE TABLE IF NOT EXISTS evaluations_sociales_v2 (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    model_id TEXT NOT NULL,     
    no_leopard TEXT NOT NULL,
    payload_crypte TEXT NOT NULL, 
    statut TEXT DEFAULT 'brouillon', 
    signature_nom TEXT,
    date_signature DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(client_id) REFERENCES clients(id),
    FOREIGN KEY(model_id) REFERENCES evaluation_definitions(id)
);`

var EvalDefinitionsSeed = `
 
INSERT OR IGNORE INTO evaluation_definitions (id, nom, icone, couleur, schema_json, version, active) VALUES (
  'psychosociale_initiale',
  'Évaluation Psychosociale Initiale',
  'ClipboardList',
  'teal',
  '{
    "sections": [
      {
        "id": "contexte",
        "label": "Contexte et Référence",
        "icone": "FileText",
        "couleur": "blue",
        "fields": [
          {
            "id": "source_reference",
            "type": "select",
            "label": "Source de référence",
            "required": true,
            "options": [
              "Auto-référence",
              "Médecin traitant",
              "CLSC / CISSS",
              "Centre hospitalier",
              "Tribunal",
              "Curateur public",
              "Famille / Proche",
              "Autre professionnel"
            ]
          },
          {
            "id": "date_reference",
            "type": "date",
            "label": "Date de la référence",
            "required": true
          },
          {
            "id": "date_evaluation",
            "type": "date",
            "label": "Date(s) de l''évaluation",
            "required": true
          },
          {
            "id": "lieu_evaluation",
            "type": "text",
            "label": "Lieu(x) de l''évaluation",
            "placeholder": "Domicile, bureau, CHSLD, hôpital..."
          },
          {
            "id": "motif_reference",
            "type": "textarea",
            "label": "Motif de référence",
            "required": true,
            "rows": 4,
            "placeholder": "Raison principale de la demande d''évaluation..."
          },
          {
            "id": "objet_evaluation",
            "type": "textarea",
            "label": "Objet de l''évaluation",
            "required": true,
            "rows": 4,
            "placeholder": "Ce qui doit spécifiquement être évalué ou déterminé..."
          },
          {
            "id": "contexte_evaluation",
            "type": "textarea",
            "label": "Contexte de l''évaluation",
            "rows": 5,
            "placeholder": "Circonstances, historique de la demande, informations pertinentes..."
          }
        ]
      },
      {
        "id": "personnes_consultees",
        "label": "Personnes Consultées",
        "icone": "Users",
        "couleur": "indigo",
        "fields": [
          {
            "id": "personnes_consultees",
            "type": "table",
            "label": "Liste des personnes consultées",
            "help": "Toutes les personnes rencontrées ou contactées dans le cadre de l''évaluation",
            "columns": [
              { "id": "nom", "label": "Nom", "type": "text" },
              { "id": "lien", "label": "Lien avec le client", "type": "text" },
              { "id": "date_contact", "label": "Date", "type": "date" },
              { "id": "mode", "label": "Mode (entrevue/tél.)", "type": "text" },
              { "id": "notes", "label": "Notes", "type": "text" }
            ]
          },
          {
            "id": "documents_consultes",
            "type": "textarea",
            "label": "Documents et dossiers consultés",
            "rows": 4,
            "placeholder": "Dossiers médicaux, rapports antérieurs, évaluations, jugements..."
          }
        ]
      },
      {
        "id": "situation_sociale",
        "label": "Situation Sociale & Réseau",
        "icone": "Network",
        "couleur": "emerald",
        "fields": [
          {
            "id": "situation_familiale",
            "type": "select",
            "label": "Situation familiale",
            "options": [
              "Célibataire",
              "Marié(e) / Union civile",
              "Conjoint(e) de fait",
              "Séparé(e) / Divorcé(e)",
              "Veuf(ve)"
            ]
          },
          {
            "id": "milieu_vie",
            "type": "select",
            "label": "Milieu de vie actuel",
            "options": [
              "Domicile privé seul",
              "Domicile privé avec conjoint",
              "Domicile de la famille",
              "Résidence privée pour aînés (RPA)",
              "CHSLD",
              "Ressource intermédiaire (RI)",
              "Ressource de type familial (RTF)",
              "Autre"
            ]
          },
          {
            "id": "description_milieu",
            "type": "textarea",
            "label": "Description du milieu de vie",
            "rows": 4,
            "placeholder": "Conditions de logement, accessibilité, sécurité, adaptation..."
          },
          {
            "id": "reseau_social",
            "type": "textarea",
            "label": "Réseau social et soutien",
            "rows": 5,
            "placeholder": "Personnes significatives, qualité des relations, disponibilité du soutien naturel..."
          },
          {
            "id": "services_en_place",
            "type": "textarea",
            "label": "Services professionnels déjà en place",
            "rows": 4,
            "placeholder": "SAD, soins infirmiers, médecin de famille, autres intervenants..."
          },
          {
            "id": "situation_financiere",
            "type": "textarea",
            "label": "Situation financière",
            "rows": 4,
            "placeholder": "Sources de revenus, gestion financière, dettes, capacité à gérer les finances..."
          }
        ]
      },
      {
        "id": "sante_physique",
        "label": "Santé Physique & Autonomie",
        "icone": "Heart",
        "couleur": "red",
        "fields": [
          {
            "id": "diagnostics_medicaux",
            "type": "textarea",
            "label": "Diagnostics médicaux connus",
            "rows": 4,
            "placeholder": "Maladies chroniques, conditions médicales, antécédents pertinents..."
          },
          {
            "id": "medication",
            "type": "textarea",
            "label": "Médication",
            "rows": 3,
            "placeholder": "Médication actuelle, observance, gestion de la médication..."
          },
          {
            "id": "autonomie_avq",
            "type": "textarea",
            "label": "Autonomie — Activités de la vie quotidienne (AVQ)",
            "rows": 4,
            "placeholder": "Hygiène personnelle, habillage, alimentation, déplacements intérieurs, continence..."
          },
          {
            "id": "autonomie_avd",
            "type": "textarea",
            "label": "Autonomie — Activités de la vie domestique (AVD)",
            "rows": 4,
            "placeholder": "Préparation des repas, entretien ménager, gestion finances, transport, courses..."
          },
          {
            "id": "etat_sante_physique",
            "type": "textarea",
            "label": "Observations générales sur l''état physique",
            "rows": 4,
            "placeholder": "Observations cliniques: mobilité, présentation physique, nutrition, sommeil..."
          }
        ]
      },
      {
        "id": "sante_cognitive",
        "label": "Santé Cognitive & Psychologique",
        "icone": "Brain",
        "couleur": "purple",
        "fields": [
          {
            "id": "orientation",
            "type": "checkboxes",
            "label": "Orientation",
            "options": [
              "Orienté(e) dans le temps",
              "Orienté(e) dans l''espace",
              "Orienté(e) quant aux personnes",
              "Orienté(e) quant à sa propre personne"
            ]
          },
          {
            "id": "capacites_cognitives",
            "type": "textarea",
            "label": "Observations des capacités cognitives",
            "rows": 6,
            "placeholder": "Mémoire (court/long terme), attention, concentration, jugement, raisonnement, compréhension, langage, praxies..."
          },
          {
            "id": "conscience_limitations",
            "type": "select",
            "label": "Conscience de ses limitations",
            "options": [
              "Bonne conscience de ses limitations",
              "Conscience partielle",
              "Peu de conscience (anosognosie partielle)",
              "Aucune conscience (anosognosie)"
            ]
          },
          {
            "id": "etat_psychologique",
            "type": "textarea",
            "label": "État psychologique",
            "rows": 5,
            "placeholder": "Humeur, anxiété, signes dépressifs, comportements, mécanismes d''adaptation, personnalité..."
          },
          {
            "id": "outils_utilises",
            "type": "textarea",
            "label": "Outils d''évaluation utilisés",
            "rows": 3,
            "placeholder": "MoCA, MMSE, évaluations ergothérapiques, rapports neuropsychologiques consultés..."
          }
        ]
      },
      {
        "id": "roles_sociaux",
        "label": "Rôles Sociaux & Habitudes de Vie",
        "icone": "Briefcase",
        "couleur": "amber",
        "fields": [
          {
            "id": "roles_familiaux",
            "type": "textarea",
            "label": "Rôles familiaux",
            "rows": 4,
            "placeholder": "Rôle de parent, conjoint, grand-parent — exercice de ces rôles, changements observés..."
          },
          {
            "id": "loisirs_activites",
            "type": "textarea",
            "label": "Loisirs, activités, participation sociale",
            "rows": 4,
            "placeholder": "Intérêts, activités de loisirs, participation communautaire, bénévolat, sorties..."
          },
          {
            "id": "historique_professionnel",
            "type": "textarea",
            "label": "Historique professionnel pertinent",
            "rows": 3,
            "placeholder": "Occupation passée, retraite, signification du travail dans l''identité..."
          }
        ]
      },
      {
        "id": "analyse_clinique",
        "label": "Analyse Clinique",
        "icone": "Layers",
        "couleur": "slate",
        "fields": [
          {
            "id": "facteurs_risque",
            "type": "textarea",
            "label": "Facteurs de risque identifiés",
            "rows": 5,
            "placeholder": "Facteurs de vulnérabilité, risques pour la sécurité, pour les droits, pour le bien-être..."
          },
          {
            "id": "facteurs_protection",
            "type": "textarea",
            "label": "Facteurs de protection",
            "rows": 5,
            "placeholder": "Forces de la personne, ressources disponibles, réseau de soutien, services en place..."
          },
          {
            "id": "analyse_clinique",
            "type": "textarea",
            "label": "Analyse clinique intégrée",
            "rows": 8,
            "placeholder": "Synthèse intégrée des dimensions évaluées, dynamiques, interactions entre les facteurs, diagnostic fonctionnel social..."
          }
        ]
      },
      {
        "id": "opinion_recommandations",
        "label": "Opinion & Recommandations",
        "icone": "CheckCircle",
        "couleur": "teal",
        "fields": [
          {
            "id": "opinion_professionnelle",
            "type": "textarea",
            "label": "Opinion professionnelle du T.S.",
            "required": true,
            "rows": 7,
            "help": "L''opinion professionnelle doit être clairement formulée, fondée sur les observations et l''analyse, et distinguée des faits rapportés (OTSTCFQ)",
            "placeholder": "En ma qualité de travailleur social, mon opinion professionnelle est que..."
          },
          {
            "id": "recommandations",
            "type": "textarea",
            "label": "Recommandations",
            "required": true,
            "rows": 7,
            "placeholder": "Sur la base de mon évaluation, je recommande: (services, ressources, régime de protection, suivi, démarches)..."
          },
          {
            "id": "plan_action",
            "type": "textarea",
            "label": "Plan d''action et prochaines étapes",
            "rows": 5,
            "placeholder": "Objectifs prioritaires, démarches à entreprendre, échéancier, prochain contact prévu..."
          }
        ]
      },
      {
        "id": "signature",
        "label": "Signature & Verrouillage",
        "icone": "ShieldCheck",
        "couleur": "emerald",
        "fields": [
          {
            "id": "info_signature",
            "type": "info",
            "contenu": "En scellant cette évaluation, vous attestez que son contenu reflète fidèlement votre évaluation professionnelle. Le document sera verrouillé et ne pourra plus être modifié. Cette action est conforme aux obligations déontologiques de l''OTSTCFQ."
          },
          {
            "id": "nom_signataire",
            "type": "text",
            "label": "Nom du signataire",
            "required": true,
            "placeholder": "Prénom Nom, T.S."
          },
          {
            "id": "no_membre_otstcfq",
            "type": "text",
            "label": "N° de membre OTSTCFQ",
            "placeholder": "Ex: TS-12345"
          },
          {
            "id": "date_rapport",
            "type": "date",
            "label": "Date du rapport"
          }
        ]
      }
    ]
  }',
  1,
  1
);
`

var EvalFonctionnementSocialSeed = `

INSERT OR IGNORE INTO evaluation_definitions (id, nom, icone, couleur, schema_json, version, active) VALUES (
  'fonctionnement_social',
  'Évaluation du Fonctionnement Social (OTSTCFQ)',
  'Users',
  'indigo',
  '{
    "sections": [
      {
        "id": "partie1_demande_contexte",
        "label": "Partie 1 — Demande et contexte de l''évaluation",
        "icone": "FileSearch",
        "couleur": "blue",
        "fields": [
          {
            "id": "info_p1",
            "type": "info",
            "contenu": "Cette partie situe le contexte dans lequel s''inscrit l''évaluation. Elle décrit les raisons et les circonstances pour lesquelles une évaluation du fonctionnement social est demandée ou nécessaire (OTSTCFQ, 2010)."
          },
          {
            "id": "identification_sociodemo",
            "type": "textarea",
            "label": "Identification sociodémographique du client",
            "rows": 5,
            "help": "Sexe, âge, scolarité, état civil, origine ethnique, statut d''immigration, pratique religieuse/spirituelle, langue parlée, occupation, sources de revenus, contexte familial, responsabilités parentales et sociales.",
            "placeholder": "Ex: Femme de 68 ans, francophone, née au Québec, retraitée (enseignante), veuve depuis 2019, catholique pratiquante, revenus: RRQ + régime de retraite. Deux enfants adultes, quatre petits-enfants..."
          },
          {
            "id": "source_demande",
            "type": "select",
            "label": "Origine de la demande de services",
            "required": true,
            "options": [
              "La personne elle-même (auto-référence)",
              "Un proche (famille, ami, voisin)",
              "Un professionnel du réseau de la santé",
              "Un établissement ou organisme",
              "Un tribunal / instance légale",
              "Le Curateur public",
              "Autre"
            ]
          },
          {
            "id": "contexte_trajectoire_demande",
            "type": "textarea",
            "label": "Contexte et trajectoire de la demande",
            "required": true,
            "rows": 5,
            "placeholder": "Décrire qui formule la demande, pourquoi à ce moment-ci, historique de la demande, services antérieurs sollicités ou reçus..."
          },
          {
            "id": "mandat_travailleur_social",
            "type": "textarea",
            "label": "Mandat confié au travailleur social",
            "required": true,
            "rows": 4,
            "placeholder": "Préciser le mandat explicite: ce qui est attendu de l''évaluation et à qui le rapport est destiné (juge, équipe, mandataire, client lui-même...)..."
          },
          {
            "id": "sources_informations",
            "type": "table",
            "label": "Sources d''informations consultées",
            "help": "Énumérer toutes les sources: personnes rencontrées, dates, documents consultés, autorisations obtenues.",
            "columns": [
              { "id": "source", "label": "Source (personne / document)", "type": "text" },
              { "id": "lien", "label": "Lien avec le client", "type": "text" },
              { "id": "date", "label": "Date", "type": "date" },
              { "id": "mode", "label": "Mode (entrevue / tél. / document)", "type": "text" },
              { "id": "autorisation", "label": "Autorisation obtenue", "type": "text" }
            ]
          },
          {
            "id": "situation_actuelle",
            "type": "textarea",
            "label": "Situation actuelle — Description des problèmes",
            "required": true,
            "rows": 6,
            "help": "Description selon la perspective de la personne, de ses proches et du réseau. Inclure la durée, l''évolution et le niveau de gravité des problèmes. Distinguer les faits des présuppositions.",
            "placeholder": "Décrire les difficultés identifiées ou rapportées, leur durée, leur évolution, leur gravité, selon la perspective des personnes impliquées..."
          },
          {
            "id": "besoins_personne",
            "type": "textarea",
            "label": "Besoins de la personne",
            "required": true,
            "rows": 5,
            "help": "Besoins psychosociaux, de base, de bien-être, d''intégration sociale et de participation citoyenne. Distinguer les besoins perçus par la personne et ceux observés par le T.S.",
            "placeholder": "Besoins identifiés par la personne elle-même, par ses proches, et par l''intervenant. Besoins de base, de sécurité, d''appartenance, de participation sociale..."
          },
          {
            "id": "attentes_personne",
            "type": "textarea",
            "label": "Attentes et objectifs de la personne",
            "rows": 4,
            "placeholder": "Ce que la personne espère de l''évaluation et de l''intervention. Niveau de motivation, obstacles à l''engagement..."
          }
        ]
      },
      {
        "id": "partie2_caracteristiques_personne",
        "label": "Partie 2 — Caractéristiques de la personne",
        "icone": "User",
        "couleur": "violet",
        "fields": [
          {
            "id": "info_p2",
            "type": "info",
            "contenu": "Cette partie décrit les dimensions biologique, intellectuelle, émotionnelle, sociale, familiale, spirituelle, économique et communautaire de la personne, qui s''interinfluencent (Sheafor et Horejsi, 2006). Elle met en valeur les forces de la personne."
          },
          {
            "id": "dimension_physique_biologique",
            "type": "textarea",
            "label": "Dimension physique et biologique",
            "rows": 5,
            "placeholder": "État de santé physique général, diagnostics, médication, capacités fonctionnelles, AVQ/AVD, mobilité, nutrition, sommeil, historique médical pertinent..."
          },
          {
            "id": "dimension_cognitive_intellectuelle",
            "type": "textarea",
            "label": "Dimension cognitive et intellectuelle",
            "rows": 5,
            "placeholder": "Capacités de compréhension, de jugement, de raisonnement, de mémoire (court et long terme), d''attention. Scolarité, habiletés d''apprentissage, outils utilisés (MoCA, MMSE, etc.)..."
          },
          {
            "id": "dimension_emotionnelle_psychologique",
            "type": "textarea",
            "label": "Dimension émotionnelle et psychologique",
            "rows": 5,
            "placeholder": "Humeur, affects, anxiété, signes dépressifs, estime de soi, mécanismes d''adaptation, résilience, vécu des pertes, réactions au stress, personnalité..."
          },
          {
            "id": "dimension_sociale_relationnelle",
            "type": "textarea",
            "label": "Dimension sociale et relationnelle",
            "rows": 5,
            "placeholder": "Habiletés relationnelles, qualité des liens interpersonnels, capacité à demander de l''aide, ouverture aux services, mode de communication..."
          },
          {
            "id": "dimension_familiale",
            "type": "textarea",
            "label": "Dimension familiale",
            "rows": 5,
            "placeholder": "Dynamiques familiales, histoire familiale pertinente, relations avec les membres de la famille, rôles familiaux assumés (parent, conjoint, enfant, grand-parent...), impacts des changements familiaux..."
          },
          {
            "id": "dimension_spirituelle_culturelle",
            "type": "textarea",
            "label": "Dimension spirituelle, culturelle et identitaire",
            "rows": 4,
            "placeholder": "Valeurs, croyances, pratiques religieuses ou spirituelles, appartenance culturelle, langue, identité, sens donné à la vie, ressources spirituelles..."
          },
          {
            "id": "dimension_economique",
            "type": "textarea",
            "label": "Dimension économique et financière",
            "rows": 4,
            "placeholder": "Sources et niveau de revenus, gestion financière, dettes, capacité à subvenir à ses besoins, impacts de la situation financière sur le fonctionnement social..."
          },
          {
            "id": "roles_sociaux",
            "type": "checkboxes",
            "label": "Rôles sociaux assumés",
            "help": "Cocher les rôles exercés. L''évaluation porte sur la qualité de l''exercice de ces rôles et les changements observés.",
            "options": [
              "Rôles familiaux (parent, conjoint, enfant, frère/sœur, grand-parent)",
              "Rôles interpersonnels (ami, voisin, membre d''un groupe)",
              "Rôles reliés à l''occupation (travailleur, étudiant, bénévole)",
              "Rôles associés à des situations particulières (usager, immigrant, personne hébergée...)"
            ]
          },
          {
            "id": "exercice_roles_sociaux",
            "type": "textarea",
            "label": "Exercice des rôles sociaux — observations",
            "rows": 5,
            "placeholder": "Décrire comment la personne assume ses différents rôles sociaux, les forces observées, les difficultés, les changements survenus, les impacts sur son bien-être..."
          },
          {
            "id": "avq_avd",
            "type": "textarea",
            "label": "Activités de la vie quotidienne (AVQ) et domestique (AVD)",
            "rows": 5,
            "placeholder": "AVQ: hygiène, habillage, alimentation, mobilité, continence. AVD: préparation des repas, ménage, gestion finances, transport, courses, prise de médication..."
          },
          {
            "id": "forces_personne",
            "type": "textarea",
            "label": "Forces, ressources et capacités de la personne",
            "required": true,
            "rows": 5,
            "help": "Section essentielle: identifier les forces sur lesquelles l''intervention pourra s''appuyer.",
            "placeholder": "Forces personnelles, habiletés, expériences passées de résolution de problèmes, capacité d''adaptation, motivation, intérêts, ressources personnelles..."
          },
          {
            "id": "strategies_anterieures",
            "type": "textarea",
            "label": "Stratégies antérieures de résolution de problèmes",
            "rows": 4,
            "placeholder": "Ce que la personne et ses proches ont tenté dans le passé pour résoudre la situation problématique. Résultats obtenus. Ce qui a fonctionné ou non..."
          }
        ]
      },
      {
        "id": "partie3_caracteristiques_environnement",
        "label": "Partie 3 — Caractéristiques de l''environnement",
        "icone": "Network",
        "couleur": "emerald",
        "fields": [
          {
            "id": "info_p3",
            "type": "info",
            "contenu": "Cette partie décrit l''environnement immédiat (réseau de la personne, conditions de vie, ressources formelles) et l''environnement sociétal (valeurs, normes, politiques sociales, opportunités et obstacles structurels). C''est la dimension distinctive de l''évaluation du T.S."
          },
          {
            "id": "milieu_vie",
            "type": "textarea",
            "label": "Milieu de vie — conditions matérielles et physiques",
            "rows": 5,
            "placeholder": "Type de logement, conditions physiques (sécurité, accessibilité, salubrité, adaptation), quartier, voisinage, transport disponible, conditions de vie générales..."
          },
          {
            "id": "reseau_immédiat_famille",
            "type": "textarea",
            "label": "Réseau immédiat — famille et proches",
            "rows": 5,
            "placeholder": "Composition du réseau familial, qualité des liens réciproques, disponibilité et nature du soutien (émotionnel, instrumental, financier), présence de conflits, limites du réseau..."
          },
          {
            "id": "reseau_social_elargi",
            "type": "textarea",
            "label": "Réseau social élargi",
            "rows": 4,
            "placeholder": "Amis, voisins, collègues, bénévolat, communauté religieuse, groupes d''appartenance, loisirs et participation sociale, isolement ou intégration..."
          },
          {
            "id": "ressources_formelles",
            "type": "table",
            "label": "Ressources formelles en place ou mobilisables",
            "help": "Services professionnels, organismes, ressources communautaires impliqués ou disponibles.",
            "columns": [
              { "id": "ressource", "label": "Ressource / Service", "type": "text" },
              { "id": "type", "label": "Type (CISSS, communautaire, privé...)", "type": "text" },
              { "id": "statut", "label": "En place / À mobiliser", "type": "text" },
              { "id": "frequence", "label": "Fréquence / Intensité", "type": "text" },
              { "id": "notes", "label": "Notes", "type": "text" }
            ]
          },
          {
            "id": "environnement_societal",
            "type": "textarea",
            "label": "Environnement sociétal — valeurs, normes et politiques",
            "rows": 5,
            "help": "Dimension distinctive du T.S.: identifier les politiques sociales, mesures de protection, valeurs culturelles et normes sociales qui influencent la situation.",
            "placeholder": "Politiques sociales applicables, mesures de protection disponibles, normes culturelles et sociales en jeu, appartenance communautaire, valeurs collectives qui influencent la situation..."
          },
          {
            "id": "facteurs_oppression_discrimination",
            "type": "textarea",
            "label": "Facteurs d''oppression, discrimination et injustice sociale",
            "rows": 5,
            "help": "Identifier les situations passées ou actuelles d''oppression, discrimination, exclusion, injustice sociale, iniquité économique ou stigmatisation qui influencent le fonctionnement social.",
            "placeholder": "Discrimination vécue (âge, genre, origine, handicap...), exclusion sociale, stigmatisation, barrières systémiques, iniquités économiques, expériences d''oppression passées ou actuelles..."
          },
          {
            "id": "opportunites_obstacles_environnement",
            "type": "textarea",
            "label": "Opportunités et obstacles dans l''environnement",
            "rows": 4,
            "placeholder": "Ressources accessibles vs inaccessibles, opportunités de participation sociale, obstacles structurels ou géographiques, disponibilité et accessibilité des services..."
          }
        ]
      },
      {
        "id": "partie4_analyse_synthese",
        "label": "Partie 4 — Analyse et synthèse",
        "icone": "Layers",
        "couleur": "amber",
        "fields": [
          {
            "id": "info_p4",
            "type": "info",
            "contenu": "Le cœur clinique du rapport. Le T.S. met en interrelation les caractéristiques de la personne et celles de son environnement pour comprendre le fonctionnement social dans ses aspects dynamiques et multicausals. Il distingue les faits des présuppositions et formule des hypothèses cliniques."
          },
          {
            "id": "comprehension_situation_problematique",
            "type": "textarea",
            "label": "Compréhension de la situation problématique",
            "required": true,
            "rows": 7,
            "placeholder": "Expliquer les problèmes rencontrés: leur gravité, leur durée, leurs conséquences sociales et leurs effets sur la personne et ses proches. Identifier les facteurs qui ont contribué à l''émergence et au maintien de ces problèmes..."
          },
          {
            "id": "interactions_personne_environnement",
            "type": "textarea",
            "label": "Interactions personne ↔ environnement",
            "required": true,
            "rows": 7,
            "help": "Dimension centrale et distinctive de l''EFS. Analyser comment les caractéristiques de la personne et celles de son environnement s''interinfluencent pour produire le fonctionnement social observé.",
            "placeholder": "Comment les caractéristiques de la personne (forces, limites, besoins) interagissent-elles avec les ressources, obstacles et opportunités de son environnement immédiat et sociétal? Quels sont les impacts de ces interactions sur son fonctionnement social?"
          },
          {
            "id": "hypotheses_cliniques",
            "type": "textarea",
            "label": "Hypothèses cliniques",
            "rows": 6,
            "help": "Formulations théoriques et analytiques appuyées sur des assises scientifiques. Distinguer clairement les faits observés des hypothèses interprétatives.",
            "placeholder": "Sur la base des informations recueillies et de mes observations, j''émets les hypothèses suivantes quant aux facteurs explicatifs de la situation: ..."
          },
          {
            "id": "facteurs_risque_vulnerabilite",
            "type": "textarea",
            "label": "Facteurs de risque et de vulnérabilité",
            "rows": 5,
            "placeholder": "Facteurs (personnels et environnementaux) qui exacerbent les problèmes, augmentent la vulnérabilité, représentent des risques pour la sécurité, les droits ou le bien-être de la personne ou de ses proches..."
          },
          {
            "id": "facteurs_protection_forces",
            "type": "textarea",
            "label": "Facteurs de protection et forces mobilisables",
            "required": true,
            "rows": 5,
            "placeholder": "Forces personnelles, ressources du réseau, ressources de l''environnement qui constituent des leviers pour l''intervention et pourraient atténuer les effets des problèmes..."
          },
          {
            "id": "impact_fonctionnement_social",
            "type": "textarea",
            "label": "Impact sur le fonctionnement social",
            "required": true,
            "rows": 6,
            "help": "Évaluer l''impact global sur la capacité de la personne à assurer son bien-être, réaliser ses activités de la vie quotidienne et exercer ses rôles sociaux.",
            "placeholder": "Comment les problèmes identifiés affectent-ils concrètement le fonctionnement social de la personne? Impacts sur ses AVQ/AVD, ses rôles sociaux, sa participation citoyenne, sa qualité de vie, ses relations..."
          }
        ]
      },
      {
        "id": "partie5_opinion_recommandations",
        "label": "Partie 5 — Opinion et recommandations",
        "icone": "CheckCircle",
        "couleur": "teal",
        "fields": [
          {
            "id": "info_p5",
            "type": "info",
            "contenu": "L''opinion professionnelle découle de l''analyse. Le T.S. doit nommer, cibler et prioriser les problèmes tout en qualifiant leur sévérité, leur gravité et leur intensité. Elle est partagée avec la personne pour valider sa perception et l''enrichir de son point de vue (OTSTCFQ, 2010)."
          },
          {
            "id": "opinion_professionnelle",
            "type": "textarea",
            "label": "Opinion professionnelle du T.S.",
            "required": true,
            "rows": 8,
            "help": "Formuler l''opinion de façon claire, fondée sur l''analyse. Nommer et prioriser les problèmes, qualifier leur gravité et leur intensité. L''opinion doit être distinguée des faits rapportés.",
            "placeholder": "En ma qualité de travailleur social, et sur la base de mon évaluation du fonctionnement social, mon opinion professionnelle est que..."
          },
          {
            "id": "recommandations",
            "type": "textarea",
            "label": "Recommandations",
            "required": true,
            "rows": 7,
            "placeholder": "Je recommande: (stratégies d''intervention, services à mettre en place, ressources à mobiliser, plan de protection, démarches légales, orientations, régime de protection, suivi envisagé)..."
          },
          {
            "id": "plan_intervention_strategies",
            "type": "textarea",
            "label": "Plan d''intervention et stratégies proposées",
            "rows": 6,
            "placeholder": "Objectifs d''intervention prioritaires, stratégies envisagées, ressources à mobiliser, rôles des différents acteurs, échéancier approximatif, indicateurs de changement..."
          },
          {
            "id": "opinion_partagee_client",
            "type": "select",
            "label": "Opinion partagée avec la personne",
            "help": "L''OTSTCFQ recommande que l''opinion professionnelle soit partagée avec la personne pour valider sa perception.",
            "options": [
              "Oui — la personne adhère à l''opinion",
              "Oui — avec nuances ou désaccords partiels de la personne",
              "Oui — la personne est en désaccord",
              "Non — situation ne le permettait pas (préciser)",
              "Non — autre raison (préciser dans les notes)"
            ]
          },
          {
            "id": "point_vue_client_sur_opinion",
            "type": "textarea",
            "label": "Point de vue de la personne sur l''opinion professionnelle",
            "rows": 4,
            "placeholder": "Réaction de la personne à l''opinion partagée: accord, nuances, désaccords, éléments qu''elle souhaitait ajouter ou corriger..."
          }
        ]
      },
      {
        "id": "partie6_interventions_realisees",
        "label": "Partie 6 — Interventions réalisées",
        "icone": "Activity",
        "couleur": "orange",
        "fields": [
          {
            "id": "info_p6",
            "type": "info",
            "contenu": "Dans certains contextes, les interventions réalisées dans le cadre du processus d''évaluation sont décrites dans le rapport. Cette rubrique est à compléter selon le mandat et le milieu de pratique."
          },
          {
            "id": "interventions_realisees",
            "type": "table",
            "label": "Interventions réalisées dans le cadre de l''évaluation",
            "columns": [
              { "id": "date", "label": "Date", "type": "date" },
              { "id": "type_intervention", "label": "Type d''intervention", "type": "text" },
              { "id": "personnes_impliquees", "label": "Personnes impliquées", "type": "text" },
              { "id": "duree", "label": "Durée (min)", "type": "text" },
              { "id": "objectif", "label": "Objectif / résultat", "type": "text" }
            ]
          },
          {
            "id": "demarches_realisees",
            "type": "textarea",
            "label": "Démarches et concertations réalisées",
            "rows": 4,
            "placeholder": "Consultations avec d''autres professionnels, transmissions d''information, démarches de référence, réunions d''équipe, contacts avec les ressources..."
          }
        ]
      },
      {
        "id": "aspects_deontologiques",
        "label": "Aspects déontologiques",
        "icone": "ShieldCheck",
        "couleur": "slate",
        "fields": [
          {
            "id": "info_deonto",
            "type": "info",
            "contenu": "Le T.S. s''assure que son évaluation respecte les droits de la personne garantis par les lois et les chartes, les principes d''autonomie et d''autodétermination, ainsi que les obligations déontologiques de l''OTSTCFQ (art. 3.01.05 du code de déontologie)."
          },
          {
            "id": "consentement_obtenu",
            "type": "select",
            "label": "Consentement éclairé obtenu",
            "required": true,
            "options": [
              "Oui — de la personne elle-même",
              "Oui — du représentant légal (préciser)",
              "Partiellement — situation de consentement limité (préciser)",
              "Non applicable — contexte légal (préciser)"
            ]
          },
          {
            "id": "confidentialite_notes",
            "type": "textarea",
            "label": "Notes sur la confidentialité et les autorisations",
            "rows": 3,
            "placeholder": "Autorisations obtenues pour partager les informations, limites à la confidentialité communiquées au client, tiers impliqués et autorisations..."
          },
          {
            "id": "limites_evaluation",
            "type": "textarea",
            "label": "Limites de l''évaluation",
            "rows": 4,
            "placeholder": "Informations manquantes, sources non consultées, facteurs ayant limité la collecte de données, biais potentiels à mentionner, éléments nécessitant un complément d''évaluation..."
          },
          {
            "id": "assises_theoriques",
            "type": "textarea",
            "label": "Assises théoriques et cadres de référence utilisés",
            "rows": 3,
            "placeholder": "Approches théoriques ayant guidé l''analyse: approche systémique, écologique, forces, empowerment, féministe, interculturelle, etc."
          }
        ]
      },
      {
        "id": "signature_efs",
        "label": "Signature & Scellement",
        "icone": "Lock",
        "couleur": "indigo",
        "fields": [
          {
            "id": "info_signature",
            "type": "info",
            "contenu": "En scellant ce rapport d''évaluation du fonctionnement social, vous attestez que son contenu reflète fidèlement votre évaluation professionnelle réalisée conformément au Cadre de référence de l''OTSTCFQ (2010) et au code de déontologie. Le document sera verrouillé et ne pourra plus être modifié."
          },
          {
            "id": "nom_signataire",
            "type": "text",
            "label": "Nom du signataire",
            "required": true,
            "placeholder": "Prénom Nom, T.S."
          },
          {
            "id": "no_membre_otstcfq",
            "type": "text",
            "label": "N° de membre OTSTCFQ",
            "placeholder": "Ex: TS-12345"
          },
          {
            "id": "date_rapport",
            "type": "date",
            "label": "Date du rapport"
          },
          {
            "id": "type_rapport",
            "type": "select",
            "label": "Type de rapport",
            "options": [
              "Rapport détaillé",
              "Rapport sommaire"
            ]
          }
        ]
      }
    ]
  }',
  1,
  1
);
`
