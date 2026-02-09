// internal/model/user.go

package models

type User struct {
	ID            int64  `db:"id"         json:"id"`
	Username      string `db:"username"   json:"username"`
	Password      string `db:"password"   json:"-"` // Mappe la colonne TEXT password
	Salt          string `db:"salt"       json:"-"`
	FullName      string `db:"full_name"  json:"fullName"`
	Role          string `db:"role"       json:"role"`
	NoMembreOrdre string `db:"NoMembreOrdre" json:"noMembreOrdre"`
	Email         string `db:"email"      json:"email"`
	Telephone     string `db:"telephone"  json:"telephone"`
	Cellulaire    string `db:"cellulaire" json:"cellulaire"`
	Telecopieur   string `db:"telecopieur" json:"telecopieur"`
	Adresse       string `db:"adresse"    json:"adresse"`
	CodePostal    string `db:"code_postal" json:"codePostal"`
	Ville         string `db:"ville"      json:"ville"`
	Pays          string `db:"pays"       json:"pays"`
	CreatedAt     string `db:"created_at" json:"createdAt"`
}
