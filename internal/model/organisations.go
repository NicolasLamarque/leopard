package models

import "time"

// ═══════════════════════════════════════════════════════════════════════════════
// MODÈLES — Organisations & Payeurs
// ═══════════════════════════════════════════════════════════════════════════════

// ─────────────────────────────────────────────────────────────────────────────
// Organisation
// ─────────────────────────────────────────────────────────────────────────────

type Organisation struct {
	ID int `db:"id"`

	// EN CLAIR
	TypeOrg            string `db:"type_org"`
	Statut             string `db:"statut"`
	NoFournisseur      string `db:"no_fournisseur"`
	ConditionsPaiement int    `db:"conditions_paiement"`
	TauxTPSExempt      int    `db:"taux_tps_exempt"`
	ModePaiementPref   string `db:"mode_paiement_pref"`
	Actif              int    `db:"actif"`

	// CHIFFRÉ — identité
	Nom         *string `db:"nom"`
	Acronyme    *string `db:"acronyme"`
	NoOrganisme *string `db:"no_organisme"`

	// CHIFFRÉ — coordonnées
	Adresse      *string `db:"adresse"`
	Ville        *string `db:"ville"`
	CodePostal   *string `db:"code_postal"`
	Province     *string `db:"province"`
	Pays         *string `db:"pays"`
	Telephone    *string `db:"telephone"`
	Telecopieur  *string `db:"telecopieur"`
	EmailGeneral *string `db:"email_general"`
	SiteWeb      *string `db:"site_web"`

	// CHIFFRÉ — personne-ressource
	ContactNom       *string `db:"contact_nom"`
	ContactTitre     *string `db:"contact_titre"`
	ContactTelephone *string `db:"contact_telephone"`
	ContactEmail     *string `db:"contact_email"`

	// CHIFFRÉ — notes
	Notes               *string `db:"notes"`
	ConditionsSpeciales *string `db:"conditions_speciales"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateOrganisationRequest struct {
	// EN CLAIR
	TypeOrg            string `json:"type_org"`
	NoFournisseur      string `json:"no_fournisseur,omitempty"`
	ConditionsPaiement int    `json:"conditions_paiement"`
	TauxTPSExempt      int    `json:"taux_tps_exempt"`
	ModePaiementPref   string `json:"mode_paiement_pref"`

	// CHIFFRÉ
	Nom                 *string `json:"nom"`
	Acronyme            *string `json:"acronyme,omitempty"`
	NoOrganisme         *string `json:"no_organisme,omitempty"`
	Adresse             *string `json:"adresse,omitempty"`
	Ville               *string `json:"ville,omitempty"`
	CodePostal          *string `json:"code_postal,omitempty"`
	Province            *string `json:"province,omitempty"`
	Pays                *string `json:"pays,omitempty"`
	Telephone           *string `json:"telephone,omitempty"`
	Telecopieur         *string `json:"telecopieur,omitempty"`
	EmailGeneral        *string `json:"email_general,omitempty"`
	SiteWeb             *string `json:"site_web,omitempty"`
	ContactNom          *string `json:"contact_nom,omitempty"`
	ContactTitre        *string `json:"contact_titre,omitempty"`
	ContactTelephone    *string `json:"contact_telephone,omitempty"`
	ContactEmail        *string `json:"contact_email,omitempty"`
	Notes               *string `json:"notes,omitempty"`
	ConditionsSpeciales *string `json:"conditions_speciales,omitempty"`
}

type UpdateOrganisationRequest struct {
	ID int `json:"id"`
	CreateOrganisationRequest
}

// ─────────────────────────────────────────────────────────────────────────────
// Payeur — vue unifiée "qui paie"
// ─────────────────────────────────────────────────────────────────────────────

type Payeur struct {
	ID int `db:"id"`

	// EN CLAIR
	TypePayeur     string `db:"type_payeur"` // client | contact | organisation
	ClientID       *int   `db:"client_id"`
	ContactID      *int   `db:"contact_id"`
	OrganisationID *int   `db:"organisation_id"`
	Actif          int    `db:"actif"`
	EstDefaut      int    `db:"est_defaut"`

	// CHIFFRÉ
	NomFacturation     *string `db:"nom_facturation"`
	AdresseFacturation *string `db:"adresse_facturation"`
	EmailFacturation   *string `db:"email_facturation"`
	Notes              *string `db:"notes"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	// Champs calculés (non DB) — remplis par le repo après JOIN
	AffichageNom    string `db:"-" json:"affichage_nom"`
	AffichageType   string `db:"-" json:"affichage_type"`
	AffichageDetail string `db:"-" json:"affichage_detail"`
}

type CreatePayeurRequest struct {
	// EN CLAIR
	TypePayeur     string `json:"type_payeur"`
	ClientID       *int   `json:"client_id,omitempty"`
	ContactID      *int   `json:"contact_id,omitempty"`
	OrganisationID *int   `json:"organisation_id,omitempty"`
	EstDefaut      int    `json:"est_defaut"`

	// CHIFFRÉ
	NomFacturation     *string `json:"nom_facturation,omitempty"`
	AdresseFacturation *string `json:"adresse_facturation,omitempty"`
	EmailFacturation   *string `json:"email_facturation,omitempty"`
	Notes              *string `json:"notes,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
// PayeurResolu — DTO plat pour le frontend
// Contient toutes les infos du payeur déchiffrées, peu importe sa source
// ─────────────────────────────────────────────────────────────────────────────

type PayeurResolu struct {
	PayeurID      int    `json:"payeur_id"`
	TypePayeur    string `json:"type_payeur"`
	Nom           string `json:"nom"`
	Detail        string `json:"detail"` // titre, type org, lien au client
	Email         string `json:"email"`
	Telephone     string `json:"telephone"`
	Adresse       string `json:"adresse"`
	ModePaiement  string `json:"mode_paiement"`
	JoursPaiement int    `json:"jours_paiement"`
	ExempteTPSTVQ bool   `json:"exempte_tps_tvq"`
}

// ─────────────────────────────────────────────────────────────────────────────
// OrganisationListItem — version légère pour les selects et listes
// ─────────────────────────────────────────────────────────────────────────────

type OrganisationListItem struct {
	ID         int    `json:"id"`
	TypeOrg    string `json:"type_org"`
	Nom        string `json:"nom"`
	Acronyme   string `json:"acronyme"`
	ContactNom string `json:"contact_nom"`
	Telephone  string `json:"telephone"`
	Actif      int    `json:"actif"`
}
