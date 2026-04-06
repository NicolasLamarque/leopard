package main

import (
	"context"
	"fmt"
	"leopard/internal/crypto"
	database "leopard/internal/db"
	models "leopard/internal/model"
)

type EvalHandler struct {
	ctx       context.Context
	db        *database.Database
	cryptoSvc *crypto.CryptoService
}

func NewEvalHandler(db *database.Database, cryptoSvc *crypto.CryptoService) *EvalHandler {
	return &EvalHandler{
		db:        db,
		cryptoSvc: cryptoSvc,
	}
}

func (h *EvalHandler) Startup(ctx context.Context) {
	h.ctx = ctx
}

// --- MÉTHODES V2 (Nouveau Système JSON) ---

// GetDefinitions retourne les modèles de formulaires disponibles (Annexe, Curateur, etc.)
func (h *EvalHandler) GetDefinitions() ([]models.EvaluationDefinition, error) {
	return h.db.GetEvaluationDefinitions()
}

// CreateEvaluationV2 crée une évaluation à partir du gros JSON de Vue
func (h *EvalHandler) CreateEvaluationV2(eval models.EvaluationV2) (int64, error) {
	if eval.ClientID == 0 {
		return 0, fmt.Errorf("ID client manquant")
	}
	// On passe l'objet au repo qui s'occupera du chiffrement
	return h.db.CreateEvaluationV2(eval, h.cryptoSvc)
}

// GetClientEvaluationsV2 récupère toutes les évals d'un client (décryptées)
func (h *EvalHandler) GetClientEvaluationsV2(clientID int) ([]models.EvaluationV2, error) {
	return h.db.GetEvaluationsByClientIDV2(clientID, h.cryptoSvc)
}

// UpdateEvaluationV2 met à jour le payload JSON d'un brouillon
func (h *EvalHandler) UpdateEvaluationV2(id int, payload string) error {
	return h.db.UpdateEvaluationV2(id, payload, h.cryptoSvc)
}

// DeleteEvaluationV2 supprime un brouillon V2
func (h *EvalHandler) DeleteEvaluationV2(id int) error {
	return h.db.DeleteEvaluationV2(id)
}

func (h *EvalHandler) SignerEvaluation(id int, nomSignature string) error {
	if nomSignature == "" {
		return fmt.Errorf("le nom du signataire est obligatoire")
	}
	// On appelle la fonction du Repo qu'on vient d'ajouter
	return h.db.VerrouillerEvaluation(id, nomSignature)
}

// GetEvaluationByID : Pour charger une éval spécifique (ex: pour le PDF ou l'édition)
func (h *EvalHandler) GetEvaluationByID(id int) (*models.EvaluationV2, error) {
	return h.db.GetEvaluationByIDV2(id, h.cryptoSvc)
}

func (h *EvalHandler) SaveDefinition(def models.EvaluationDefinition) error {
	return h.db.SaveEvaluationDefinition(def)
}

// DeleteDefinition : Désactive un modèle (soft delete)
func (h *EvalHandler) DeleteDefinition(id string) error {
	return h.db.DeleteEvaluationDefinition(id)
}
