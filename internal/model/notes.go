package models

import "time"

// Note représente la structure complète de la table notes en base de données
type Note struct {
	ID                  int        `db:"id" json:"id"`
	ClientID            int        `db:"client_id" json:"client_id"`
	ClientNoRAMQ        *string    `db:"client_NoRAMQ" json:"client_NoRAMQ,omitempty"` // CHIFFRÉ
	ClientNom           *string    `db:"client_Nom" json:"client_Nom,omitempty"`       // CHIFFRÉ
	ClientPrenom        *string    `db:"client_Prenom" json:"client_Prenom,omitempty"` // CHIFFRÉ
	ClientDateNaissance *time.Time `db:"client_date_naissance" json:"client_date_naissance,omitempty"`
	ClientTelephone     *string    `db:"client_Telephone" json:"client_Telephone,omitempty"`
	ClientCellulaire    *string    `db:"client_Cellulaire" json:"client_Cellulaire,omitempty"`
	ClientNoLeopard     *string    `db:"client_NoLeopard" json:"client_NoLeopard,omitempty"`
	ClientAdresse       *string    `db:"client_Adresse" json:"client_Adresse,omitempty"`
	ClientAppartement   *string    `db:"client_appartement" json:"client_appartement,omitempty"`
	ClientCodePostal    *string    `db:"client_code_postal" json:"client_code_postal,omitempty"`
	ClientVille         *string    `db:"client_ville" json:"client_ville,omitempty"`
	ClientPays          *string    `db:"client_pays" json:"client_pays,omitempty"`
	ClientProvince      *string    `db:"client_province" json:"client_province,omitempty"`
	UserID              int        `db:"user_id" json:"user_id"`
	DateNote            time.Time  `db:"date_note" json:"date_note"`
	DateIntervention    *time.Time `db:"date_intervention" json:"date_intervention,omitempty"`
	HeureIntervention   *string    `db:"heure_intervention" json:"heure_intervention,omitempty"`
	DureeIntervention   *string    `db:"duree_intervention" json:"duree_intervention,omitempty"`
	ModeIntervention    *string    `db:"mode_intervention" json:"mode_intervention,omitempty"`
	TypeIntervention    *string    `db:"type_intervention" json:"type_intervention,omitempty"`
	TypeNote            *string    `db:"type_note" json:"type_note,omitempty"`
	Titre               *string    `db:"titre" json:"titre,omitempty"`     // CHIFFRÉ
	Contenu             *string    `db:"contenu" json:"contenu,omitempty"` // CHIFFRÉ
	Verrouille          int        `db:"verrouille" json:"verrouille"`     // 0 = brouillon, 1 = verrouillé
	NoteTardive         int        `db:"note_tardive" json:"note_tardive"`
	NoteDeTier          int        `db:"note_de_tier" json:"note_de_tier"`
	SignatureNom        *string    `db:"signature_nom" json:"signature_nom,omitempty"` // CHIFFRÉ
	SignatureDate       *time.Time `db:"signature_date" json:"signature_date,omitempty"`
	NoteLieeID          *int       `db:"note_liee_id" json:"note_liee_id,omitempty"`
	TypeLien            *string    `db:"type_lien" json:"type_lien,omitempty"`
	CreatedAt           time.Time  `db:"created_at" json:"created_at"`
}

