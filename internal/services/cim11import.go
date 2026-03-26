package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	database "leopard/internal/db"
	models "leopard/internal/model"
)

// ── Structs API OMS ───────────────────────────────────────────────

type icdTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type icdSearchResp struct {
	ResultChoices []icdSearchEntity `json:"destinationEntities"`
	ResultChopped bool              `json:"resultChopped"`
	Error         bool              `json:"error"`
	ErrorMessage  string            `json:"errorMessage"`
}

type icdSearchEntity struct {
	TheCode string `json:"theCode"`
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
}

// ── Service ───────────────────────────────────────────────────────

type Cim11ImportService struct {
	db          *database.Database
	http        *http.Client
	token       string
	tokenExpiry time.Time
	mu          sync.Mutex
}

const (
	icdTokenURL  = "https://icdaccessmanagement.who.int/connect/token"
	icdSearchURL = "https://id.who.int/icd/release/11/2024-01/mms/search"
	icdMmsURL    = "https://id.who.int/icd/release/11/2024-01/mms"
)

func NewCim11ImportService(db *database.Database) *Cim11ImportService {
	return &Cim11ImportService{
		db:   db,
		http: &http.Client{Timeout: 30 * time.Second},
	}
}

// ── Auth ──────────────────────────────────────────────────────────

func (s *Cim11ImportService) getToken(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.token != "" && time.Now().Before(s.tokenExpiry) {
		return s.token, nil
	}
	cid := os.Getenv("ICD_CLIENT_ID")
	csec := os.Getenv("ICD_CLIENT_SECRET")
	if cid == "" || csec == "" {
		return "", fmt.Errorf("ICD_CLIENT_ID / ICD_CLIENT_SECRET manquants dans .env")
	}
	body := url.Values{}
	body.Set("client_id", cid)
	body.Set("client_secret", csec)
	body.Set("scope", "icdapi_access")
	body.Set("grant_type", "client_credentials")

	req, _ := http.NewRequestWithContext(ctx, "POST", icdTokenURL, strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := s.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token OMS refusé (%d): %s", resp.StatusCode, b)
	}
	var t icdTokenResp
	json.NewDecoder(resp.Body).Decode(&t)
	s.token = t.AccessToken
	s.tokenExpiry = time.Now().Add(time.Duration(t.ExpiresIn-60) * time.Second)
	return s.token, nil
}

// ── Recherche avec détection de troncature ────────────────────────

func (s *Cim11ImportService) searchRaw(ctx context.Context, query, lang string) (*icdSearchResp, error) {
	tok, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("q", query)
	params.Set("includeKeywordResult", "false")
	params.Set("useFlexisearch", "false")
	params.Set("flatResults", "true")
	params.Set("highlightingEnabled", "false")
	params.Set("medicalCodingMode", "true")

	req, err := http.NewRequestWithContext(ctx, "GET", icdSearchURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", lang)
	req.Header.Set("API-Version", "v2")

	resp, err := s.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("search %d: %s", resp.StatusCode, b)
	}
	var result icdSearchResp
	json.NewDecoder(resp.Body).Decode(&result)
	// DEBUG TEMPORAIRE
	if len(result.ResultChoices) > 0 {
		fmt.Printf("🔍 SAMPLE: code=%s title=%s chapter=%s\n",
			result.ResultChoices[0].TheCode,
			result.ResultChoices[0].Title,
			result.ResultChoices[0].Chapter)
	}
	return &result, nil
}

// searchAll récupère TOUS les résultats pour une query.
// Si ResultChopped=true, subdivise automatiquement (ex: "a" → "aa","ab"..."az")
func (s *Cim11ImportService) searchAll(ctx context.Context, query, lang string, dest map[string]*codeEntry) error {
	result, err := s.searchRaw(ctx, query, lang)
	if err != nil {
		return err
	}

	// Ajouter les résultats reçus
	for _, e := range result.ResultChoices {
		if e.TheCode == "" {
			continue
		}
		if _, ok := dest[e.TheCode]; !ok {
			dest[e.TheCode] = &codeEntry{}
		}
		if lang == "fr" {
			dest[e.TheCode].titreFR = e.Title
		} else {
			dest[e.TheCode].titreEN = e.Title
			dest[e.TheCode].chapter = e.Chapter
		}
	}

	// Si tronqué ET query courte → subdiviser
	if result.ResultChopped && len(query) < 3 {
		alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"
		for _, c := range alphabet {
			subQuery := query + string(c)
			if err := s.searchAll(ctx, subQuery, lang, dest); err != nil {
				fmt.Printf("    ⚠️  '%s': %v\n", subQuery, err)
			}
		}
	}

	return nil
}

type codeEntry struct {
	titreFR string
	titreEN string
	chapter string
}

// ── Import principal ──────────────────────────────────────────────

