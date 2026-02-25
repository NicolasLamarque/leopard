package main

import (
	models "leopard/internal/model"
)

/// ========== PAYS HANDLERS ==========

// GetPaysCount retourne le nombre de pays en base
func (a *App) GetPaysCount() (int, error) {
	return a.db.GetPaysCount()
}

// GetAllPays retourne tous les pays
func (a *App) GetAllPays() ([]models.Pays, error) {
	return a.db.GetAllPays()
}

// GetPaysByID retourne un pays par son ID
func (a *App) GetPaysByID(id int) (*models.Pays, error) {
	return a.db.GetPaysByID(id)
}

// GetPaysByAlpha2 retourne un pays par son code Alpha-2
func (a *App) GetPaysByAlpha2(alpha2 string) (*models.Pays, error) {
	return a.db.GetPaysByAlpha2(alpha2)
}

// GetPaysByAlpha3 retourne un pays par son code Alpha-3
func (a *App) GetPaysByAlpha3(alpha3 string) (*models.Pays, error) {
	return a.db.GetPaysByAlpha3(alpha3)
}

// SearchPays recherche des pays par terme
func (a *App) SearchPays(searchTerm string) ([]models.Pays, error) {
	return a.db.SearchPays(searchTerm)
}

// GetPaysForSelect retourne les pays pour un select dropdown
func (a *App) GetPaysForSelect() ([]models.PaysListItemForSelect, error) {
	return a.db.GetPaysForSelect()
}
