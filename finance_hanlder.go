// finance_handler.go
// Wrappers Wails pour le module Finances — séparé de app.go
// Pattern identique à notaire_handler.go
package main

import (
	"errors"
	models "leopard/internal/model"
)

// ═══════════════════════════════════════════════════════════════════════════════
// PARAMÈTRES FINANCE
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetParametresFinance() (*models.ParametresFinance, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetParametresFinance(a.cryptoSvc)
}

func (a *App) SaveParametresFinance(p models.ParametresFinance) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.SaveParametresFinance(p, a.cryptoSvc)
}

// ═══════════════════════════════════════════════════════════════════════════════
// SERVICES — catalogue des prestations
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetAllServices(categorie string, actifSeulement bool) ([]models.Service, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllServices(categorie, actifSeulement)
}

func (a *App) GetServiceByID(id int) (*models.Service, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetServiceByID(id)
}

func (a *App) CreateService(req models.CreateServiceRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateService(req, int(a.currentUser.ID))
}

func (a *App) UpdateService(req models.UpdateServiceRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateService(req)
}

func (a *App) ArchiverService(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.ArchiverService(id)
}

func (a *App) DeleteService(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteService(id)
}

// ═══════════════════════════════════════════════════════════════════════════════
// REVENUS
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetAllRevenus(dateDebut, dateFin, statut, categorie string) ([]models.Revenu, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllRevenus(a.cryptoSvc, dateDebut, dateFin, statut, categorie)
}

func (a *App) GetRevenusByClient(clientID int) ([]models.Revenu, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetRevenusByClient(clientID, a.cryptoSvc)
}

func (a *App) CreateRevenu(req models.CreateRevenuRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateRevenu(req, int(a.currentUser.ID), a.cryptoSvc)
}

func (a *App) UpdateRevenu(req models.UpdateRevenuRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateRevenu(req, a.cryptoSvc)
}

func (a *App) DeleteRevenu(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteRevenu(id)
}

// ═══════════════════════════════════════════════════════════════════════════════
// DÉPENSES
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetAllDepenses(dateDebut, dateFin, categorie string, deductibleSeulement bool) ([]models.Depense, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllDepenses(dateDebut, dateFin, categorie, deductibleSeulement)
}

func (a *App) CreateDepense(req models.CreateDepenseRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateDepense(req, int(a.currentUser.ID))
}

func (a *App) UpdateDepense(req models.UpdateDepenseRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateDepense(req)
}

func (a *App) DeleteDepense(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteDepense(id)
}

// ═══════════════════════════════════════════════════════════════════════════════
// FACTURES
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetAllFactures(statut, dateDebut, dateFin string) ([]models.Facture, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllFactures(a.cryptoSvc, statut, dateDebut, dateFin)
}

func (a *App) GetFactureByID(id int) (*models.Facture, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetFactureByID(id, a.cryptoSvc)
}

func (a *App) CreateFacture(req models.CreateFactureRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateFacture(req, int(a.currentUser.ID), a.cryptoSvc)
}

func (a *App) UpdateFactureStatut(id int, statut string) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateFactureStatut(id, statut)
}

func (a *App) EnregistrerPaiement(req models.CreatePaiementRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.EnregistrerPaiement(req, a.cryptoSvc)
}

func (a *App) VerifierFacturesEnRetard() error {
	return a.db.VerifierFacturesEnRetard()
}

// ═══════════════════════════════════════════════════════════════════════════════
// CONTRATS
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetAllContrats(clientID int, statut string) ([]models.Contrat, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllContrats(clientID, statut)
}

func (a *App) GetContratByID(id int) (*models.Contrat, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetContratByID(id)
}

func (a *App) CreateContrat(req models.CreateContratRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateContrat(req, int(a.currentUser.ID))
}

func (a *App) UpdateContrat(req models.UpdateContratRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateContrat(req)
}

func (a *App) SignerContrat(id int, dateSignature string) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.SignerContrat(id, dateSignature)
}

// ═══════════════════════════════════════════════════════════════════════════════
// STATS & RAPPORTS — zéro crypto, tout en SQL clair
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetDashboardData() (*models.DashboardData, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetDashboardData()
}

func (a *App) GetStatsMensuelles(annee, mois int) (*models.StatsMensuelles, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetStatsMensuelles(annee, mois)
}

func (a *App) GetRapportFiscalAnnuel(annee int) (*models.RapportFiscalAnnuel, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetRapportFiscalAnnuel(annee)
}

func (a *App) GetRentabiliteParClient(annee int) ([]models.RentabiliteParClient, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetRentabiliteParClient(a.cryptoSvc, annee)
}
