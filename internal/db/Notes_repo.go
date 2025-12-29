package database

import (
	"fmt"
	models "leopard/internal/model"
	"time"
)

// CreateNote : Crée une note avec les paramètres nommés (lisible)
func (db *Database) CreateNote(req models.CreateNoteRequest, userID int, userName string) (int64, error) {
	// Conversion de la string "YYYY-MM-DD" en objet time.Time
	parsedDate, err := time.Parse("2006-01-02", req.DateIntervention)
	if err != nil {
		// Si la date est vide ou invalide, on met la date du jour par défaut
		parsedDate = time.Now()
	}

	query := `
        INSERT INTO notes (
            client_id, user_id, date_intervention, mode_intervention, 
            type_intervention, type_note, titre, contenu, 
            signature_nom, signature_date, verrouille
        ) VALUES (
            :client_id, :user_id, :date_intervention, :mode_intervention, 
            :type_intervention, :type_note, :sujet, :contenu, 
            :signature_nom, :signature_date, 1
        )`

	data := map[string]interface{}{
		"client_id":         req.ClientID,
		"user_id":           userID,
		"date_intervention": parsedDate, // On envoie l'objet Time parsé
		"mode_intervention": req.ModeIntervention,
		"type_intervention": req.TypeIntervention,
		"type_note":         req.TypeNote,
		"sujet":             req.Sujet,
		"contenu":           req.Contenu,
		"signature_nom":     userName,
		"signature_date":    time.Now(),
	}

	result, err := db.NamedExec(query, data)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// GetNoteByID : (Celle qui manquait !)
func (db *Database) GetNoteByID(id int) (*models.Note, error) {
	var note models.Note
	err := db.Get(&note, "SELECT * FROM notes WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// GetClientNotes : Récupère la liste allégée
func (db *Database) GetClientNotes(clientID int) ([]models.NoteListItem, error) {
	var notes []models.NoteListItem
	query := `SELECT id, type_note, titre, date_note, signature_nom, verrouille 
          FROM notes WHERE client_id = ? ORDER BY date_note DESC`
	err := db.Select(&notes, query, clientID)
	return notes, err
}

func (db *Database) GetClientNotesFiltered(filter models.NotesFilter) ([]models.NoteListItem, error) {
	var notes []models.NoteListItem
	query := `SELECT id, type_note, titre, date_note, signature_nom, verrouille 
              FROM notes 
              WHERE client_id = ? AND (titre LIKE ? OR contenu LIKE ?) 
              ORDER BY date_note DESC`

	search := "%" + filter.SearchQuery + "%"
	err := db.Select(&notes, query, filter.ClientID, search, search)
	return notes, err
}

// UpdateNote : (Interdit en clinique, mais nécessaire pour compiler)
func (db *Database) UpdateNote(req models.Note) error {
	return fmt.Errorf("modification impossible : créez une correction à la place")
}

// DeleteNote : Supprime si non verrouillé
func (db *Database) DeleteNote(id int) error {
	_, err := db.Exec("DELETE FROM notes WHERE id = ? AND verrouille = 0", id)
	return err
}

// LockNote : Verrouille manuellement
func (db *Database) LockNote(id int) error {
	_, err := db.Exec("UPDATE notes SET verrouille = 1 WHERE id = ?", id)
	return err
}

// GetNotesStats : Stats pour le drawer
func (db *Database) GetNotesStats(clientID int) (map[string]interface{}, error) {
	var total int
	err := db.Get(&total, "SELECT COUNT(*) FROM notes WHERE client_id = ?", clientID)
	return map[string]interface{}{"total": total}, err
}
