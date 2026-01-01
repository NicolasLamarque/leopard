package database

import (
	"database/sql"
	"fmt" // Ajoute cet import en haut du fichier
	"strconv"

	models "leopard/internal/model"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
)

// Ajoute ceci pour lire le body

// Ajoute ceci pour les accents

// ========== CRUD BASIQUE ==========

// GetAllResidences récupère toutes les résidences
func (db *Database) GetAllResidences() ([]models.Residence, error) {
	var residences []models.Residence
	query := `SELECT * FROM residences ORDER BY municipalite, titre`

	err := db.Select(&residences, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des résidences: %w", err)
	}

	return residences, nil
}

// GetResidenceByID récupère une résidence par ID
func (db *Database) GetResidenceByID(id int) (*models.Residence, error) {
	var residence models.Residence
	query := `SELECT * FROM residences WHERE id = ?`

	err := db.Get(&residence, query, id)
	if err != nil {
		return nil, fmt.Errorf("résidence ID %d non trouvée: %w", id, err)
	}
	return &residence, nil
}
func (db *Database) GetResidenceByRegistre(registre string, sync bool) (*models.Residence, error) {
	var res models.Residence

	// 1. Chercher en base de données d'abord
	query := `SELECT * FROM residences WHERE registre = ?`
	err := db.Get(&res, query, registre)
	if err != nil {
		return nil, err
	}

	// 2. Si on demande la synchro et qu'on a l'URL
	if sync && res.SourceURL != "" {
		// ICI : Ton code de scraping qui remplit res.Services (Section 7)
		// updatedRes, err := db.ScrapeAndSave(res.SourceURL)
		// if err == nil { res = *updatedRes }
	}

	return &res, nil
}

// InsertResidence insère une nouvelle résidence avec tous les détails extraits
func (db *Database) InsertResidence(residence *models.Residence) error {
	query := `
        INSERT INTO residences (
            region, registre, titre, municipalite, adresse, ville, code_postal,
            telephone, capacite, type_resid, proprietaires, services, 
            date_certification, statut, source_url, notes, 
            date_ajout, date_maj, derniere_verification
        )
        VALUES (
            :region, :registre, :titre, :municipalite, :adresse, :ville, :code_postal,
            :telephone, :capacite, :type_resid, :proprietaires, :services, 
            :date_certification, :statut, :source_url, :notes, 
            :date_ajout, :date_maj, :derniere_verification
        )
        ON CONFLICT(registre) DO UPDATE SET
            titre = excluded.titre,
            municipalite = excluded.municipalite,
            source_url = excluded.source_url,
            derniere_verification = excluded.derniere_verification,
            date_maj = excluded.date_maj
    `

	_, err := db.NamedExec(query, residence)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion/maj de la résidence: %w", err)
	}
	return nil
}

// UpdateResidence met à jour une résidence existante (utilisé lors de la synchro)
func (db *Database) UpdateResidence(residence *models.Residence) error {
	// On force la date de mise à jour à maintenant
	residence.DateMAJ = time.Now().Format("2006-01-02 15:04:05")

	query := `
		UPDATE residences
		SET region = :region,
			titre = :titre,
			municipalite = :municipalite,
			adresse = :adresse,
			ville = :ville,
			code_postal = :code_postal,
			telephone = :telephone,
			capacite = :capacite,
			type_resid = :type_resid,
			proprietaires = :proprietaires,
			services = :services,
			date_certification = :date_certification,
			statut = :statut,
			source_url = :source_url,
			notes = :notes,
			date_maj = :date_maj,
			derniere_verification = :derniere_verification
		WHERE registre = :registre
	`

	_, err := db.NamedExec(query, residence)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de la résidence %s: %w", residence.Registre, err)
	}
	return nil
}