// CreateNoteRequest pour la création d'une nouvelle note (brouillon)
type CreateNoteRequest struct {
	ClientID            int     `json:"client_id" db:"client_id"`
	ClientNoRAMQ        *string `json:"client_NoRAMQ,omitempty" db:"client_NoRAMQ"`
	ClientNom           *string `json:"client_Nom,omitempty" db:"client_Nom"`
	ClientPrenom        *string `json:"client_Prenom,omitempty" db:"client_Prenom"`
	ClientDateNaissance *string `json:"client_date_naissance,omitempty" db:"client_date_naissance"`
	ClientTelephone     *string `json:"client_Telephone,omitempty" db:"client_Telephone"`
	ClientCellulaire    *string `json:"client_Cellulaire,omitempty" db:"client_Cellulaire"`
	ClientNoLeopard     *string `json:"client_NoLeopard,omitempty" db:"client_NoLeopard"`
	ClientAdresse       *string `json:"client_Adresse,omitempty" db:"client_Adresse"`
	ClientAppartement   *string `json:"client_appartement,omitempty" db:"client_appartement"`
	ClientCodePostal    *string `json:"client_code_postal,omitempty" db:"client_code_postal"`
	ClientVille         *string `json:"client_ville,omitempty" db:"client_ville"`
	ClientPays          *string `json:"client_pays,omitempty" db:"client_pays"`
	ClientProvince      *string `json:"client_province,omitempty" db:"client_province"`
	UserID              int     `json:"user_id" db:"user_id"`
	DateIntervention    *string `json:"date_intervention,omitempty" db:"date_intervention"`
	HeureIntervention   *string `json:"heure_intervention,omitempty" db:"heure_intervention"`
	DureeIntervention   *string `json:"duree_intervention,omitempty" db:"duree_intervention"`
	ModeIntervention    *string `json:"mode_intervention,omitempty" db:"mode_intervention"`
	TypeIntervention    *string `json:"type_intervention,omitempty" db:"type_intervention"`
	TypeNote            *string `json:"type_note,omitempty" db:"type_note"`
	Titre               *string `json:"titre,omitempty" db:"titre"`
	Contenu             *string `json:"contenu,omitempty" db:"contenu"`
	NoteTardive         int     `json:"note_tardive" db:"note_tardive"`
	NoteDeTier          int     `json:"note_de_tier" db:"note_de_tier"`
	NoteLieeID          *int    `json:"note_liee_id,omitempty" db:"note_liee_id"`
	TypeLien            *string `json:"type_lien,omitempty" db:"type_lien"`
}

// UpdateNoteRequest pour la mise à jour d'un brouillon
type UpdateNoteRequest struct {
	ID                int     `json:"id" db:"id"`
	DateIntervention  *string `json:"date_intervention,omitempty" db:"date_intervention"`
	HeureIntervention *string `json:"heure_intervention,omitempty" db:"heure_intervention"`
	DureeIntervention *string `json:"duree_intervention,omitempty" db:"duree_intervention"`
	ModeIntervention  *string `json:"mode_intervention,omitempty" db:"mode_intervention"`
	TypeIntervention  *string `json:"type_intervention,omitempty" db:"type_intervention"`
	TypeNote          *string `json:"type_note,omitempty" db:"type_note"`
	Titre             *string `json:"titre,omitempty" db:"titre"`
	Contenu           *string `json:"contenu,omitempty" db:"contenu"`
	NoteTardive       int     `json:"note_tardive" db:"note_tardive"`
	NoteDeTier        int     `json:"note_de_tier" db:"note_de_tier"`
	NoteLieeID        *int    `json:"note_liee_id,omitempty" db:"note_liee_id"`
	TypeLien          *string `json:"type_lien,omitempty" db:"type_lien"`
}

// NoteListItem pour l'affichage dans la sidebar (version légère)
type NoteListItem struct {
	ID               int        `db:"id" json:"id"`
	DateNote         time.Time  `db:"date_note" json:"date_note"`
	DateIntervention *time.Time `db:"date_intervention" json:"date_intervention,omitempty"`
	TypeNote         *string    `db:"type_note" json:"type_note,omitempty"`
	Titre            string     `db:"titre" json:"titre"`
	Verrouille       int        `db:"verrouille" json:"verrouille"`
	NoteTardive      int        `db:"note_tardive" json:"note_tardive"`
	TypeLien         *string    `db:"type_lien" json:"type_lien,omitempty"`
	NoteLieeID       *int       `db:"note_liee_id" json:"note_liee_id,omitempty"`
	NoteLieeTitre    string     `db:"note_liee_titre" json:"note_liee_titre,omitempty"`
	SignatureNom     *string    `db:"signature_nom" json:"signature_nom,omitempty"`
}
