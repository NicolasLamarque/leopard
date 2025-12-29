package models

import "time"

// Note complète pour la DB
type Note struct {
	ID               int        `db:"id" json:"id"`
	ClientID         int        `db:"client_id" json:"client_id"`
	UserID           int        `db:"user_id" json:"user_id"`
	DateNote         time.Time  `db:"date_note" json:"date_note"`
	DateIntervention *time.Time `db:"date_intervention" json:"date_intervention,omitempty"`
	ModeIntervention string     `db:"mode_intervention" json:"mode_intervention"`
	TypeIntervention string     `db:"type_intervention" json:"type_intervention"`
	TypeNote         string     `db:"type_note" json:"type_note"`
	Sujet            string     `db:"titre" json:"sujet"`
	Contenu          string     `db:"contenu" json:"contenu"`
	Verrouille       bool       `db:"verrouille" json:"verrouille"`
	SignatureNom     string     `db:"signature_nom" json:"signature_nom"`
	SignatureDate    *time.Time `db:"signature_date" json:"signature_date"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
}

// Pour la création depuis Vue
type CreateNoteRequest struct {
	ClientID         int    `json:"client_id"`
	DateIntervention string `json:"date_intervention"`
	ModeIntervention string `json:"mode_intervention"`
	TypeIntervention string `json:"type_intervention"`
	TypeNote         string `json:"type_note"`
	Sujet            string `json:"sujet"`
	Contenu          string `json:"contenu"`
}

// L'OBJET QUI MANQUE : NotesFilter
type NotesFilter struct {
	ClientID    int    `json:"client_id"`
	SearchQuery string `json:"search_query"`
}

// Pour les listes légères
type NoteListItem struct {
	ID           int       `db:"id" json:"id"`
	TypeNote     string    `db:"type_note" json:"type_note"`
	Sujet        string    `db:"titre" json:"sujet"`
	DateNote     time.Time `db:"date_note" json:"date_note"`
	SignatureNom string    `db:"signature_nom" json:"signature_nom"`
	Verrouille   bool      `db:"verrouille" json:"verrouille"`
}
