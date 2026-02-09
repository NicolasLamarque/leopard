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

	CREATE TABLE IF NOT EXISTS "T_pays" (
		"ID_Pays" INTEGER,
		"PAYS" TEXT,
		"Alpha2" TEXT,
		"Alpha3" TEXT,
		"numerique" INTEGER
	);

	

	
`
