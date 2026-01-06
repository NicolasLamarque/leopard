package models

// CHSLD représente un Centre d'hébergement et de soins de longue durée
type CHSLD struct {
	ID                   int    `json:"id" db:"id"`
	Region               string `json:"Region" db:"Region"`
	TitreCHSLD           string `json:"TitreCHSLD" db:"TitreCHSLD"`
	Adresse              string `json:"Adresse" db:"Adresse"`
	Municipalite         string `json:"Municipalite" db:"Municipalite"`
	CodePostal           string `json:"CodePostal" db:"CodePostal"`
	Telephone            string `json:"Telephone" db:"Telephone"`
	Telecopieur          string `json:"telecopieur" db:"telecopieur"`
	PosteGardeInfirmiere string `json:"Poste_Garde_infirmiere" db:"Poste_Garde_infirmiere"`
	Capacite             int    `json:"Capacite" db:"Capacite"`
	TypeCHSLD            string `json:"TypeCHSLD" db:"TypeCHSLD"`
	Proprietaire         string `json:"Proprietaire" db:"Proprietaire"`
	DateCertification    string `json:"DateCertification" db:"DateCertification"`
	Statut               string `json:"Statut" db:"Statut"`
	SourceURL            string `json:"SourceURL" db:"SourceURL"`
	InfosCHSLD           string `json:"InfosCHSLD" db:"InfosCHSLD"`
	Notes                string `json:"Notes" db:"Notes"`
	DateAjout            string `json:"DateAjout" db:"DateAjout"`
	DateMaj              string `json:"DateMaj" db:"DateMaj"`
}
