package models

type Client struct {
	ID                      int    `db:"id" json:"id"`
	Nom                     string `db:"nom" json:"nom"`
	Prenom                  string `db:"prenom" json:"prenom"`
	DateNaissance           string `db:"date_naissance" json:"date_naissance,omitempty"`
	Telephone               string `db:"telephone" json:"telephone"`
	Cellulaire              string `db:"cellulaire" json:"cellulaire,omitempty"`
	Email                   string `db:"email" json:"email"`
	Adresse                 string `db:"adresse" json:"adresse"`
	CodePostal              string `db:"code_postal" json:"code_postal,omitempty"`
	Ville                   string `db:"ville" json:"ville,omitempty"`
	Pays                    string `db:"pays" json:"pays,omitempty"`
	NumeroAssuranceMaladie  string `db:"numero_assurance_maladie" json:"numero_assurance_maladie"`
	NumeroSecuriteSociale   string `db:"numero_securite_sociale" json:"numero_securite_sociale"`
	NoHCM                   string `db:"no_hcm" json:"no_hcm"`
	NoCHAUR                 string `db:"no_chaur" json:"no_chaur"`
	NoDossierLeopard        string `db:"no_dossier_leopard" json:"no_dossier_leopard"`
	MedecinFamilleNoLicence string `db:"medecin_famille_No_Licence" json:"medecin_famille_No_Licence,omitempty"`
	NotaireID               int    `db:"notaire_id" json:"notaire_id,omitempty"`
	PivotID                 int    `db:"pivot_id" json:"pivot_id,omitempty"`
	RPAID                   int    `db:"rpa_id" json:"rpa_id,omitempty"`
	CHSLDID                 int    `db:"chsld_id" json:"chsld_id,omitempty"`
	RIID                    int    `db:"ri_id" json:"ri_id,omitempty"`
	NoteFixe                string `db:"note_fixe" json:"note_fixe,omitempty"`
	Actif                   int    `db:"Actif" json:"actif"`
	DCD                     int    `db:"dcd" json:"dcd"`
	CreatedBy               int    `db:"created_by" json:"created_by"`
	CreatedAt               string `db:"created_at" json:"created_at"`
}

// Request pour créer un client
type CreateClientRequest struct {
	Nom                     string `db:"nom" json:"nom"`
	Prenom                  string `db:"prenom" json:"prenom"`
	DateNaissance           string `db:"date_naissance" json:"date_naissance,omitempty"`
	Telephone               string `db:"telephone" json:"telephone"`
	Cellulaire              string `db:"cellulaire" json:"cellulaire,omitempty"`
	Email                   string `db:"email" json:"email"`
	Adresse                 string `db:"adresse" json:"adresse"`
	CodePostal              string `db:"code_postal" json:"code_postal,omitempty"`
	Ville                   string `db:"ville" json:"ville,omitempty"`
	Pays                    string `db:"pays" json:"pays,omitempty"`
	NumeroAssuranceMaladie  string `db:"numero_assurance_maladie" json:"numero_assurance_maladie"`
	NumeroSecuriteSociale   string `db:"numero_securite_sociale" json:"numero_securite_sociale"`
	NoHCM                   string `db:"no_hcm" json:"no_hcm"`
	NoCHAUR                 string `db:"no_chaur" json:"no_chaur"`
	NoDossierLeopard        string `db:"no_dossier_leopard" json:"no_dossier_leopard"`
	MedecinFamilleNoLicence string `db:"medecin_famille_No_Licence" json:"medecin_famille_No_Licence,omitempty"`
	NotaireID               int    `db:"notaire_id" json:"notaire_id,omitempty"`
	PivotID                 int    `db:"pivot_id" json:"pivot_id,omitempty"`
	RPAID                   int    `db:"rpa_id" json:"rpa_id,omitempty"`
	CHSLDID                 int    `db:"chsld_id" json:"chsld_id,omitempty"`
	RIID                    int    `db:"ri_id" json:"ri_id,omitempty"`
	NoteFixe                string `db:"note_fixe" json:"note_fixe,omitempty"`
	Actif                   int    `db:"Actif" json:"actif"`
	DCD                     int    `db:"dcd" json:"dcd"`
}

// Request pour mettre à jour un client
type UpdateClientRequest struct {
	ID                      int    `db:"id" json:"id"`
	Nom                     string `db:"nom" json:"nom"`
	Prenom                  string `db:"prenom" json:"prenom"`
	DateNaissance           string `db:"date_naissance" json:"date_naissance,omitempty"`
	Telephone               string `db:"telephone" json:"telephone"`
	Cellulaire              string `db:"cellulaire" json:"cellulaire,omitempty"`
	Email                   string `db:"email" json:"email"`
	Adresse                 string `db:"adresse" json:"adresse"`
	CodePostal              string `db:"code_postal" json:"code_postal,omitempty"`
	Ville                   string `db:"ville" json:"ville,omitempty"`
	Pays                    string `db:"pays" json:"pays,omitempty"`
	NumeroAssuranceMaladie  string `db:"numero_assurance_maladie" json:"numero_assurance_maladie"`
	NumeroSecuriteSociale   string `db:"numero_securite_sociale" json:"numero_securite_sociale"`
	NoHCM                   string `db:"no_hcm" json:"no_hcm"`
	NoCHAUR                 string `db:"no_chaur" json:"no_chaur"`
	NoDossierLeopard        string `db:"no_dossier_leopard" json:"no_dossier_leopard"`
	MedecinFamilleNoLicence string `db:"medecin_famille_No_Licence" json:"medecin_famille_No_Licence,omitempty"`
	NotaireID               int    `db:"notaire_id" json:"notaire_id,omitempty"`
	PivotID                 int    `db:"pivot_id" json:"pivot_id,omitempty"`
	RPAID                   int    `db:"rpa_id" json:"rpa_id,omitempty"`
	CHSLDID                 int    `db:"chsld_id" json:"chsld_id,omitempty"`
	RIID                    int    `db:"ri_id" json:"ri_id,omitempty"`
	NoteFixe                string `db:"note_fixe" json:"note_fixe,omitempty"`
	Actif                   int    `db:"Actif" json:"actif"`
	DCD                     int    `db:"dcd" json:"dcd"`
}
