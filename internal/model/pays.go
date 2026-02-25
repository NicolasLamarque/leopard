package models

// Pays - Structure pour la gestion des pays
type Pays struct {
	ID        int    `db:"id" json:"id"`
	Pays      string `db:"pays" json:"pays"`
	Alpha2    string `db:"alpha2" json:"alpha2"`
	Alpha3    string `db:"alpha3" json:"alpha3"`
	Numerique int    `db:"numerique" json:"numerique"`
}

type PaysListItem struct {
	ID        int    `db:"id" json:"id"`
	Pays      string `db:"pays" json:"pays"`
	Alpha2    string `db:"alpha2" json:"alpha2"`
	Alpha3    string `db:"alpha3" json:"alpha3"`
	Numerique int    `db:"numerique" json:"numerique"`
}

type PaysListItemForSelect struct {
	ID        int    `db:"id" json:"id"`
	Pays      string `db:"pays" json:"pays"`
	Alpha2    string `db:"alpha2" json:"alpha2"`
	Alpha3    string `db:"alpha3" json:"alpha3"`
	Numerique int    `db:"numerique" json:"numerique"`
}
