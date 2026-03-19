package schema

var tablePremiereNation = `
CREATE TABLE IF NOT EXISTS ref_nations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nation_nom TEXT,       -- ex: "Atikamekw", "Innu"
    communaute_nom TEXT,   -- ex: "Manawan", "Mashteuiatsh"
    famille_linguistique TEXT,      -- ex: "Algonquienne", "Iroquoienne", "Eskimo-aléoute"
    logo_path TEXT,                 -- Pour afficher le blason de la bande si disponible
    site_web TEXT,                  -- Lien vers le conseil de bande
    mcode TEXT,                     -- Lien vers votre table MUN (municipalités)
    actif INTEGER DEFAULT 1
);

CREATE TABLE ref_communautes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nation_id INTEGER,
    communaute_nom TEXT,
    logo_path TEXT,
    mcode TEXT,
    site_web TEXT,
    actif INTEGER DEFAULT 1,
);

`
