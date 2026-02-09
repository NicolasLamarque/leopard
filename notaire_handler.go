// folders.go
// Fichier optionnel pour séparer la logique des dossiers
// À placer dans le même package que app.go
package main

import (
	"errors"
	"fmt"
	models "leopard/internal/model"
)

// ========== NOTAIRES ==========
// À ajouter dans app.go après la section CONTACTS

// GetAllNotaires retourne tous les notaires actifs
func (a *App) GetAllNotaires() ([]models.Notaire, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetAllNotaires(a.cryptoSvc)
}

// GetNotaireByID retourne un notaire spécifique par son ID
func (a *App) GetNotaireByID(id int) (*models.Notaire, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}
	return a.db.GetNotaireByID(id, a.cryptoSvc)
}

// CreateNotaire crée un nouveau notaire
func (a *App) CreateNotaire(req models.CreateNotaireRequest) (int64, error) {
	if a.currentUser == nil {
		return 0, errors.New("non authentifié")
	}

	// Validation basique
	if req.Prenom == "" || req.Nom == "" {
		return 0, errors.New("le prénom et le nom sont requis")
	}

	// Définir le created_by avec l'utilisateur actuel
	req.CreatedBy = int(a.currentUser.ID)

	return a.db.CreateNotaire(req, a.cryptoSvc)
}

// UpdateNotaire met à jour un notaire existant
func (a *App) UpdateNotaire(notaire models.Notaire) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if notaire.ID == 0 {
		return errors.New("ID notaire invalide")
	}

	return a.db.UpdateNotaire(notaire, a.cryptoSvc)
}

// DeleteNotaire supprime définitivement un notaire
func (a *App) DeleteNotaire(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if id == 0 {
		return errors.New("ID notaire invalide")
	}

	return a.db.DeleteNotaire(id)
}

// ArchiveNotaire désactive un notaire (soft delete)
func (a *App) ArchiveNotaire(id int) error {
	if a.currentUser == nil {
		return errors.New("non authentifié")
	}

	if id == 0 {
		return errors.New("ID notaire invalide")
	}

	return a.db.ArchiveNotaire(id)
}

// GetClientsByNotaire retourne tous les clients liés à un notaire
func (a *App) GetClientsByNotaire(notaireID int) ([]models.Client, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}

	if notaireID == 0 {
		return nil, errors.New("ID notaire invalide")
	}

	return a.db.GetClientsByNotaire(notaireID, a.cryptoSvc)
}

// ImportNotaires déclenche l'importation Excel pour les notaires
func (a *App) ImportNotaires(filePath string) (string, error) {
	// On récupère l'ID de l'utilisateur actuel (à adapter selon ta gestion de session)
	createdBy := 1

	// On appelle la fonction que tu as écrite dans internal/db/notaire_import.go
	count, err := a.db.ImportNotairesFromExcel(filePath, createdBy, a.cryptoSvc)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d notaires ont été importés avec succès", count), nil
}
