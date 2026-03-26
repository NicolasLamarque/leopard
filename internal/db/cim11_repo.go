package database

import (
	"fmt"
	models "leopard/internal/model"
)

// ── LECTURE ───────────────────────────────────────────────────────

func (db *Database) SearchCim11(query string, lang string, limit int) ([]models.Cim11Code, error) {
	if limit <= 0 {
		limit = 20
	}
	var sqlQuery string
	if lang == "fr" {
		sqlQuery = `
			SELECT id, code, titre_fr, titre_en,
			       chapitre_code, chapitre_titre,
			       bloc_code, bloc_titre,
			       parent_code, niveau, is_billable, actif
			FROM cim11_codes
			WHERE actif = 1
			  AND (code LIKE ? OR titre_fr LIKE ?)
			ORDER BY
			  CASE WHEN code = ?    THEN 0
			       WHEN code LIKE ? THEN 1
			       ELSE 2 END,
			  niveau, titre_fr
			LIMIT ?`
	} else {
		sqlQuery = `
			SELECT id, code, titre_fr, titre_en,
			       chapitre_code, chapitre_titre,
			       bloc_code, bloc_titre,
			       parent_code, niveau, is_billable, actif
			FROM cim11_codes
			WHERE actif = 1
			  AND (code LIKE ? OR titre_en LIKE ?)
			ORDER BY
			  CASE WHEN code = ?    THEN 0
			       WHEN code LIKE ? THEN 1
			       ELSE 2 END,
			  niveau, titre_en
			LIMIT ?`
	}
	q := "%" + query + "%"
	var codes []models.Cim11Code
	err := db.Select(&codes, sqlQuery, q, q, query, query+"%", limit)
	if err != nil {
		return nil, fmt.Errorf("SearchCim11: %w", err)
	}
	return codes, nil
}

func (db *Database) GetCim11ByCode(code string) (*models.Cim11Code, error) {
	var c models.Cim11Code
	err := db.Get(&c, `
		SELECT id, code, titre_fr, titre_en,
		       chapitre_code, chapitre_titre,
		       bloc_code, bloc_titre,
		       parent_code, niveau, is_billable, actif
		FROM cim11_codes
		WHERE code = ? AND actif = 1`, code)
	if err != nil {
		return nil, fmt.Errorf("GetCim11ByCode %s: %w", code, err)
	}
	return &c, nil
}
func (db *Database) GetCim11Chapitres() ([]models.Cim11Code, error) {
	var chapitres []models.Cim11Code
	err := db.Select(&chapitres, `
		SELECT 
			0 AS id,
			chapitre_code AS code,
			chapitre_titre AS titre_fr,
			chapitre_titre AS titre_en,
			chapitre_code AS chapitre_code,
			chapitre_titre AS chapitre_titre,
			'' AS bloc_code, '' AS bloc_titre,
			'' AS parent_code,
			0 AS niveau, 0 AS is_billable, 1 AS actif
		FROM cim11_codes
		WHERE actif = 1 
  AND chapitre_code != '' 
  AND chapitre_titre != ''
  AND chapitre_titre != chapitre_code
GROUP BY chapitre_code
ORDER BY chapitre_code`)
	if err != nil {
		return nil, fmt.Errorf("GetCim11Chapitres: %w", err)
	}
	return chapitres, nil
}
func (db *Database) GetCim11Enfants(parentCode string) ([]models.Cim11Code, error) {
	var enfants []models.Cim11Code
	err := db.Select(&enfants, `
		SELECT id, code, titre_fr, titre_en,
		       chapitre_code, chapitre_titre,
		       bloc_code, bloc_titre,
		       parent_code, niveau, is_billable, actif
		FROM cim11_codes
		WHERE parent_code = ? AND actif = 1
		ORDER BY code`, parentCode)
	if err != nil {
		return nil, fmt.Errorf("GetCim11Enfants %s: %w", parentCode, err)
	}
	return enfants, nil
}

// ── PAGINATION ────────────────────────────────────────────────────

type Cim11PageResult struct {
	Codes []models.Cim11Code `json:"codes"`
	Total int                `json:"total"`
}

func (db *Database) GetCim11Page(query, chapitreCode, lang string, limit, offset int) (*Cim11PageResult, error) {
	col := "titre_fr"
	if lang == "en" {
		col = "titre_en"
	}

	conditions := "WHERE actif = 1"
	args := []interface{}{}

	if query != "" && query != "%" {
		conditions += fmt.Sprintf(" AND (code LIKE ? OR %s LIKE ?)", col)
		like := "%" + query + "%"
		args = append(args, like, like)
	}
	if chapitreCode != "" {
		conditions += " AND chapitre_code = ?"
		args = append(args, chapitreCode)
	}

	var total int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)
	db.QueryRow("SELECT COUNT(*) FROM cim11_codes "+conditions, countArgs...).Scan(&total)

	sqlQuery := fmt.Sprintf(`
		SELECT id, code, titre_fr, titre_en, chapitre_code, chapitre_titre,
		       bloc_code, bloc_titre, parent_code, niveau, is_billable, actif
		FROM cim11_codes %s ORDER BY %s LIMIT ? OFFSET ?`, conditions, col)

	args = append(args, limit, offset)
	var codes []models.Cim11Code
	err := db.Select(&codes, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("GetCim11Page: %w", err)
	}
	if codes == nil {
		codes = []models.Cim11Code{}
	}
	return &Cim11PageResult{Codes: codes, Total: total}, nil
}

