// internal/model/user.go

package models

type User struct {
	ID        int64  `db:"id"         json:"id"`
	Username  string `db:"username"   json:"username"`
	Password  string `db:"password"   json:"-"` // Mappe la colonne TEXT password
	Salt      string `db:"salt"       json:"-"`
	FullName  string `db:"full_name"  json:"fullName"`
	Role      string `db:"role"       json:"role"`
	CreatedAt string `db:"created_at" json:"createdAt"`
}
