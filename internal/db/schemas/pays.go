package schema

var TablePays = `
CREATE TABLE IF NOT EXISTS pays (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	pays TEXT NOT NULL,
	alpha2 TEXT NOT NULL,
	alpha3 TEXT NOT NULL,
	numerique INTEGER NOT NULL
);
`
