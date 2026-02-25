package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetAllEvaluationsByClientID récupère les évaluations via la VUE
func (db *Database) GetAllEvaluationsByClientID(clientID int, cryptoSvc *crypto.CryptoService) ([]models.EvaluationSocialeDetail, error) {
	var list []models.EvaluationSocialeDetail

	query := `SELECT * FROM v_evaluation_details WHERE client_id = ? ORDER BY created_at DESC`

	err := db.Select(&list, query, clientID)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des évaluations: %w", err)
	}

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

// CreateEvaluation crée l'évaluation dans la TABLE
// no_leopard et type_evaluation sont en clair — pas de chiffrement
func (db *Database) CreateEvaluation(req models.CreateEvaluationRequest, createdBy int, cryptoSvc *crypto.CryptoService) (int64, error) {
	encryptedReq, err := encryptEvaluationRequest(req, cryptoSvc)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO evaluations_sociales (
			client_id, created_by, no_leopard, no_eval_leopard, type_evaluation,
			contexte_evaluation, motif_reference, objet_evaluation,
			capacites_cognitives, etat_sante_physique, dimensions_psycho_sociales,
			roles_sociaux, reseau_social_soutien, analyse_clinique,
			opinion_professionnelle, recommandations, verrouille, isDraft
		) VALUES (
			:client_id, :created_by, :no_leopard, :no_eval_leopard, :type_evaluation,
			:contexte_evaluation, :motif_reference, :objet_evaluation,
			:capacites_cognitives, :etat_sante_physique, :dimensions_psycho_sociales,
			:roles_sociaux, :reseau_social_soutien, :analyse_clinique,
			:opinion_professionnelle, :recommandations, :verrouille, :isDraft
		)`

	data := map[string]interface{}{
		"client_id":                  req.ClientID,
		"created_by":                 createdBy,
		"no_leopard":                 req.NoLeopard,
		"no_eval_leopard":            req.NoEvalLeopard,  // EN CLAIR
		"type_evaluation":            req.TypeEvaluation, // EN CLAIR
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
		"verrouille":                 0,
		"isDraft":                    1,
	}

	result, err := db.NamedExec(query, data)
	if err != nil {
		return 0, fmt.Errorf("erreur insertion évaluation: %w", err)
	}

	return result.LastInsertId()
}

func (db *Database) UpdateEvaluationBrouillon(id int, req models.CreateEvaluationRequest, cryptoSvc *crypto.CryptoService) error {
	enc, err := encryptEvaluationRequest(req, cryptoSvc)
	if err != nil {
		return err
	}

	query := `
        UPDATE evaluations_sociales SET
            type_evaluation            = :type_evaluation,
            contexte_evaluation        = :contexte_evaluation,
            motif_reference            = :motif_reference,
            objet_evaluation           = :objet_evaluation,
            capacites_cognitives       = :capacites_cognitives,
            etat_sante_physique        = :etat_sante_physique,
            dimensions_psycho_sociales = :dimensions_psycho_sociales,
            roles_sociaux              = :roles_sociaux,
            reseau_social_soutien      = :reseau_social_soutien,
            analyse_clinique           = :analyse_clinique,
            opinion_professionnelle    = :opinion_professionnelle,
            recommandations            = :recommandations,
            updated_at                 = CURRENT_TIMESTAMP
        WHERE id = :id AND verrouille = 0`

	data := map[string]interface{}{
		"id":                         id,
		"type_evaluation":            req.TypeEvaluation, // EN CLAIR
		"contexte_evaluation":        enc.ContexteEvaluation,
		"motif_reference":            enc.MotifReference,
		"objet_evaluation":           enc.ObjetEvaluation,
		"capacites_cognitives":       enc.CapacitesCognitives,
		"etat_sante_physique":        enc.EtatSantePhysique,
		"dimensions_psycho_sociales": enc.DimensionsPsychoSociales,
		"roles_sociaux":              enc.RolesSociaux,
		"reseau_social_soutien":      enc.ReseauSocialSoutien,
		"analyse_clinique":           enc.AnalyseClinique,
		"opinion_professionnelle":    enc.OpinionProfessionnelle,
		"recommandations":            enc.Recommandations,
	}

	_, err = db.NamedExec(query, data)
	return err
}

// VerrouillerEvaluation signe et verrouille — non modifiable après
func (db *Database) VerrouillerEvaluation(id int, signatureNom string) error {
	query := `
        UPDATE evaluations_sociales SET
            verrouille     = 1,
            isDraft        = 0,
            signature_nom  = :signature_nom,
            date_signature = CURRENT_TIMESTAMP
        WHERE id = :id AND verrouille = 0`

	data := map[string]interface{}{
		"id":            id,
		"signature_nom": signatureNom,
	}

	_, err := db.NamedExec(query, data)
	return err
}

// DeleteEvaluation supprime un brouillon seulement
func (db *Database) DeleteEvaluation(id int) error {
	var verrouille int
	err := db.Get(&verrouille, "SELECT verrouille FROM evaluations_sociales WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("évaluation non trouvée: %w", err)
	}

	if verrouille == 1 {
		return fmt.Errorf("impossible de supprimer une évaluation verrouillée")
	}

	_, err = db.Exec("DELETE FROM evaluations_sociales WHERE id = ?", id)
	return err
}

// --- Helpers chiffrement ---

func decryptEvaluationDetail(e *models.EvaluationSocialeDetail, cryptoSvc *crypto.CryptoService) error {
	// no_leopard et type_evaluation : EN CLAIR, on ne déchiffre pas
	e.ContexteEvaluation, _ = cryptoSvc.DecryptStringPtr(e.ContexteEvaluation)
	e.MotifReference, _ = cryptoSvc.DecryptStringPtr(e.MotifReference)
	e.ObjetEvaluation, _ = cryptoSvc.DecryptStringPtr(e.ObjetEvaluation)
	e.CapacitesCognitives, _ = cryptoSvc.DecryptStringPtr(e.CapacitesCognitives)
	e.EtatSantePhysique, _ = cryptoSvc.DecryptStringPtr(e.EtatSantePhysique)
	e.DimensionsPsychoSociales, _ = cryptoSvc.DecryptStringPtr(e.DimensionsPsychoSociales)
	e.RolesSociaux, _ = cryptoSvc.DecryptStringPtr(e.RolesSociaux)
	e.ReseauSocialSoutien, _ = cryptoSvc.DecryptStringPtr(e.ReseauSocialSoutien)
	e.AnalyseClinique, _ = cryptoSvc.DecryptStringPtr(e.AnalyseClinique)
	e.OpinionProfessionnelle, _ = cryptoSvc.DecryptStringPtr(e.OpinionProfessionnelle)
	e.Recommandations, _ = cryptoSvc.DecryptStringPtr(e.Recommandations)
	return nil
}

func encryptEvaluationRequest(req models.CreateEvaluationRequest, cryptoSvc *crypto.CryptoService) (models.CreateEvaluationRequest, error) {
	enc := req
	// no_leopard et type_evaluation restent EN CLAIR
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
