package main

import (
	"errors"
	"fmt"
	database "leopard/internal/db"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ========== EXCEL FILE HANDLERS ==========

// ExcelFileType définit le type de fichier Excel attendu
type ExcelFileType string

const (
	ExcelTypeMedecins     ExcelFileType = "medecins"
	ExcelTypeIntervenants ExcelFileType = "intervenants"
	ExcelTypeClients      ExcelFileType = "clients"
	ExcelTypeContacts     ExcelFileType = "contacts"
	ExcelTypeRPA          ExcelFileType = "rpa"
	ExcelTypeCHSLD        ExcelFileType = "chsld"
	ExcelTypeGeneric      ExcelFileType = "generic"
)

// SelectExcelFile ouvre un dialogue pour sélectionner un fichier Excel générique
func (a *App) SelectExcelFile() (string, error) {
	return a.selectExcelFileWithType(ExcelTypeGeneric)
}

// SelectExcelFileForMedecins sélectionne un fichier Excel pour import de médecins
func (a *App) SelectExcelFileForMedecins() (string, error) {
	return a.selectExcelFileWithType(ExcelTypeMedecins)
}

// SelectExcelFileForIntervenants sélectionne un fichier Excel pour import d'intervenants
func (a *App) SelectExcelFileForIntervenants() (string, error) {
	return a.selectExcelFileWithType(ExcelTypeIntervenants)
}

// SelectExcelFileForClients sélectionne un fichier Excel pour import de clients
func (a *App) SelectExcelFileForClients() (string, error) {
	return a.selectExcelFileWithType(ExcelTypeClients)
}

// SelectExcelFileForRPA sélectionne un fichier Excel pour import de RPA
func (a *App) SelectExcelFileForRPA() (string, error) {
	return a.selectExcelFileWithType(ExcelTypeRPA)
}

// selectExcelFileWithType fonction interne pour gérer la sélection selon le type
func (a *App) selectExcelFileWithType(fileType ExcelFileType) (string, error) {
	// Créer le dossier data s'il n'existe pas
	dataDir := "./data"
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			return "", fmt.Errorf("impossible de créer le dossier data: %w", err)
		}
	}

	// Sous-dossier spécifique selon le type
	var subDir string
	var title string

	switch fileType {
	case ExcelTypeMedecins:
		subDir = "medecins"
		title = "Sélectionner un fichier Excel - Médecins"
	case ExcelTypeIntervenants:
		subDir = "intervenants"
		title = "Sélectionner un fichier Excel - Intervenants"
	case ExcelTypeClients:
		subDir = "clients"
		title = "Sélectionner un fichier Excel - Clients"
	case ExcelTypeRPA:
		subDir = "rpa"
		title = "Sélectionner un fichier Excel - Résidences"
	case ExcelTypeCHSLD:
		subDir = "chsld"
		title = "Sélectionner un fichier Excel - CHSLD"
	default:
		subDir = ""
		title = "Sélectionner un fichier Excel"
	}

	// Créer le sous-dossier si nécessaire
	targetDir := dataDir
	if subDir != "" {
		targetDir = filepath.Join(dataDir, subDir)
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				return "", fmt.Errorf("impossible de créer le dossier %s: %w", targetDir, err)
			}
		}
	}

	absTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		return "", fmt.Errorf("erreur chemin absolu: %w", err)
	}

	// Ouvrir le dialogue
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            title,
		DefaultDirectory: absTargetDir,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Fichiers Excel",
				Pattern:     "*.xlsx;*.xls",
			},
			{
				DisplayName: "Tous les fichiers",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		return "", err
	}

	// Vérifier que le fichier existe
	if filePath == "" {
		return "", errors.New("aucun fichier sélectionné")
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("le fichier n'existe pas: %s", filePath)
	}

	return filePath, nil
}

// ValidateExcelFile valide qu'un fichier est bien un Excel valide
func (a *App) ValidateExcelFile(filePath string) error {
	// Vérifier l'extension
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext != ".xlsx" && ext != ".xls" {
		return fmt.Errorf("format de fichier invalide: %s (attendu .xlsx ou .xls)", ext)
	}

	// Vérifier que le fichier existe
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("le fichier n'existe pas: %s", filePath)
	}

	// Vérifier que c'est bien un fichier (pas un dossier)
	if info.IsDir() {
		return fmt.Errorf("le chemin est un dossier, pas un fichier: %s", filePath)
	}

	// Vérifier que le fichier n'est pas vide
	if info.Size() == 0 {
		return errors.New("le fichier est vide")
	}

	// Vérifier les permissions de lecture
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("impossible d'ouvrir le fichier: %w", err)
	}
	defer file.Close()

	return nil
}

