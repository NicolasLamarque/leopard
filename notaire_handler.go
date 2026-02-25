// folders.go
// Fichier optionnel pour séparer la logique des dossiers
// À placer dans le même package que app.go
package main

import (
	"errors"
	"fmt"
	models "leopard/internal/model"
	"leopard/internal/services/importer"
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

func (a *App) ImportNotairesWithUpdate(filePath string) (string, error) {
	// On redirige simplement vers la fonction de base
	return a.ImportNotaires(filePath)
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

// ImportNotaires déclenche l'importation via le nouveau moteur Pipeline
func (a *App) ImportNotaires(filePath string) (string, error) {
	if a.currentUser == nil {
		return "", errors.New("non authentifié")
	}

	// 1. On utilise le "Notaire complet" (ton nouveau dossier importer)
	// On passe le filePath reçu du frontend
	p := importer.NewPipeline(filePath)

	// 2. Lecture dynamique (Règle le bug Sheet1 car il prend sheets[0])
	rows, err := p.ReadExcel()
	if err != nil {
		return "", fmt.Errorf("erreur de lecture Excel : %w", err)
	}

	// 3. Transformation des données selon ton image (mappage des colonnes)
	notaires, err := p.ProcessNotaires(rows)
	if err != nil {
		return "", fmt.Errorf("erreur de traitement des données : %w", err)
	}

	// 4. Injection en base de données
	// On envoie la slice de NotaireRow à ton repository SQLite
	count, err := a.db.SaveNotaireList(notaires, a.currentUserID, a.cryptoSvc)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la sauvegarde en base : %w", err)
	}

	return fmt.Sprintf("✅ %d notaires importés avec succès", count), nil
}
