package main

import (
	"errors"
	"fmt"
	models "leopard/internal/model"
	"leopard/internal/services/importer"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ========== REF LISTES ==========

func (a *App) GetRefListeByCategorie(categorie string) ([]models.RefListe, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	if categorie == "" {
		return nil, errors.New("catégorie requise")
	}
	return a.db.GetRefListeByCategorie(categorie)
}

func (a *App) GetAllRefCategories() ([]string, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllRefCategories()
}

func (a *App) GetRefListeByID(id int) (*models.RefListe, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	if id == 0 {
		return nil, errors.New("ID invalide")
	}
	return a.db.GetRefListeByID(id)
}

func (a *App) CreateRefListe(req models.CreateRefListeRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}
	if req.Categorie == "" || req.Libelle == "" {
		return 0, errors.New("catégorie et libellé requis")
	}
	return a.db.CreateRefListe(req)
}

func (a *App) UpdateRefListe(id int, req models.CreateRefListeRequest) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	if id == 0 {
		return errors.New("ID invalide")
	}
	return a.db.UpdateRefListe(id, req)
}

func (a *App) DeleteRefListe(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	if id == 0 {
		return errors.New("ID invalide")
	}
	return a.db.DeleteRefListe(id)
}

func (a *App) UpdateRefListeOrdre(id int, ordre int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}
	return a.db.UpdateRefListeOrdre(id, ordre)
}

// OuvrirDialogCSVRefListes ouvre le sélecteur fichier (Wails v2) et lance l'import.
func (a *App) OuvrirDialogCSVRefListes() (string, error) {
	if a.currentUser == nil {
		return "", errors.New("non authentifié")
	}

	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Importer un fichier CSV - Listes de référence",
		Filters: []runtime.FileFilter{
			{DisplayName: "Fichiers CSV", Pattern: "*.csv"},
		},
	})
	if err != nil {
		return "", fmt.Errorf("erreur dialog: %w", err)
	}
	if filePath == "" {
		return "", nil
	}

	// Pipeline comme les notaires
	p := importer.NewPipeline(filePath)

	rows, err := p.ReadCSV()
	if err != nil {
		return "", fmt.Errorf("erreur lecture CSV: %w", err)
	}

	items, err := p.ProcessRefListes(rows)
	if err != nil {
		return "", fmt.Errorf("erreur traitement: %w", err)
	}

	count, ignored, err := a.db.SaveRefListeList(items, a.currentUserID)
	if err != nil {
		return "", fmt.Errorf("erreur sauvegarde: %w", err)
	}

	return fmt.Sprintf("✅ %d entrées importées, %d doublons ignorés", count, ignored), nil
}
