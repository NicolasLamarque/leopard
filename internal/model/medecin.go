package models

import "time"

type Medecin struct {
	ID                     int       `db:"id" json:"id"`
	Licence                string    `db:"licence" json:"licence"`
	Civilite               string    `db:"civilite" json:"civilite"`
	Prenom                 string    `db:"prenom" json:"prenom"`
	Nom                    string    `db:"nom" json:"nom"`
	NomComplet             string    `db:"nomComplet" json:"nomComplet"`
	Statut                 string    `db:"statut" json:"statut"`
	Specialisation         string    `db:"specialisation" json:"specialisation"`
	Service                string    `db:"service" json:"service"`
	Departement            string    `db:"departement" json:"departement"`
	InstallationPrincipale string    `db:"installation_principale" json:"installationPrincipale"`
	RLS                    string    `db:"rls" json:"rls"`
	Telephone              string    `db:"telephone" json:"telephone"`
	Extension              string    `db:"extension" json:"extension"`
	Cellulaire             string    `db:"cellulaire" json:"cellulaire"`
	Email                  string    `db:"email" json:"email"`
	Adresse                string    `db:"adresse" json:"adresse"`
	CodePostal             string    `db:"code_postal" json:"codePostal"`
	Ville                  string    `db:"ville" json:"ville"`
	Pays                   string    `db:"pays" json:"pays"`
	NoteFixe               string    `db:"Note_fixe" json:"noteFixe"`
	Actif                  int       `db:"actif" json:"actif"`
	CreatedBy              int       `db:"created_by" json:"createdBy"`
	CreatedAt              time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt              time.Time `db:"updated_at" json:"updatedAt"`
}

type CreateMedecinRequest struct {
	Licence                string `json:"licence" validate:"required"`
	Civilite               string `json:"civilite"`
	Prenom                 string `json:"prenom"`
	Nom                    string `json:"nom"`
	NomComplet             string `json:"nomComplet" validate:"required"`
	Statut                 string `json:"statut"`
	Specialisation         string `json:"specialisation"`
	Service                string `json:"service"`
	Departement            string `json:"departement"`
	InstallationPrincipale string `json:"installationPrincipale"`
	RLS                    string `json:"rls"`
	Telephone              string `json:"telephone"`
	Extension              string `json:"extension"`
	Cellulaire             string `json:"cellulaire"`
	Email                  string `json:"email"`
	Adresse                string `json:"adresse"`
	CodePostal             string `json:"codePostal"`
	Ville                  string `json:"ville"`
	Pays                   string `json:"pays"`
	NoteFixe               string `json:"noteFixe"`
	Actif                  int    `json:"actif"`
}

type UpdateMedecinRequest struct {
	ID                     int    `json:"id" validate:"required"`
	Licence                string `json:"licence" validate:"required"`
	Civilite               string `json:"civilite"`
	Prenom                 string `json:"prenom"`
	Nom                    string `json:"nom"`
	NomComplet             string `json:"nomComplet" validate:"required"`
	Statut                 string `json:"statut"`
	Specialisation         string `json:"specialisation"`
	Service                string `json:"service"`
	Departement            string `json:"departement"`
	InstallationPrincipale string `json:"installationPrincipale"`
	RLS                    string `json:"rls"`
	Telephone              string `json:"telephone"`
	Extension              string `json:"extension"`
	Cellulaire             string `json:"cellulaire"`
	Email                  string `json:"email"`
	Adresse                string `json:"adresse"`
	CodePostal             string `json:"codePostal"`
	Ville                  string `json:"ville"`
	Pays                   string `json:"pays"`
	NoteFixe               string `json:"noteFixe"`
	Actif                  int    `json:"actif"`
}
