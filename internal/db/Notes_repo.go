package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"

	"github.com/jmoiron/sqlx"
)

// ========== RÉCUPÉRATION DES NOTES ==========

// GetAllNotesByClientID récupère toutes les notes d'un client (liste légère pour sidebar)
func (db *Database) GetAllNotesByClientID(clientID int, cryptoSvc *crypto.CryptoService) ([]models.NoteListItem, error) {
	var notes []models.NoteListItem

	query := `
		SELECT 
			id, date_note, date_intervention, type_note, titre, 
			verrouille, note_tardive, type_lien
		FROM notes 
		WHERE client_id = ? 
		ORDER BY date_note DESC, created_at DESC
	`

	err := db.Select(&notes, query, clientID)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération notes client %d: %w", clientID, err)
	}

	// Déchiffrer uniquement le titre pour l'affichage
	for i := range notes {
		if notes[i].Titre != "" {
			decrypted, err := cryptoSvc.Decrypt(notes[i].Titre)
			if err != nil {
				// Si erreur de déchiffrement, on garde le titre original
				continue
			}
			notes[i].Titre = decrypted
		}
	}

	return notes, nil
}

// GetNoteByID récupère une note complète avec déchiffrement
func (db *Database) GetNoteByID(noteID int, cryptoSvc *crypto.CryptoService) (*models.Note, error) {
	var note models.Note

	query := `SELECT * FROM notes WHERE id = ?`

	err := db.Get(&note, query, noteID)
	if err != nil {
		return nil, fmt.Errorf("note ID %d non trouvée: %w", noteID, err)
	}

	// Déchiffrer toutes les données sensibles
	if err := decryptNote(&note, cryptoSvc); err != nil {
		return nil, fmt.Errorf("erreur déchiffrement note: %w", err)
	}

	return &note, nil
}

// GetNotesByIDs récupère plusieurs notes par leurs IDs (pour export PDF)
func (db *Database) GetNotesByIDs(noteIDs []int, cryptoSvc *crypto.CryptoService) ([]models.Note, error) {
	if len(noteIDs) == 0 {
		return []models.Note{}, nil
	}

	var notes []models.Note

	// Construire la query avec sqlx.In
	query := `SELECT * FROM notes WHERE id IN (?) ORDER BY date_note ASC, created_at ASC`

	// Convertir []int en []interface{}
	args := make([]interface{}, len(noteIDs))
	for i, id := range noteIDs {
		args[i] = id
	}

	// Utiliser sqlx.In pour gérer le IN clause
	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erreur construction requête: %w", err)
	}

	// Rebind pour le driver SQLite
	query = db.Rebind(query)

	err = db.Select(&notes, query, args...)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération notes: %w", err)
	}

	// Déchiffrer toutes les notes
	for i := range notes {
		if err := decryptNote(&notes[i], cryptoSvc); err != nil {
			return nil, fmt.Errorf("erreur déchiffrement note %d: %w", notes[i].ID, err)
		}
	}

	return notes, nil
}

// ========== CRÉATION ET MODIFICATION ==========

// CreateNote crée une nouvelle note (brouillon) avec chiffrement
func (db *Database) CreateNote(req models.CreateNoteRequest, cryptoSvc *crypto.CryptoService) (int64, error) {
	// Chiffrer les données sensibles
	encryptedReq, err := encryptCreateNoteRequest(req, cryptoSvc)
	if err != nil {
		return 0, fmt.Errorf("erreur chiffrement: %w", err)
	}

	query := `
		INSERT INTO notes (
			client_id, client_NoRAMQ, client_Nom, client_Prenom, 
			client_Telephone, client_Cellulaire, client_NoLeopard,
			client_Adresse, client_appartement, client_code_postal, 
			client_ville, client_pays, client_province,
			user_id, date_intervention, heure_intervention, duree_intervention,
			mode_intervention, type_intervention, type_note,
			titre, contenu, note_tardive, note_de_tier,
			note_liee_id, type_lien, verrouille
		) VALUES (
			:client_id, :client_NoRAMQ, :client_Nom, :client_Prenom,
			:client_Telephone, :client_Cellulaire, :client_NoLeopard,
			:client_Adresse, :client_appartement, :client_code_postal,
			:client_ville, :client_pays, :client_province,
			:user_id, :date_intervention, :heure_intervention, :duree_intervention,
			:mode_intervention, :type_intervention, :type_note,
			:titre, :contenu, :note_tardive, :note_de_tier,
			:note_liee_id, :type_lien, 0
		)
	`

	result, err := db.NamedExec(query, encryptedReq)
	if err != nil {
		return 0, fmt.Errorf("erreur création note: %w", err)
	}

	return result.LastInsertId()
}

