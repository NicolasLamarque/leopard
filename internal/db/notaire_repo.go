package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"

	"github.com/xuri/excelize/v2"
)

// GetAllNotaires récupère tous les notaires actifs
func (db *Database) GetAllNotaires(cryptoSvc *crypto.CryptoService) ([]models.Notaire, error) {
	var liste []models.Notaire
	query := `SELECT * FROM notaires WHERE actif = 1 ORDER BY nom, prenom`

	err := db.Select(&liste, query)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération notaires: %w", err)
	}

	for i := range liste {
		decryptNotaire(&liste[i], cryptoSvc)
	}
	return liste, nil
}

// CreateNotaire insère un nouveau notaire
func (db *Database) CreateNotaire(req models.CreateNotaireRequest, cryptoSvc *crypto.CryptoService) (int64, error) {
	// Chiffrement des données sensibles
	encryptedEmail, _ := cryptoSvc.EncryptStringPtr(req.Email)
	encryptedTel, _ := cryptoSvc.EncryptStringPtr(req.Telephone)
	encryptedCell, _ := cryptoSvc.EncryptStringPtr(req.Cellulaire)
	encryptedAddr, _ := cryptoSvc.EncryptStringPtr(req.Adresse)

	query := `
		INSERT INTO notaires (
			civilite, prenom, nom, telephone, cellulaire, telecopieur,
			adresse, code_postal, ville, email, secteur_activite, 
			note_fixe, created_by
		) VALUES (
			:civilite, :prenom, :nom, :telephone, :cellulaire, :telecopieur,
			:adresse, :code_postal, :ville, :email, :secteur_activite, 
			:note_fixe, :created_by
		)`

	data := map[string]interface{}{
		"civilite":         req.Civilite,
		"prenom":           req.Prenom,
		"nom":              req.Nom,
		"telephone":        encryptedTel,
		"cellulaire":       encryptedCell,
		"telecopieur":      req.Telecopieur, // Souvent laissé en clair ou chiffré selon besoin
		"adresse":          encryptedAddr,
		"code_postal":      req.CodePostal,
		"ville":            req.Ville,
		"email":            encryptedEmail,
		"secteur_activite": req.SecteurActivite,
		"note_fixe":        req.NoteFixe,
		"created_by":       req.CreatedBy,
	}

	resultat, err := db.NamedExec(query, data)
	if err != nil {
		return 0, fmt.Errorf("erreur création notaire: %w", err)
	}

	return resultat.LastInsertId()
}

// --- Helpers de Chiffrement (privés) ---

func decryptNotaire(n *models.Notaire, cryptoSvc *crypto.CryptoService) {
	n.Email, _ = cryptoSvc.DecryptStringPtr(n.Email)
	n.Telephone, _ = cryptoSvc.DecryptStringPtr(n.Telephone)
	n.Cellulaire, _ = cryptoSvc.DecryptStringPtr(n.Cellulaire)
	n.Adresse, _ = cryptoSvc.DecryptStringPtr(n.Adresse)
	n.NoteFixe, _ = cryptoSvc.DecryptStringPtr(n.NoteFixe)
}

// UpdateNotaire met à jour les informations d'un notaire existant
func (db *Database) UpdateNotaire(n models.Notaire, cryptoSvc *crypto.CryptoService) error {
	// Chiffrement des données avant la mise à jour (similaire à UpdateClient)
	encryptedEmail, _ := cryptoSvc.EncryptStringPtr(n.Email)
	encryptedTel, _ := cryptoSvc.EncryptStringPtr(n.Telephone)
	encryptedCell, _ := cryptoSvc.EncryptStringPtr(n.Cellulaire)
	encryptedAddr, _ := cryptoSvc.EncryptStringPtr(n.Adresse)
	encryptedNote, _ := cryptoSvc.EncryptStringPtr(n.NoteFixe)

	query := `
		UPDATE notaires SET
			civilite = :civilite, prenom = :prenom, nom = :nom, 
			telephone = :telephone, cellulaire = :cellulaire, telecopieur = :telecopieur,
			adresse = :adresse, code_postal = :code_postal, ville = :ville, 
			email = :email, secteur_activite = :secteur_activite, 
			note_fixe = :note_fixe, actif = :actif, 
			updated_at = CURRENT_TIMESTAMP
		WHERE id = :id`

	data := map[string]interface{}{
		"id":               n.ID,
		"civilite":         n.Civilite,
		"prenom":           n.Prenom,
		"nom":              n.Nom,
		"telephone":        encryptedTel,
		"cellulaire":       encryptedCell,
		"telecopieur":      n.Telecopieur,
		"adresse":          encryptedAddr,
		"code_postal":      n.CodePostal,
		"ville":            n.Ville,
		"email":            encryptedEmail,
		"secteur_activite": n.SecteurActivite,
		"note_fixe":        encryptedNote,
		"actif":            n.Actif,
	}

	_, err := db.NamedExec(query, data)
	if err != nil {
		return fmt.Errorf("erreur mise à jour notaire %d: %w", n.ID, err)
	}
	return nil
}

