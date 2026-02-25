package models

// Pharmacie - Modèle 100% aligné avec la table SQLite
type Pharmacie struct {
	ID              int     `json:"id" db:"id"`
	NomOrganisation string  `json:"NomOrganisation" db:"NomOrganisation"`
	Banniere        *string `json:"Banniere" db:"Banniere"`
	Adresse         *string `json:"Adresse" db:"Adresse"`
	Ville           *string `json:"Ville" db:"Ville"`
	Region          *string `json:"Region" db:"Region"`
	Tel             *string `json:"Tel" db:"Tel"`
	Fax             *string `json:"Fax" db:"Fax"`

	// Dimanche - INTEGER en base → int en Go
	DimancheOuvert         int     `json:"DimancheOuvert" db:"DimancheOuvert"`
	HeureOuvertureDimanche *string `json:"HeureOuvertureDimanche" db:"HeureOuvertureDimanche"`
	HeureFermetureDimanche *string `json:"HeureFermetureDimanche" db:"HeureFermetureDimanche"`

	// Lundi - INTEGER en base → int en Go
	LundiOuvert         int     `json:"LundiOuvert" db:"LundiOuvert"`
	HeureOuvertureLundi *string `json:"HeureOuvertureLundi" db:"HeureOuvertureLundi"`
	HeureFermetureLundi *string `json:"HeureFermetureLundi" db:"HeureFermetureLundi"`

	// Mardi - INTEGER en base → int en Go
	MardiOuvert         int     `json:"MardiOuvert" db:"MardiOuvert"`
	HeureOuvertureMardi *string `json:"HeureOuvertureMardi" db:"HeureOuvertureMardi"`
	HeureFermetureMardi *string `json:"HeureFermetureMardi" db:"HeureFermetureMardi"`

	// Mercredi - INTEGER en base → int en Go
	MercrediOuvert         int     `json:"MercrediOuvert" db:"MercrediOuvert"`
	HeureOuvertureMercredi *string `json:"HeureOuvertureMercredi" db:"HeureOuvertureMercredi"`
	HeureFermetureMercredi *string `json:"HeureFermetureMercredi" db:"HeureFermetureMercredi"`

	// Jeudi - INTEGER en base → int en Go
	JeudiOuvert         int     `json:"JeudiOuvert" db:"JeudiOuvert"`
	HeureOuvertureJeudi *string `json:"HeureOuvertureJeudi" db:"HeureOuvertureJeudi"`
	HeureFermetureJeudi *string `json:"HeureFermetureJeudi" db:"HeureFermetureJeudi"`

	// Vendredi - INTEGER en base → int en Go
	VendrediOuvert         int     `json:"VendrediOuvert" db:"VendrediOuvert"`
	HeureOuvertureVendredi *string `json:"HeureOuvertureVendredi" db:"HeureOuvertureVendredi"`
	HeureFermetureVendredi *string `json:"HeureFermetureVendredi" db:"HeureFermetureVendredi"`

	// Samedi - INTEGER en base → int en Go
	SamediOuvert         int     `json:"SamediOuvert" db:"SamediOuvert"`
	HeureOuvertureSamedi *string `json:"HeureOuvertureSamedi" db:"HeureOuvertureSamedi"`
	HeureFermetureSamedi *string `json:"HeureFermetureSamedi" db:"HeureFermetureSamedi"`

	// Métadonnées
	DateMaj *string `json:"DateMaj" db:"DateMaj"`
	Note    *string `json:"note" db:"note"`
}

// ClientPharmacieInfo - Pour la liste des clients d'une pharmacie
type ClientPharmacieInfo struct {
	ID               int    `json:"id" db:"id"`
	Nom              string `json:"nom" db:"nom"`
	Prenom           string `json:"prenom" db:"prenom"`
	NoDossierLeopard string `json:"no_dossier_leopard" db:"no_dossier_leopard"`
	DCD              int    `json:"dcd" db:"dcd"`
	Actif            int    `json:"actif" db:"actif"`
}
