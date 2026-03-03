package main

import (
	models "leopard/internal/model"
)

func (a *App) GetMunicipalities() ([]models.Municipality, error) {
	return a.db.GetAllMunicipalities()
}

func (a *App) GetMunicipalityByID(mcode string) (*models.Municipality, error) {
	return a.db.GetMunicipalityByID(mcode)
}

func (a *App) GetMunicipalityByName(name string) (*models.Municipality, error) {
	return a.db.GetMunicipalityByName(name)
}

func (a *App) GetArrondissements(mcode string) ([]models.Arrondissement, error) {
	arrs, err := a.db.GetArrondissementsByMun(mcode)
	if err != nil {
		return nil, err
	}
	if arrs == nil {
		return []models.Arrondissement{}, nil
	}
	return arrs, nil
}

func (a *App) ImportMunicipalitesCSV(filePath string) (int, error) {
	return a.db.ImportMunCSV(filePath)
}

func (a *App) ImportArrondissementsCSV(filePath string) (int, error) {
	return a.db.ImportArrCSV(filePath)
}
