package schema

// Tables de référence et données statiques
var TableMedecins = `
	CREATE TABLE IF NOT EXISTS medecins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		licence TEXT NOT NULL UNIQUE,
		civilite TEXT,
		prenom TEXT,
		nom TEXT,
		nomComplet TEXT,
		statut TEXT,
		specialisation TEXT,
		service TEXT,
		departement TEXT,
		installation_principale TEXT,
		rls TEXT,
		telephone TEXT,
		extension TEXT,
		cellulaire TEXT,
		email TEXT,
		adresse TEXT,
		code_postal TEXT,
		ville TEXT,
		pays TEXT DEFAULT 'Canada',
		Note_fixe TEXT,
		actif INTEGER DEFAULT 1,
		created_by INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
`
