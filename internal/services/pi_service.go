package services

import (
	"errors"
	"leopard/internal/crypto"
	database "leopard/internal/db"
	models "leopard/internal/model"
)

// PIService centralise la logique des Plans d'Intervention
type PIService struct {
	db        *database.Database
	cryptoSvc *crypto.CryptoService
}

// NewPIService initialise le service
func NewPIService(db *database.Database, cryptoSvc *crypto.CryptoService) *PIService {
	return &PIService{
		db:        db,
		cryptoSvc: cryptoSvc,
	}
}

// GetPlans récupère et déchiffre les plans
func (s *PIService) GetPlans(clientID int) ([]models.PlanInterventionDetail, error) {
	if clientID <= 0 {
		return nil, errors.New("ID client invalide")
	}
	return s.db.GetAllPlansByClientID(clientID, s.cryptoSvc)
}

// Create crée un plan avec audit trail
func (s *PIService) Create(req models.CreatePlanRequest, userID int64) (int64, error) {
	if req.Titre == "" {
		return 0, errors.New("le titre est obligatoire")
	}
	// On passe l'ID de l'utilisateur (travailleur social) pour la Loi 25
	return s.db.CreatePlan(req, int(userID), s.cryptoSvc)
}

// Lock scelle le plan (action irréversible)
func (s *PIService) Lock(planID int, signature string) error {
	plan, err := s.db.GetPlanByID(planID, s.cryptoSvc)
	if err != nil {
		return err
	}
	if plan.Verrouille == 1 {
		return errors.New("ce plan est déjà scellé")
	}
	return s.db.LockPlan(planID, signature)
}

// pi_service.go

// / pi_service.go
// pi_service.go
// pi_service.go

// On définit l'ordre : 1. l'ID, 2. la Structure de données
func (s *PIService) Update(id int, req models.CreatePlanRequest) error {
	// 1. Vérification de sécurité avant modif
	plan, err := s.db.GetPlanByID(id, s.cryptoSvc)
	if err != nil {
		return err
	}
	if plan.Verrouille == 1 {
		return errors.New("ce plan est scellé et ne peut plus être modifié")
	}

	// 2. On appelle la DB avec l'ordre exact qu'elle demande :
	// (id, req, cryptoSvc)
	return s.db.UpdatePlan(id, req, s.cryptoSvc)
}
