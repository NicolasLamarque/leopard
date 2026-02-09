package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	models "leopard/internal/model"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
)

// 🛠️ Helper pour convertir string → *string
// Si la string est vide, on retourne nil pour que SQLite enregistre NULL
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// 🛠️ Helper pour déréférencer *string → string (utile pour les comparaisons)
func stringVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ========== CRUD BASIQUE ==========

func (db *Database) GetAllResidences() ([]models.Residence, error) {
	var residences []models.Residence
	query := `SELECT * FROM residences ORDER BY municipalite, titre`

	err := db.Select(&residences, query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des résidences: %w", err)
	}

	return residences, nil
}

func (db *Database) GetResidenceByID(id int) (*models.Residence, error) {
	var residence models.Residence
	query := `SELECT * FROM residences WHERE id = ?`

	err := db.Get(&residence, query, id)
	if err != nil {
		return nil, fmt.Errorf("résidence ID %d non trouvée: %w", id, err)
	}
	return &residence, nil
}

// GetResidenceByRegistre - Fonction appelée par app.go
func (db *Database) GetResidenceByRegistre(registre string, sync bool) (*models.Residence, error) {
	fmt.Printf("\n📖 [REPO] GetResidenceByRegistre appelé\n")
	fmt.Printf("   - Registre: %s\n", registre)
	fmt.Printf("   - Sync demandé: %v\n", sync)

	var res models.Residence

	query := `SELECT * FROM residences WHERE registre = ?`
	err := db.Get(&res, query, registre)
	if err != nil {
		fmt.Printf("❌ [REPO] Résidence non trouvée en DB: %v\n", err)
		return nil, err
	}

	fmt.Printf("✅ [REPO] Résidence trouvée: %s\n", res.Titre)

	// 🔄 SYNCHRONISATION EN TEMPS RÉEL (page abrégée)
	if sync {
		fmt.Println("🔄 [REPO] Actualisation en direct depuis MSSS...")

		client := &http.Client{Timeout: 15 * time.Second}

		// On passe la référence pour mettre à jour l'objet
		db.scrapeRPADetailsSimple(client, &res)

		// Mise à jour des timestamps avec stringPtr
		nowStr := time.Now().Format("2006-01-02 15:04:05")
		res.DateMAJ = stringPtr(nowStr)
		res.DerniereVerification = stringPtr(nowStr)

		// Sauvegarde en DB
		err = db.InsertResidence(&res)
		if err != nil {
			fmt.Printf("⚠️ [REPO] Erreur sauvegarde après sync: %v\n", err)
		} else {
			fmt.Println("✅ [REPO] Données actualisées et sauvegardées en DB")
		}
	} else {
		fmt.Println("📴 [REPO] Mode cache - Pas de synchronisation demandée")
	}

	return &res, nil
}

// ========== INSERT / UPDATE (Utilisent les NamedExec de sqlx) ==========

func (db *Database) InsertResidence(res *models.Residence) error {
	query := `
        INSERT INTO residences (
            region, registre, numero_interne, titre, municipalite, adresse, ville, code_postal,
            telephone, capacite, type_resid, proprietaires, services, 
            date_certification, statut, source_url, source_url_detaillee, 
            date_ajout, date_maj, derniere_verification
        ) VALUES (
            :region, :registre, :numero_interne, :titre, :municipalite, :adresse, :ville, :code_postal,
            :telephone, :capacite, :type_resid, :proprietaires, :services, 
            :date_certification, :statut, :source_url, :source_url_detaillee, 
            :date_ajout, :date_maj, :derniere_verification
        ) ON CONFLICT(registre) DO UPDATE SET
            numero_interne = excluded.numero_interne,
            adresse = excluded.adresse,
            ville = excluded.ville,
            code_postal = excluded.code_postal,
            telephone = excluded.telephone,
            capacite = excluded.capacite,
            services = excluded.services,
            date_maj = excluded.date_maj
    `
	_, err := db.NamedExec(query, res)
	return err
}

