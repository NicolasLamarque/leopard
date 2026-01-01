package models

import "time"

type Medecin struct {
	ID             int       `db:"id" json:"id"`
	Licence        string    `db:"licence" json:"licence"`
	NomComplet     string    `db:"nomComplet" json:"nomComplet"`
	Specialisation string    `db:"specialisation" json:"specialisation"`
	Telephone      string    `db:"telephone" json:"telephone"`
	Extension      string    `db:"extension" json:"extension"`
	Cellulaire     string    `db:"cellulaire" json:"cellulaire"`
	Email          string    `db:"email" json:"email"`
	Adresse        string    `db:"adresse" json:"adresse"`
	CodePostal     string    `db:"code_postal" json:"code_postal"`
	Ville          string    `db:"ville" json:"ville"`
	Pays           string    `db:"pays" json:"pays"`
	NoteFixe       string    `db:"Note_fixe" json:"note_fixe"`
	Actif          int       `db:"actif" json:"actif"`
	CreatedBy      int       `db:"created_by" json:"created_by"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

type CreateMedecinRequest struct {
	Licence        string `json:"licence" validate:"required"`
	NomComplet     string `json:"nomComplet" validate:"required"`
	Specialisation string `json:"specialisation"`
	Telephone      string `json:"telephone"`
	Extension      string `json:"extension"`
	Cellulaire     string `json:"cellulaire"`
	Email          string `json:"email"`
	Adresse        string `json:"adresse"`
	CodePostal     string `json:"code_postal"`
	Ville          string `json:"ville"`
	Pays           string `json:"pays"`
	NoteFixe       string `json:"note_fixe"`
	Actif          int    `json:"actif"`
}

type UpdateMedecinRequest struct {
	ID             int    `json:"id" validate:"required"`
	Licence        string `json:"licence" validate:"required"`
	NomComplet     string `json:"nomComplet" validate:"required"`
	Specialisation string `json:"specialisation"`
	Telephone      string `json:"telephone"`
	Extension      string `json:"extension"`
	Cellulaire     string `json:"cellulaire"`
	Email          string `json:"email"`
	Adresse        string `json:"adresse"`
	CodePostal     string `json:"code_postal"`
	Ville          string `json:"ville"`
	Pays           string `json:"pays"`
	NoteFixe       string `json:"note_fixe"`
	Actif          int    `json:"actif"`
}
