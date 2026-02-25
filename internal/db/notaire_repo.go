package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
	"leopard/internal/services/importer"

	"github.com/xuri/excelize/v2"
)

// ==========================================
//          LOGIQUE DE RÉCUPÉRATION
// ==========================================

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

// ==========================================
//          LOGIQUE DE MODIFICATION (CRUD)
// ==========================================

func (db *Database) CreateNotaire(req models.CreateNotaireRequest, cryptoSvc *crypto.CryptoService) (int64, error) {
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
		"telecopieur":      req.Telecopieur,
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

func (db *Database) UpdateNotaire(n models.Notaire, cryptoSvc *crypto.CryptoService) error {
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
	return err
}

func (db *Database) ArchiveNotaire(id int) error {
	query := `UPDATE notaires SET actif = 0 WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

func (db *Database) DeleteNotaire(id int) error {
	query := `DELETE FROM notaires WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// ==========================================
//          LOGIQUE D'IMPORTATION (EXCEL)
// ==========================================

func (db *Database) ImportNotairesFromExcel(path string, userID int, cryptoSvc *crypto.CryptoService) (int64, error) {
	p := importer.NewPipeline(path)
	rows, err := p.ReadExcel()
	if err != nil {
		return 0, err
	}
	notaires, err := p.ProcessNotaires(rows)
	if err != nil {
		return 0, err
	}
	return db.SaveNotaireList(notaires, userID, cryptoSvc)
}

func (db *Database) SaveNotaireList(notaires []importer.NotaireRow, userID int, cryptoSvc *crypto.CryptoService) (int64, error) {
	var count int64 = 0
	for _, n := range notaires {
		req := models.CreateNotaireRequest{
			Civilite:        n.Civilite,
			Prenom:          n.Prenom,
			Nom:             n.Nom,
			Telephone:       stringToPtr(n.Telephone),
			Cellulaire:      stringToPtr(n.Cellulaire),
			Telecopieur:     stringToPtr(n.Telecopieur),
			Email:           stringToPtr(n.Email),
			Adresse:         stringToPtr(n.Adresse),
			Ville:           stringToPtr(n.Ville),
			CodePostal:      stringToPtr(n.CodePostal),
			SecteurActivite: stringToPtr(n.SecteurActivite),
			NoteFixe:        stringToPtr(n.Note),
			CreatedBy:       userID,
		}
		_, err := db.CreateNotaire(req, cryptoSvc)
		if err == nil {
			count++
		}
	}
	return count, nil
}

// ==========================================
//                HELPERS
// ==========================================

func decryptNotaire(n *models.Notaire, cryptoSvc *crypto.CryptoService) {
	n.Email, _ = cryptoSvc.DecryptStringPtr(n.Email)
	n.Telephone, _ = cryptoSvc.DecryptStringPtr(n.Telephone)
	n.Cellulaire, _ = cryptoSvc.DecryptStringPtr(n.Cellulaire)
	n.Adresse, _ = cryptoSvc.DecryptStringPtr(n.Adresse)
	n.NoteFixe, _ = cryptoSvc.DecryptStringPtr(n.NoteFixe)
}

func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (db *Database) CreateNotaireTemplate(path string) error {
	f := excelize.NewFile()
	defer f.Close()
	f.SetCellValue("Sheet1", "A1", "Civilite") // Mis à jour pour ton image
	f.SetCellValue("Sheet1", "B1", "Prenom")
	f.SetCellValue("Sheet1", "C1", "Nom")
	return f.SaveAs(path)
}

func (db *Database) GetClientsByNotaire(notaireID int, cryptoSvc *crypto.CryptoService) ([]models.Client, error) {
	var clients []models.Client
	query := `SELECT * FROM clients WHERE notaire_id = ? AND actif = 1`

	err := db.Select(&clients, query, notaireID)
	if err != nil {
		return nil, err
	}

	// Si tu n'as pas encore de logique de décryptage ici,
	// on utilise le blanc souligné "_" pour dire à Go d'ignorer l'index
	for _ = range clients {
		// decryptClient(&clients[i], cryptoSvc) // À décommenter quand tu seras prêt
	}

	return clients, nil
}
