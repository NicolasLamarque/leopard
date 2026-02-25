package models

// ═══════════════════════════════════════════════════════════════════════════════
// MODÈLES FINANCIERS — Leopard — version finale
//
// Convention de type :
//   float64 / int / string  →  EN CLAIR en DB (calculs SQL directs)
//   *string                 →  CHIFFRÉ en DB (identifiants personnels uniquement)
//
// Champs chiffrés restants :
//   ParametresFinance : nom_complet, titre, no_membre, email, téléphone,
//                       adresse, no_tps, no_tvq, interac_email
//   Revenu            : reference_paiement
//   Facture           : tiers_payant_nom
//   Paiement          : reference
// ═══════════════════════════════════════════════════════════════════════════════

import "time"

// ─────────────────────────────────────────────────────────────────────────────
// ParametresFinance
// ─────────────────────────────────────────────────────────────────────────────

type ParametresFinance struct {
	ID int `db:"id"`

	// CHIFFRÉ — identifiants personnels de la praticienne
	NomComplet             *string `db:"nom_complet"`
	TitreProfessionnel     *string `db:"titre_professionnel"`
	NoMembreOrdre          *string `db:"no_membre_ordre"`
	EmailProfessionnel     *string `db:"email_professionnel"`
	TelephoneProfessionnel *string `db:"telephone_professionnel"`
	AdressePratique        *string `db:"adresse_pratique"`
	NoTPS                  *string `db:"no_tps"`
	NoTVQ                  *string `db:"no_tvq"`
	InteracEmail           *string `db:"interac_email"`

	// EN CLAIR — configuration
	TauxHoraireDefaut     float64 `db:"taux_horaire_defaut"`
	TauxForfaitEvaluation float64 `db:"taux_forfait_evaluation"`
	TauxRapportExpertise  float64 `db:"taux_rapport_expertise"`
	TauxKM                float64 `db:"taux_km"`
	SousSeuilTPS          int     `db:"sous_seuil_tps"`
	InscriteTPS           int     `db:"inscrite_tps"`
	InscriteTVQ           int     `db:"inscrite_tvq"`
	TauxTPS               float64 `db:"taux_tps"`
	TauxTVQ               float64 `db:"taux_tvq"`
	PrefixeFacture        string  `db:"prefixe_facture"`
	AnneeFacture          int     `db:"annee_facture"`
	ProchainNoFacture     int     `db:"prochain_no_facture"`
	JoursPaiementDefaut   int     `db:"jours_paiement_defaut"`
	InfosPaiement         string  `db:"infos_paiement"`
	PolitiqueAnnulation   string  `db:"politique_annulation"`
	MentionsLegalesPDF    string  `db:"mentions_legales_pdf"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Service — catalogue complet, rien de chiffré
// ─────────────────────────────────────────────────────────────────────────────

type Service struct {
	ID int `db:"id"`

	// Tout EN CLAIR
	Code                   string   `db:"code"`
	Categorie              string   `db:"categorie"`
	TypeLivraison          string   `db:"type_livraison"`
	Actif                  int      `db:"actif"`
	OrdreAffichage         int      `db:"ordre_affichage"`
	Nom                    string   `db:"nom"`
	DescriptionCourte      string   `db:"description_courte"`
	DescriptionLongue      string   `db:"description_longue"`
	PublicCible            string   `db:"public_cible"`
	NotesInternes          string   `db:"notes_internes"`
	ModeTarification       string   `db:"mode_tarification"`
	DureeMinutes           *int     `db:"duree_minutes"`
	TauxHoraire            *float64 `db:"taux_horaire"`
	TarifUnitaire          *float64 `db:"tarif_unitaire"`
	TarifMin               *float64 `db:"tarif_min"`
	TarifMax               *float64 `db:"tarif_max"`
	ExempteTaxes           int      `db:"exempte_taxes"`
	TPSApplicable          int      `db:"tps_applicable"`
	TVQApplicable          int      `db:"tvq_applicable"`
	NbPlacesMax            *int     `db:"nb_places_max"`
	NbSeancesForfait       *int     `db:"nb_seances_forfait"`
	DureeProgrammeSemaines *int     `db:"duree_programme_semaines"`
	FormatTract            string   `db:"format_tract"`
	CouleurTract           string   `db:"couleur_tract"`
	DureeVideoMinutes      *int     `db:"duree_video_minutes"`
	URLRessource           string   `db:"url_ressource"`
	Tags                   string   `db:"tags"`
	Icone                  string   `db:"icone"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateServiceRequest struct {
	Code                   string   `json:"code"`
	Categorie              string   `json:"categorie"`
	TypeLivraison          string   `json:"type_livraison"`
	Actif                  int      `json:"actif"`
	OrdreAffichage         int      `json:"ordre_affichage"`
	Nom                    string   `json:"nom"`
	DescriptionCourte      string   `json:"description_courte,omitempty"`
	DescriptionLongue      string   `json:"description_longue,omitempty"`
	PublicCible            string   `json:"public_cible,omitempty"`
	NotesInternes          string   `json:"notes_internes,omitempty"`
	ModeTarification       string   `json:"mode_tarification"`
	DureeMinutes           *int     `json:"duree_minutes,omitempty"`
	TauxHoraire            *float64 `json:"taux_horaire,omitempty"`
	TarifUnitaire          *float64 `json:"tarif_unitaire,omitempty"`
	TarifMin               *float64 `json:"tarif_min,omitempty"`
	TarifMax               *float64 `json:"tarif_max,omitempty"`
	ExempteTaxes           int      `json:"exempte_taxes"`
	TPSApplicable          int      `json:"tps_applicable"`
	TVQApplicable          int      `json:"tvq_applicable"`
	NbPlacesMax            *int     `json:"nb_places_max,omitempty"`
	NbSeancesForfait       *int     `json:"nb_seances_forfait,omitempty"`
	DureeProgrammeSemaines *int     `json:"duree_programme_semaines,omitempty"`
	FormatTract            string   `json:"format_tract,omitempty"`
	CouleurTract           string   `json:"couleur_tract,omitempty"`
	DureeVideoMinutes      *int     `json:"duree_video_minutes,omitempty"`
	URLRessource           string   `json:"url_ressource,omitempty"`
	Tags                   string   `json:"tags,omitempty"`
	Icone                  string   `json:"icone,omitempty"`
}

type UpdateServiceRequest struct {
	ID int `json:"id"`
	CreateServiceRequest
}

// ─────────────────────────────────────────────────────────────────────────────
// Revenu — un seul champ chiffré : reference_paiement
// ─────────────────────────────────────────────────────────────────────────────

type Revenu struct {
	ID        int  `db:"id"`
	ClientID  int  `db:"client_id"`
	ServiceID *int `db:"service_id"`
	FactureID *int `db:"facture_id"`

	// EN CLAIR — tout
	DateService         string  `db:"date_service"`
	Categorie           string  `db:"categorie"`
	ModeFacturation     string  `db:"mode_facturation"`
	StatutPaiement      string  `db:"statut_paiement"`
	ModePaiement        string  `db:"mode_paiement"`
	DatePaiement        string  `db:"date_paiement"`
	DureeHeures         float64 `db:"duree_heures"`
	TauxHoraireApplique float64 `db:"taux_horaire_applique"`
	MontantBrut         float64 `db:"montant_brut"`
	MontantTPS          float64 `db:"montant_tps"`
	MontantTVQ          float64 `db:"montant_tvq"`
	MontantTotal        float64 `db:"montant_total"`
	Description         string  `db:"description"`
	Notes               string  `db:"notes"`

	// CHIFFRÉ — confirmation Interac (lié à un compte bancaire)
	ReferencePaiement *string `db:"reference_paiement"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateRevenuRequest struct {
	ClientID            int     `json:"client_id"`
	ServiceID           *int    `json:"service_id,omitempty"`
	FactureID           *int    `json:"facture_id,omitempty"`
	DateService         string  `json:"date_service"`
	Categorie           string  `json:"categorie"`
	ModeFacturation     string  `json:"mode_facturation"`
	StatutPaiement      string  `json:"statut_paiement"`
	ModePaiement        string  `json:"mode_paiement"`
	DatePaiement        string  `json:"date_paiement,omitempty"`
	DureeHeures         float64 `json:"duree_heures"`
	TauxHoraireApplique float64 `json:"taux_horaire_applique"`
	MontantBrut         float64 `json:"montant_brut"`
	MontantTPS          float64 `json:"montant_tps"`
	MontantTVQ          float64 `json:"montant_tvq"`
	MontantTotal        float64 `json:"montant_total"`
	Description         string  `json:"description,omitempty"`
	Notes               string  `json:"notes,omitempty"`
	ReferencePaiement   *string `json:"reference_paiement,omitempty"` // CHIFFRÉ
}

type UpdateRevenuRequest struct {
	ID int `json:"id"`
	CreateRevenuRequest
}

// ─────────────────────────────────────────────────────────────────────────────
// Depense — rien de chiffré
// ─────────────────────────────────────────────────────────────────────────────

type Depense struct {
	ID int `db:"id"`

	// EN CLAIR — tout
	DateDepense            string  `db:"date_depense"`
	Categorie              string  `db:"categorie"`
	Deductible             int     `db:"deductible"`
	PctUtilisationAffaires float64 `db:"pct_utilisation_affaires"`
	EstKilometrage         int     `db:"est_kilometrage"`
	SousTotal              float64 `db:"sous_total"`
	MontantTPS             float64 `db:"montant_tps"`
	MontantTVQ             float64 `db:"montant_tvq"`
	MontantTotal           float64 `db:"montant_total"`
	MontantDeductible      float64 `db:"montant_deductible"`
	KMParcourus            float64 `db:"km_parcourus"`
	TauxKMUtilise          float64 `db:"taux_km_utilise"`
	Fournisseur            string  `db:"fournisseur"`
	NoRecu                 string  `db:"no_recu"`
	Description            string  `db:"description"`
	Notes                  string  `db:"notes"`
	CheminRecu             string  `db:"chemin_recu"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateDepenseRequest struct {
	DateDepense            string  `json:"date_depense"`
	Categorie              string  `json:"categorie"`
	Deductible             int     `json:"deductible"`
	PctUtilisationAffaires float64 `json:"pct_utilisation_affaires"`
	EstKilometrage         int     `json:"est_kilometrage"`
	SousTotal              float64 `json:"sous_total"`
	MontantTPS             float64 `json:"montant_tps"`
	MontantTVQ             float64 `json:"montant_tvq"`
	MontantTotal           float64 `json:"montant_total"`
	MontantDeductible      float64 `json:"montant_deductible"`
	KMParcourus            float64 `json:"km_parcourus"`
	TauxKMUtilise          float64 `json:"taux_km_utilise"`
	Fournisseur            string  `json:"fournisseur,omitempty"`
	NoRecu                 string  `json:"no_recu,omitempty"`
	Description            string  `json:"description"`
	Notes                  string  `json:"notes,omitempty"`
}

type UpdateDepenseRequest struct {
	ID int `json:"id"`
	CreateDepenseRequest
}

// ─────────────────────────────────────────────────────────────────────────────
// Facture + FactureLigne — seul tiers_payant_nom est chiffré
// ─────────────────────────────────────────────────────────────────────────────

type Facture struct {
	ID       int `db:"id"`
	ClientID int `db:"client_id"`

	// EN CLAIR — tout
	Numero             string  `db:"numero"`
	DateEmission       string  `db:"date_emission"`
	DateEcheance       string  `db:"date_echeance"`
	DatePaiement       string  `db:"date_paiement"`
	Statut             string  `db:"statut"`
	ExempteTaxes       int     `db:"exempte_taxes"`
	AvecTPS            int     `db:"avec_tps"`
	AvecTVQ            int     `db:"avec_tvq"`
	JoursPaiement      int     `db:"jours_paiement"`
	MontantSousTotal   float64 `db:"montant_sous_total"`
	MontantTPS         float64 `db:"montant_tps"`
	MontantTVQ         float64 `db:"montant_tvq"`
	MontantTotal       float64 `db:"montant_total"`
	MontantPaye        float64 `db:"montant_paye"`
	MontantSolde       float64 `db:"montant_solde"`
	MontantTiersPayant float64 `db:"montant_tiers_payant"`
	MontantDuClient    float64 `db:"montant_du_client"`
	TitreFacture       string  `db:"titre_facture"`
	NotesClient        string  `db:"notes_client"`
	NotesInternes      string  `db:"notes_internes"`

	// CHIFFRÉ — nom de l'assureur / organisme
	TiersPayantNom *string `db:"tiers_payant_nom"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Lignes []FactureLigne `db:"-"`
}

type FactureLigne struct {
	ID        int  `db:"id"`
	FactureID int  `db:"facture_id"`
	ServiceID *int `db:"service_id"`
	RevenuID  *int `db:"revenu_id"`

	// EN CLAIR — tout
	Ordre         int     `db:"ordre"`
	Description   string  `db:"description"`
	DateService   string  `db:"date_service"`
	Quantite      float64 `db:"quantite"`
	Unite         string  `db:"unite"`
	TarifUnitaire float64 `db:"tarif_unitaire"`
	MontantLigne  float64 `db:"montant_ligne"`
}

type CreateFactureRequest struct {
	ClientID           int                         `json:"client_id"`
	DateEmission       string                      `json:"date_emission"`
	DateEcheance       string                      `json:"date_echeance"`
	ExempteTaxes       int                         `json:"exempte_taxes"`
	AvecTPS            int                         `json:"avec_tps"`
	AvecTVQ            int                         `json:"avec_tvq"`
	JoursPaiement      int                         `json:"jours_paiement"`
	MontantSousTotal   float64                     `json:"montant_sous_total"`
	MontantTPS         float64                     `json:"montant_tps"`
	MontantTVQ         float64                     `json:"montant_tvq"`
	MontantTotal       float64                     `json:"montant_total"`
	TitreFacture       string                      `json:"titre_facture,omitempty"`
	NotesClient        string                      `json:"notes_client,omitempty"`
	NotesInternes      string                      `json:"notes_internes,omitempty"`
	TiersPayantNom     *string                     `json:"tiers_payant_nom,omitempty"` // CHIFFRÉ
	MontantTiersPayant float64                     `json:"montant_tiers_payant"`
	MontantDuClient    float64                     `json:"montant_du_client"`
	Lignes             []CreateFactureLigneRequest `json:"lignes"`
}

type CreateFactureLigneRequest struct {
	ServiceID     *int    `json:"service_id,omitempty"`
	RevenuID      *int    `json:"revenu_id,omitempty"`
	Ordre         int     `json:"ordre"`
	Description   string  `json:"description"`
	DateService   string  `json:"date_service,omitempty"`
	Quantite      float64 `json:"quantite"`
	Unite         string  `json:"unite,omitempty"`
	TarifUnitaire float64 `json:"tarif_unitaire"`
	MontantLigne  float64 `json:"montant_ligne"`
}

type UpdateFactureRequest struct {
	ID int `json:"id"`
	CreateFactureRequest
}

// ─────────────────────────────────────────────────────────────────────────────
// Paiement — seule la référence de transaction est chiffrée
// ─────────────────────────────────────────────────────────────────────────────

type Paiement struct {
	ID           int       `db:"id"`
	FactureID    int       `db:"facture_id"`
	DatePaiement string    `db:"date_paiement"`
	Source       string    `db:"source"`
	ModePaiement string    `db:"mode_paiement"`
	Montant      float64   `db:"montant"`
	Notes        string    `db:"notes"`
	Reference    *string   `db:"reference"` // CHIFFRÉ : confirmation Interac
	CreatedAt    time.Time `db:"created_at"`
}

type CreatePaiementRequest struct {
	FactureID    int     `json:"facture_id"`
	DatePaiement string  `json:"date_paiement"`
	Source       string  `json:"source"`
	ModePaiement string  `json:"mode_paiement"`
	Montant      float64 `json:"montant"`
	Notes        string  `json:"notes,omitempty"`
	Reference    *string `json:"reference,omitempty"` // CHIFFRÉ
}

// ─────────────────────────────────────────────────────────────────────────────
// Contrat — rien de chiffré
// ─────────────────────────────────────────────────────────────────────────────

type Contrat struct {
	ID        int  `db:"id"`
	ClientID  int  `db:"client_id"`
	ServiceID *int `db:"service_id"`

	// EN CLAIR — tout
	TypeContrat                  string  `db:"type_contrat"`
	Statut                       string  `db:"statut"`
	DateDebut                    string  `db:"date_debut"`
	DateFin                      string  `db:"date_fin"`
	RenouvellementAuto           int     `db:"renouvellement_auto"`
	ConsentementSigne            int     `db:"consentement_signe"`
	DateSignature                string  `db:"date_signature"`
	ModeFacturation              string  `db:"mode_facturation"`
	TauxHoraire                  float64 `db:"taux_horaire"`
	TauxForfait                  float64 `db:"taux_forfait"`
	DureeSeanceMin               *int    `db:"duree_seance_min"`
	FrequenceSeances             string  `db:"frequence_seances"`
	ClauseObjet                  string  `db:"clause_objet"`
	ClauseObjectifs              string  `db:"clause_objectifs"`
	ClauseServicesInclus         string  `db:"clause_services_inclus"`
	ClauseServicesExclus         string  `db:"clause_services_exclus"`
	ClauseTarification           string  `db:"clause_tarification"`
	ClausePaiementRetard         string  `db:"clause_paiement_retard"`
	ClauseAnnulation             string  `db:"clause_annulation"`
	ClauseAbsenceProlongee       string  `db:"clause_absence_prolongee"`
	ClauseConfidentialite        string  `db:"clause_confidentialite"`
	ClauseLimitesConfidentialite string  `db:"clause_limites_confidentialite"`
	ClauseDossierClient          string  `db:"clause_dossier_client"`
	ClauseCommunication          string  `db:"clause_communication"`
	ClauseUrgences               string  `db:"clause_urgences"`
	ClauseResiliationClient      string  `db:"clause_resiliation_client"`
	ClauseResiliationTS          string  `db:"clause_resiliation_ts"`
	ClausesSpecifiques           string  `db:"clauses_specifiques"`
	NotesInternes                string  `db:"notes_internes"`

	CreatedBy *int      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateContratRequest struct {
	ClientID                     int     `json:"client_id"`
	ServiceID                    *int    `json:"service_id,omitempty"`
	TypeContrat                  string  `json:"type_contrat"`
	DateDebut                    string  `json:"date_debut"`
	DateFin                      string  `json:"date_fin,omitempty"`
	RenouvellementAuto           int     `json:"renouvellement_auto"`
	ModeFacturation              string  `json:"mode_facturation"`
	TauxHoraire                  float64 `json:"taux_horaire"`
	TauxForfait                  float64 `json:"taux_forfait"`
	DureeSeanceMin               *int    `json:"duree_seance_min,omitempty"`
	FrequenceSeances             string  `json:"frequence_seances,omitempty"`
	ClauseObjet                  string  `json:"clause_objet,omitempty"`
	ClauseObjectifs              string  `json:"clause_objectifs,omitempty"`
	ClauseServicesInclus         string  `json:"clause_services_inclus,omitempty"`
	ClauseServicesExclus         string  `json:"clause_services_exclus,omitempty"`
	ClauseTarification           string  `json:"clause_tarification,omitempty"`
	ClausePaiementRetard         string  `json:"clause_paiement_retard,omitempty"`
	ClauseAnnulation             string  `json:"clause_annulation,omitempty"`
	ClauseAbsenceProlongee       string  `json:"clause_absence_prolongee,omitempty"`
	ClauseConfidentialite        string  `json:"clause_confidentialite,omitempty"`
	ClauseLimitesConfidentialite string  `json:"clause_limites_confidentialite,omitempty"`
	ClauseDossierClient          string  `json:"clause_dossier_client,omitempty"`
	ClauseCommunication          string  `json:"clause_communication,omitempty"`
	ClauseUrgences               string  `json:"clause_urgences,omitempty"`
	ClauseResiliationClient      string  `json:"clause_resiliation_client,omitempty"`
	ClauseResiliationTS          string  `json:"clause_resiliation_ts,omitempty"`
	ClausesSpecifiques           string  `json:"clauses_specifiques,omitempty"`
	NotesInternes                string  `json:"notes_internes,omitempty"`
}

type UpdateContratRequest struct {
	ID int `json:"id"`
	CreateContratRequest
}

// ─────────────────────────────────────────────────────────────────────────────
// StatsMensuelles + DTOs dashboard
// ─────────────────────────────────────────────────────────────────────────────

type StatsMensuelles struct {
	ID    int `db:"id"`
	Annee int `db:"annee"`
	Mois  int `db:"mois"`

	RevTotal         float64 `db:"rev_total"`
	RevConsultations float64 `db:"rev_consultations"`
	RevEvaluations   float64 `db:"rev_evaluations"`
	RevRapports      float64 `db:"rev_rapports"`
	RevForfaits      float64 `db:"rev_forfaits"`
	RevAteliers      float64 `db:"rev_ateliers"`
	RevAutres        float64 `db:"rev_autres"`
	NbSeances        int     `db:"nb_seances"`
	NbClientsUniques int     `db:"nb_clients_uniques"`

	DepTotal       float64 `db:"dep_total"`
	DepDeductible  float64 `db:"dep_deductible"`
	DepBureau      float64 `db:"dep_bureau"`
	DepFormation   float64 `db:"dep_formation"`
	DepAssurance   float64 `db:"dep_assurance"`
	DepLogiciel    float64 `db:"dep_logiciel"`
	DepMateriel    float64 `db:"dep_materiel"`
	DepDeplacement float64 `db:"dep_deplacement"`
	DepHonoraires  float64 `db:"dep_honoraires"`
	DepMarketing   float64 `db:"dep_marketing"`
	DepAutres      float64 `db:"dep_autres"`

	BeneficeNet        float64 `db:"benefice_net"`
	CumulAnnuelRevenus float64 `db:"cumul_annuel_revenus"`
	PctMarge           float64 `db:"pct_marge"`

	NbFacturesEmises   int     `db:"nb_factures_emises"`
	NbFacturesPayees   int     `db:"nb_factures_payees"`
	NbFacturesEnRetard int     `db:"nb_factures_en_retard"`
	MontantAPercevoir  float64 `db:"montant_a_percevoir"`

	UpdatedAt time.Time `db:"updated_at"`
}

type DashboardData struct {
	StatsMois    StatsMensuelles   `json:"stats_mois"`
	StatsAnnee   StatsAnnee        `json:"stats_annee"`
	DerniersMois []StatsMensuelles `json:"derniers_mois"`
	Alertes      []AlerteFinance   `json:"alertes"`
}

type StatsAnnee struct {
	Annee             int     `json:"annee"`
	RevenuBrutTotal   float64 `json:"revenu_brut_total"`
	TotalDepenses     float64 `json:"total_depenses"`
	TotalDeductible   float64 `json:"total_deductible"`
	BeneficeNet       float64 `json:"benefice_net"`
	PctVersSeuilTPS   float64 `json:"pct_vers_seuil_tps"`
	MontantAPercevoir float64 `json:"montant_a_percevoir"`
	FacturesEnRetard  int     `json:"factures_en_retard"`
}

type AlerteFinance struct {
	Type    string  `json:"type"`
	Niveau  string  `json:"niveau"` // info | warn | danger
	Message string  `json:"message"`
	Valeur  float64 `json:"valeur,omitempty"`
}

type RapportFiscalAnnuel struct {
	Annee int `json:"annee"`

	RevenuBrutTotal  float64 `json:"revenu_brut_total"`
	RevConsultations float64 `json:"rev_consultations"`
	RevEvaluations   float64 `json:"rev_evaluations"`
	RevRapports      float64 `json:"rev_rapports"`
	RevForfaits      float64 `json:"rev_forfaits"`
	RevAteliers      float64 `json:"rev_ateliers"`
	RevAutres        float64 `json:"rev_autres"`

	TotalDepensesDeduct float64 `json:"total_depenses_deduct"`
	DepBureau           float64 `json:"dep_bureau"`      // L.8810 Loyer
	DepFormation        float64 `json:"dep_formation"`   // L.8540
	DepAssurance        float64 `json:"dep_assurance"`   // L.8690
	DepLogiciel         float64 `json:"dep_logiciel"`    // L.8810 Bureau
	DepMateriel         float64 `json:"dep_materiel"`    // L.8811 Fournitures
	DepDeplacement      float64 `json:"dep_deplacement"` // L.9281 Km
	DepHonoraires       float64 `json:"dep_honoraires"`  // L.8860
	DepMarketing        float64 `json:"dep_marketing"`   // L.8520

	RevenuNetEntreprise float64 `json:"revenu_net_entreprise"`
	CotisationsRRQ      float64 `json:"cotisations_rrq"` // Net × 11.9%
	DeductionRRQ        float64 `json:"deduction_rrq"`   // ÷ 2
	RevenuImposable     float64 `json:"revenu_imposable"`
	SeuilTPSAtteint     bool    `json:"seuil_tps_atteint"`

	ParMois []StatsMensuelles `json:"par_mois"`
}

type RentabiliteParClient struct {
	ClientID         int     `json:"client_id"`
	ClientNom        string  `json:"client_nom"`
	ClientPrenom     string  `json:"client_prenom"`
	NbSeances        int     `json:"nb_seances"`
	RevenuTotal      float64 `json:"revenu_total"`
	RevenuMoyen      float64 `json:"revenu_moyen"`
	DernierContact   string  `json:"dernier_contact"`
	MontantEnAttente float64 `json:"montant_en_attente"`
}