// UpdateNote met à jour un brouillon existant (seulement si verrouille = 0)
func (db *Database) UpdateNote(req models.UpdateNoteRequest, cryptoSvc *crypto.CryptoService) error {
	// Vérifier que la note est bien un brouillon
	var verrouille int
	err := db.Get(&verrouille, "SELECT verrouille FROM notes WHERE id = ?", req.ID)
	if err != nil {
		return fmt.Errorf("note ID %d non trouvée: %w", req.ID, err)
	}

	if verrouille == 1 {
		return fmt.Errorf("impossible de modifier une note verrouillée")
	}

	// Chiffrer les données modifiables
	encryptedReq, err := encryptUpdateNoteRequest(req, cryptoSvc)
	if err != nil {
		return fmt.Errorf("erreur chiffrement: %w", err)
	}

	query := `
		UPDATE notes SET
			date_intervention = :date_intervention,
			heure_intervention = :heure_intervention,
			duree_intervention = :duree_intervention,
			mode_intervention = :mode_intervention,
			type_intervention = :type_intervention,
			type_note = :type_note,
			titre = :titre,
			contenu = :contenu,
			note_tardive = :note_tardive,
			note_de_tier = :note_de_tier,
			note_liee_id = :note_liee_id,
			type_lien = :type_lien
		WHERE id = :id AND verrouille = 0
	`

	_, err = db.NamedExec(query, encryptedReq)
	if err != nil {
		return fmt.Errorf("erreur mise à jour note %d: %w", req.ID, err)
	}

	return nil
}