func (db *Database) UpdateResidence(residence *models.Residence) error {
	residence.DateMAJ = stringPtr(time.Now().Format("2006-01-02 15:04:05"))

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

func (db *Database) DeleteResidence(id int) error {
	query := `DELETE FROM residences WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de la résidence: %w", err)
	}
	return nil
}

// ========== RECHERCHE ==========

func (db *Database) SearchResidences(municipalite, nom, statut string) ([]models.Residence, error) {
	var residences []models.Residence

	query := `
        SELECT * FROM residences 
        WHERE municipalite LIKE ? 
        AND titre LIKE ? 
        AND statut LIKE ?
        ORDER BY titre ASC`

	err := db.Select(&residences, query,
		"%"+municipalite+"%",
		"%"+nom+"%",
		"%"+statut+"%",
	)
	return residences, err
}

// ========== SYNCHRONISATION RPA ==========

type RPAStats struct {
	TotalScraped int       `json:"total_scraped"`
	Nouveaux     int       `json:"nouveaux"`
	MisAJour     int       `json:"mis_a_jour"`
	Fermes       int       `json:"fermes"`
	Erreurs      []string  `json:"erreurs"`
	DateSync     time.Time `json:"date_sync"`
}

func (db *Database) SyncRPAFromMSSS() (*RPAStats, error) {
	stats := &RPAStats{
		DateSync: time.Now(),
		Erreurs:  []string{},
	}

	fmt.Println("🚀 DÉMARRAGE SYNCHRONISATION MSSS")

	rpasScraped, err := db.scrapeRPAList()
	if err != nil {
		return nil, fmt.Errorf("erreur scraping: %w", err)
	}

	// --- LA SOLUTION EST ICI ---
	// On retire maintenant les 4 premières lignes
	if len(rpasScraped) > 3 {
		rpasScraped = rpasScraped[3:]
	} else {
		return stats, nil
	}

	stats.TotalScraped = len(rpasScraped)
	fmt.Printf("✅ %d RPA récupérés du MSSS\n", stats.TotalScraped)

	existingRPAs, err := db.GetAllResidences()
	if err != nil {
		return nil, fmt.Errorf("erreur récupération existants: %w", err)
	}

	existingMap := make(map[string]*models.Residence)
	for i := range existingRPAs {
		existingMap[existingRPAs[i].Registre] = &existingRPAs[i]
	}

	for _, rpaNew := range rpasScraped {
		tx, err := db.Begin()
		if err != nil {
			continue
		}

		if existing, found := existingMap[rpaNew.Registre]; found {
			if db.needsUpdate(existing, &rpaNew) {
				err := db.updateRPAInTx(tx, &rpaNew)
				if err == nil {
					tx.Commit()
					stats.MisAJour++
					fmt.Printf("🔄 MAJ: %s\n", rpaNew.Titre)
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
				stats.Nouveaux++
				fmt.Printf("✨ NOUVEAU: %s\n", rpaNew.Titre)
			} else {
				tx.Rollback()
				fmt.Printf("❌ Erreur Ajout: %v\n", err)
			}
		}
	}

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

	fmt.Println("🎉 SYNCHRONISATION TERMINÉE")
	return stats, nil
}

func (db *Database) scrapeRPAList() ([]models.Residence, error) {
	baseURL := "http://k10.pub.msss.rtss.qc.ca"
	searchURL := baseURL + "/public/K10FormRecherche.asp?hidPasseParFormulaireRecherche=1&cert=&act=Rechercher&nmResid=&nomMunicipalite=&cdRSS=&cdCLSC=&cdRLS=&cdMRC=&refTpResid=&refStForm=&noResReg=&boolLogeRepas=&boolSoin=&boolAssistance=&boolAlimentation=&boolLoisir=&boolSecurite=&nmProprio=&pnmProprio=&refStRes=F"

	client := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	utf8Reader := charmap.Windows1252.NewDecoder().Reader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		return nil, err
	}

	var rpas []models.Residence
	now := time.Now().Format("2006-01-02 15:04:05")

	doc.Find("table").Each(func(index int, table *goquery.Selection) {
		if strings.Contains(strings.ToLower(table.Text()), "registre") {
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				cols := s.Find("td")

				if cols.Length() >= 5 {
					reg := cleanText(cols.Eq(2).Text())
					nom := cleanText(cols.Eq(3).Text())

					if reg != "" && reg != "Registre" && !strings.Contains(nom, "Nom de la") {
						rpa := models.Residence{
							Region:               stringPtr(cleanText(cols.Eq(1).Text())),
							Registre:             reg,
							Titre:                nom,
							Municipalite:         stringPtr(cleanText(cols.Eq(4).Text())),
							Statut:               "actif",
							DateAjout:            stringPtr(now),
							DerniereVerification: stringPtr(now),
						}

						url := "http://k10.pub.msss.rtss.qc.ca/public/formulaire/K10FormCons.asp?noForm=" + reg
						rpa.SourceURL = stringPtr(url)
						rpas = append(rpas, rpa)
					}
				}
			})
		}
	})

	return rpas, nil
}

func (db *Database) scrapeRPADetails(client *http.Client, rpa *models.Residence) {
	if rpa.SourceURL == nil {
		return
	}

	resp, err := client.Get(*rpa.SourceURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	utf8Reader := charmap.Windows1252.NewDecoder().Reader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		return
	}

	extract := func(txt string, s *goquery.Selection, label string) string {
		val := cleanText(strings.Replace(txt, label, "", 1))
		if val == "" {
			val = cleanText(s.Next().Text())
		}
		return val
	}

	var servicesBuilder strings.Builder
	inServicesSection := false

	doc.Find("td, b, span, div, p").Each(func(i int, s *goquery.Selection) {
		txt := s.Text()

		if strings.Contains(txt, "7 - ") && strings.Contains(txt, "services offerts") {
			inServicesSection = true
			return
		}

		if inServicesSection && (strings.Contains(txt, "8 -") || strings.Contains(txt, "Commentaires")) {
			inServicesSection = false
			return
		}

		if inServicesSection && txt != "" && !strings.Contains(txt, "7 -") {
			cleaned := cleanText(txt)
			if len(cleaned) > 3 {
				servicesBuilder.WriteString(cleaned)
				servicesBuilder.WriteString("\n")
			}
		}

		switch {
		case strings.Contains(txt, "Adresse :"):
			rpa.Adresse = stringPtr(extract(txt, s, "Adresse :"))
		case strings.Contains(txt, "Ville :"):
			rpa.Ville = stringPtr(extract(txt, s, "Ville :"))
		case strings.Contains(txt, "Code postal :"):
			rpa.CodePostal = stringPtr(extract(txt, s, "Code postal :"))
		case strings.Contains(txt, "Téléphone :"):
			rpa.Telephone = stringPtr(extract(txt, s, "Téléphone :"))
		case strings.Contains(txt, "Capacité d'accueil :"):
			valStr := extract(txt, s, "Capacité d'accueil :")
			valInt, _ := strconv.Atoi(valStr)
			rpa.Capacite = valInt
		case strings.Contains(txt, "Propriétaire(s) :"):
			rpa.Proprietaires = stringPtr(extract(txt, s, "Propriétaire(s) :"))
		case strings.Contains(txt, "Type de résidence :"):
			rpa.TypeResid = stringPtr(extract(txt, s, "Type de résidence :"))
		case strings.Contains(txt, "Date de certification :"):
			rpa.DateCertification = stringPtr(extract(txt, s, "Date de certification :"))
		}
	})

	if servicesBuilder.Len() > 0 {
		rpa.Services = stringPtr(strings.TrimSpace(servicesBuilder.String()))
	}
}

func (db *Database) GetResidenceForDetails(registre string, isConnected bool) (*models.Residence, error) {
	var rpa models.Residence
	err := db.Get(&rpa, "SELECT * FROM residences WHERE registre = ?", registre)
	if err != nil {
		return nil, err
	}

	if isConnected && rpa.SourceURL != nil {
		client := &http.Client{Timeout: 15 * time.Second}
		db.scrapeRPADetails(client, &rpa)

		nowStr := time.Now().Format("2006-01-02 15:04:05")
		rpa.DateMAJ = stringPtr(nowStr)
		rpa.DerniereVerification = stringPtr(nowStr)

		db.InsertResidence(&rpa)
	}

	return &rpa, nil
}

// ========== NOUVEAU SCRAPER OPTIMISÉ (La page abrégée) ==========

func (db *Database) scrapeRPADetailsSimple(client *http.Client, rpa *models.Residence) {
	url := fmt.Sprintf("http://k10.pub.msss.rtss.qc.ca/public/K10ConsFormAbg.asp?Registre=%s", rpa.Registre)
	fmt.Printf("\n--- 🌐 SCRAPING RPA #%s ---\n", rpa.Registre)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ ERREUR HTTP: %v\n", err)
		return
	}
	defer resp.Body.Close()

	utf8Reader := charmap.Windows1252.NewDecoder().Reader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		return
	}

	reCodePostal := regexp.MustCompile(`(?i)[A-Z][0-9][A-Z]\s?[0-9][A-Z][0-9]`)
	reTelephone := regexp.MustCompile(`[0-9]{3}-[0-9]{3}-[0-9]{4}`)

	// --- BLOC 1 : TABLE D'ADRESSE (textedanstableau) ---
	doc.Find("table.textedanstableau tr td").Each(func(i int, s *goquery.Selection) {
		line := cleanText(s.Text())
		if line == "" || strings.Contains(strings.ToLower(line), "certification") {
			return
		}

		// --- 1. CODE POSTAL ---
		if reCodePostal.MatchString(line) {
			cp := reCodePostal.FindString(line)
			rpa.CodePostal = stringPtr(strings.ToUpper(cp))
			fmt.Printf("   ✅ CP: %s\n", cp)
		}

		// --- 2. TÉLÉPHONE ---
		if reTelephone.MatchString(line) {
			tel := reTelephone.FindString(line)
			rpa.Telephone = stringPtr(tel)
			fmt.Printf("   ✅ TEL: %s\n", tel)
		}

		// --- 3. VILLE ---
		if strings.Contains(line, "(Québec)") || strings.Contains(line, "(Quebec)") {
			v := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(line, "(Québec)", ""), "(Quebec)", ""))
			rpa.Ville = stringPtr(v)
			fmt.Printf("   ✅ VILLE: %s\n", v)
		}

		// --- 4. RÉGION ---
		if strings.Contains(line, "-") && !reTelephone.MatchString(line) && !unicode.IsDigit(rune(line[0])) {
			rpa.Region = stringPtr(line)
			fmt.Printf("   ✅ RÉGION: %s\n", line)
		}

		// --- 5. ADRESSE ---
		if unicode.IsDigit(rune(line[0])) && !reTelephone.MatchString(line) && !reCodePostal.MatchString(line) {
			if !strings.Contains(line, " - ") {
				rpa.Adresse = stringPtr(line)
				fmt.Printf("   ✅ ADRESSE: %s\n", line)
			}
		}
	})

	// --- BLOC 2 : DÉTAILS TECHNIQUES (tableau) ---
	doc.Find("table.tableau tr").Each(func(i int, s *goquery.Selection) {
		label := strings.ToLower(cleanText(s.Find("td").First().Text()))
		val := cleanText(s.Find("td").Last().Text())

		if strings.Contains(label, "unités locatives de la rpa") {
			if capVal, err := strconv.Atoi(val); err == nil {
				rpa.Capacite = capVal
				fmt.Printf("      ✅ CAPACITÉ: %d\n", capVal)
			}
		}
		if strings.Contains(label, "type de résidence") {
			rpa.TypeResid = stringPtr(val)
			fmt.Printf("      ✅ TYPE: %s\n", val)
		}
		// Ajout du numéro interne si présent
		if strings.Contains(label, "numéro interne") {
			rpa.NumeroInterne = stringPtr(val)
			fmt.Printf("      ✅ NUMÉRO INTERNE: %s\n", val)
		}
	})

	fmt.Printf("--- 🏁 FIN SCRAPING #%s ---\n\n", rpa.Registre)
}

// Fonction utilitaire pour vérifier le début d'une string
func anyPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
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
	_, err := tx.Exec(query, now, stringPtr(now), id)
	return err
}

func cleanText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\t", " ")
	return strings.Join(strings.Fields(s), " ")
}

func (db *Database) needsUpdate(existing, new *models.Residence) bool {
	return existing.Titre != new.Titre ||
		stringVal(existing.Municipalite) != stringVal(new.Municipalite) ||
		stringVal(existing.Adresse) != stringVal(new.Adresse) ||
		stringVal(existing.Telephone) != stringVal(new.Telephone) ||
		stringVal(existing.Services) != stringVal(new.Services) ||
		stringVal(existing.SourceURL) != stringVal(new.SourceURL)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
