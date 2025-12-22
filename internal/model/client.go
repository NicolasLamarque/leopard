package models

type Client struct {
	ID                     int    `db:"id" json:"id"`
	Nom                    string `db:"nom" json:"nom"`
	Prenom                 string `db:"prenom" json:"prenom"`
	DateNaissance          string `db:"date_naissance" json:"date_naissance,omitempty"`
	Telephone              string `db:"telephone" json:"telephone"`
	Email                  string `db:"email" json:"email"`
	Adresse                string `db:"adresse" json:"adresse"`
	NumeroAssuranceMaladie string `db:"numero_assurance_maladie" json:"numero_assurance_maladie"`
	NumeroSecuriteSociale  string `db:"numero_securite_sociale" json:"numero_securite_sociale"`
	NoHCM                  string `db:"no_hcm" json:"no_hcm"`
	NoCHAUR                string `db:"no_chaur" json:"no_chaur"`
	NoDossierLeopard       string `db:"no_dossier_leopard" json:"no_dossier_leopard"`
	CreatedBy              int    `db:"created_by" json:"created_by"`
	CreatedAt              string `db:"created_at" json:"created_at"`
}

type CreateClientRequest struct {
	Nom                    string `db:"nom" json:"nom"`
	Prenom                 string `db:"prenom" json:"prenom"`
	DateNaissance          string `db:"date_naissance" json:"date_naissance,omitempty"`
	Telephone              string `db:"telephone" json:"telephone"`
	Email                  string `db:"email" json:"email"`
	Adresse                string `db:"adresse" json:"adresse"`
	NumeroAssuranceMaladie string `db:"numero_assurance_maladie" json:"numero_assurance_maladie"`
	NumeroSecuriteSociale  string `db:"numero_securite_sociale" json:"numero_securite_sociale"`
	NoHCM                  string `db:"no_hcm" json:"no_hcm"`
	NoCHAUR                string `db:"no_chaur" json:"no_chaur"`
	NoDossierLeopard       string `db:"no_dossier_leopard" json:"no_dossier_leopard"`
}

type UpdateClientRequest struct {
	ID                     int    `db:"id" json:"id"`
	Nom                    string `db:"nom" json:"nom"`
	Prenom                 string `db:"prenom" json:"prenom"`
	DateNaissance          string `db:"date_naissance" json:"date_naissance,omitempty"`
	Telephone              string `db:"telephone" json:"telephone"`
	Email                  string `db:"email" json:"email"`
	Adresse                string `db:"adresse" json:"adresse"`
	NumeroAssuranceMaladie string `db:"numero_assurance_maladie" json:"numero_assurance_maladie"`
	NumeroSecuriteSociale  string `db:"numero_securite_sociale" json:"numero_securite_sociale"`
	NoHCM                  string `db:"no_hcm" json:"no_hcm"`
	NoCHAUR                string `db:"no_chaur" json:"no_chaur"`
	NoDossierLeopard       string `db:"no_dossier_leopard" json:"no_dossier_leopard"`
}
