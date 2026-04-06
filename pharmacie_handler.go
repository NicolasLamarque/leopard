// ========================================
// WRAPPERS PHARMACIES POUR APP.GO
// ========================================

package main

import (
	"fmt"
	models "leopard/internal/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// -----------------------------------------------------------------------------
// READ
// -----------------------------------------------------------------------------

func (a *App) GetAllPharmacies() ([]models.Pharmacie, error) {
	return a.db.GetAllPharmacies()
}

func (a *App) GetPharmacieByID(id int) (*models.Pharmacie, error) {
	return a.db.GetPharmacieByID(id)
}

func (a *App) GetPharmacieForClient(clientID int) (*models.Pharmacie, error) {
	return a.db.GetPharmacieForClient(clientID)
}

func (a *App) GetClientsForPharmacie(pharmacieID int) ([]models.ClientPharmacieInfo, error) {
	return a.db.GetClientsByPharmacie(pharmacieID, a.cryptoSvc)
}

// -----------------------------------------------------------------------------
// CREATE / UPDATE
// -----------------------------------------------------------------------------

func (a *App) SavePharmacie(p models.Pharmacie) error {
	if p.ID == 0 {
		_, err := a.db.CreatePharmacie(p)
		return err
	}
	return a.db.UpdatePharmacie(p)
}

// -----------------------------------------------------------------------------
// DELETE
// -----------------------------------------------------------------------------

func (a *App) DeletePharmacie(id int) error {
	return a.db.DeletePharmacie(id)
}

// -----------------------------------------------------------------------------
// IMPORT CSV
// -----------------------------------------------------------------------------

// ImportPharmaciesCSVResult est le résultat retourné au frontend
type ImportPharmaciesCSVResult struct {
	Imported int    `json:"imported"`
	Skipped  int    `json:"skipped"`
	Message  string `json:"message"`
}

// ImportPharmaciesCSV ouvre un dialog de fichier natif puis importe le CSV.
// Retourne le résultat avec les compteurs ou une erreur.
func (a *App) ImportPharmaciesCSV() (*ImportPharmaciesCSVResult, error) {
	// 1. Dialog natif pour sélectionner le fichier
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Importer des pharmacies (CSV)",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Fichiers CSV (*.csv)",
				Pattern:     "*.csv",
			},
			{
				DisplayName: "Tous les fichiers",
				Pattern:     "*",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// L'utilisateur a annulé
	if filePath == "" {
		return nil, nil
	}

	// 2. Import
	imported, skipped, err := a.db.ImportPharmaciesFromCSV(filePath)
	if err != nil {
		return nil, err
	}

	return &ImportPharmaciesCSVResult{
		Imported: imported,
		Skipped:  skipped,
		Message:  formatImportMessage(imported, skipped),
	}, nil
}

func formatImportMessage(imported, skipped int) string {
	msg := ""
	if imported > 0 {
		msg += fmt.Sprintf("%d pharmacie(s) importée(s)", imported)
	}
	if skipped > 0 {
		if msg != "" {
			msg += ", "
		}
		msg += fmt.Sprintf("%d ligne(s) ignorée(s)", skipped)
	}
	if msg == "" {
		msg = "Aucune pharmacie importée"
	}
	return msg
}
