package models

import "time"

// Appel - Structure simplifiée pour la gestion des appels
type Appel struct {
	ID         int       `db:"id" json:"id"`
	DateAppel  time.Time `db:"date_appel" json:"date_appel"`
	HeureAppel string    `db:"heure_appel" json:"heure_appel"`

	// Appelant (crypté)
	AppelantNom       string `db:"appelant_nom" json:"appelant_nom"`
	AppelantPrenom    string `db:"appelant_prenom" json:"appelant_prenom"`
	AppelantTelephone string `db:"appelant_telephone" json:"appelant_telephone"`
	AppelantRelation  string `db:"appelant_relation" json:"appelant_relation"`

	// Client lié (optionnel)
	ClientID *int `db:"client_id" json:"client_id"`

	// Prospect si pas de client (crypté)
	ProspectNom       string `db:"prospect_nom" json:"prospect_nom"`
	ProspectPrenom    string `db:"prospect_prenom" json:"prospect_prenom"`
	ProspectTelephone string `db:"prospect_telephone" json:"prospect_telephone"`

	// Détails de l'appel
	TypeDemande string `db:"type_demande" json:"type_demande"`
	MotifAppel  string `db:"motif_appel" json:"motif_appel"` // Crypté
	Priorite    string `db:"priorite" json:"priorite"`

	// Suivi
	Statut        string `db:"statut" json:"statut"`
	NotesInternes string `db:"notes_internes" json:"notes_internes"` // Crypté

	// RDV
	RDVDate  *string `db:"rdv_date" json:"rdv_date"`
	RDVHeure string  `db:"rdv_heure" json:"rdv_heure"`
	RDVLieu  string  `db:"rdv_lieu" json:"rdv_lieu"`

	// Gestion
	RecuPar  int  `db:"recu_par" json:"recu_par"`
	AssigneA *int `db:"assigne_a" json:"assigne_a"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// CreateAppelRequest - Requête de création
type CreateAppelRequest struct {
	DateAppel  string `json:"date_appel"`
	HeureAppel string `json:"heure_appel"`

	// Appelant
	AppelantNom       string `json:"appelant_nom"`
	AppelantPrenom    string `json:"appelant_prenom"`
	AppelantTelephone string `json:"appelant_telephone"`
	AppelantRelation  string `json:"appelant_relation"`

	// Client (optionnel)
	ClientID *int `json:"client_id"`

	// Prospect (si pas de client)
	ProspectNom       string `json:"prospect_nom"`
	ProspectPrenom    string `json:"prospect_prenom"`
	ProspectTelephone string `json:"prospect_telephone"`

	// Appel
	TypeDemande string `json:"type_demande"`
	MotifAppel  string `json:"motif_appel"`
	Priorite    string `json:"priorite"`

	// Suivi
	Statut        string `json:"statut"`
	NotesInternes string `json:"notes_internes"`

	// RDV
	RDVDate  string `json:"rdv_date"`
	RDVHeure string `json:"rdv_heure"`
	RDVLieu  string `json:"rdv_lieu"`

	// Assignation
	AssigneA *int `json:"assigne_a"`
}

// AppelListItem - Pour la sidebar (liste)
type AppelListItem struct {
	ID         int       `db:"id" json:"id"`
	DateAppel  time.Time `db:"date_appel" json:"date_appel"`
	HeureAppel string    `db:"heure_appel" json:"heure_appel"`

	// Déchiffrés pour affichage
	AppelantNom       string `db:"appelant_nom" json:"appelant_nom"`
	AppelantPrenom    string `db:"appelant_prenom" json:"appelant_prenom"`
	AppelantTelephone string `db:"appelant_telephone" json:"appelant_telephone"`

	ProspectNom    string `db:"prospect_nom" json:"prospect_nom"`
	ProspectPrenom string `db:"prospect_prenom" json:"prospect_prenom"`

	TypeDemande string `db:"type_demande" json:"type_demande"`
	Priorite    string `db:"priorite" json:"priorite"`
	Statut      string `db:"statut" json:"statut"`

	ClientID  *int      `db:"client_id" json:"client_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// StatsAppels - Statistiques
type StatsAppels struct {
	Total        int `json:"total"`
	Nouveaux     int `json:"nouveaux"`
	EnAttente    int `json:"enAttente"`
	RDVPlanifies int `json:"rdvPlanifies"`
	Urgents      int `json:"urgents"`
	Aujourdhui   int `json:"aujourdhui"`
}