// DeleteResidence supprime une résidence
func (db *Database) DeleteResidence(id int) error {
	query := `DELETE FROM residences WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de la résidence: %w", err)
	}
	return nil
}

// ========== RECHERCHE ==========

// SearchResidences recherche des résidences avec filtres
func (db *Database) SearchResidences(municipalite, nom, statut string) ([]models.Residence, error) {
	var residences []models.Residence

	// On utilise % pour que si la chaîne est vide, SQL ignore le filtre
	query := `
        SELECT * FROM residences 
        WHERE municipalite LIKE ? 
        AND titre_rpa LIKE ? 
        AND statut LIKE ?
        ORDER BY titre_rpa ASC`

	err := db.Select(&residences, query,
		"%"+municipalite+"%",
		"%"+nom+"%",
		"%"+statut+"%",
	)
	return residences, err
}

// ========== SYNCHRONISATION RPA ==========

// RPAStats statistiques de synchronisation
type RPAStats struct {
	TotalScraped int       `json:"total_scraped"`
	Nouveaux     int       `json:"nouveaux"`
	MisAJour     int       `json:"mis_a_jour"`
	Fermes       int       `json:"fermes"`
	Erreurs      []string  `json:"erreurs"`
	DateSync     time.Time `json:"date_sync"`
}

// SyncRPAFromMSSS scrape et synchronise les RPA depuis le site du MSSS
func (db *Database) SyncRPAFromMSSS() (*RPAStats, error) {
	stats := &RPAStats{
		DateSync: time.Now(),
		Erreurs:  []string{},
	}

	// 1. On aspire la liste brute (très rapide)
	rpasScraped, err := db.scrapeRPAList()
	if err != nil {
		return nil, fmt.Errorf("erreur scraping: %w", err)
	}
	stats.TotalScraped = len(rpasScraped)

	// 2. On charge les existants pour comparer
	existingRPAs, err := db.GetAllResidences()
	if err != nil {
		return nil, fmt.Errorf("erreur récupération existants: %w", err)
	}

	existingMap := make(map[string]*models.Residence)
	for i := range existingRPAs {
		existingMap[existingRPAs[i].Registre] = &existingRPAs[i]
	}

	// 3. Traitement avec COMMIT IMMÉDIAT pour chaque ligne
	for _, rpaNew := range rpasScraped {
		tx, err := db.Begin()
		if err != nil {
			continue
		}

		// On ignore temporairement scrapeRPADetails pour voir si l'insertion marche
		// db.scrapeRPADetails(client, &rpaNew)

		if existing, found := existingMap[rpaNew.Registre]; found {
			if db.needsUpdate(existing, &rpaNew) {
				err := db.updateRPAInTx(tx, &rpaNew)
				if err == nil {
					tx.Commit()
					fmt.Printf("✅ MAJ en DB: %s\n", rpaNew.Titre)
				} else {
					tx.Rollback()
					fmt.Printf("❌ Erreur MAJ: %v\n", err)
				}
			} else {
				tx.Rollback()
			}
			delete(existingMap, rpaNew.Registre)
		} else {
			err := db.insertRPAInTx(tx, &rpaNew)
			if err == nil {
				tx.Commit()
				fmt.Printf("✨ NOUVEAU en DB: %s\n", rpaNew.Titre)
			} else {
				tx.Rollback()
				fmt.Printf("❌ Erreur Ajout: %v\n", err)
			}
		}
	}

	// 4. Marquer les disparus comme Fermés
	for _, rpa := range existingMap {
		if rpa.Statut == "actif" {
			tx, _ := db.Begin()
			err := db.markRPAFermeInTx(tx, rpa.ID)
			if err == nil {
				tx.Commit()
				stats.Fermes++
				fmt.Printf("🚫 FERMÉ: %s\n", rpa.Titre)
			} else {
				tx.Rollback()
			}
		}
	}

	return stats, nil
}
func (db *Database) scrapeRPAList() ([]models.Residence, error) {
	baseURL := "http://k10.pub.msss.rtss.qc.ca"
	// On utilise l'URL complète que tu as confirmée
	searchURL := baseURL + "/public/K10FormRecherche.asp?hidPasseParFormulaireRecherche=1&cert=&act=Rechercher&nmResid=&nomMunicipalite=&cdRSS=&cdCLSC=&cdRLS=&cdMRC=&refTpResid=&refStForm=&noResReg=&boolLogeRepas=&boolSoin=&boolAssistance=&boolAlimentation=&boolLoisir=&boolSecurite=&nmProprio=&pnmProprio=&refStRes=F"

	client := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	// Simulation d'un navigateur réel
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Correction des accents (Windows-1252 vers UTF-8)
	utf8Reader := charmap.Windows1252.NewDecoder().Reader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		return nil, err
	}

	var rpas []models.Residence
	now := time.Now().Format("2006-01-02 15:04:05")

	// On cible la table. Le site du MSSS n'a pas de classe, on cherche l'entête.
	doc.Find("table").Each(func(index int, table *goquery.Selection) {
		if strings.Contains(strings.ToLower(table.Text()), "registre") {
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				cols := s.Find("td")

				if cols.Length() >= 5 {
					reg := cleanText(cols.Eq(2).Text())
					nom := cleanText(cols.Eq(3).Text())

					if reg != "" && reg != "Registre" && !strings.Contains(nom, "Nom de la") {
						rpa := models.Residence{
							Region:               cleanText(cols.Eq(1).Text()),
							Registre:             reg,
							Titre:                nom,
							Municipalite:         cleanText(cols.Eq(4).Text()),
							Statut:               "actif",
							DateAjout:            now,
							DerniereVerification: now,
						}

						// On stocke le LIEN pour usage futur
						if link, exists := cols.Eq(3).Find("a").Attr("href"); exists {
							// On s'assure que le lien est complet
							rpa.SourceURL = baseURL + strings.TrimSpace(link)
						}

						fmt.Printf("🎯 [%d] %s (URL capturée)\n", len(rpas)+1, rpa.Titre)

						// --- STRATÉGIE : ON NE FAIT PAS ÇA MAINTENANT ---
						// db.scrapeRPADetails(client, &rpa)
						// ------------------------------------------------

						rpas = append(rpas, rpa)
					}
				}
			})
		}
	})

	return rpas, nil
}

func (db *Database) scrapeRPADetails(client *http.Client, rpa *models.Residence) {
	resp, err := client.Get(rpa.SourceURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// Utilitaire pour extraire la valeur proprement
	extract := func(txt string, s *goquery.Selection, label string) string {
		val := cleanText(strings.Replace(txt, label, "", 1))
		if val == "" {
			val = cleanText(s.Next().Text())
		}
		return val
	}

	doc.Find("td, b, span, div").Each(func(i int, s *goquery.Selection) {
		txt := s.Text()

		switch {
		case strings.Contains(txt, "Adresse :"):
			rpa.Adresse = extract(txt, s, "Adresse :")

		case strings.Contains(txt, "Ville :"):
			rpa.Ville = extract(txt, s, "Ville :")

		case strings.Contains(txt, "Code postal :"):
			rpa.CodePostal = extract(txt, s, "Code postal :")

		case strings.Contains(txt, "Téléphone :"):
			rpa.Telephone = extract(txt, s, "Téléphone :")

		case strings.Contains(txt, "Capacité d'accueil :"):
			valStr := extract(txt, s, "Capacité d'accueil :")
			valInt, _ := strconv.Atoi(valStr)
			rpa.Capacite = valInt

		case strings.Contains(txt, "Propriétaire(s) :"):
			rpa.Proprietaires = extract(txt, s, "Propriétaire(s) :")

		case strings.Contains(txt, "Type de résidence :"):
			rpa.TypeResid = extract(txt, s, "Type de résidence :")

		case strings.Contains(txt, "Date de certification :"):
			rpa.DateCertification = extract(txt, s, "Date de certification :")

		// SECTION 7 - LES SERVICES
		case strings.Contains(txt, "7 - Services offerts"):
			// Souvent les services sont dans le bloc qui suit immédiatement ce titre
			rpa.Services = extract(txt, s, "7 - Services offerts par la résidence")
			if rpa.Services == "" {
				// Si c'est un gros bloc en dessous, on prend le parent ou le voisin
				rpa.Services = cleanText(s.Parent().Next().Text())
			}
		}
	})
}

// GetResidenceForDetails récupère une résidence et actualise les détails si connecté
func (db *Database) GetResidenceForDetails(registre string, isConnected bool) (*models.Residence, error) {
	var rpa models.Residence
	err := db.Get(&rpa, "SELECT * FROM residences WHERE registre = ?", registre)
	if err != nil {
		return nil, err
	}

	if isConnected && rpa.SourceURL != "" {
		fmt.Printf("🔍 Actualisation en direct : %s\n", rpa.Titre)

		client := &http.Client{Timeout: 10 * time.Second}
		db.scrapeRPADetails(client, &rpa)

		// On transforme le temps présent en STRING
		rpa.DateMAJ = time.Now().Format("2006-01-02 15:04:05")
		rpa.DerniereVerification = time.Now().Format("2006-01-02 15:04:05")

		db.InsertResidence(&rpa)
	}

	return &rpa, nil
}

// ========== FONCTIONS UTILITAIRES PRIVÉES ==========

func cleanText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	return strings.Join(strings.Fields(s), " ")
}

func (db *Database) needsUpdate(existing, new *models.Residence) bool {
	return existing.Titre != new.Titre ||
		existing.Municipalite != new.Municipalite ||
		existing.Adresse != new.Adresse ||
		existing.Telephone != new.Telephone || // Ajouté
		existing.Services != new.Services // Ajouté
}

func (db *Database) insertRPAInTx(tx *sql.Tx, rpa *models.Residence) error {
	query := `
		INSERT INTO residences (
			region, registre, titre, municipalite, adresse, ville, code_postal,
			telephone, capacite, type_resid, proprietaires, services, 
			date_certification, statut, source_url, notes, 
			date_ajout, date_maj, derniere_verification
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := tx.Exec(query,
		rpa.Region, rpa.Registre, rpa.Titre, rpa.Municipalite, rpa.Adresse, rpa.Ville, rpa.CodePostal,
		rpa.Telephone, rpa.Capacite, rpa.TypeResid, rpa.Proprietaires, rpa.Services,
		rpa.DateCertification, rpa.Statut, rpa.SourceURL, rpa.Notes,
		rpa.DateAjout, rpa.DateMAJ, rpa.DerniereVerification,
	)
	return err
}

func (db *Database) updateRPAInTx(tx *sql.Tx, rpa *models.Residence) error {
	query := `
		UPDATE residences SET
			region = ?, titre = ?, municipalite = ?, adresse = ?, ville = ?, code_postal = ?,
			telephone = ?, capacite = ?, type_resid = ?, proprietaires = ?, services = ?,
			date_certification = ?, statut = ?, source_url = ?, notes = ?,
			date_maj = ?, derniere_verification = ?
		WHERE registre = ?
	`
	_, err := tx.Exec(query,
		rpa.Region, rpa.Titre, rpa.Municipalite, rpa.Adresse, rpa.Ville, rpa.CodePostal,
		rpa.Telephone, rpa.Capacite, rpa.TypeResid, rpa.Proprietaires, rpa.Services,
		rpa.DateCertification, rpa.Statut, rpa.SourceURL, rpa.Notes,
		rpa.DateMAJ, rpa.DerniereVerification,
		rpa.Registre,
	)
	return err
}

func (db *Database) markRPAFermeInTx(tx *sql.Tx, id int) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	query := `
		UPDATE residences 
		SET statut = 'ferme', date_fermeture = ?, date_maj = ? 
		WHERE id = ?
	`
	_, err := tx.Exec(query, now, now, id)
	return err
}
