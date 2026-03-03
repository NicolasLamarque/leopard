package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/html"
)

// =============================
// Types pour le sync/crawl
// =============================

// SyncProgress est émis via Wails EventsEmit à chaque étape
type SyncProgress struct {
	Region   string `json:"region"`
	Current  int    `json:"current"`
	Total    int    `json:"total"`
	Message  string `json:"message"`
	Inserted int    `json:"inserted"`
	Skipped  int    `json:"skipped"`
	Errors   int    `json:"errors"`
	Done     bool   `json:"done"`
}

// SyncResult est le résultat final retourné à Vue
type SyncResult struct {
	Inserted int    `json:"inserted"`
	Skipped  int    `json:"skipped"`
	Errors   int    `json:"errors"`
	Message  string `json:"message"`
}

// Régions MSSS
var msssRegions = []struct{ Code, Name string }{
	{"01", "Bas-Saint-Laurent"},
	{"02", "Saguenay–Lac-Saint-Jean"},
	{"03", "Capitale-Nationale"},
	{"04", "Mauricie et Centre-du-Québec"},
	{"05", "Estrie"},
	{"06", "Montréal"},
	{"07", "Outaouais"},
	{"08", "Abitibi-Témiscamingue"},
	{"09", "Côte-Nord"},
	{"10", "Nord-du-Québec"},
	{"11", "Gaspésie–Îles-de-la-Madeleine"},
	{"12", "Chaudière-Appalaches"},
	{"13", "Laval"},
	{"14", "Lanaudière"},
	{"15", "Laurentides"},
	{"16", "Montérégie"},
	{"17", "Nunavik"},
	{"18", "Terres-Cries-de-la-Baie-James"},
}

// Mots-clés pour filtrer les CHSLD dans la liste
var chsldKeywords = []string{
	"HÉBERGEMENT",
	"HEBERGEMENT",      // sans accent (décodage imparfait)
	"H\u00c9BERGEMENT", // unicode exact
	"CHSLD",
	"SOINS DE LONGUE",
	"MAISON DES A", // couvre AÎNÉS et variantes
	"FOYER",        // Foyer Mamo, Foyer Saints-Anges...
	"PAVILLON",     // Pavillon Sainte-Marie...
	"RÉSIDENCE",
	"RESIDENCE",
	"VILLA ",
	"MANOIR",
}

const msssBaseURL = "https://m02.pub.msss.rtss.qc.ca"

// =============================
// Méthodes exposées à Wails
// =============================

// SyncCHSLDFromMSSS crawle le site MSSS et insère les CHSLD dans la DB.
// Émet des événements "chsld:sync:progress" en temps réel vers Vue.
// Appelable depuis Vue : SyncCHSLDFromMSSS()
func (a *App) SyncCHSLDFromMSSS() (SyncResult, error) {
	result := SyncResult{}

	totalRegions := len(msssRegions)

	for idx, region := range msssRegions {
		// Émettre la progression région
		a.emitSyncProgress(SyncProgress{
			Region:  region.Name,
			Current: idx + 1,
			Total:   totalRegions,
			Message: fmt.Sprintf("Scan région %s...", region.Name),
		})

		links, err := msssListInstallations(region.Code, region.Name)
		if err != nil {
			log.Printf("Région %s: erreur liste: %v", region.Code, err)
			result.Errors++
			continue
		}

		// Filtrer les CHSLD
		var chsldLinks []msssInstallLink
		for _, lk := range links {
			nameUp := strings.ToUpper(lk.Name)
			for _, kw := range chsldKeywords {
				if strings.Contains(nameUp, kw) {
					chsldLinks = append(chsldLinks, lk)
					break
				}
			}
		}

		a.emitSyncProgress(SyncProgress{
			Region:  region.Name,
			Current: idx + 1,
			Total:   totalRegions,
			Message: fmt.Sprintf("%s — %d CHSLD trouvés", region.Name, len(chsldLinks)),
		})

		for _, lk := range chsldLinks {
			// Déjà en base ?
			exists, _ := a.chsldURLExists(lk.URL)
			if exists {
				result.Skipped++
				continue
			}

			data, err := msssFetchDetail(lk)
			if err != nil {
				log.Printf("Erreur détail %s: %v", lk.Name, err)
				result.Errors++
				continue
			}

			if err := a.insertCHSLDFromCrawl(data); err != nil {
				log.Printf("Erreur insertion %s: %v", lk.Name, err)
				result.Errors++
				continue
			}

			result.Inserted++
			a.emitSyncProgress(SyncProgress{
				Region:   region.Name,
				Current:  idx + 1,
				Total:    totalRegions,
				Message:  fmt.Sprintf("✅ %s", data.TitreCHSLD),
				Inserted: result.Inserted,
				Skipped:  result.Skipped,
				Errors:   result.Errors,
			})

			time.Sleep(300 * time.Millisecond)
		}
	}

	result.Message = fmt.Sprintf(
		"Sync terminé — %d insérés, %d déjà présents, %d erreurs",
		result.Inserted, result.Skipped, result.Errors,
	)

	// Émettre l'événement "done"
	a.emitSyncProgress(SyncProgress{
		Done:     true,
		Inserted: result.Inserted,
		Skipped:  result.Skipped,
		Errors:   result.Errors,
		Message:  result.Message,
	})

	return result, nil
}

