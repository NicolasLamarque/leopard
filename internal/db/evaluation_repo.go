package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllEvaluationsByClientID récupère les évaluations via la VUE pour avoir les infos client
func (db *Database) GetAllEvaluationsByClientID(clientID int, cryptoSvc *crypto.CryptoService) ([]models.EvaluationSocialeDetail, error) {
	var list []models.EvaluationSocialeDetail

	// On interroge la VUE pour avoir le nom, prénom et NoLeopard sans effort
	query := `SELECT * FROM v_evaluation_details WHERE client_id = ? ORDER BY created_at DESC`

	err := db.Select(&list, query, clientID)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des évaluations: %w", err)
	}

	// Déchiffrement du contenu clinique
	for i := range list {
		if err := decryptEvaluationDetail(&list[i], cryptoSvc); err != nil {
			return nil, err
		}
	}

	return list, nil
}

// GetEvaluationByID récupère une évaluation spécifique via la VUE
func (db *Database) GetEvaluationByID(id int, cryptoSvc *crypto.CryptoService) (*models.EvaluationSocialeDetail, error) {
	var eval models.EvaluationSocialeDetail
	query := `SELECT * FROM v_evaluation_details WHERE id = ?`

	err := db.Get(&eval, query, id)
	if err != nil {
		return nil, fmt.Errorf("évaluation %d non trouvée: %w", id, err)
	}

	if err := decryptEvaluationDetail(&eval, cryptoSvc); err != nil {
		return nil, err
	}

	return &eval, nil
}

// CreateEvaluation crée l'évaluation dans la TABLE (pas la vue)
func (db *Database) CreateEvaluation(req models.CreateEvaluationRequest, createdBy int, cryptoSvc *crypto.CryptoService) (int64, error) {
	// Chiffrement des blocs de texte clinique
	encryptedReq, err := encryptEvaluationRequest(req, cryptoSvc)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO evaluations_sociales (
			client_id, created_by, contexte_evaluation, motif_reference, 
			objet_evaluation, capacites_cognitives, etat_sante_physique, 
			dimensions_psycho_sociales, roles_sociaux, reseau_social_soutien, 
			analyse_clinique, opinion_professionnelle, recommandations
		) VALUES (
			:client_id, :created_by, :contexte_evaluation, :motif_reference, 
			:objet_evaluation, :capacites_cognitives, :etat_sante_physique, 
			:dimensions_psycho_sociales, :roles_sociaux, :reseau_social_soutien, 
			:analyse_clinique, :opinion_professionnelle, :recommandations
		)`

	data := map[string]interface{}{
		"client_id":                  req.ClientID,
		"created_by":                 createdBy,
		"contexte_evaluation":        encryptedReq.ContexteEvaluation,
		"motif_reference":            encryptedReq.MotifReference,
		"objet_evaluation":           encryptedReq.ObjetEvaluation,
		"capacites_cognitives":       encryptedReq.CapacitesCognitives,
		"etat_sante_physique":        encryptedReq.EtatSantePhysique,
		"dimensions_psycho_sociales": encryptedReq.DimensionsPsychoSociales,
		"roles_sociaux":              encryptedReq.RolesSociaux,
		"reseau_social_soutien":      encryptedReq.ReseauSocialSoutien,
		"analyse_clinique":           encryptedReq.AnalyseClinique,
		"opinion_professionnelle":    encryptedReq.OpinionProfessionnelle,
		"recommandations":            encryptedReq.Recommandations,
	}

	result, err := db.NamedExec(query, data)
	if err != nil {
		return 0, fmt.Errorf("erreur insertion évaluation: %w", err)
	}

	return result.LastInsertId()
}

func (db *Database) UpdateEvaluationBrouillon(id int, req models.CreateEvaluationRequest, cryptoSvc *crypto.CryptoService) error {
	// ON UTILISE TON SYSTÈME : On crypte avant d'envoyer à la DB
	enc, err := encryptEvaluationRequest(req, cryptoSvc)
	if err != nil {
		return err
	}

	query := `
		UPDATE evaluations_sociales SET 
			contexte_evaluation = ?, motif_reference = ?, objet_evaluation = ?,
			capacites_cognitives = ?, etat_sante_physique = ?, dimensions_psycho_sociales = ?,
			roles_sociaux = ?, reseau_social_soutien = ?, analyse_clinique = ?,
			opinion_professionnelle = ?, recommandations = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND verrouille = 0` // Sécurité : on ne modifie JAMAIS si c'est déjà verrouillé

	_, err = db.Exec(query,
		enc.ContexteEvaluation, enc.MotifReference, enc.ObjetEvaluation,
		enc.CapacitesCognitives, enc.EtatSantePhysique, enc.DimensionsPsychoSociales,
		enc.RolesSociaux, enc.ReseauSocialSoutien, enc.AnalyseClinique,
		enc.OpinionProfessionnelle, enc.Recommandations,
		id)
	return err
}

// VerrouillerEvaluation signe et verrouille l'évaluation (non modifiable après)
func (db *Database) VerrouillerEvaluation(id int, signatureNom string) error {
	query := `UPDATE evaluations_sociales SET verrouille = 1, signature_nom = ?, date_signature = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := db.Exec(query, signatureNom, id)
	return err
}

