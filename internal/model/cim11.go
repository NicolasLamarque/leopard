package models

// Cim11Code représente un code diagnostique CIM-11 dans la DB locale.
// Données publiques OMS — pas de chiffrement.
type Cim11Code struct {
	ID            int    `db:"id" json:"id"`
	Code          string `db:"code" json:"code"`
	TitreFr       string `db:"titre_fr" json:"titre_fr"`
	TitreEn       string `db:"titre_en" json:"titre_en"`
	ChapitreCode  string `db:"chapitre_code" json:"chapitre_code"`
	ChapitreTitre string `db:"chapitre_titre" json:"chapitre_titre"`
	BlocCode      string `db:"bloc_code" json:"bloc_code"`
	BlocTitre     string `db:"bloc_titre" json:"bloc_titre"`
	ParentCode    string `db:"parent_code" json:"parent_code"`
	Niveau        int    `db:"niveau" json:"niveau"`
	IsBillable    int    `db:"is_billable" json:"is_billable"`
	Actif         int    `db:"actif" json:"actif"`
}

// Cim11SyncMeta représente les métadonnées de la dernière sync OMS.
type Cim11SyncMeta struct {
	DerniereSync *string `db:"derniere_sync" json:"derniere_sync"`
	VersionOMS   string  `db:"version_oms" json:"version_oms"`
	NbCodesFr    int     `db:"nb_codes_fr" json:"nb_codes_fr"`
	NbCodesEn    int     `db:"nb_codes_en" json:"nb_codes_en"`
	Statut       string  `db:"statut" json:"statut"`
}

// Cim11ImportResult résultat retourné au frontend après un import OMS.
type Cim11ImportResult struct {
	NbInseres  int    `json:"nb_inseres"`
	NbMisAJour int    `json:"nb_mis_a_jour"`
	Duree      string `json:"duree"`
	Erreur     string `json:"erreur,omitempty"`
}
