package models

// Contact représente une personne gravitant autour d'un client
type Contact struct {
	ID                  int     `db:"id" json:"id"`
	Nom                 string  `db:"Nom" json:"nom"`
	Prenom              string  `db:"Prenom" json:"prenom"`
	Telephone           *string `db:"Telephone" json:"telephone,omitempty"`
	Cellulaire          *string `db:"Cellulaire" json:"cellulaire,omitempty"`
	Adresse             *string `db:"Adresse" json:"adresse,omitempty"`
	CodePostal          *string `db:"Code_Postal" json:"code_postal,omitempty"`
	Ville               *string `db:"Ville" json:"ville,omitempty"`
	Pays                *string `db:"Pays" json:"pays,omitempty"`
	Email               *string `db:"Email" json:"email,omitempty"`
	Employeur           *string `db:"employeur" json:"employeur,omitempty"`
	Profession          *string `db:"Profession" json:"profession,omitempty"`
	RelationClient      *string `db:"Relation_Client" json:"relation_client,omitempty"`
	LienFamilial        *string `db:"lien_familial" json:"lien_familial,omitempty"`
	ForceLien           int     `db:"force_lien" json:"force_lien"`
	ContactUrgence      int     `db:"contact_urgence" json:"contact_urgence"`
	AidantNaturel       int     `db:"aidant_naturel" json:"aidant_naturel"`
	TypeDeContact       *string `db:"type_de_contact" json:"type_de_contact,omitempty"`
	ProcurationBancaire int     `db:"procuration_bancaire" json:"procuration_bancaire"`
	ProcurationNotariee int     `db:"procuration_notariee" json:"procuration_notariee"`
	NoteFixe            *string `db:"note_fixe" json:"note_fixe,omitempty"`
	RelationType        string  `db:"relation_type" json:"relation_type"`
	ClientID            int     `db:"client_id" json:"client_id"`
	CreatedAt           string  `db:"created_at" json:"created_at"`
}

// CreateContactRequest pour créer un contact
type CreateContactRequest struct {
	Nom                 string  `json:"nom" binding:"required"`
	Prenom              string  `json:"prenom" binding:"required"`
	Telephone           *string `json:"telephone,omitempty"`
	Cellulaire          *string `json:"cellulaire,omitempty"`
	Adresse             *string `json:"adresse,omitempty"`
	CodePostal          *string `json:"code_postal,omitempty"`
	Ville               *string `json:"ville,omitempty"`
	Pays                *string `json:"pays,omitempty"`
	Email               *string `json:"email,omitempty"`
	Employeur           *string `json:"employeur,omitempty"`
	Profession          *string `json:"profession,omitempty"`
	RelationClient      *string `json:"relation_client,omitempty"`
	LienFamilial        *string `json:"lien_familial,omitempty"`
	ForceLien           int     `json:"force_lien"`
	ContactUrgence      int     `json:"contact_urgence"`
	AidantNaturel       int     `json:"aidant_naturel"`
	TypeDeContact       *string `json:"type_de_contact,omitempty"`
	ProcurationBancaire int     `json:"procuration_bancaire"`
	ProcurationNotariee int     `json:"procuration_notariee"`
	NoteFixe            *string `json:"note_fixe,omitempty"`
	RelationType        string  `json:"relation_type" binding:"required"`
	ClientID            int     `json:"client_id" binding:"required"`
}

// UpdateContactRequest pour mettre à jour un contact
type UpdateContactRequest struct {
	ID                  int     `json:"id" binding:"required"`
	Nom                 string  `json:"nom" binding:"required"`
	Prenom              string  `json:"prenom" binding:"required"`
	Telephone           *string `json:"telephone,omitempty"`
	Cellulaire          *string `json:"cellulaire,omitempty"`
	Adresse             *string `json:"adresse,omitempty"`
	CodePostal          *string `json:"code_postal,omitempty"`
	Ville               *string `json:"ville,omitempty"`
	Pays                *string `json:"pays,omitempty"`
	Email               *string `json:"email,omitempty"`
	Employeur           *string `json:"employeur,omitempty"`
	Profession          *string `json:"profession,omitempty"`
	RelationClient      *string `json:"relation_client,omitempty"`
	LienFamilial        *string `json:"lien_familial,omitempty"`
	ForceLien           int     `json:"force_lien"`
	ContactUrgence      int     `json:"contact_urgence"`
	AidantNaturel       int     `json:"aidant_naturel"`
	TypeDeContact       *string `json:"type_de_contact,omitempty"`
	ProcurationBancaire int     `json:"procuration_bancaire"`
	ProcurationNotariee int     `json:"procuration_notariee"`
	NoteFixe            *string `json:"note_fixe,omitempty"`
	RelationType        string  `json:"relation_type" binding:"required"`
	ClientID            int     `json:"client_id" binding:"required"`
}
