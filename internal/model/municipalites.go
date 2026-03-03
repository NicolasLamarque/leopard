package models

type Municipality struct {
	MCode     string  `db:"mcode"     json:"mcode"`
	Munnom    string  `db:"munnom"    json:"munnom"`
	Madr1     string  `db:"madr1"     json:"madr1"`
	Mcodpos   string  `db:"mcodpos"   json:"mcodpos"`
	Mtel      string  `db:"mtel"      json:"mtel"`
	Mfax      string  `db:"mfax"      json:"mfax"`
	Mcourriel string  `db:"mcourriel" json:"mcourriel"`
	Mweb      string  `db:"mweb"      json:"mweb"`
	Regadm    string  `db:"regadm"    json:"regadm"`
	Mrc       string  `db:"mrc"       json:"mrc"`
	Maire     string  `db:"maire"     json:"maire"`
	Dirgen    string  `db:"dirgen"    json:"dirgen"`
	Msuperf   float64 `db:"msuperf"   json:"msuperf"`
	Mpopul    int     `db:"mpopul"    json:"mpopul"`
	Polic     string  `db:"polic"     json:"polic"`
	Incen     string  `db:"incen"     json:"incen"`
	Loisir    string  `db:"loisir"    json:"loisir"`
	Trvpub    string  `db:"trvpub"    json:"trvpub"`
	Mesurg    string  `db:"mesurg"    json:"mesurg"`
	Urban     string  `db:"urban"     json:"urban"`
	Communic  string  `db:"communic"  json:"communic"`
	Permis    string  `db:"permis"    json:"permis"`
	Batim     string  `db:"batim"     json:"batim"`
}

type Arrondissement struct {
	Arrcod  string `db:"arrcod"  json:"arrcod"`
	Arrnom  string `db:"arrnom"  json:"arrnom"`
	Arrcode string `db:"arrcode" json:"arrcode"`
}
