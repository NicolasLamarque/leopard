package schema

var TableMunicipalites = `CREATE TABLE IF NOT EXISTS MUN (
    mcode     TEXT PRIMARY KEY,
    munnom    TEXT NOT NULL,
    madr1     TEXT,
    mcodpos   TEXT,
    mtel      TEXT,
    mfax      TEXT,
    mcourriel TEXT,
    mweb      TEXT,
    regadm    TEXT,
    mrc       TEXT,
    maire     TEXT,
    dirgen    TEXT,
    msuperf   REAL,
    mpopul    INTEGER,
    polic     TEXT,
    incen     TEXT,
    loisir    TEXT,
    trvpub    TEXT,
    mesurg    TEXT,
    urban     TEXT,
    communic  TEXT,
    permis    TEXT,
    batim     TEXT
);`

var TableSecteurs = `CREATE TABLE IF NOT EXISTS ARR (
    arrcod  TEXT PRIMARY KEY,
    arrnom  TEXT NOT NULL,
    arrcode TEXT NOT NULL,
    FOREIGN KEY(arrcode) REFERENCES MUN(mcode)
);`
