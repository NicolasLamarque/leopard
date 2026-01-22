// internal/db/medecin_import.go
package database

import (
	"fmt"
	models "leopard/internal/model"
	"strings"

	"github.com/xuri/excelize/v2"
)

// ImportMedecinsFromExcel importe les médecins depuis un fichier Excel
func (db *Database) ImportMedecinsFromExcel(filePath string, createdBy int) (int, error) {
	fmt.Printf("\n🚀 DÉMARRAGE IMPORT MÉDECINS\n")
	fmt.Printf("   📁 Fichier: %s\n", filePath)

	// Ouvrir le fichier Excel
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("erreur ouverture fichier Excel: %w", err)
	}
	defer f.Close()

	// Lire la première feuille
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return 0, fmt.Errorf("erreur lecture des lignes: %w", err)
	}

	if len(rows) < 2 {
		return 0, fmt.Errorf("le fichier Excel est vide ou ne contient pas de données")
	}

	fmt.Printf("   📊 %d lignes trouvées (en-tête inclus)\n", len(rows))

	// Ignorer la ligne d'en-tête
	dataRows := rows[1:]
	importCount := 0
	errors := []string{}

	// Traiter chaque ligne
	for i, row := range dataRows {
		// S'assurer qu'il y a assez de colonnes (11 colonnes minimum selon votre exemple)
		if len(row) < 11 {
			errors = append(errors, fmt.Sprintf("Ligne %d: nombre de colonnes insuffisant (%d/11)", i+2, len(row)))
			continue
		}

		// Extraire les données (index correspond à l'ordre des colonnes Excel)
		// Permis | Civilité | Prénom | Nom | Statut | Spécialité | Service | Département | Installation principale | RLS | Courriel
		permis := strings.TrimSpace(row[0])
		civilite := strings.TrimSpace(row[1])
		prenom := strings.TrimSpace(row[2])
		nom := strings.TrimSpace(row[3])
		statut := strings.TrimSpace(row[4])
		specialite := strings.TrimSpace(row[5])
		service := strings.TrimSpace(row[6])
		departement := strings.TrimSpace(row[7])
		installation := strings.TrimSpace(row[8])
		rls := strings.TrimSpace(row[9])
		email := strings.TrimSpace(row[10])

		// Validation minimale
		if permis == "" || permis == "Permis" {
			errors = append(errors, fmt.Sprintf("Ligne %d: permis manquant ou ligne d'en-tête", i+2))
			continue
		}

		// Construire le nom complet
		nomComplet := fmt.Sprintf("%s %s %s", civilite, prenom, nom)
		nomComplet = strings.TrimSpace(nomComplet)

		// Déterminer si actif (1 si statut = "Actif", 0 sinon)
		actif := 0
		if strings.ToLower(statut) == "actif" {
			actif = 1
		}

		// Créer la requête
		req := models.CreateMedecinRequest{
			Licence:                permis,
			Civilite:               civilite,
			Prenom:                 prenom,
			Nom:                    nom,
			NomComplet:             nomComplet,
			Statut:                 statut,
			Specialisation:         specialite,
			Service:                service,
			Departement:            departement,
			InstallationPrincipale: installation,
			RLS:                    rls,
			Email:                  email,
			Pays:                   "Canada", // Valeur par défaut
			Actif:                  actif,
		}

		// Insérer dans la base de données
		_, err := db.CreateMedecin(req, createdBy)
		if err != nil {
			// Si erreur de duplicata (licence existe déjà), on peut l'ignorer ou mettre à jour
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				errors = append(errors, fmt.Sprintf("Ligne %d (Permis: %s): médecin existe déjà", i+2, permis))
			} else {
				errors = append(errors, fmt.Sprintf("Ligne %d (Permis: %s): %v", i+2, permis, err))
			}
			continue
		}

		importCount++
		if importCount%10 == 0 {
			fmt.Printf("   ⏳ %d médecins importés...\n", importCount)
		}
	}

	fmt.Printf("\n✅ IMPORT TERMINÉ\n")
	fmt.Printf("   ✔️  Succès: %d médecins\n", importCount)
	fmt.Printf("   ❌ Erreurs: %d\n", len(errors))

	// Retourner un résumé
	if len(errors) > 0 {
		errorMsg := strings.Join(errors, "\n")
		return importCount, fmt.Errorf("import terminé avec %d succès et %d erreurs:\n%s",
			importCount, len(errors), errorMsg)
	}

	return importCount, nil
}

// ImportMedecinsFromExcelWithUpdate importe et met à jour les médecins existants
func (db *Database) ImportMedecinsFromExcelWithUpdate(filePath string, createdBy int) (*ImportStats, error) {
	stats := &ImportStats{}

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur ouverture fichier: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("erreur lecture: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("fichier vide")
	}

	dataRows := rows[1:]

	for i, row := range dataRows {
		if len(row) < 11 {
			stats.Errors = append(stats.Errors, fmt.Sprintf("Ligne %d: colonnes insuffisantes", i+2))
			continue
		}

		permis := strings.TrimSpace(row[0])
		if permis == "" || permis == "Permis" {
			continue
		}

		civilite := strings.TrimSpace(row[1])
		prenom := strings.TrimSpace(row[2])
		nom := strings.TrimSpace(row[3])
		statut := strings.TrimSpace(row[4])
		specialite := strings.TrimSpace(row[5])
		service := strings.TrimSpace(row[6])
		departement := strings.TrimSpace(row[7])
		installation := strings.TrimSpace(row[8])
		rls := strings.TrimSpace(row[9])
		email := strings.TrimSpace(row[10])

		nomComplet := strings.TrimSpace(fmt.Sprintf("%s %s %s", civilite, prenom, nom))
		actif := 0
		if strings.ToLower(statut) == "actif" {
			actif = 1
		}

		// Vérifier si le médecin existe déjà
		existing, err := db.GetMedecinByLicence(permis)

		if err != nil {
			// N'existe pas, on crée
			req := models.CreateMedecinRequest{
				Licence:                permis,
				Civilite:               civilite,
				Prenom:                 prenom,
				Nom:                    nom,
				NomComplet:             nomComplet,
				Statut:                 statut,
				Specialisation:         specialite,
				Service:                service,
				Departement:            departement,
				InstallationPrincipale: installation,
				RLS:                    rls,
				Email:                  email,
				Pays:                   "Canada",
				Actif:                  actif,
			}

			_, err := db.CreateMedecin(req, createdBy)
			if err != nil {
				stats.Errors = append(stats.Errors, fmt.Sprintf("Ligne %d: %v", i+2, err))
			} else {
				stats.Nouveaux++
			}
		} else {
			// Existe, on met à jour
			updateReq := models.UpdateMedecinRequest{
				ID:                     existing.ID,
				Licence:                permis,
				Civilite:               civilite,
				Prenom:                 prenom,
				Nom:                    nom,
				NomComplet:             nomComplet,
				Statut:                 statut,
				Specialisation:         specialite,
				Service:                service,
				Departement:            departement,
				InstallationPrincipale: installation,
				RLS:                    rls,
				Email:                  email,
				Pays:                   "Canada",
				Actif:                  actif,
			}

			err := db.UpdateMedecin(updateReq)
			if err != nil {
				stats.Errors = append(stats.Errors, fmt.Sprintf("Ligne %d: %v", i+2, err))
			} else {
				stats.MisAJour++
			}
		}
	}

	return stats, nil
}

// ImportStats statistiques d'import
type ImportStats struct {
	Nouveaux int      `json:"nouveaux"`
	MisAJour int      `json:"mis_a_jour"`
	Errors   []string `json:"errors"`
}
