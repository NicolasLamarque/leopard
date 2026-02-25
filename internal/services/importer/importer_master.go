package importer

import "fmt"

// ProcessNotaires transforme les lignes brutes en structures NotaireRow
func (p *Pipeline) ProcessNotaires(rows [][]string) ([]NotaireRow, error) {
	if len(rows) < 2 {
		return nil, fmt.Errorf("le fichier est vide ou ne contient que des entêtes")
	}

	// 1. Extraction et validation des entêtes
	headers := rows[0]
	if err := p.validateNotaireHeaders(headers); err != nil {
		return nil, err
	}

	var finalData []NotaireRow

	// 2. Boucle de traitement
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		// Protection contre les lignes qui n'ont pas le bon nombre de colonnes
		rowData := make(map[string]string)
		for j, headerName := range headers {
			if j < len(row) {
				rowData[headerName] = row[j]
			} else {
				rowData[headerName] = "" // Cellule manquante en fin de ligne
			}
		}

		// 3. Transformation (Mapping)
		notaire := MapToNotaire(rowData)

		// 4. Validation métier minimale (Loi 25 : On ne stocke pas de vide inutile)
		if notaire.Nom != "" && notaire.Prenom != "" {
			finalData = append(finalData, notaire)
		}
	}

	return finalData, nil
}

// validateNotaireHeaders vérifie que l'utilisateur a bien fourni le bon fichier
func (p *Pipeline) validateNotaireHeaders(headers []string) error {
	required := []string{"Nom", "Prenom", "Email"} // Le strict minimum pour que ce soit utile

	headerMap := make(map[string]bool)
	for _, h := range headers {
		headerMap[h] = true
	}

	for _, req := range required {
		if !headerMap[req] {
			return fmt.Errorf("colonne manquante dans l'Excel : %s", req)
		}
	}
	return nil
}
