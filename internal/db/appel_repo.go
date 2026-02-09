package database

import (
	"fmt"
	"time"

	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// AppelRepository gère les opérations sur les appels
type AppelRepository struct {
	db *Database
}

// NewAppelRepository crée un nouveau repository
func NewAppelRepository(db *Database) *AppelRepository {
	return &AppelRepository{db: db}
}

// Create crée un nouvel appel avec cryptage
// ========== CRÉATION ==========

// Create crée un nouvel appel avec cryptage et log d'audit
func (r *AppelRepository) Create(req models.CreateAppelRequest, recuParUserID int, cryptoSvc *crypto.CryptoService) (*models.Appel, error) {
	// 1. Crypter les données sensibles
	appelantNom, _ := cryptoSvc.Encrypt(req.AppelantNom)
	appelantPrenom, _ := cryptoSvc.Encrypt(req.AppelantPrenom)
	appelantTelephone, _ := cryptoSvc.Encrypt(req.AppelantTelephone)
	prospectNom, _ := cryptoSvc.Encrypt(req.ProspectNom)
	prospectPrenom, _ := cryptoSvc.Encrypt(req.ProspectPrenom)
	prospectTelephone, _ := cryptoSvc.Encrypt(req.ProspectTelephone)
	motifAppel, _ := cryptoSvc.Encrypt(req.MotifAppel)
	notesInternes, _ := cryptoSvc.Encrypt(req.NotesInternes)

	// 2. Gérer les IDs optionnels (Pointeurs *int) pour éviter l'erreur sur le "0"
	var cID interface{} = nil
	if req.ClientID != nil && *req.ClientID > 0 {
		cID = *req.ClientID
	}

	var aA interface{} = nil
	if req.AssigneA != nil && *req.AssigneA > 0 {
		aA = *req.AssigneA
	}

	// 3. Préparer les dates
	dateAppel := time.Now()
	if req.DateAppel != "" {
		if parsed, err := time.Parse("2006-01-02", req.DateAppel); err == nil {
			dateAppel = parsed
		}
	}

	var rdvDate interface{} = nil
	if req.RDVDate != "" {
		rdvDate = req.RDVDate
	}

	query := `
		INSERT INTO appels (
			date_appel, heure_appel,
			appelant_nom, appelant_prenom, appelant_telephone, appelant_relation,
			client_id,
			prospect_nom, prospect_prenom, prospect_telephone,
			type_demande, motif_appel, priorite,
			statut, notes_internes,
			rdv_date, rdv_heure, rdv_lieu,
			recu_par, assigne_a,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`

	// 4. Insertion en base de données
	result, err := r.db.Exec(query,
		dateAppel, req.HeureAppel,
		appelantNom, appelantPrenom, appelantTelephone, req.AppelantRelation,
		cID, // Utilise NULL si 0 ou nil
		prospectNom, prospectPrenom, prospectTelephone,
		req.TypeDemande, motifAppel, req.Priorite,
		req.Statut, notesInternes,
		rdvDate, req.RDVHeure, req.RDVLieu,
		recuParUserID,
		aA, // Utilise NULL si 0 ou nil
	)

	if err != nil {
		return nil, fmt.Errorf("erreur création appel: %w", err)
	}

	newID, _ := result.LastInsertId()

	// 5. LOG D'AUDIT (Loi 25) - On utilise une goroutine pour la performance
	go func() {
		logRepo := NewLogRepo(r.db.DB)
		// On log qui a créé quoi

		_ = logRepo.PushLog(recuParUserID, "CREATE", "APPELS", int(newID), "INFO", req)
	}()

	return r.GetByID(int(newID), cryptoSvc)
}

// ========== MISE À JOUR ==========

// Update met à jour un appel existant
// ========== MISE À JOUR ==========

func (r *AppelRepository) Update(id int, req models.CreateAppelRequest, recuParUserID int, cryptoSvc *crypto.CryptoService) error {
	// Crypter les données
	appelantNom, _ := cryptoSvc.Encrypt(req.AppelantNom)
	appelantPrenom, _ := cryptoSvc.Encrypt(req.AppelantPrenom)
	appelantTelephone, _ := cryptoSvc.Encrypt(req.AppelantTelephone)
	prospectNom, _ := cryptoSvc.Encrypt(req.ProspectNom)
	prospectPrenom, _ := cryptoSvc.Encrypt(req.ProspectPrenom)
	prospectTelephone, _ := cryptoSvc.Encrypt(req.ProspectTelephone)
	motifAppel, _ := cryptoSvc.Encrypt(req.MotifAppel)
	notesInternes, _ := cryptoSvc.Encrypt(req.NotesInternes)

	// Gestion NULL pour les IDs
	var cID interface{} = nil
	if req.ClientID != nil && *req.ClientID > 0 {
		cID = *req.ClientID
	}
	var aA interface{} = nil
	if req.AssigneA != nil && *req.AssigneA > 0 {
		aA = *req.AssigneA
	}

	query := `
		UPDATE appels SET
			heure_appel = ?,
			appelant_nom = ?, appelant_prenom = ?, appelant_telephone = ?, appelant_relation = ?,
			client_id = ?,
			prospect_nom = ?, prospect_prenom = ?, prospect_telephone = ?,
			type_demande = ?, motif_appel = ?, priorite = ?,
			statut = ?, notes_internes = ?,
			rdv_date = ?, rdv_heure = ?, rdv_lieu = ?,
			assigne_a = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	var rdvDate interface{} = nil
	if req.RDVDate != "" {
		rdvDate = req.RDVDate
	}

	_, err := r.db.Exec(query,
		req.HeureAppel,
		appelantNom, appelantPrenom, appelantTelephone, req.AppelantRelation,
		cID,
		prospectNom, prospectPrenom, prospectTelephone,
		req.TypeDemande, motifAppel, req.Priorite,
		req.Statut, notesInternes,
		rdvDate, req.RDVHeure, req.RDVLieu,
		aA,
		id,
	)

	if err == nil {
		// Log de modification
		go func() {
			logRepo := NewLogRepo(r.db.DB)
			_ = logRepo.PushLog(recuParUserID, "UPDATE", "APPELS", id, "INFO", req)
		}()
	}

	return err
}

// ========== LECTURE ==========

// GetAll récupère tous les appels (liste) avec déchiffrement
func (r *AppelRepository) GetAll(cryptoSvc *crypto.CryptoService) ([]models.AppelListItem, error) {
	query := `
		SELECT 
			id, date_appel, heure_appel,
			appelant_nom, appelant_prenom, appelant_telephone,
			prospect_nom, prospect_prenom,
			type_demande, priorite, statut,
			client_id, created_at
		FROM appels
		ORDER BY date_appel DESC, created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appels []models.AppelListItem
	for rows.Next() {
		var item models.AppelListItem
		var appelantNom, appelantPrenom, appelantTel, prospectNom, prospectPrenom string

		err := rows.Scan(
			&item.ID, &item.DateAppel, &item.HeureAppel,
			&appelantNom, &appelantPrenom, &appelantTel,
			&prospectNom, &prospectPrenom,
			&item.TypeDemande, &item.Priorite, &item.Statut,
			&item.ClientID, &item.CreatedAt,
		)
		if err != nil {
			continue
		}

		// Déchiffrer
		item.AppelantNom, _ = cryptoSvc.Decrypt(appelantNom)
		item.AppelantPrenom, _ = cryptoSvc.Decrypt(appelantPrenom)
		item.AppelantTelephone, _ = cryptoSvc.Decrypt(appelantTel)
		item.ProspectNom, _ = cryptoSvc.Decrypt(prospectNom)
		item.ProspectPrenom, _ = cryptoSvc.Decrypt(prospectPrenom)

		appels = append(appels, item)
	}

	return appels, nil
}

// GetByID récupère un appel complet par ID avec déchiffrement
func (r *AppelRepository) GetByID(id int, cryptoSvc *crypto.CryptoService) (*models.Appel, error) {
	query := `
		SELECT 
			id, date_appel, heure_appel,
			appelant_nom, appelant_prenom, appelant_telephone, appelant_relation,
			client_id,
			prospect_nom, prospect_prenom, prospect_telephone,
			type_demande, motif_appel, priorite,
			statut, notes_internes,
			rdv_date, rdv_heure, rdv_lieu,
			recu_par, assigne_a,
			created_at, updated_at
		FROM appels WHERE id = ?
	`

	var appel models.Appel
	var appelantNom, appelantPrenom, appelantTel string
	var prospectNom, prospectPrenom, prospectTel string
	var motif, notes string
	var rdvDate *string

	err := r.db.QueryRow(query, id).Scan(
		&appel.ID, &appel.DateAppel, &appel.HeureAppel,
		&appelantNom, &appelantPrenom, &appelantTel, &appel.AppelantRelation,
		&appel.ClientID,
		&prospectNom, &prospectPrenom, &prospectTel,
		&appel.TypeDemande, &motif, &appel.Priorite,
		&appel.Statut, &notes,
		&rdvDate, &appel.RDVHeure, &appel.RDVLieu,
		&appel.RecuPar, &appel.AssigneA,
		&appel.CreatedAt, &appel.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Déchiffrer
	appel.AppelantNom, _ = cryptoSvc.Decrypt(appelantNom)
	appel.AppelantPrenom, _ = cryptoSvc.Decrypt(appelantPrenom)
	appel.AppelantTelephone, _ = cryptoSvc.Decrypt(appelantTel)
	appel.ProspectNom, _ = cryptoSvc.Decrypt(prospectNom)
	appel.ProspectPrenom, _ = cryptoSvc.Decrypt(prospectPrenom)
	appel.ProspectTelephone, _ = cryptoSvc.Decrypt(prospectTel)
	appel.MotifAppel, _ = cryptoSvc.Decrypt(motif)
	appel.NotesInternes, _ = cryptoSvc.Decrypt(notes)
	appel.RDVDate = rdvDate

	return &appel, nil
}

// ========== STATISTIQUES ==========

// GetStats retourne les statistiques des appels
func (r *AppelRepository) GetStats() (*models.StatsAppels, error) {
	stats := &models.StatsAppels{}

	// Total
	r.db.QueryRow("SELECT COUNT(*) FROM appels").Scan(&stats.Total)

	// Nouveaux
	r.db.QueryRow("SELECT COUNT(*) FROM appels WHERE statut = 'nouveau'").Scan(&stats.Nouveaux)

	// En attente
	r.db.QueryRow("SELECT COUNT(*) FROM appels WHERE statut = 'en_attente'").Scan(&stats.EnAttente)

	// RDV planifiés
	r.db.QueryRow("SELECT COUNT(*) FROM appels WHERE statut = 'rdv_planifie'").Scan(&stats.RDVPlanifies)

	// Urgents
	r.db.QueryRow("SELECT COUNT(*) FROM appels WHERE priorite = 'urgente'").Scan(&stats.Urgents)

	// Aujourd'hui
	today := time.Now().Format("2006-01-02")
	r.db.QueryRow("SELECT COUNT(*) FROM appels WHERE DATE(date_appel) = ?", today).Scan(&stats.Aujourdhui)

	return stats, nil
}

// ========== SUPPRESSION ==========

// Delete supprime un appel
func (r *AppelRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM appels WHERE id = ?", id)
	return err
}