// ── SYNC META ─────────────────────────────────────────────────────

func (db *Database) GetCim11SyncMeta() (*models.Cim11SyncMeta, error) {
	var meta models.Cim11SyncMeta
	err := db.Get(&meta, `
		SELECT derniere_sync, version_oms, nb_codes_fr, nb_codes_en, statut
		FROM cim11_sync_meta WHERE id = 1`)
	if err != nil {
		return nil, fmt.Errorf("GetCim11SyncMeta: %w", err)
	}
	var total int
	db.QueryRow(`SELECT COUNT(*) FROM cim11_codes WHERE actif = 1`).Scan(&total)
	if total > 0 {
		meta.NbCodesFr = total
		meta.NbCodesEn = total
		if meta.Statut == "jamais" {
			meta.Statut = "ok"
		}
	}
	return &meta, nil
}

func (db *Database) UpdateCim11SyncMeta(statut string, nbFr, nbEn int, version string) error {
	_, err := db.Exec(`
		UPDATE cim11_sync_meta SET
			derniere_sync = CURRENT_TIMESTAMP,
			version_oms   = ?,
			nb_codes_fr   = ?,
			nb_codes_en   = ?,
			statut        = ?
		WHERE id = 1`, version, nbFr, nbEn, statut)
	if err != nil {
		return fmt.Errorf("UpdateCim11SyncMeta: %w", err)
	}
	return nil
}

// ── CRUD MANUEL ───────────────────────────────────────────────────

func (db *Database) CreateCim11Code(c models.Cim11Code) (int64, error) {
	result, err := db.NamedExec(`
		INSERT INTO cim11_codes (
			code, titre_fr, titre_en,
			chapitre_code, chapitre_titre,
			bloc_code, bloc_titre,
			parent_code, niveau, is_billable, actif, updated_at
		) VALUES (
			:code, :titre_fr, :titre_en,
			:chapitre_code, :chapitre_titre,
			:bloc_code, :bloc_titre,
			:parent_code, :niveau, :is_billable, 1, CURRENT_TIMESTAMP
		)`, c)
	if err != nil {
		return 0, fmt.Errorf("CreateCim11Code %s: %w", c.Code, err)
	}
	return result.LastInsertId()
}

func (db *Database) UpdateCim11Code(c models.Cim11Code) error {
	_, err := db.NamedExec(`
		UPDATE cim11_codes SET
			titre_fr       = :titre_fr,
			titre_en       = :titre_en,
			chapitre_code  = :chapitre_code,
			chapitre_titre = :chapitre_titre,
			bloc_code      = :bloc_code,
			bloc_titre     = :bloc_titre,
			parent_code    = :parent_code,
			niveau         = :niveau,
			is_billable    = :is_billable,
			updated_at     = CURRENT_TIMESTAMP
		WHERE id = :id`, c)
	if err != nil {
		return fmt.Errorf("UpdateCim11Code %d: %w", c.ID, err)
	}
	return nil
}

func (db *Database) DeleteCim11Code(id int) error {
	_, err := db.Exec(`
		UPDATE cim11_codes
		SET actif = 0, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("DeleteCim11Code %d: %w", id, err)
	}
	return nil
}

func (db *Database) UpsertCim11Code(c models.Cim11Code) (int, int, error) {
	result, err := db.NamedExec(`
		INSERT INTO cim11_codes (
			code, titre_fr, titre_en,
			chapitre_code, chapitre_titre,
			bloc_code, bloc_titre,
			parent_code, niveau, is_billable, actif, updated_at
		) VALUES (
			:code, :titre_fr, :titre_en,
			:chapitre_code, :chapitre_titre,
			:bloc_code, :bloc_titre,
			:parent_code, :niveau, :is_billable, 1, CURRENT_TIMESTAMP
		)
		ON CONFLICT(code) DO UPDATE SET
			titre_fr       = :titre_fr,
			titre_en       = :titre_en,
			chapitre_code  = :chapitre_code,
			chapitre_titre = :chapitre_titre,
			bloc_code      = :bloc_code,
			bloc_titre     = :bloc_titre,
			parent_code    = :parent_code,
			niveau         = :niveau,
			is_billable    = :is_billable,
			updated_at     = CURRENT_TIMESTAMP`, c)
	if err != nil {
		return 0, 0, fmt.Errorf("UpsertCim11Code %s: %w", c.Code, err)
	}
	rows, _ := result.RowsAffected()
	if rows == 1 {
		return 1, 0, nil
	}
	return 0, 1, nil
}
