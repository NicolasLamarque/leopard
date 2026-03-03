package database

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	models "leopard/internal/model"
)

// ─── MUNICIPALITÉS ────────────────────────────────────────────────────────────

func (d *Database) GetAllMunicipalities() ([]models.Municipality, error) {
	var muns []models.Municipality
	err := d.Select(&muns, `
		SELECT mcode, munnom, madr1, mcodpos, mtel, mfax, mcourriel, mweb,
		       regadm, mrc, maire, dirgen, msuperf, mpopul,
		       polic, incen, loisir, trvpub, mesurg, urban, communic, permis, batim
		FROM MUN ORDER BY munnom`)
	return muns, err
}

func (d *Database) GetMunicipalityByID(mcode string) (*models.Municipality, error) {
	var m models.Municipality
	err := d.Get(&m, `
		SELECT mcode, munnom, madr1, mcodpos, mtel, mfax, mcourriel, mweb,
		       regadm, mrc, maire, dirgen, msuperf, mpopul,
		       polic, incen, loisir, trvpub, mesurg, urban, communic, permis, batim
		FROM MUN WHERE mcode = ?`, mcode)
	return &m, err
}

func (d *Database) GetMunicipalityByName(name string) (*models.Municipality, error) {
	var m models.Municipality
	err := d.Get(&m, `
		SELECT mcode, munnom, madr1, mcodpos, mtel, mfax, mcourriel, mweb,
		       regadm, mrc, maire, dirgen, msuperf, mpopul,
		       polic, incen, loisir, trvpub, mesurg, urban, communic, permis, batim
		FROM MUN WHERE munnom LIKE ? ORDER BY munnom LIMIT 1`, "%"+name+"%")
	return &m, err
}

// ─── ARRONDISSEMENTS ──────────────────────────────────────────────────────────

func (d *Database) GetArrondissementsByMun(mcode string) ([]models.Arrondissement, error) {
	var arrs []models.Arrondissement
	err := d.Select(&arrs,
		`SELECT arrcod, arrnom, arrcode FROM ARR WHERE arrcode = ? ORDER BY arrnom`,
		mcode)
	return arrs, err
}

func (d *Database) AddArrondissement(arr models.Arrondissement) error {
	_, err := d.NamedExec(
		`INSERT INTO ARR (arrcod, arrnom, arrcode) 
         VALUES (:arrcod, :arrnom, :arrcode)`, arr)
	return err
}

func (d *Database) UpdateArrondissement(arr models.Arrondissement) error {
	_, err := d.NamedExec(
		`UPDATE ARR SET arrnom = :arrnom 
         WHERE arrcod = :arrcod`, arr)
	return err
}

func (d *Database) DeleteArrondissement(arrcod string) error {
	_, err := d.Exec(
		`DELETE FROM ARR WHERE arrcod = ?`, arrcod)
	return err
}

// ─── IMPORT CSV ───────────────────────────────────────────────────────────────

