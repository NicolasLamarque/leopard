package models

import "time"

// EvaluationV2 représente une instance d'évaluation en DB
type EvaluationV2 struct {
	ID        int    `json:"id" db:"id"`
	ClientID  int    `json:"client_id" db:"client_id"`
	ModelID   string `json:"model_id" db:"model_id"`
	NoLeopard string `json:"no_leopard" db:"no_leopard"`

	// On dit à Go : "Ne cherche pas de colonne 'payload' en DB, je le gère moi-même"
	Payload string `json:"payload" db:"-"`

	// On dit à Go : "Va lire/écrire le crypté dans la colonne 'payload_crypte'"
	PayloadCrypte string `json:"-" db:"payload_crypte"`

	Statut        string     `json:"statut" db:"statut"`
	SignatureNom  *string    `json:"signature_nom" db:"signature_nom"`
	DateSignature *time.Time `json:"date_signature" db:"date_signature"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

// EvaluationDefinition définit le "moule" du formulaire
type EvaluationDefinition struct {
	ID         string `json:"id" db:"id"`
	Nom        string `json:"nom" db:"nom"`
	Icone      string `json:"icone" db:"icone"`
	Couleur    string `json:"couleur" db:"couleur"`
	SchemaJSON string `json:"schema_json" db:"schema_json"` // AJOUTE db:"schema_json"
}
