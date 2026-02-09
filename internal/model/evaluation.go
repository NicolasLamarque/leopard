package models

import (
	"time"
)

// EvaluationSociale est le reflet exact de ta table SQL
type EvaluationSociale struct {
	ID        int `db:"id" json:"id"`
	ClientID  int `db:"client_id" json:"client_id"`
	CreatedBy int `db:"created_by" json:"created_by"`

	// Contenu clinique (Champs chiffrés en DB)
	ContexteEvaluation       *string `db:"contexte_evaluation" json:"contexte_evaluation"`
	MotifReference           *string `db:"motif_reference" json:"motif_reference"`
	ObjetEvaluation          *string `db:"objet_evaluation" json:"objet_evaluation"`
	CapacitesCognitives      *string `db:"capacites_cognitives" json:"capacites_cognitives"`
	EtatSantePhysique        *string `db:"etat_sante_physique" json:"etat_sante_physique"`
	DimensionsPsychoSociales *string `db:"dimensions_psycho_sociales" json:"dimensions_psycho_sociales"`
	RolesSociaux             *string `db:"roles_sociaux" json:"roles_sociaux"`
	ReseauSocialSoutien      *string `db:"reseau_social_soutien" json:"reseau_social_soutien"`
	AnalyseClinique          *string `db:"analyse_clinique" json:"analyse_clinique"`
	OpinionProfessionnelle   *string `db:"opinion_professionnelle" json:"opinion_professionnelle"`
	Recommandations          *string `db:"recommandations" json:"recommandations"`

	// Workflow
	SignatureNom  *string    `db:"signature_nom" json:"signature_nom"`
	Verrouille    int        `db:"verrouille" json:"verrouille"`
	DateSignature *time.Time `db:"date_signature" json:"date_signature"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}

// EvaluationSocialeDetail représente les données venant de la VUE v_evaluation_details
// C'est ce que tu vas envoyer à ton Frontend Vue 3
type EvaluationSocialeDetail struct {
	EvaluationSociale

	// Champs additionnels venant de la jointure dans la Vue
	ClientNom     string  `db:"client_nom" json:"client_nom"`
	ClientPrenom  string  `db:"client_prenom" json:"client_prenom"`
	ClientDN      *string `db:"client_dn" json:"client_dn"`
	ClientNAM     *string `db:"client_nam" json:"client_nam"`
	ClientLeopard *string `db:"client_leopard" json:"client_leopard"`
	AuteurNom     string  `db:"auteur_nom" json:"auteur_nom"`
}

// CreateEvaluationRequest est utilisé pour le binding FormKit -> Wails -> Go
type CreateEvaluationRequest struct {
	ClientID                 int     `json:"client_id"`
	ContexteEvaluation       *string `json:"contexte_evaluation"`
	MotifReference           *string `json:"motif_reference"`
	ObjetEvaluation          *string `json:"objet_evaluation"`
	CapacitesCognitives      *string `json:"capacites_cognitives"`
	EtatSantePhysique        *string `json:"etat_sante_physique"`
	DimensionsPsychoSociales *string `json:"dimensions_psycho_sociales"`
	RolesSociaux             *string `json:"roles_sociaux"`
	ReseauSocialSoutien      *string `json:"reseau_social_soutien"`
	AnalyseClinique          *string `json:"analyse_clinique"`
	OpinionProfessionnelle   *string `json:"opinion_professionnelle"`
	Recommandations          *string `json:"recommandations"`
}
