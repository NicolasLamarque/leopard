package models

// Notaire représente la structure de la table notaires
type Notaire struct {
	ID              int     `db:"id" json:"id"`
	Civilite        string  `db:"civilite" json:"civilite"`
	Prenom          string  `db:"prenom" json:"prenom"`
	Nom             string  `db:"nom" json:"nom"`
	Telephone       *string `db:"telephone" json:"telephone,omitempty"`
	Cellulaire      *string `db:"cellulaire" json:"cellulaire,omitempty"`
	Telecopieur     *string `db:"telecopieur" json:"telecopieur,omitempty"`
	Adresse         *string `db:"adresse" json:"adresse,omitempty"`
	CodePostal      *string `db:"code_postal" json:"code_postal,omitempty"`
	Ville           *string `db:"ville" json:"ville,omitempty"`
	Pays            string  `db:"pays" json:"pays"`
	Email           *string `db:"email" json:"email,omitempty"`
	SecteurActivite *string `db:"secteur_activite" json:"secteur_activite,omitempty"`
	NoteFixe        *string `db:"note_fixe" json:"note_fixe,omitempty"`
	Actif           int     `db:"actif" json:"actif"`
	CreatedBy       *int    `db:"created_by" json:"created_by,omitempty"`
	CreatedAt       string  `db:"created_at" json:"created_at"`
	UpdatedAt       string  `db:"updated_at" json:"updated_at"`
}

type CreateNotaireRequest struct {
	Civilite        string  `json:"civilite"`
	Prenom          string  `json:"prenom"`
	Nom             string  `json:"nom"`
	Telephone       *string `json:"telephone,omitempty"`
	Cellulaire      *string `json:"cellulaire,omitempty"`
	Telecopieur     *string `json:"telecopieur,omitempty"`
	Adresse         *string `json:"adresse,omitempty"`
	CodePostal      *string `json:"code_postal,omitempty"`
	Ville           *string `json:"ville,omitempty"`
	Email           *string `json:"email,omitempty"`
	SecteurActivite *string `json:"secteur_activite,omitempty"`
	NoteFixe        *string `json:"note_fixe,omitempty"`
	CreatedBy       int     `json:"created_by"`
}
