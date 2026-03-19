package schema

// Tables de référence et données statiques
var System = `
	CREATE TABLE IF NOT EXISTS T_Mun (
		IDMun INTEGER,
		MunCode INTEGER,
		Désignation TEXT,
		MunName TEXT,
		MRC TEXT,
		RegAdm TEXT,
		AjoutListeMun TEXT
	);

	
`