// GetSyncStatus retourne combien de CHSLD sont en base (pour afficher dans Settings)
func (a *App) GetCHSLDCount() (int, error) {
	var count int
	err := a.db.QueryRow(`SELECT COUNT(*) FROM T_CHSLD`).Scan(&count)
	return count, err
}

// =============================
// Helpers internes
// =============================

func (a *App) emitSyncProgress(p SyncProgress) {
	// Wails v2 : runtime.EventsEmit
	// Wails v3 : a.ctx events — ajuster selon ta version
	if a.ctx != nil {
		// Wails v2
		runtime.EventsEmit(a.ctx, "chsld:sync:progress", p)

	}
}

func (a *App) chsldURLExists(sourceURL string) (bool, error) {
	var count int
	err := a.db.QueryRow(`SELECT COUNT(*) FROM T_CHSLD WHERE SourceURL = ?`, sourceURL).Scan(&count)
	return count > 0, err
}

func (a *App) insertCHSLDFromCrawl(d *msssData) error {
	_, err := a.db.Exec(`
		INSERT OR IGNORE INTO T_CHSLD
			(Region, TitreCHSLD, Adresse, Municipalite, CodePostal, Telephone,
			 telecopieur, Poste_Garde_infirmiere, Capacite, TypeCHSLD, Proprietaire,
			 DateCertification, Statut, SourceURL, InfosCHSLD, Notes, DateAjout, DateMaj)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		d.Region, d.TitreCHSLD, d.Adresse, d.Municipalite, d.CodePostal,
		d.Telephone, d.Telecopieur, d.PosteGardeInfirmiere, d.Capacite,
		d.TypeCHSLD, d.Proprietaire, d.DateCertification, d.Statut,
		d.SourceURL, d.InfosCHSLD, d.Notes,
		time.Now().Format("2006-01-02"),
		time.Now().Format("2006-01-02"),
	)
	return err
}

// =============================
// Crawler MSSS (logique pure)
// =============================

type msssInstallLink struct {
	Region string
	Name   string
	URL    string
}

type msssData struct {
	Region               string
	TitreCHSLD           string
	Adresse              string
	Municipalite         string
	CodePostal           string
	Telephone            string
	Telecopieur          string
	PosteGardeInfirmiere string
	Capacite             int
	TypeCHSLD            string
	Proprietaire         string
	DateCertification    string
	Statut               string
	SourceURL            string
	InfosCHSLD           string
	Notes                string
}

func msssListInstallations(regionCode, regionName string) ([]msssInstallLink, error) {
	url := fmt.Sprintf("%s/M02ListeInstall.asp?cdRss=%s&CodeTri=&Install=", msssBaseURL, regionCode)
	body, err := msssFetchURL(url)
	if err != nil {
		return nil, err
	}

	var links []msssInstallLink
	re := regexp.MustCompile(`href="(M02Installation\.asp\?CdIntervSocSan=\d+)"[^>]*>\s*([^<]+)`)
	for _, m := range re.FindAllStringSubmatch(body, -1) {
		if len(m) < 3 {
			continue
		}
		name := msssCleanText(m[2])
		name = regexp.MustCompile(`\s*\(\d{4}-\d{4}\)\s*$`).ReplaceAllString(name, "")
		links = append(links, msssInstallLink{
			Region: regionName,
			Name:   strings.TrimSpace(name),
			URL:    msssBaseURL + "/" + m[1],
		})
	}
	return links, nil
}

func msssFetchDetail(lk msssInstallLink) (*msssData, error) {
	body, err := msssFetchURL(lk.URL)
	if err != nil {
		return nil, err
	}

	d := &msssData{
		Region:     lk.Region,
		TitreCHSLD: lk.Name,
		SourceURL:  lk.URL,
	}
	msssParsePage(body, d)
	return d, nil
}

// msssParsePage parse la page de détail MSSS via DOM (span.titreColonne → td suivante)
func msssParsePage(body string, d *msssData) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return
	}

	type tdCell struct {
		text    string
		isLabel bool
		label   string
	}
	var cells []tdCell

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.EqualFold(n.Data, "td") {
			cell := tdCell{text: msssCleanText(msssExtractText(n))}
			if span := msssFindSpan(n, "titreColonne"); span != nil {
				cell.isLabel = true
				cell.label = msssCleanText(msssExtractText(span))
			}
			cells = append(cells, cell)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	for i := 0; i < len(cells)-1; i++ {
		if !cells[i].isLabel {
			continue
		}
		labelUp := strings.ToUpper(cells[i].label)
		// Valeur = prochaine td non-label dans les 3 suivantes
		value := ""
		for j := i + 1; j < len(cells) && j <= i+3; j++ {
			if !cells[j].isLabel && cells[j].text != "" {
				value = cells[j].text
				break
			}
		}
		if value == "" {
			continue
		}
		msssApply(labelUp, value, d)
	}
}

func msssApply(labelUp, value string, d *msssData) {
	reCode := regexp.MustCompile(`\s*\(\d+\)\s*$`)
	switch {
	case strings.Contains(labelUp, "NOM LÉGAL") || strings.Contains(labelUp, "NOM LEGAL"):
		if d.TitreCHSLD == "" {
			d.TitreCHSLD = value
		}
	case strings.Contains(labelUp, "ADRESSE"):
		d.Adresse = value
	case strings.Contains(labelUp, "MUNICIPALIT"):
		d.Municipalite = reCode.ReplaceAllString(value, "")
	case strings.Contains(labelUp, "CODE POSTAL"):
		d.CodePostal = value
	case strings.Contains(labelUp, "TÉLÉPHONE") || strings.Contains(labelUp, "TELEPHONE"):
		if d.Telephone == "" {
			d.Telephone = value
		}
	case strings.Contains(labelUp, "TÉLÉCOPIEUR") || strings.Contains(labelUp, "TELECOPIEUR"):
		d.Telecopieur = value
	case strings.Contains(labelUp, "STATUT"):
		d.Statut = value
	case strings.Contains(labelUp, "TYPE"):
		d.TypeCHSLD = value
	case strings.Contains(labelUp, "PROPRIÉTAIRE") || strings.Contains(labelUp, "PROPRIETAIRE"):
		d.Proprietaire = value
	case strings.Contains(labelUp, "CERTIF"):
		d.DateCertification = value
	case strings.Contains(labelUp, "CAPACIT"):
		reNum := regexp.MustCompile(`\d+`)
		if n := reNum.FindString(value); n != "" {
			fmt.Sscanf(n, "%d", &d.Capacite)
		}
	case strings.Contains(labelUp, "GARDE") || strings.Contains(labelUp, "POSTE"):
		d.PosteGardeInfirmiere = value
	case strings.Contains(labelUp, "INFORMATION") || strings.Contains(labelUp, "INFO"):
		d.InfosCHSLD = value
	case strings.Contains(labelUp, "NOTE"):
		d.Notes = value
	}
}

func msssFindSpan(n *html.Node, class string) *html.Node {
	if n.Type == html.ElementNode && strings.EqualFold(n.Data, "span") {
		for _, a := range n.Attr {
			if strings.EqualFold(a.Key, "class") && strings.Contains(a.Val, class) {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := msssFindSpan(c, class); found != nil {
			return found
		}
	}
	return nil
}

func msssExtractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var b strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		b.WriteString(msssExtractText(c))
	}
	return b.String()
}

func msssCleanText(s string) string {
	s = strings.ReplaceAll(s, "\u00a0", " ")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

func msssFetchURL(url string) (string, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * time.Second)
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; CHSLD-App/1.0)")
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = err
			continue
		}
		return msssDecodeWin1252(b), nil
	}
	return "", fmt.Errorf("fetch failed after 3 attempts: %w", lastErr)
}

func msssDecodeWin1252(b []byte) string {
	extra := map[byte]rune{
		0x80: '\u20AC', 0x82: '\u201A', 0x83: '\u0192', 0x84: '\u201E',
		0x85: '\u2026', 0x86: '\u2020', 0x87: '\u2021', 0x88: '\u02C6',
		0x89: '\u2030', 0x8A: '\u0160', 0x8B: '\u2039', 0x8C: '\u0152',
		0x8E: '\u017D', 0x91: '\u2018', 0x92: '\u2019', 0x93: '\u201C',
		0x94: '\u201D', 0x95: '\u2022', 0x96: '\u2013', 0x97: '\u2014',
		0x98: '\u02DC', 0x99: '\u2122', 0x9A: '\u0161', 0x9B: '\u203A',
		0x9C: '\u0153', 0x9E: '\u017E', 0x9F: '\u0178',
	}
	runes := make([]rune, 0, len(b))
	for _, c := range b {
		if c < 0x80 {
			runes = append(runes, rune(c))
		} else if r, ok := extra[c]; ok {
			runes = append(runes, r)
		} else {
			runes = append(runes, rune(c))
		}
	}
	return string(runes)
}

// Assurer que context est disponible sur App — à adapter si déjà présent dans app.go
// type App struct {
//     ctx context.Context
//     db  *sql.DB
// }
// func (a *App) startup(ctx context.Context) { a.ctx = ctx }

// Pour Wails v2, décommenter dans emitSyncProgress :
// runtime.EventsEmit(a.ctx, "chsld:sync:progress", p)
// Et importer : "github.com/wailsapp/wails/v2/pkg/runtime"

// Pour Wails v3 : utiliser a.runtime.Events.Emit(...)
var _ context.Context // évite import inutilisé si runtime non encore importé
func isChsldCandidate(name string) bool {
	// Normaliser : majuscules + enlever accents pour comparaison robuste
	normalized := strings.ToUpper(name)
	normalized = strings.NewReplacer(
		"É", "E", "È", "E", "Ê", "E", "Ë", "E",
		"À", "A", "Â", "A", "Ä", "A",
		"Î", "I", "Ï", "I",
		"Ô", "O", "Ö", "O",
		"Û", "U", "Ù", "U", "Ü", "U",
		"Ç", "C",
	).Replace(normalized)

	keywords := []string{
		"HEBERGEMENT", // CENTRE D'HÉBERGEMENT, CENTRE D'HEBERGEMENT
		"CHSLD",
		"SOINS DE LONGUE",  // SOINS DE LONGUE DURÉE
		"MAISON DES AINES", // MAISON DES AÎNÉS
		"FOYER",            // FOYER MAMO, FOYER SAINTS-ANGES
		"PAVILLON",         // PAVILLON SAINTE-MARIE
		"RESIDENCE",        // RÉSIDENCE
		"VILLA ",           // VILLA BONHEUR etc (espace pour éviter faux positifs)
		"MANOIR",
		"CARPE DIEM", // maisons Alzheimer
		"MAISON DE SOINS",
	}

	for _, kw := range keywords {
		if strings.Contains(normalized, kw) {
			return true
		}
	}
	return false
}