// GetNotaireByID récupère un notaire unique par son ID et déchiffre ses données
func (db *Database) GetNotaireByID(id int, cryptoSvc *crypto.CryptoService) (*models.Notaire, error) {
	var n models.Notaire
	query := `SELECT * FROM notaires WHERE id = ?`

	err := db.Get(&n, query, id)
	if err != nil {
		return nil, fmt.Errorf("notaire avec l'ID %d non trouvé: %w", id, err)
	}

	decryptNotaire(&n, cryptoSvc)
	return &n, nil
}

// ArchiveNotaire désactive un notaire (soft delete)
func (db *Database) ArchiveNotaire(id int) error {
	query := `UPDATE notaires SET actif = 0 WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// DeleteNotaire supprime définitivement un notaire de la base
func (db *Database) DeleteNotaire(id int) error {
	query := `DELETE FROM notaires WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
func (db *Database) GetClientsByNotaire(notaireID int, cryptoSvc *crypto.CryptoService) ([]models.Client, error) {
	var clients []models.Client
	// On cherche les clients qui ont le notaire_id correspondant
	query := `SELECT * FROM clients WHERE notaire_id = ? AND actif = 1`

	err := db.Select(&clients, query, notaireID)
	if err != nil {
		return nil, err
	}

	// N'oublie pas de déchiffrer chaque client comme dans client_repo.go
	for i := range clients {
		decryptClient(&clients[i], cryptoSvc)
	}
	return clients, nil
}

// ImportNotairesFromExcel lit le fichier Excel et insère les données en DB
// On ajoute (userID et cryptoSvc) pour matcher l'appel du handler
// ImportNotairesFromExcel lit le fichier Excel et insère les données en DB
// ImportNotairesFromExcel - La version originale qui fonctionnait
func (db *Database) ImportNotairesFromExcel(path string, userID int, cryptoSvc *crypto.CryptoService) (int64, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return 0, fmt.Errorf("impossible d'ouvrir l'excel: %w", err)
	}
	defer f.Close()

	// On essaie de lire "Sheet1" comme avant
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return 0, fmt.Errorf("impossible de lire la feuille: %w", err)
	}

	var count int64 = 0
	for i, row := range rows {
		// On skip l'entête et on s'assure d'avoir les 3 colonnes (Prénom, Nom, Ville)
		if i == 0 || len(row) < 3 {
			continue
		}

		query := `INSERT INTO notaires (prenom, nom, ville, created_by, actif) VALUES (?, ?, ?, ?, 1)`
		_, err := db.Exec(query, row[0], row[1], row[2], userID)
		if err == nil {
			count++
		}
	}

	return count, nil
}

// Et voici la fonction pour créer le modèle si il manque
func (db *Database) CreateNotaireTemplate(path string) error {
	f := excelize.NewFile()
	defer f.Close()

	// Créer l'entête
	f.SetCellValue("Sheet1", "A1", "Prénom")
	f.SetCellValue("Sheet1", "B1", "Nom")
	f.SetCellValue("Sheet1", "C1", "Ville")

	if err := f.SaveAs(path); err != nil {
		return err
	}
	return nil
}