func (d *Database) ImportMunCSV(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("impossible d'ouvrir %s : %w", filePath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // Ignore le nombre de colonnes variable

	headers, err := reader.Read()
	if err != nil {
		return 0, fmt.Errorf("erreur lecture entête: %w", err)
	}
	if len(headers) > 0 {
		headers[0] = strings.TrimPrefix(headers[0], "\xef\xbb\xbf")
	}
	idx := munBuildIndex(headers)
	log.Printf("[MUN] Entête lue: %d colonnes", len(headers))

	tx, err := d.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `
		INSERT OR IGNORE INTO MUN (
			mcode, munnom, madr1, mcodpos, mtel, mfax, mcourriel, mweb,
			regadm, mrc, maire, dirgen, msuperf, mpopul,
			polic, incen, loisir, trvpub, mesurg, urban, communic, permis, batim
		) VALUES (
			:mcode, :munnom, :madr1, :mcodpos, :mtel, :mfax, :mcourriel, :mweb,
			:regadm, :mrc, :maire, :dirgen, :msuperf, :mpopul,
			:polic, :incen, :loisir, :trvpub, :mesurg, :urban, :communic, :permis, :batim
		)`

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("[MUN] Erreur lecture ligne %d: %v", count+1, err)
			continue
		}
		data := map[string]interface{}{
			"mcode":     munCol(record, idx, "mcode"),
			"munnom":    munCol(record, idx, "munnom"),
			"madr1":     munCol(record, idx, "madr1"),
			"mcodpos":   munCol(record, idx, "mcodpos"),
			"mtel":      munCol(record, idx, "mtel"),
			"mfax":      munCol(record, idx, "mfax"),
			"mcourriel": munCol(record, idx, "mcourriel"),
			"mweb":      munCol(record, idx, "mweb"),
			"regadm":    munCol(record, idx, "regadm"),
			"mrc":       munCol(record, idx, "mrc"),
			"maire":     munCol(record, idx, "maire"),
			"dirgen":    munCol(record, idx, "dirgen"),
			"msuperf":   munParseFloat(munCol(record, idx, "msuperf")),
			"mpopul":    munParseInt(munCol(record, idx, "mpopul")),
			"polic":     munCol(record, idx, "polic"),
			"incen":     munCol(record, idx, "incen"),
			"loisir":    munCol(record, idx, "loisir"),
			"trvpub":    munCol(record, idx, "trvpub"),
			"mesurg":    munCol(record, idx, "mesurg"),
			"urban":     munCol(record, idx, "urban"),
			"communic":  munCol(record, idx, "communic"),
			"permis":    munCol(record, idx, "permis"),
			"batim":     munCol(record, idx, "batim"),
		}
		_, err = tx.NamedExec(query, data)
		if err != nil {
			log.Printf("[MUN] Erreur ligne %d (mcode=%s): %v", count+1, data["mcode"], err)
		} else {
			count++
		}
	}
	log.Printf("[MUN] Import terminé: %d lignes insérées", count)
	commitErr := tx.Commit()
	if commitErr != nil {
		log.Printf("[MUN] ERREUR COMMIT: %v", commitErr)
	} else {
		log.Printf("[MUN] Commit OK")
	}
	return count, commitErr
}

func (d *Database) ImportArrCSV(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("impossible d'ouvrir %s : %w", filePath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // Ignore le nombre de colonnes variable

	headers, err := reader.Read()
	if err != nil {
		return 0, err
	}
	if len(headers) > 0 {
		headers[0] = strings.TrimPrefix(headers[0], "\xef\xbb\xbf")
	}
	idx := munBuildIndex(headers)
	log.Printf("[ARR] Entête lue: %d colonnes", len(headers))

	tx, err := d.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	queryArr := `
		INSERT OR IGNORE INTO ARR (arrcod, arrnom, arrcode)
		VALUES (:arrcod, :arrnom, :arrcode)`

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("[ARR] Erreur lecture ligne %d: %v", count+1, err)
			continue
		}
		data := map[string]interface{}{
			"arrcod":  munCol(record, idx, "arrcod"),
			"arrnom":  munCol(record, idx, "arrnom"),
			"arrcode": munCol(record, idx, "arrcode"),
		}
		_, err = tx.NamedExec(queryArr, data)
		if err != nil {
			log.Printf("[ARR] Erreur ligne %d (arrcod=%s): %v", count+1, data["arrcod"], err)
		} else {
			count++
		}
	}
	log.Printf("[ARR] Import terminé: %d lignes insérées", count)
	commitErr := tx.Commit()
	if commitErr != nil {
		log.Printf("[ARR] ERREUR COMMIT: %v", commitErr)
	} else {
		log.Printf("[ARR] Commit OK")
	}
	return count, commitErr
}

// ─── HELPERS (préfixés mun pour éviter conflits avec les autres repos) ────────

func munBuildIndex(headers []string) map[string]int {
	m := make(map[string]int, len(headers))
	for i, h := range headers {
		// Nettoyer BOM, espaces ET guillemets doubles du CSV
		h = strings.TrimSpace(h)
		h = strings.Trim(h, `"`)
		h = strings.ToLower(h)
		m[h] = i
	}
	return m
}

func munCol(record []string, idx map[string]int, name string) string {
	i, ok := idx[name]
	if !ok || i >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[i])
}

func munParseFloat(s string) float64 {
	s = strings.ReplaceAll(s, ",", ".")
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func munParseInt(s string) int {
	v, _ := strconv.Atoi(strings.TrimSpace(s))
	return v
}