func (s *Cim11ImportService) ImportAll(ctx context.Context) (*models.Cim11ImportResult, error) {
	start := time.Now()
	res := &models.Cim11ImportResult{}

	allCodes := make(map[string]*codeEntry)

	// Lettres de base — si ResultChopped, elles se subdivisent automatiquement
	baseQueries := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	fmt.Println("🔍 Récupération EN...")
	for _, q := range baseQueries {
		before := len(allCodes)
		if err := s.searchAll(ctx, q, "en", allCodes); err != nil {
			fmt.Printf("  ⚠️  EN '%s': %v\n", q, err)
			continue
		}
		fmt.Printf("  ✓ EN '%s' +%d (total: %d)\n", q, len(allCodes)-before, len(allCodes))
	}

	fmt.Println("🔍 Récupération FR...")
	for _, q := range baseQueries {
		if err := s.searchAll(ctx, q, "fr", allCodes); err != nil {
			fmt.Printf("  ⚠️  FR '%s': %v\n", q, err)
		}
		fmt.Printf("  ✓ FR '%s'\n", q)
	}

	fmt.Printf("💾 %d codes → SQLite...\n", len(allCodes))

	// Transaction unique
	tx, err := s.db.Begin()
	if err != nil {
		res.Erreur = err.Error()
		return res, err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO cim11_codes
			(code, titre_fr, titre_en, chapitre_code, chapitre_titre,
			 bloc_code, bloc_titre, parent_code, niveau, is_billable, actif, updated_at)
		VALUES (?, ?, ?, ?, ?, '', '', '', 2, 1, 1, CURRENT_TIMESTAMP)
		ON CONFLICT(code) DO UPDATE SET
			titre_fr       = excluded.titre_fr,
			titre_en       = excluded.titre_en,
			chapitre_code  = excluded.chapitre_code,
			chapitre_titre = excluded.chapitre_titre,
			updated_at     = CURRENT_TIMESTAMP`)
	if err != nil {
		tx.Rollback()
		res.Erreur = err.Error()
		return res, err
	}
	defer stmt.Close()

	inserted, updated := 0, 0
	for code, e := range allCodes {
		titreFR := e.titreFR
		if titreFR == "" {
			titreFR = e.titreEN
		}
		r, err := stmt.Exec(code, titreFR, e.titreEN, e.chapter, e.chapter)
		if err != nil {
			continue
		}
		if rows, _ := r.RowsAffected(); rows == 1 {
			inserted++
		} else {
			updated++
		}
	}

	if err := tx.Commit(); err != nil {
		res.Erreur = err.Error()
		return res, err
	}

	total := inserted + updated
	s.db.Exec(`
		UPDATE cim11_sync_meta
		SET derniere_sync=CURRENT_TIMESTAMP, version_oms='2024-01',
		    nb_codes_fr=?, nb_codes_en=?, statut='ok'
		WHERE id=1`, total, total)

	duree := time.Since(start).Round(time.Second)
	res.NbInseres = inserted
	res.NbMisAJour = updated
	res.Duree = duree.String()
	fmt.Printf("✅ %d insérés, %d MàJ en %s\n", inserted, updated, duree)
	return res, nil
}

// ── Recherche locale SQLite ───────────────────────────────────────

func (s *Cim11ImportService) SearchCim11(query, lang string, limit int) ([]models.Cim11Code, error) {
	if limit <= 0 {
		limit = 20
	}
	col := "titre_fr"
	if lang == "en" {
		col = "titre_en"
	}
	q := fmt.Sprintf(`
		SELECT id, code, titre_fr, titre_en, chapitre_code, chapitre_titre,
		       bloc_code, bloc_titre, parent_code, niveau, is_billable, actif
		FROM cim11_codes
		WHERE actif = 1
		  AND (code LIKE ? OR %s LIKE ?)
		ORDER BY
		  CASE WHEN code = ? THEN 0 WHEN code LIKE ? THEN 1 ELSE 2 END,
		  niveau, %s
		LIMIT ?`, col, col)

	like := "%" + query + "%"
	rows, err := s.db.Query(q, like, like, query, query+"%", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Cim11Code
	for rows.Next() {
		var c models.Cim11Code
		if err := rows.Scan(
			&c.ID, &c.Code, &c.TitreFr, &c.TitreEn,
			&c.ChapitreCode, &c.ChapitreTitre,
			&c.BlocCode, &c.BlocTitre,
			&c.ParentCode, &c.Niveau, &c.IsBillable, &c.Actif,
		); err != nil {
			continue
		}
		out = append(out, c)
	}
	if out == nil {
		out = []models.Cim11Code{}
	}
	return out, nil
}

// ── SyncMeta ──────────────────────────────────────────────────────

func (s *Cim11ImportService) GetSyncMeta() (*models.Cim11SyncMeta, error) {
	return s.db.GetCim11SyncMeta()
}
