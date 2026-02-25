package importer

import "strings"

// NotaireRow représente exactement ce qu'on voit sur ton image
type NotaireRow struct {
	Civilite        string `json:"civilite"`
	Prenom          string `json:"prenom"`
	Nom             string `json:"nom"`
	Telephone       string `json:"telephone"`
	Cellulaire      string `json:"cellulaire"`
	Telecopieur     string `json:"telecopieur"`
	Email           string `json:"email"`
	Adresse         string `json:"adresse"`
	Ville           string `json:"ville"`
	CodePostal      string `json:"code_postal"`
	SecteurActivite string `json:"secteur_activite"`
	Note            string `json:"note"`
}

// MapToNotaire fait le pont entre Excel et ton App
func MapToNotaire(data map[string]string) NotaireRow {
	email := strings.TrimSpace(data["Email"])
	// Nettoyage : si l'email n'est pas une vraie adresse, on le vide
	if strings.Contains(email, "site Web") {
		email = ""
	}

	return NotaireRow{
		Civilite:        data["Civilite"],
		Prenom:          strings.TrimSpace(data["Prenom"]),
		Nom:             strings.TrimSpace(data["Nom"]),
		Telephone:       data["Telephone"],
		Cellulaire:      data["Cellulaire"],
		Telecopieur:     data["Telecopieur"],
		Email:           email,
		Adresse:         data["Adresse"],
		Ville:           data["Ville"],
		CodePostal:      data["Code Postal"], // Attention à l'espace dans l'entête Excel
		SecteurActivite: data["Secteur Activite"],
		Note:            data["Note"],
	}
}
