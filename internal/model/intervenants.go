package models

import (
	"fmt"
	"strings"
	"time"
)

// Intervenant représente un intervenant (médecin, travailleur social, etc.)
// dans le système de gestion de clientèle
type Intervenant struct {
	// Identification
	ID          int    `db:"id" json:"id"`
	NomComplet  string `db:"nom_complet" json:"nom_complet"`
	TitreEmploi string `db:"titre_emploi" json:"titre_emploi"`

	// Affectation
	Direction    string `db:"direction" json:"direction"`
	Installation string `db:"installation" json:"installation"`

	// Contact Professionnel (TOUS CRYPTÉS)
	Telephone     string `db:"telephone" json:"telephone"`           // Crypté
	Poste         string `db:"poste" json:"poste"`                   // Crypté
	CellulairePro string `db:"cellulaire_pro" json:"cellulaire_pro"` // Crypté

	// Contact Personnel (CRYPTÉ)
	CellulairePerso string `db:"cellulaire_perso" json:"cellulaire_perso"` // Crypté

	// Emails (TOUS CRYPTÉS)
	CourrielPersonnel     string `db:"courriel_personnel" json:"courriel_personnel"`         // Crypté
	CourrielProfessionnel string `db:"courriel_professionnel" json:"courriel_professionnel"` // Crypté
	CourrierInterne       string `db:"courrier_interne" json:"courrier_interne"`             // Crypté

	// Statuts
	Actif               bool `db:"actif" json:"actif"`
	IsIntervenantSocial bool `db:"is_intervenant_social" json:"is_intervenant_social"`

	// Informations Professionnelles
	NumeroPermis       string `db:"numero_permis" json:"numero_permis"`             // Crypté
	OrdreProfessionnel string `db:"ordre_professionnel" json:"ordre_professionnel"` // NON crypté
	DateNaissance      string `db:"date_naissance" json:"date_naissance"`           // Crypté (format: YYYY-MM-DD)

	// Notes
	NoteFixe string `db:"note_fixe" json:"note_fixe"` // Crypté

	// Réseau
	PersonneRessourceReseauDir string `db:"personne_ressource_reseau_dir" json:"personne_ressource_reseau_dir"`

	// Métadonnées
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// IsActive retourne true si l'intervenant est actif
func (i *Intervenant) IsActive() bool {
	return i.Actif
}

// GetFullContact retourne une chaîne avec toutes les coordonnées
func (i *Intervenant) GetFullContact() string {
	contact := ""

	if i.Telephone != "" {
		contact += "Tél: " + i.Telephone
		if i.Poste != "" {
			contact += " poste " + i.Poste
		}
		contact += "\n"
	}

	if i.CellulairePro != "" {
		contact += "Cell Pro: " + i.CellulairePro + "\n"
	}

	if i.CourrielProfessionnel != "" {
		contact += "Email: " + i.CourrielProfessionnel + "\n"
	}

	return contact
}

// GetTypeLabel retourne le label du type d'intervenant
func (i *Intervenant) GetTypeLabel() string {
	if i.IsIntervenantSocial {
		return "Intervenant Social"
	}
	return "Personnel Administratif"
}

// Validate effectue une validation de base des données
func (i *Intervenant) Validate() error {
	if i.NomComplet == "" {
		return fmt.Errorf("le nom complet est obligatoire")
	}

	// Validation email professionnel si présent
	if i.CourrielProfessionnel != "" && !isValidEmail(i.CourrielProfessionnel) {
		return fmt.Errorf("courriel professionnel invalide")
	}

	// Validation email personnel si présent
	if i.CourrielPersonnel != "" && !isValidEmail(i.CourrielPersonnel) {
		return fmt.Errorf("courriel personnel invalide")
	}

	return nil
}

// Helper pour validation email
func isValidEmail(email string) bool {
	// Validation simple - peut être améliorée avec regex
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
