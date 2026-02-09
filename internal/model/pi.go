package models

import "time"

// PlanIntervention représente la table "plans_intervention" en base de données.
// Les champs textuels sont des pointeurs (*string) pour gérer le chiffrement et les valeurs nulles.
type PlanIntervention struct {
	ID        int `db:"id" json:"id"`
	ClientID  int `db:"client_id" json:"client_id"`
	CreatedBy int `db:"created_by" json:"created_by"`

	// Métadonnées
	Titre    string `db:"titre" json:"titre"`
	TypePlan string `db:"type_plan" json:"type_plan"` // ex: psychosocial
	Statut   string `db:"statut" json:"statut"`       // brouillon, actif, termine, archive

	// Dates (stockées en TEXT dans SQLite)
	DateDebut          *string `db:"date_debut" json:"date_debut"`
	DateFinPrevue      *string `db:"date_fin_prevue" json:"date_fin_prevue"`
	DateRevisionPrevue *string `db:"date_revision_prevue" json:"date_revision_prevue"`

	// Contenu (CHIFFRÉ en base de données via le repository)
	Contexte      *string `db:"contexte" json:"contexte"`
	Problematique *string `db:"problematique" json:"problematique"`
	Forces        *string `db:"forces" json:"forces"`
	Objectifs     *string `db:"objectifs" json:"objectifs"`
	Interventions *string `db:"interventions" json:"interventions"`
	Resultats     *string `db:"resultats" json:"resultats"`
	Ententes      *string `db:"ententes" json:"ententes"`

	// Workflow de signature
	Verrouille    int        `db:"verrouille" json:"verrouille"`
	SignatureNom  *string    `db:"signature_nom" json:"signature_nom"`
	DateSignature *time.Time `db:"date_signature" json:"date_signature"`

	// Audit automatique via SQLite
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// PlanInterventionDetail correspond à la vue "v_pi_details".
// Elle combine les données du plan avec les informations liées (Client et Auteur).
type PlanInterventionDetail struct {
	PlanIntervention

	// Champs provenant de la jointure avec la table 'clients'
	ClientNom     string  `db:"client_nom" json:"client_nom"`
	ClientPrenom  string  `db:"client_prenom" json:"client_prenom"`
	ClientLeopard *string `db:"client_leopard" json:"client_leopard"`

	// Champ provenant de la jointure avec la table 'users'
	AuteurNom string `db:"auteur_nom" json:"auteur_nom"`
}

// CreatePlanRequest définit la structure attendue lors de la réception d'un nouveau plan (JSON).
type CreatePlanRequest struct {
	ClientID int    `json:"client_id"`
	Titre    string `json:"titre"`
	TypePlan string `json:"type_plan"`

	// Dates formatées (ex: "2023-10-27")
	DateDebut          *string `json:"date_debut"`
	DateFinPrevue      *string `json:"date_fin_prevue"`
	DateRevisionPrevue *string `json:"date_revision_prevue"`

	// Contenu en clair (sera chiffré par le repository)
	Contexte      *string `json:"contexte"`
	Problematique *string `json:"problematique"`
	Forces        *string `json:"forces"`
	Objectifs     *string `json:"objectifs"`
	Interventions *string `json:"interventions"`
	Resultats     *string `json:"resultats"`
	Ententes      *string `json:"ententes"`
}

// UpdatePlanRequest définit la structure attendue lors de la mise à jour d'un plan existant (JSON).
type UpdatePlanRequest struct {
	ID int `json:"id"`

	ClientID int    `json:"client_id"`
	Titre    string `json:"titre"`
	TypePlan string `json:"type_plan"`

	// Dates formatées (ex: "2023-10-27")
	DateDebut          *string `json:"date_debut"`
	DateFinPrevue      *string `json:"date_fin_prevue"`
	DateRevisionPrevue *string `json:"date_revision_prevue"`

	// Contenu en clair (sera chiffré par le repository)
	Contexte      *string `json:"contexte"`
	Problematique *string `json:"problematique"`
	Forces        *string `json:"forces"`
	Objectifs     *string `json:"objectifs"`
	Interventions *string `json:"interventions"`
	Resultats     *string `json:"resultats"`
	Ententes      *string `json:"ententes"`
}
