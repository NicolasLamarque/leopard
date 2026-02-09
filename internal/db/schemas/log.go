package schema

var Logs = `
   
    CREATE TABLE IF NOT EXISTS logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
        user_id INTEGER,           -- Lié à users.id
        action TEXT NOT NULL,      -- LOGIN, LOGOUT, CREATE, UPDATE, DELETE, VIEW
        module TEXT,               -- Appels, Clients, Factures
        target_id INTEGER,         -- L'ID de l'objet modifié
        severity TEXT DEFAULT 'INFO', -- INFO, WARN, CRITICAL
        
        -- Infos de Cohérence (Loi 25)
        machine_name TEXT,         -- os.Hostname()
        os_info TEXT,              -- runtime.GOOS
        ip_address TEXT,           -- IP locale
        
        -- Données brutes (JSON)
        payload TEXT,              -- Pour stocker l'ancien/nouveau contenu
        
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
    );

    CREATE VIEW IF NOT EXISTS v_audit_trail AS
    SELECT 
        l.id,
        l.timestamp,
        l.action,
        l.module,
        l.severity,
        l.machine_name,
        u.full_name AS intervenant,
        u.role AS intervenant_role,
        u.NoMembreOrdre,
        l.payload
    FROM logs l
    LEFT JOIN users u ON l.user_id = u.id;
`
