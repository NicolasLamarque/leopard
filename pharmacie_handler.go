// ========================================
// WRAPPERS PHARMACIES POUR APP.GO
// À ajouter à la fin de ton fichier app.go
// ========================================

package main

import (
	models "leopard/internal/model"
)

// -----------------------------------------------------------------------------
// READ : Récupérer les infos
// -----------------------------------------------------------------------------

// GetAllPharmacies récupère toutes les pharmacies
func (a *App) GetAllPharmacies() ([]models.Pharmacie, error) {
	return a.db.GetAllPharmacies()
}

// GetPharmacieByID récupère une pharmacie par son ID
func (a *App) GetPharmacieByID(id int) (*models.Pharmacie, error) {
	return a.db.GetPharmacieByID(id)
}

// GetPharmacieForClient retourne la pharmacie assignée à un client
func (a *App) GetPharmacieForClient(clientID int) (*models.Pharmacie, error) {
	return a.db.GetPharmacieForClient(clientID)
}

// GetClientsForPharmacie - Pour ta fameuse liste "Clients en commun"
func (a *App) GetClientsForPharmacie(pharmacieID int) ([]models.ClientPharmacieInfo, error) {
	return a.db.GetClientsByPharmacie(pharmacieID, a.cryptoSvc)
}

// -----------------------------------------------------------------------------
// CREATE & UPDATE : Sauvegarder (Le fameux Save)
// -----------------------------------------------------------------------------

// SavePharmacie crée ou met à jour une pharmacie
func (a *App) SavePharmacie(p models.Pharmacie) error {
	// Si l'ID est 0, c'est une nouvelle pharmacie, sinon on met à jour
	if p.ID == 0 {
		_, err := a.db.CreatePharmacie(p)
		return err
	}
	return a.db.UpdatePharmacie(p)
}

// -----------------------------------------------------------------------------
// DELETE : Ménage (Attention, c'est définitif !)
// -----------------------------------------------------------------------------

// DeletePharmacie supprime une pharmacie
func (a *App) DeletePharmacie(id int) error {
	// Note: La fonction DeletePharmacie dans le repo met déjà
	// pharmacie_id à NULL pour tous les clients liés avant de supprimer
	return a.db.DeletePharmacie(id)
}