// LockNote verrouille et signe une note (action irréversible)
func (db *Database) LockNote(noteID int, signatureNom string, cryptoSvc *crypto.CryptoService) error {
	// Vérifier que la note existe et n'est pas déjà verrouillée
	var verrouille int
	err := db.Get(&verrouille, "SELECT verrouille FROM notes WHERE id = ?", noteID)
	if err != nil {
		return fmt.Errorf("note ID %d non trouvée: %w", noteID, err)
	}

	if verrouille == 1 {
		return fmt.Errorf("la note est déjà verrouillée")
	}

	// Chiffrer le nom de signature
	encryptedSignature, err := cryptoSvc.Encrypt(signatureNom)
	if err != nil {
		return fmt.Errorf("erreur chiffrement signature: %w", err)
	}

	query := `
		UPDATE notes SET
			verrouille = 1,
			signature_nom = ?,
			signature_date = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err = db.Exec(query, encryptedSignature, noteID)
	if err != nil {
		return fmt.Errorf("erreur verrouillage note %d: %w", noteID, err)
	}

	return nil
}

// DeleteNote supprime définitivement une note (seulement si brouillon)
func (db *Database) DeleteNote(noteID int) error {
	// Vérifier que c'est un brouillon
	var verrouille int
	err := db.Get(&verrouille, "SELECT verrouille FROM notes WHERE id = ?", noteID)
	if err != nil {
		return fmt.Errorf("note ID %d non trouvée: %w", noteID, err)
	}

	if verrouille == 1 {
		return fmt.Errorf("impossible de supprimer une note verrouillée")
	}

	query := `DELETE FROM notes WHERE id = ? AND verrouille = 0`
	_, err = db.Exec(query, noteID)
	if err != nil {
		return fmt.Errorf("erreur suppression note %d: %w", noteID, err)
	}

	return nil
}

// ========== HELPERS DE CHIFFREMENT/DÉCHIFFREMENT ==========

// decryptNote déchiffre tous les champs sensibles d'une note
func decryptNote(note *models.Note, cryptoSvc *crypto.CryptoService) error {
	var err error

	// Champs client (peuvent être nil)
	note.ClientNoRAMQ, _ = cryptoSvc.DecryptStringPtr(note.ClientNoRAMQ)
	note.ClientNom, _ = cryptoSvc.DecryptStringPtr(note.ClientNom)
	note.ClientPrenom, _ = cryptoSvc.DecryptStringPtr(note.ClientPrenom)
	note.ClientTelephone, _ = cryptoSvc.DecryptStringPtr(note.ClientTelephone)
	note.ClientCellulaire, _ = cryptoSvc.DecryptStringPtr(note.ClientCellulaire)
	note.ClientNoLeopard, _ = cryptoSvc.DecryptStringPtr(note.ClientNoLeopard)
	note.ClientAdresse, _ = cryptoSvc.DecryptStringPtr(note.ClientAdresse)
	note.ClientAppartement, _ = cryptoSvc.DecryptStringPtr(note.ClientAppartement)
	note.ClientCodePostal, _ = cryptoSvc.DecryptStringPtr(note.ClientCodePostal)
	note.ClientVille, _ = cryptoSvc.DecryptStringPtr(note.ClientVille)
	note.ClientPays, _ = cryptoSvc.DecryptStringPtr(note.ClientPays)
	note.ClientProvince, _ = cryptoSvc.DecryptStringPtr(note.ClientProvince)

	// Champs de la note
	note.Titre, _ = cryptoSvc.DecryptStringPtr(note.Titre)
	note.Contenu, _ = cryptoSvc.DecryptStringPtr(note.Contenu)
	note.SignatureNom, _ = cryptoSvc.DecryptStringPtr(note.SignatureNom)

	return err
}

// encryptCreateNoteRequest chiffre les données lors de la création
func encryptCreateNoteRequest(req models.CreateNoteRequest, cryptoSvc *crypto.CryptoService) (models.CreateNoteRequest, error) {
	encrypted := req
	var err error

	// Champs client
	encrypted.ClientNoRAMQ, err = cryptoSvc.EncryptStringPtr(req.ClientNoRAMQ)
	if err != nil {
		return req, err
	}

	encrypted.ClientNom, err = cryptoSvc.EncryptStringPtr(req.ClientNom)
	if err != nil {
		return req, err
	}

	encrypted.ClientPrenom, err = cryptoSvc.EncryptStringPtr(req.ClientPrenom)
	if err != nil {
		return req, err
	}

	encrypted.ClientTelephone, err = cryptoSvc.EncryptStringPtr(req.ClientTelephone)
	if err != nil {
		return req, err
	}

	encrypted.ClientCellulaire, err = cryptoSvc.EncryptStringPtr(req.ClientCellulaire)
	if err != nil {
		return req, err
	}

	encrypted.ClientNoLeopard, err = cryptoSvc.EncryptStringPtr(req.ClientNoLeopard)
	if err != nil {
		return req, err
	}

	encrypted.ClientAdresse, err = cryptoSvc.EncryptStringPtr(req.ClientAdresse)
	if err != nil {
		return req, err
	}

	encrypted.ClientAppartement, err = cryptoSvc.EncryptStringPtr(req.ClientAppartement)
	if err != nil {
		return req, err
	}

	encrypted.ClientCodePostal, err = cryptoSvc.EncryptStringPtr(req.ClientCodePostal)
	if err != nil {
		return req, err
	}

	encrypted.ClientVille, err = cryptoSvc.EncryptStringPtr(req.ClientVille)
	if err != nil {
		return req, err
	}

	encrypted.ClientPays, err = cryptoSvc.EncryptStringPtr(req.ClientPays)
	if err != nil {
		return req, err
	}

	encrypted.ClientProvince, err = cryptoSvc.EncryptStringPtr(req.ClientProvince)
	if err != nil {
		return req, err
	}

	// Champs de la note
	encrypted.Titre, err = cryptoSvc.EncryptStringPtr(req.Titre)
	if err != nil {
		return req, err
	}

	encrypted.Contenu, err = cryptoSvc.EncryptStringPtr(req.Contenu)
	if err != nil {
		return req, err
	}

	return encrypted, nil
}

// encryptUpdateNoteRequest chiffre les données lors de la mise à jour
func encryptUpdateNoteRequest(req models.UpdateNoteRequest, cryptoSvc *crypto.CryptoService) (models.UpdateNoteRequest, error) {
	encrypted := req
	var err error

	encrypted.Titre, err = cryptoSvc.EncryptStringPtr(req.Titre)
	if err != nil {
		return req, err
	}

	encrypted.Contenu, err = cryptoSvc.EncryptStringPtr(req.Contenu)
	if err != nil {
		return req, err
	}

	return encrypted, nil
}