// --- Helpers de chiffrement ---

func decryptEvaluationDetail(e *models.EvaluationSocialeDetail, cryptoSvc *crypto.CryptoService) error {
	var err error
	e.AnalyseClinique, _ = cryptoSvc.DecryptStringPtr(e.AnalyseClinique)
	e.OpinionProfessionnelle, _ = cryptoSvc.DecryptStringPtr(e.OpinionProfessionnelle)
	e.MotifReference, _ = cryptoSvc.DecryptStringPtr(e.MotifReference)
	e.CapacitesCognitives, _ = cryptoSvc.DecryptStringPtr(e.CapacitesCognitives)
	e.EtatSantePhysique, _ = cryptoSvc.DecryptStringPtr(e.EtatSantePhysique)
	e.DimensionsPsychoSociales, _ = cryptoSvc.DecryptStringPtr(e.DimensionsPsychoSociales)
	e.RolesSociaux, _ = cryptoSvc.DecryptStringPtr(e.RolesSociaux)
	e.ReseauSocialSoutien, _ = cryptoSvc.DecryptStringPtr(e.ReseauSocialSoutien)
	e.Recommandations, _ = cryptoSvc.DecryptStringPtr(e.Recommandations)
	e.ContexteEvaluation, _ = cryptoSvc.DecryptStringPtr(e.ContexteEvaluation)
	e.ObjetEvaluation, _ = cryptoSvc.DecryptStringPtr(e.ObjetEvaluation)
	return err
}

func encryptEvaluationRequest(req models.CreateEvaluationRequest, cryptoSvc *crypto.CryptoService) (models.CreateEvaluationRequest, error) {
	enc := req
	// On crypte chaque champ individuellement
	enc.ContexteEvaluation, _ = cryptoSvc.EncryptStringPtr(req.ContexteEvaluation)
	enc.MotifReference, _ = cryptoSvc.EncryptStringPtr(req.MotifReference)
	enc.ObjetEvaluation, _ = cryptoSvc.EncryptStringPtr(req.ObjetEvaluation)
	enc.CapacitesCognitives, _ = cryptoSvc.EncryptStringPtr(req.CapacitesCognitives)
	enc.EtatSantePhysique, _ = cryptoSvc.EncryptStringPtr(req.EtatSantePhysique)
	enc.DimensionsPsychoSociales, _ = cryptoSvc.EncryptStringPtr(req.DimensionsPsychoSociales)
	enc.RolesSociaux, _ = cryptoSvc.EncryptStringPtr(req.RolesSociaux)
	enc.ReseauSocialSoutien, _ = cryptoSvc.EncryptStringPtr(req.ReseauSocialSoutien)
	enc.AnalyseClinique, _ = cryptoSvc.EncryptStringPtr(req.AnalyseClinique)
	enc.OpinionProfessionnelle, _ = cryptoSvc.EncryptStringPtr(req.OpinionProfessionnelle)
	enc.Recommandations, _ = cryptoSvc.EncryptStringPtr(req.Recommandations)
	return enc, nil
}

func (db *Database) DeleteEvaluation(id int) error {
	// Vérifier que c'est un brouillon
	var verrouille int
	err := db.Get(&verrouille, "SELECT verrouille FROM evaluations_sociales WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("évaluation non trouvée: %w", err)
	}

	if verrouille == 1 {
		return fmt.Errorf("impossible de supprimer une évaluation verrouillée")
	}

	// Supprimer
	_, err = db.Exec("DELETE FROM evaluations_sociales WHERE id = ?", id)
	return err
}
