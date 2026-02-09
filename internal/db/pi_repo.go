package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllPlansByClientID récupère tous les plans d'un client via la vue v_pi_details
func (db *Database) GetAllPlansByClientID(clientID int, cryptoSvc *crypto.CryptoService) ([]models.PlanInterventionDetail, error) {
	var plans []models.PlanInterventionDetail

	// On utilise la vue définie dans db.go pour avoir les noms du client et de l'auteur
	query := `SELECT * FROM v_pi_details WHERE client_id = ? ORDER BY created_at DESC`

	err := db.Select(&plans, query, clientID)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération plans: %w", err)
	}

	// Déchiffrer les champs sensibles de chaque plan
	for i := range plans {
		if err := decryptPlan(&plans[i].PlanIntervention, cryptoSvc); err != nil {
			return nil, fmt.Errorf("erreur déchiffrement plan %d: %w", plans[i].ID, err)
		}
	}

	return plans, nil
}

// GetPlanByID récupère un plan spécifique par son ID
func (db *Database) GetPlanByID(id int, cryptoSvc *crypto.CryptoService) (*models.PlanInterventionDetail, error) {
	var plan models.PlanInterventionDetail
	query := `SELECT * FROM v_pi_details WHERE id = ?`

	err := db.Get(&plan, query, id)
	if err != nil {
		return nil, fmt.Errorf("plan %d non trouvé: %w", id, err)
	}

	if err := decryptPlan(&plan.PlanIntervention, cryptoSvc); err != nil {
		return nil, fmt.Errorf("erreur déchiffrement plan: %w", err)
	}

	return &plan, nil
}

// CreatePlan insère un nouveau plan en base (Statut par défaut: brouillon)
func (db *Database) CreatePlan(req models.CreatePlanRequest, createdBy int, cryptoSvc *crypto.CryptoService) (int64, error) {
	// On chiffre les données avant l'insertion
	encryptedReq, err := encryptPlanRequest(req, cryptoSvc)
	if err != nil {
		return 0, fmt.Errorf("erreur chiffrement: %w", err)
	}

	query := `
		INSERT INTO plans_intervention (
			client_id, created_by, titre, type_plan, statut,
			date_debut, date_fin_prevue, date_revision_prevue,
			contexte, problematique, forces, objectifs, 
			interventions, resultats, ententes
		) VALUES (
			:client_id, :created_by, :titre, :type_plan, 'brouillon',
			:date_debut, :date_fin_prevue, :date_revision_prevue,
			:contexte, :problematique, :forces, :objectifs,
			:interventions, :resultats, :ententes
		)`

	data := map[string]interface{}{
		"client_id":            req.ClientID,
		"created_by":           createdBy,
		"titre":                req.Titre,
		"type_plan":            req.TypePlan,
		"date_debut":           encryptedReq.DateDebut,
		"date_fin_prevue":      encryptedReq.DateFinPrevue,
		"date_revision_prevue": encryptedReq.DateRevisionPrevue,
		"contexte":             encryptedReq.Contexte,
		"problematique":        encryptedReq.Problematique,
		"forces":               encryptedReq.Forces,
		"objectifs":            encryptedReq.Objectifs,
		"interventions":        encryptedReq.Interventions,
		"resultats":            encryptedReq.Resultats,
		"ententes":             encryptedReq.Ententes,
	}

	result, err := db.NamedExec(query, data)
	if err != nil {
		return 0, fmt.Errorf("erreur création plan: %w", err)
	}

	return result.LastInsertId()
}

// UpdatePlan met à jour le plan SI il n'est pas verrouillé
func (db *Database) UpdatePlan(id int, req models.CreatePlanRequest, cryptoSvc *crypto.CryptoService) error {
	encryptedReq, err := encryptPlanRequest(req, cryptoSvc)
	if err != nil {
		return fmt.Errorf("erreur chiffrement: %w", err)
	}

	query := `
		UPDATE plans_intervention SET
			titre = :titre, type_plan = :type_plan,
			date_debut = :date_debut, date_fin_prevue = :date_fin_prevue, 
			date_revision_prevue = :date_revision_prevue,
			contexte = :contexte, problematique = :problematique, forces = :forces,
			objectifs = :objectifs, interventions = :interventions, 
			resultats = :resultats, ententes = :ententes,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = :id AND verrouille = 0`

	data := map[string]interface{}{
		"id":                   id,
		"titre":                req.Titre,
		"type_plan":            req.TypePlan,
		"date_debut":           encryptedReq.DateDebut,
		"date_fin_prevue":      encryptedReq.DateFinPrevue,
		"date_revision_prevue": encryptedReq.DateRevisionPrevue,
		"contexte":             encryptedReq.Contexte,
		"problematique":        encryptedReq.Problematique,
		"forces":               encryptedReq.Forces,
		"objectifs":            encryptedReq.Objectifs,
		"interventions":        encryptedReq.Interventions,
		"resultats":            encryptedReq.Resultats,
		"ententes":             encryptedReq.Ententes,
	}

	result, err := db.NamedExec(query, data)
	if err != nil {
		return fmt.Errorf("erreur mise à jour plan %d: %w", id, err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("plan %d non trouvé ou verrouillé", id)
	}

	return nil
}

// LockPlan verrouille définitivement le plan (Signature)
func (db *Database) LockPlan(id int, signatureNom string) error {
	query := `
		UPDATE plans_intervention 
		SET verrouille = 1, statut = 'actif', signature_nom = ?, date_signature = CURRENT_TIMESTAMP 
		WHERE id = ? AND verrouille = 0`

	result, err := db.Exec(query, signatureNom, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("plan déjà verrouillé ou inexistant")
	}
	return nil
}

// Helpers de Chiffrement interne
func decryptPlan(p *models.PlanIntervention, cryptoSvc *crypto.CryptoService) error {
	p.Contexte, _ = cryptoSvc.DecryptStringPtr(p.Contexte)
	p.Problematique, _ = cryptoSvc.DecryptStringPtr(p.Problematique)
	p.Forces, _ = cryptoSvc.DecryptStringPtr(p.Forces)
	p.Objectifs, _ = cryptoSvc.DecryptStringPtr(p.Objectifs)
	p.Interventions, _ = cryptoSvc.DecryptStringPtr(p.Interventions)
	p.Resultats, _ = cryptoSvc.DecryptStringPtr(p.Resultats)
	p.Ententes, _ = cryptoSvc.DecryptStringPtr(p.Ententes)
	return nil
}

func encryptPlanRequest(req models.CreatePlanRequest, cryptoSvc *crypto.CryptoService) (models.CreatePlanRequest, error) {
	enc := req
	enc.Contexte, _ = cryptoSvc.EncryptStringPtr(req.Contexte)
	enc.Problematique, _ = cryptoSvc.EncryptStringPtr(req.Problematique)
	enc.Forces, _ = cryptoSvc.EncryptStringPtr(req.Forces)
	enc.Objectifs, _ = cryptoSvc.EncryptStringPtr(req.Objectifs)
	enc.Interventions, _ = cryptoSvc.EncryptStringPtr(req.Interventions)
	enc.Resultats, _ = cryptoSvc.EncryptStringPtr(req.Resultats)
	enc.Ententes, _ = cryptoSvc.EncryptStringPtr(req.Ententes)
	return enc, nil
}
