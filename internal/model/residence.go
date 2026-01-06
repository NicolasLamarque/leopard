package models

type Residence struct {
	ID                   int     `db:"id" json:"id"`
	Region               *string `db:"region" json:"region"`                 // ⬅️
	Registre             string  `db:"registre" json:"registre"`             // NOT NULL
	NumeroInterne        *string `db:"numero_interne" json:"numero_interne"` // ⬅️
	Titre                string  `db:"titre" json:"titre"`                   // NOT NULL
	Municipalite         *string `db:"municipalite" json:"municipalite"`     // ⬅️
	Adresse              *string `db:"adresse" json:"adresse"`               // ⬅️
	Ville                *string `db:"ville" json:"ville"`                   // ⬅️
	CodePostal           *string `db:"code_postal" json:"code_postal"`       // ⬅️
	Telephone            *string `db:"telephone" json:"telephone"`           // ⬅️
	Capacite             int     `db:"capacite" json:"capacite"`
	TypeResid            *string `db:"type_resid" json:"type_resid"`                       // ⬅️
	Proprietaires        *string `db:"proprietaires" json:"proprietaires"`                 // ⬅️
	Services             *string `db:"services" json:"services"`                           // ⬅️
	DateCertification    *string `db:"date_certification" json:"date_certification"`       // ⬅️
	Statut               string  `db:"statut" json:"statut"`                               // NOT NULL (default)
	SourceURL            *string `db:"source_url" json:"source_url"`                       // ⬅️
	SourceURLDetaillee   *string `db:"source_url_detaillee" json:"source_url_detaillee"`   // ⬅️
	Notes                *string `db:"notes" json:"notes"`                                 // ⬅️
	DerniereVerification *string `db:"derniere_verification" json:"derniere_verification"` // ⬅️
	DateAjout            *string `db:"date_ajout" json:"date_ajout"`                       // ⬅️
	DateMAJ              *string `db:"date_maj" json:"date_maj"`                           // ⬅️
	DateFermeture        *string `db:"date_fermeture" json:"date_fermeture"`               // ⬅️
}
