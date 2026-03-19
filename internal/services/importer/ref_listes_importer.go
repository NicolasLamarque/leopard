package importer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RefListeRow représente une ligne CSV transformée et validée.
type RefListeRow struct {
	Categorie string
	Libelle   string
	Couleur   string
	Icone     string
	Ordre     int
}

// ProcessRefListes transforme les lignes brutes CSV en structures RefListeRow.
// Format attendu (en-tête obligatoire) :
//
//	categorie,libelle,couleur,icone,ordre
//
// Les colonnes couleur / icone / ordre sont optionnelles.
func (p *Pipeline) ProcessRefListes(rows [][]string) ([]RefListeRow, error) {
	if len(rows) < 2 {
		return nil, fmt.Errorf("le fichier est vide ou ne contient que des entêtes")
	}

	// 1. Extraction et validation des entêtes
	headers := rows[0]
	if err := p.validateRefListesHeaders(headers); err != nil {
		return nil, err
	}

	// 2. Mapping dynamique des colonnes (insensible à l'ordre)
	colIndex := make(map[string]int)
	for i, h := range headers {
		colIndex[strings.ToLower(strings.TrimSpace(h))] = i
	}

	var finalData []RefListeRow

	// 3. Boucle de traitement
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		// Protection contre les lignes courtes
		rowData := make(map[string]string)
		for j, headerName := range headers {
			key := strings.ToLower(strings.TrimSpace(headerName))
			if j < len(row) {
				rowData[key] = strings.TrimSpace(row[j])
			} else {
				rowData[key] = ""
			}
		}

		// 4. Transformation
		item := mapToRefListe(rowData)

		// 5. Validation métier minimale
		if item.Categorie != "" && item.Libelle != "" {
			finalData = append(finalData, item)
		}
	}

	return finalData, nil
}

// validateRefListesHeaders vérifie que les colonnes obligatoires sont présentes.
func (p *Pipeline) validateRefListesHeaders(headers []string) error {
	required := []string{"categorie", "libelle"}

	headerMap := make(map[string]bool)
	for _, h := range headers {
		headerMap[strings.ToLower(strings.TrimSpace(h))] = true
	}

	for _, req := range required {
		if !headerMap[req] {
			return fmt.Errorf("colonne manquante dans le CSV : %s", req)
		}
	}
	return nil
}

// mapToRefListe transforme une ligne (map colonne→valeur) en RefListeRow.
func mapToRefListe(row map[string]string) RefListeRow {
	couleur := row["couleur"]
	if couleur == "" {
		couleur = "slate"
	}

	ordre := 0
	if raw := row["ordre"]; raw != "" {
		if n, err := strconv.Atoi(raw); err == nil {
			ordre = n
		}
	}

	return RefListeRow{
		Categorie: row["categorie"],
		Libelle:   row["libelle"],
		Couleur:   couleur,
		Icone:     row["icone"],
		Ordre:     ordre,
	}
}

// ReadCSV lit un fichier CSV et retourne les lignes brutes.
// Utilisé à la place de ReadExcel pour les fichiers .csv
func (p *Pipeline) ReadCSV() ([][]string, error) {
	f, err := os.Open(p.FilePath)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier CSV: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true
	reader.Comment = '#'

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erreur lecture CSV: %w", err)
	}

	return rows, nil
}