// SaveExcelFile sauvegarde un fichier téléchargé dans le bon dossier data
func (a *App) SaveExcelFile(sourcePath string, fileType ExcelFileType) (string, error) {
	if err := a.ValidateExcelFile(sourcePath); err != nil {
		return "", err
	}

	// Déterminer le dossier de destination
	var subDir string
	switch fileType {
	case ExcelTypeMedecins:
		subDir = "medecins"
	case ExcelTypeIntervenants:
		subDir = "intervenants"
	case ExcelTypeClients:
		subDir = "clients"
	case ExcelTypeRPA:
		subDir = "rpa"
	case ExcelTypeCHSLD:
		subDir = "chsld"
	default:
		subDir = "imports"
	}

	destDir := filepath.Join("./data", subDir)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", fmt.Errorf("impossible de créer le dossier destination: %w", err)
	}

	// Nom du fichier avec timestamp pour éviter les conflits
	fileName := filepath.Base(sourcePath)
	destPath := filepath.Join(destDir, fileName)

	// Copier le fichier
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return "", fmt.Errorf("erreur ouverture source: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("erreur création destination: %w", err)
	}
	defer destFile.Close()

	if _, err := destFile.ReadFrom(sourceFile); err != nil {
		return "", fmt.Errorf("erreur copie fichier: %w", err)
	}

	return destPath, nil
}

// GetExcelFileInfo retourne des informations sur un fichier Excel
func (a *App) GetExcelFileInfo(filePath string) (map[string]interface{}, error) {
	if err := a.ValidateExcelFile(filePath); err != nil {
		return nil, err
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"path":     filePath,
		"name":     info.Name(),
		"size":     info.Size(),
		"modified": info.ModTime().Format("2006-01-02 15:04:05"),
		"ext":      filepath.Ext(filePath),
	}, nil
}

// ListExcelFiles liste tous les fichiers Excel dans un dossier data
func (a *App) ListExcelFiles(fileType ExcelFileType) ([]map[string]interface{}, error) {
	var subDir string
	switch fileType {
	case ExcelTypeMedecins:
		subDir = "medecins"
	case ExcelTypeIntervenants:
		subDir = "intervenants"
	case ExcelTypeClients:
		subDir = "clients"
	case ExcelTypeRPA:
		subDir = "rpa"
	case ExcelTypeCHSLD:
		subDir = "chsld"
	default:
		subDir = ""
	}

	targetDir := "./data"
	if subDir != "" {
		targetDir = filepath.Join(targetDir, subDir)
	}

	// Créer le dossier s'il n'existe pas
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			return nil, err
		}
		return []map[string]interface{}{}, nil
	}

	files, err := os.ReadDir(targetDir)
	if err != nil {
		return nil, err
	}

	var excelFiles []map[string]interface{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext != ".xlsx" && ext != ".xls" {
			continue
		}

		info, err := file.Info()
		if err != nil {
			continue
		}

		excelFiles = append(excelFiles, map[string]interface{}{
			"name":     file.Name(),
			"path":     filepath.Join(targetDir, file.Name()),
			"size":     info.Size(),
			"modified": info.ModTime().Format("2006-01-02 15:04:05"),
		})
	}

	return excelFiles, nil
}

// DeleteExcelFile supprime un fichier Excel
func (a *App) DeleteExcelFile(filePath string) error {
	if err := a.ValidateExcelFile(filePath); err != nil {
		return err
	}

	return os.Remove(filePath)
}

// ImportMedecins importe les médecins depuis un fichier Excel
func (a *App) ImportMedecins(filePath string) (string, error) {
	if a.currentUser == nil {
		return "", errors.New("non authentifié")
	}

	count, err := a.db.ImportMedecinsFromExcel(filePath, int(a.currentUser.ID))
	if err != nil {
		return fmt.Sprintf("Import partiel: %d médecins importés avec erreurs", count), err
	}

	return fmt.Sprintf("✅ Import réussi: %d médecins importés", count), nil
}

// ImportMedecinsWithUpdate importe avec mise à jour des existants
func (a *App) ImportMedecinsWithUpdate(filePath string) (*database.ImportStats, error) {
	if a.currentUser == nil {
		return nil, errors.New("non authentifié")
	}

	return a.db.ImportMedecinsFromExcelWithUpdate(filePath, int(a.currentUser.ID))
}
