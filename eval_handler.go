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

func (h *EvalHandler) startup(ctx context.Context) {
	h.ctx = ctx
}

// CreateEvaluation crée une nouvelle évaluation sociale
func (h *EvalHandler) CreateEvaluation(req models.CreateEvaluationRequest) (int64, error) {
	// Récupérer l'ID de l'utilisateur connecté depuis le contexte ou une session
	// Ici je mets 1 par défaut, à adapter selon ta gestion de session
	userID := 1

	if req.ClientID == 0 {
		return 0, fmt.Errorf("ID client manquant")
	}

	id, err := h.db.CreateEvaluation(req, userID, h.cryptoSvc)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetEvaluationsByClient récupère la liste via la VUE (donc avec infos clients)
func (h *EvalHandler) GetEvaluationsByClient(clientID int) ([]models.EvaluationSocialeDetail, error) {
	return h.db.GetAllEvaluationsByClientID(clientID, h.cryptoSvc)
}

// GetEvaluationByID récupère une évaluation précise (via la VUE)
func (h *EvalHandler) GetEvaluationByID(id int) (*models.EvaluationSocialeDetail, error) {
	eval, err := h.db.GetEvaluationByID(id, h.cryptoSvc)
	if err != nil {
		return nil, err
	}
	return eval, nil
}

// SignerEvaluation verrouille le document pour l'OTSTCFQ
func (h *EvalHandler) SignerEvaluation(id int, nomSignature string) error {
	if nomSignature == "" {
		return fmt.Errorf("Le nom du signataire est obligatoire")
	}

	err := h.db.VerrouillerEvaluation(id, nomSignature)
	if err != nil {
		return fmt.Errorf("Erreur lors de la signature: %w", err)
	}

	return nil
}

// LockEvaluation verrouille une évaluation sociale
func (h *EvalHandler) LockEvaluation(evalID int, signatureNom string) error {
	return h.db.VerrouillerEvaluation(evalID, signatureNom)
}

// DeleteEvaluation supprime une évaluation sociale
func (h *EvalHandler) DeleteEvaluation(id int) error {
	return h.db.DeleteEvaluation(id)
}
