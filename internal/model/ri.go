package models

import "time"

type RI struct {
	ID           int       `json:"id"`
	TitreRI      string    `json:"titre_ri"`
	Region       string    `json:"region"`
	Municipalite string    `json:"municipalite"`
	Statut       string    `json:"statut"`
	TypeRI       string    `json:"type_ri"`
	DateAjout    time.Time `json:"date_ajout"`
	DateMaj      time.Time `json:"date_maj"`
}
