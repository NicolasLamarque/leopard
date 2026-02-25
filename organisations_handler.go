// organisations_handler.go
package main

import (
	"errors"
	models "leopard/internal/model"
)

// ═══════════════════════════════════════════════════════════════════════════════
// ORGANISATIONS
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) GetAllOrganisations(typeOrg string, actifSeulement bool) ([]models.Organisation, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllOrganisations(a.cryptoSvc, typeOrg, actifSeulement)
}

func (a *App) GetOrganisationsListItems() ([]models.OrganisationListItem, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetOrganisationsListItems(a.cryptoSvc)
}

func (a *App) GetOrganisationByID(id int) (*models.Organisation, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetOrganisationByID(id, a.cryptoSvc)
}

func (a *App) CreateOrganisation(req models.CreateOrganisationRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateOrganisation(req, int(a.currentUser.ID), a.cryptoSvc)
}

func (a *App) UpdateOrganisation(req models.UpdateOrganisationRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateOrganisation(req, a.cryptoSvc)
}

func (a *App) ArchiverOrganisation(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.ArchiverOrganisation(id)
}

func (a *App) DeleteOrganisation(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteOrganisation(id)
}

// ═══════════════════════════════════════════════════════════════════════════════
// PAYEURS
// ═══════════════════════════════════════════════════════════════════════════════

func (a *App) CreatePayeur(req models.CreatePayeurRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreatePayeur(req, int(a.currentUser.ID), a.cryptoSvc)
}

func (a *App) GetPayeursByClient(clientID int) ([]models.PayeurResolu, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetPayeursByClient(clientID, a.cryptoSvc)
}

func (a *App) ResoudrePayeur(payeurID int) (*models.PayeurResolu, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.ResoudrePayeur(payeurID, a.cryptoSvc)
}
