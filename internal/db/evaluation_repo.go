package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// GetEvaluationDefinitions : Match h.db.GetEvaluationDefinitions()
func (db *Database) GetEvaluationDefinitions() ([]models.EvaluationDefinition, error) {
	var defs []models.EvaluationDefinition
	query := `SELECT id, nom, icone, couleur, schema_json FROM evaluation_definitions WHERE active = 1`

	err := db.Select(&defs, query)
	if err != nil {
		return nil, fmt.Errorf("repo: erreur lecture definitions: %w", err)
	}
	return defs, nil
}

// CreateEvaluationV2 : Match h.db.CreateEvaluationV2(eval, h.cryptoSvc)
func (db *Database) CreateEvaluationV2(eval models.EvaluationV2, cryptoSvc *crypto.CryptoService) (int64, error) {
	// Loi 25 : On chiffre le payload JSON complet
	encrypted, err := cryptoSvc.EncryptString(eval.Payload)
	if err != nil {
		return 0, fmt.Errorf("repo: erreur crypto: %w", err)
	}

	query := `INSERT INTO evaluations_sociales_v2 
			  (client_id, model_id, no_leopard, payload_crypte, statut) 
			  VALUES (?, ?, ?, ?, 'brouillon')`

	res, err := db.Exec(query, eval.ClientID, eval.ModelID, eval.NoLeopard, encrypted)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetEvaluationsByClientIDV2 : Match h.db.GetEvaluationsByClientIDV2(clientID, h.cryptoSvc)
func (db *Database) GetEvaluationsByClientIDV2(clientID int, cryptoSvc *crypto.CryptoService) ([]models.EvaluationV2, error) {
	var list []models.EvaluationV2
	query := `SELECT id, client_id, model_id, no_leopard, payload_crypte, statut, created_at 
			  FROM evaluations_sociales_v2 WHERE client_id = ? ORDER BY created_at DESC`

	err := db.Select(&list, query, clientID)
	if err != nil {
		return nil, err
	}

	for i := range list {
		if list[i].PayloadCrypte != "" {
			decrypted, err := cryptoSvc.DecryptString(list[i].PayloadCrypte)
			if err == nil {
				list[i].Payload = decrypted
			}
		}
	}
	return list, nil
}

// UpdateEvaluationV2 : Match h.db.UpdateEvaluationV2(id, payload, h.cryptoSvc)
func (db *Database) UpdateEvaluationV2(id int, payload string, cryptoSvc *crypto.CryptoService) error {
	encrypted, err := cryptoSvc.EncryptString(payload)
	if err != nil {
		return err
	}

	query := `UPDATE evaluations_sociales_v2 SET payload_crypte = ?, updated_at = CURRENT_TIMESTAMP 
			  WHERE id = ? AND statut = 'brouillon'`

	_, err = db.Exec(query, encrypted, id)
	return err
}

// DeleteEvaluationV2 : Match h.db.DeleteEvaluationV2(id)
func (db *Database) DeleteEvaluationV2(id int) error {
	_, err := db.Exec("DELETE FROM evaluations_sociales_v2 WHERE id = ? AND statut = 'brouillon'", id)
	return err
}

// VerrouillerEvaluation : Scelle l'évaluation (V2) et ajoute la signature
func (db *Database) VerrouillerEvaluation(id int, nomSignature string) error {
	query := `
		UPDATE evaluations_sociales_v2 
		SET statut = 'verrouille', 
		    signature_nom = ?, 
		    date_signature = CURRENT_TIMESTAMP,
		    updated_at = CURRENT_TIMESTAMP 
		WHERE id = ?`

	_, err := db.Exec(query, nomSignature, id)
	if err != nil {
		return fmt.Errorf("repo: erreur lors du verrouillage: %w", err)
	}
	return nil
}

// GetEvaluationByIDV2 : Récupère une seule évaluation complète et décryptée
func (db *Database) GetEvaluationByIDV2(id int, cryptoSvc *crypto.CryptoService) (*models.EvaluationV2, error) {
	var eval models.EvaluationV2
	query := `SELECT * FROM evaluations_sociales_v2 WHERE id = ?`

	err := db.Get(&eval, query, id)
	if err != nil {
		return nil, fmt.Errorf("repo: évaluation introuvable: %w", err)
	}

	// Décryptage du payload pour le Front
	if eval.PayloadCrypte != "" {
		decrypted, err := cryptoSvc.DecryptString(eval.PayloadCrypte)
		if err == nil {
			eval.Payload = decrypted
		}
	}

	return &eval, nil
}

// SaveEvaluationDefinition : INSERT OR REPLACE d'un modèle
func (db *Database) SaveEvaluationDefinition(def models.EvaluationDefinition) error {
	query := `
		INSERT INTO evaluation_definitions (id, nom, icone, couleur, schema_json, active)
		VALUES (?, ?, ?, ?, ?, 1)
		ON CONFLICT(id) DO UPDATE SET
			nom        = excluded.nom,
			icone      = excluded.icone,
			couleur    = excluded.couleur,
			schema_json = excluded.schema_json,
			version    = version + 1
	`
	_, err := db.Exec(query,
		def.ID,
		def.Nom,
		def.Icone,
		def.Couleur,
		def.SchemaJSON,
	)
	if err != nil {
		return fmt.Errorf("repo: erreur sauvegarde definition: %w", err)
	}
	return nil
}

// DeleteEvaluationDefinition : Soft delete (active = 0)
func (db *Database) DeleteEvaluationDefinition(id string) error {
	_, err := db.Exec(
		`UPDATE evaluation_definitions SET active = 0 WHERE id = ? AND is_systeme = 0`,
		id,
	)
	return err
}
