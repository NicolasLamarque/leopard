package models

// RefListe représente une entrée dans la table universelle des listes de référence.
// Utilisée pour tous les menus déroulants de l'application (type_intervenant, statut_dossier, etc.)
type RefListe struct {
	ID        int    `db:"id"`
	Categorie string `db:"categorie"`
	Libelle   string `db:"libelle"`
	Couleur   string `db:"couleur"`
	Icone     string `db:"icone"`
	IsSysteme bool   `db:"is_systeme"`
	Ordre     int    `db:"ordre"`
	Actif     bool   `db:"actif"`
}

// CreateRefListeRequest est utilisé pour créer une nouvelle entrée utilisateur.
type CreateRefListeRequest struct {
	Categorie string `json:"categorie"`
	Libelle   string `json:"libelle"`
	Couleur   string `json:"couleur"`
	Icone     string `json:"icone"`
	Ordre     int    `json:"ordre"`
}

// UpdateOrdreRequest est utilisé pour réordonner les entrées d'une catégorie.
type UpdateOrdreRequest struct {
	ID    int `json:"id"`
	Ordre int `json:"ordre"`
}
