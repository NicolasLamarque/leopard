package schema

var TableEval = `
CREATE TABLE IF NOT EXISTS evaluations_sociales (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    contexte_evaluation TEXT,
    motif_reference TEXT,
    objet_evaluation TEXT,
    capacites_cognitives TEXT,
    etat_sante_physique TEXT,
    dimensions_psycho_sociales TEXT,
    roles_sociaux TEXT,
    reseau_social_soutien TEXT,
    analyse_clinique TEXT,
    opinion_professionnelle TEXT,
    recommandations TEXT,
	signature_nom TEXT,
    verrouille INTEGER DEFAULT 0,
    date_signature DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(client_id) REFERENCES clients(id),
    FOREIGN KEY(created_by) REFERENCES users(id)
);

CREATE VIEW IF NOT EXISTS v_evaluation_details AS
SELECT 
    e.*,
    c.nom AS client_nom,
    c.prenom AS client_prenom,
    c.date_naissance AS client_dn,
    c.numero_assurance_maladie AS client_nam,
    c.no_dossier_leopard AS client_leopard,
    u.full_name AS auteur_nom
FROM evaluations_sociales e
JOIN clients c ON e.client_id = c.id
JOIN users u ON e.created_by = u.id;
`
