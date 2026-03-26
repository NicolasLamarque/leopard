// cim11_handler.go
// Handlers Wails pour le module CIM-11
// package main — même package que app.go
package main

import (
	"context"
	"errors"
	"time"

	models "leopard/internal/model"
	"leopard/internal/services"
)

// ── LECTURE ───────────────────────────────────────────────────────

func (a *App) SearchCim11(query string, lang string, limit int) ([]models.Cim11Code, error) {
	if query == "" {
		return []models.Cim11Code{}, nil
	}
	if lang == "" {
		lang = "fr"
	}
	return a.db.SearchCim11(query, lang, limit)
}

func (a *App) GetCim11ByCode(code string) (*models.Cim11Code, error) {
	return a.db.GetCim11ByCode(code)
}

func (a *App) GetCim11Chapitres() ([]models.Cim11Code, error) {
	return a.db.GetCim11Chapitres()
}

func (a *App) GetCim11Enfants(parentCode string) ([]models.Cim11Code, error) {
	return a.db.GetCim11Enfants(parentCode)
}

func (a *App) GetCim11Page(query, chapitreCode, lang string, limit, offset int) (*Cim11PageResult, error) {
	result, err := a.db.GetCim11Page(query, chapitreCode, lang, limit, offset)
	if err != nil {
		return nil, err
	}
	return &Cim11PageResult{Codes: result.Codes, Total: result.Total}, nil
}

func (a *App) GetCim11DetailOMS(code string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"definition": "Enrichissement OMS disponible prochainement.",
	}, nil
}

// ── CRUD MANUEL ───────────────────────────────────────────────────

func (a *App) CreateCim11Code(c models.Cim11Code) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	return a.db.CreateCim11Code(c)
}

func (a *App) UpdateCim11Code(c models.Cim11Code) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateCim11Code(c)
}

func (a *App) DeleteCim11Code(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.DeleteCim11Code(id)
}

// ── SYNC OMS ──────────────────────────────────────────────────────

func (a *App) GetCim11SyncMeta() (*models.Cim11SyncMeta, error) {
	return a.db.GetCim11SyncMeta()
}

func (a *App) ImportCim11() (*models.Cim11ImportResult, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	svc := services.NewCim11ImportService(a.db)
	ctx, cancel := context.WithTimeout(a.ctx, 30*time.Minute)
	defer cancel()
	return svc.ImportAll(ctx)
}

// ── Types ─────────────────────────────────────────────────────────

type Cim11PageResult struct {
	Codes []models.Cim11Code `json:"codes"`
	Total int                `json:"total"`
}
