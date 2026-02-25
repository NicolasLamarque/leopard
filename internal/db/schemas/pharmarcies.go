package schema

var TablePharmacies = `
CREATE TABLE IF NOT EXISTS pharmacies (
    "id"                    INTEGER PRIMARY KEY AUTOINCREMENT,
    "NomOrganisation"       TEXT,
    "Banniere"              TEXT,
    "Adresse"               TEXT,
    "Ville"                 TEXT,
    "Region"                TEXT,
    "Tel"                   TEXT,
    "Fax"                   TEXT,

    "DimancheOuvert"         INTEGER, 
    "HeureOuvertureDimanche" TEXT,
    "HeureFermetureDimanche" TEXT, 

    "LundiOuvert"            INTEGER, 
    "HeureOuvertureLundi"    TEXT,
    "HeureFermetureLundi"    TEXT,

    "MardiOuvert"            INTEGER,
    "HeureOuvertureMardi"    TEXT,
    "HeureFermetureMardi"    TEXT,
    

    "MercrediOuvert"         INTEGER,
    "HeureOuvertureMercredi" TEXT,
    "HeureFermetureMercredi" TEXT,

    "JeudiOuvert"            INTEGER,
    "HeureOuvertureJeudi"    TEXT,
    "HeureFermetureJeudi"    TEXT,
 
    "VendrediOuvert"         INTEGER,
    "HeureOuvertureVendredi" TEXT,
    "HeureFermetureVendredi" TEXT,

    "SamediOuvert"           INTEGER,
    "HeureOuvertureSamedi"   TEXT,
    "HeureFermetureSamedi"   TEXT,
    
    "DateMaj"                TEXT,
    "note"                   TEXT
);
`
