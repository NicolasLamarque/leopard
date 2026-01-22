// src/utils/noteTemplates.js

export const noteTemplates = {
  Ouverture: {
    sujet: 'Ouverture de dossier',
    motif_consultation: `Motif de la demande de services :
- Demande volontaire / Référence par : ___
- Problématique principale : ___`,
    
    situation_actuelle: `Situation actuelle :
- Contexte familial et social : ___
- Conditions de vie : ___
- Réseau de soutien : ___
- Facteurs de risque identifiés : ___
- Facteurs de protection : ___`,
    
    observations: `Observations lors de la rencontre :
- État émotionnel : ___
- Niveau de collaboration : ___
- Compréhension de la situation : ___
- Motivation au changement : ___`,
    
    interventions: `Interventions réalisées :
- Accueil et mise en confiance
- Explication du mandat et du cadre d'intervention
- Évaluation préliminaire des besoins
- Signature des consentements`,
    
    plan_action: `Plan d'action initial :
1. ___
2. ___
3. ___`,
    
    objectifs: `Objectifs identifiés :
- Court terme : ___
- Moyen terme : ___
- Long terme : ___`,
    
    suivi_prevu: 'Prochaine rencontre prévue le : ___\nModalité : ___'
  },

  Suivi: {
    sujet: 'Note de suivi',
    motif_consultation: 'Suivi régulier dans le cadre du plan d\'intervention',
    
    situation_actuelle: `Évolution depuis la dernière rencontre :
- Changements observés : ___
- Progrès réalisés : ___
- Difficultés rencontrées : ___`,
    
    observations: `Observations :
- État général : ___
- Niveau de fonctionnement : ___
- Éléments significatifs : ___`,
    
    interventions: `Interventions réalisées aujourd'hui :
- ___
- ___
- ___`,
    
    plan_action: `Ajustements au plan d'action :
- ___`,
    
    objectifs: `Révision des objectifs :
- Objectifs maintenus : ___
- Objectifs atteints : ___
- Nouveaux objectifs : ___`,
    
    suivi_prevu: 'Prochain suivi prévu : ___'
  },

  Plan: {
    sujet: 'Plan d\'intervention',
    motif_consultation: 'Élaboration du plan d\'intervention',
    
    situation_actuelle: `Analyse de la situation :
- Besoins identifiés : ___
- Ressources disponibles : ___
- Obstacles potentiels : ___`,
    
    observations: `Évaluation :
- Forces de la personne : ___
- Défis à relever : ___
- Niveau d'engagement : ___`,
    
    interventions: `Services et interventions planifiés :
1. ___
2. ___
3. ___`,
    
    plan_action: `Actions concrètes :
- À court terme (0-3 mois) : ___
- À moyen terme (3-6 mois) : ___
- À long terme (6-12 mois) : ___`,
    
    objectifs: `Objectifs SMART :
1. Objectif : ___
   Indicateur de succès : ___
   Échéancier : ___

2. Objectif : ___
   Indicateur de succès : ___
   Échéancier : ___`,
    
    suivi_prevu: 'Révision du plan prévue le : ___\nFréquence des suivis : ___'
  },

  Fermeture: {
    sujet: 'Note de fermeture de dossier',
    motif_consultation: 'Fermeture du dossier - Raison : ___',
    
    situation_actuelle: `Situation au moment de la fermeture :
- Évolution globale : ___
- Situation stabilisée : ___
- Autonomie acquise : ___`,
    
    observations: `Bilan de l'intervention :
- Durée totale du suivi : ___
- Nombre de rencontres : ___
- Services rendus : ___`,
    
    interventions: `Synthèse des interventions :
- Principales actions réalisées : ___
- Collaborations établies : ___
- Ressources mises en place : ___`,
    
    plan_action: `Recommandations pour l'avenir :
- ___
- ___`,
    
    objectifs: `Atteinte des objectifs :
- Objectifs atteints : ___
- Objectifs partiellement atteints : ___
- Objectifs non atteints : ___
  Raisons : ___`,
    
    suivi_prevu: `Modalités de réouverture :
- Possibilité de reprendre contact : ___
- Références vers d'autres ressources : ___`
  },

  Incident: {
    sujet: 'Note d\'incident / Événement significatif',
    motif_consultation: 'Documentation d\'un événement significatif',
    
    situation_actuelle: `Contexte de l'événement :
- Date et heure : ___
- Lieu : ___
- Personnes présentes : ___`,
    
    observations: `Description factuelle de l'événement :
- Déroulement chronologique : ___
- Comportements observés : ___
- Paroles échangées : ___`,
    
    interventions: `Interventions immédiates :
- Actions posées : ___
- Personnes contactées : ___
- Mesures de sécurité : ___`,
    
    plan_action: `Actions à court terme :
- Suivi requis : ___
- Mesures préventives : ___`,
    
    objectifs: `Analyse de l'incident :
- Facteurs déclencheurs : ___
- Impacts potentiels : ___
- Leçons à retenir : ___`,
    
    suivi_prevu: 'Suivi particulier prévu : ___\nPersonnes à informer : ___'
  }
}

// Fonction helper pour obtenir un template
export const getTemplate = (type) => {
  return noteTemplates[type] || {}
}

// Types de notes disponibles
export const noteTypesList = [
  { value: 'Ouverture', label: 'Note d\'ouverture', color: 'green' },
  { value: 'Suivi', label: 'Note de suivi', color: 'blue' },
  { value: 'Plan', label: 'Plan d\'intervention', color: 'purple' },
  { value: 'Fermeture', label: 'Note de fermeture', color: 'red' },
  { value: 'Incident', label: 'Note d\'incident', color: 'orange' }
]