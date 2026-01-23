// folders.go
// Fichier optionnel pour séparer la logique des dossiers
// À placer dans le même package que app.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"time"
)

// ClientFolderResult représente le résultat d'une opération sur dossier
type ClientFolderResult struct {
	Success bool   `json:"success"`
	Path    string `json:"path,omitempty"`
	Error   string `json:"error,omitempty"`
}

// ClientFolderInfo contient les informations d'un dossier client
type ClientFolderInfo struct {
	Path            string `json:"path"`
	SubfoldersCount int    `json:"subfoldersCount"`
	FilesCount      int    `json:"filesCount"`
}

// GetBasePath retourne le chemin de base pour les dossiers Leopard (Version Portative)
func (a *App) GetBasePath() string {
	// On récupère le chemin de l'exécutable
	exePath, err := os.Executable()
	if err != nil {
		// En cas d'erreur, on utilise le dossier courant "."
		return filepath.Join(".", "Leopard_Dossiers")
	}

	// On prend le dossier parent de l'exécutable
	exeDir := filepath.Dir(exePath)

	// On crée le chemin spécifique
	return filepath.Join(exeDir, "Leopard_Dossiers")
}

// initializeLeopardFolders crée la structure de base des dossiers Leopard
func (a *App) initializeLeopardFolders() error {
	basePath := a.GetBasePath()

	folders := []string{
		filepath.Join(basePath, "Clients"),
		filepath.Join(basePath, "Exports"),
		filepath.Join(basePath, "Sauvegardes"),
		filepath.Join(basePath, "Templates"),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, 0755); err != nil {
			return fmt.Errorf("erreur création dossier %s: %v", folder, err)
		}
	}

	fmt.Printf("✅ Structure Leopard initialisée: %s\n", basePath)
	return nil
}

func (a *App) CreateClientFolderStructure(data map[string]interface{}) ClientFolderResult {
	// 1. Récupérer le numéro Leopard
	leopardNumber, ok := data["leopardNumber"].(string)
	if !ok || leopardNumber == "" {
		return ClientFolderResult{
			Success: false,
			Error:   "Numéro Leopard manquant",
		}
	}

	// 2. Récupérer le nom complet du dossier (ou utiliser juste le numéro)
	folderName, ok := data["folderName"].(string)
	if !ok || folderName == "" {
		folderName = leopardNumber // Si pas de nom, on utilise juste le numéro
	}

	// 3. Construire le chemin complet
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")
	fullClientPath := filepath.Join(clientsPath, folderName)

	// 4. DÉBOGAGE : afficher ce qu'on essaie de créer
	fmt.Printf("🔍 Tentative de création:\n")
	fmt.Printf("   Base: %s\n", basePath)
	fmt.Printf("   Clients: %s\n", clientsPath)
	fmt.Printf("   Dossier final: %s\n", fullClientPath)

	// 5. Vérifier si le dossier existe déjà
	if _, err := os.Stat(fullClientPath); err == nil {
		return ClientFolderResult{
			Success: false,
			Error:   "Le dossier existe déjà",
			Path:    fullClientPath,
		}
	}

	// 6. Créer le dossier principal
	if err := os.MkdirAll(fullClientPath, 0755); err != nil {
		return ClientFolderResult{
			Success: false,
			Error:   fmt.Sprintf("Impossible de créer le dossier: %v", err),
		}
	}

	fmt.Printf("✅ Dossier principal créé: %s\n", fullClientPath)

	// 7. Créer les sous-dossiers
	subfolders := []string{
		"Evaluations",
		"Notes",
		"Rapports",
		"Correspondance",
		"Documents",
		"Contrats",
	}

	for _, subfolder := range subfolders {
		subPath := filepath.Join(fullClientPath, subfolder)
		if err := os.MkdirAll(subPath, 0755); err != nil {
			fmt.Printf("⚠️ Erreur création sous-dossier %s: %v\n", subfolder, err)
		} else {
			fmt.Printf("   ✅ Sous-dossier créé: %s\n", subfolder)

			// Si c'est Documents, créer la sous-structure
			if subfolder == "Documents" {
				documentsSubfolders := []string{"Medicaux", "Legaux", "Identification", "Consentements"}
				for _, docSub := range documentsSubfolders {
					docSubPath := filepath.Join(subPath, docSub)
					if err := os.MkdirAll(docSubPath, 0755); err != nil {
						fmt.Printf("      ⚠️ Erreur Documents/%s: %v\n", docSub, err)
					} else {
						fmt.Printf("      📄 Créé: Documents/%s\n", docSub)
					}
				}
			}
		}
	}
	// 8. Créer le fichier README
	readmePath := filepath.Join(fullClientPath, "README.txt")
	readmeContent := fmt.Sprintf(
		"Dossier Client Leopard\n"+
			"=====================\n\n"+
			"Numéro: %s\n"+
			"Créé le: %s\n\n"+
			"Ce dossier contient tous les documents relatifs à ce client.\n",
		leopardNumber,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		fmt.Printf("⚠️ Erreur création README: %v\n", err)
	}

	return ClientFolderResult{
		Success: true,
		Path:    fullClientPath,
	}
}

// OpenClientFolder ouvre le dossier d'un client dans l'explorateur
func (a *App) OpenClientFolder(leopardNumber string) ClientFolderResult {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// On cherche le dossier qui commence par le numéro Leopard
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return ClientFolderResult{Success: false, Error: "Impossible de lire le dossier Clients"}
	}

	var fullFolderPath string
	targetLength := len(leopardNumber)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				fullFolderPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if fullFolderPath == "" {
		return ClientFolderResult{Success: false, Error: "Dossier introuvable"}
	}

	// Ouvrir le dossier dans l'explorateur
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", fullFolderPath)
	case "darwin":
		cmd = exec.Command("open", fullFolderPath)
	case "linux":
		cmd = exec.Command("xdg-open", fullFolderPath)
	default:
		return ClientFolderResult{Success: false, Error: "OS non supporté"}
	}

	if err := cmd.Start(); err != nil {
		return ClientFolderResult{Success: false, Error: fmt.Sprintf("Erreur ouverture: %v", err)}
	}

	return ClientFolderResult{Success: true, Path: fullFolderPath}
}

// ClientFolderExists vérifie si un dossier client existe
func (a *App) ClientFolderExists(leopardNumber string) bool {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return false
	}

	targetLength := len(leopardNumber) // Toujours la même longueur (17)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			// On s'en crisse du reste : on compare juste le début sur la longueur exacte
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				return true
			}
		}
	}
	return false
}

// GetClientFolderInfo retourne les infos d'un dossier client
func (a *App) GetClientFolderInfo(leopardNumber string) *ClientFolderInfo {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return nil
	}

	var clientPath string
	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= len(leopardNumber) && name[:len(leopardNumber)] == leopardNumber {
				clientPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if clientPath == "" {
		return nil
	}

	subfoldersCount := 0
	filesCount := 0

	filepath.Walk(clientPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || path == clientPath {
			return nil
		}
		if info.IsDir() {
			subfoldersCount++
		} else {
			filesCount++
		}
		return nil
	})

	return &ClientFolderInfo{
		Path:            clientPath,
		SubfoldersCount: subfoldersCount,
		FilesCount:      filesCount,
	}
}

// ListClientFolders retourne tous les dossiers clients
func (a *App) ListClientFolders() []string {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return []string{}
	}

	var folders []string
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	return folders
}

// RenameClientFolder renomme un dossier client
func (a *App) RenameClientFolder(oldName, newName string) ClientFolderResult {
	basePath := a.GetBasePath()
	oldPath := filepath.Join(basePath, "Clients", oldName)
	newPath := filepath.Join(basePath, "Clients", newName)

	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		return ClientFolderResult{Success: false, Error: "Dossier source inexistant"}
	}

	if _, err := os.Stat(newPath); err == nil {
		return ClientFolderResult{Success: false, Error: "Le nouveau nom existe déjà"}
	}

	if err := os.Rename(oldPath, newPath); err != nil {
		return ClientFolderResult{Success: false, Error: fmt.Sprintf("Erreur renommage: %v", err)}
	}

	return ClientFolderResult{Success: true, Path: newPath}
}

// ExportToClientFolder exporte un fichier dans un dossier client
func (a *App) ExportToClientFolder(data map[string]interface{}) ClientFolderResult {
	leopardNumber, _ := data["leopardNumber"].(string)
	subfolder, _ := data["subfolder"].(string)
	filename, _ := data["filename"].(string)
	fileData, _ := data["data"].([]byte)

	if leopardNumber == "" || filename == "" {
		return ClientFolderResult{Success: false, Error: "Données manquantes"}
	}

	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")
	entries, _ := os.ReadDir(clientsPath)

	var clientPath string
	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= len(leopardNumber) && name[:len(leopardNumber)] == leopardNumber {
				if subfolder != "" {
					clientPath = filepath.Join(clientsPath, name, subfolder)
				} else {
					clientPath = filepath.Join(clientsPath, name)
				}
				break
			}
		}
	}

	if clientPath == "" {
		return ClientFolderResult{Success: false, Error: "Dossier introuvable"}
	}

	os.MkdirAll(clientPath, 0755)

	filePath := filepath.Join(clientPath, filename)
	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return ClientFolderResult{Success: false, Error: fmt.Sprintf("Erreur écriture: %v", err)}
	}

	return ClientFolderResult{Success: true, Path: filePath}
}

// OpenMainClientsFolder ouvre le dossier racine des clients
func (a *App) OpenMainClientsFolder() ClientFolderResult {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// On s'assure que le dossier existe avant d'essayer de l'ouvrir
	if err := os.MkdirAll(clientsPath, 0755); err != nil {
		return ClientFolderResult{Success: false, Error: "Impossible de créer/accéder au dossier"}
	}

	return a.OpenClientFolder("") // On appelle ton OpenClientFolder existant
}

// GetClientFolderStructure retourne la structure détaillée avec statut de chaque sous-dossier
func (a *App) GetClientFolderStructure(leopardNumber string) map[string]interface{} {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// Trouver le dossier du client
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   "Impossible de lire le dossier Clients",
		}
	}

	var clientPath string
	targetLength := len(leopardNumber)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				clientPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if clientPath == "" {
		return map[string]interface{}{
			"success": false,
			"error":   "Dossier client introuvable",
		}
	}

	// Vérifier chaque sous-dossier attendu
	expectedSubfolders := []string{
		"Evaluations",
		"Notes",
		"Rapports",
		"Correspondance",
		"Documents",
		"Contrats",
	}

	subfolderStatus := make(map[string]bool)

	for _, subfolder := range expectedSubfolders {
		subPath := filepath.Join(clientPath, subfolder)
		_, err := os.Stat(subPath)
		subfolderStatus[subfolder] = (err == nil)
	}

	return map[string]interface{}{
		"success":    true,
		"path":       clientPath,
		"subfolders": subfolderStatus,
	}
}

// RepairClientFolderStructure répare UNIQUEMENT ce qui manque (SAFE)
func (a *App) RepairClientFolderStructure(leopardNumber string) map[string]interface{} {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// 1. Trouver le dossier du client
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   "Impossible de lire le dossier Clients",
		}
	}

	var clientPath string
	targetLength := len(leopardNumber)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				clientPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if clientPath == "" {
		return map[string]interface{}{
			"success": false,
			"error":   "Dossier client introuvable. Créez d'abord le dossier principal.",
		}
	}

	fmt.Printf("🔧 Réparation sécurisée: %s\n", clientPath)

	// 2. Liste des sous-dossiers attendus
	expectedSubfolders := []string{
		"Evaluations",
		"Notes",
		"Rapports",
		"Correspondance",
		"Documents",
		"Contrats",
	}

	// 3. Vérifier et créer SEULEMENT ce qui manque
	created := []string{}
	alreadyExists := []string{}
	errors := []string{}

	for _, subfolder := range expectedSubfolders {
		subPath := filepath.Join(clientPath, subfolder)

		// Vérifier si existe
		if _, err := os.Stat(subPath); os.IsNotExist(err) {
			// N'existe PAS → Créer
			if err := os.MkdirAll(subPath, 0755); err != nil {
				errorMsg := fmt.Sprintf("%s: %v", subfolder, err)
				errors = append(errors, errorMsg)
				fmt.Printf("   ❌ Erreur: %s\n", errorMsg)
			} else {
				created = append(created, subfolder)
				fmt.Printf("   ✅ Créé: %s\n", subfolder)

				// Si c'est Documents, créer la sous-structure
				if subfolder == "Documents" {
					documentsSubfolders := []string{"Medicaux", "Legaux", "Identification", "Consentements"}
					for _, docSub := range documentsSubfolders {
						docSubPath := filepath.Join(subPath, docSub)
						if err := os.MkdirAll(docSubPath, 0755); err != nil {
							fmt.Printf("      ⚠️ Erreur Documents/%s: %v\n", docSub, err)
						} else {
							fmt.Printf("      📄 Créé: Documents/%s\n", docSub)
						}
					}
				}
			}
		} else {
			// Existe déjà → SKIP (sécurité)
			alreadyExists = append(alreadyExists, subfolder)
			fmt.Printf("   ℹ️ Ignoré (existe déjà): %s\n", subfolder)
		}
	}

	// 4. Créer le README s'il n'existe pas (SAFE)
	readmePath := filepath.Join(clientPath, "README.txt")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		readmeContent := fmt.Sprintf(
			"Dossier Client Leopard\n"+
				"=====================\n\n"+
				"Numéro: %s\n"+
				"Créé le: %s\n\n"+
				"Ce dossier contient tous les documents relatifs à ce client.\n",
			leopardNumber,
			time.Now().Format("2006-01-02 15:04:05"),
		)
		os.WriteFile(readmePath, []byte(readmeContent), 0644)
	}

	// 5. Rapport détaillé
	summary := fmt.Sprintf(
		"Réparation terminée:\n"+
			"- %d sous-dossiers créés\n"+
			"- %d déjà existants (ignorés)\n"+
			"- %d erreurs",
		len(created),
		len(alreadyExists),
		len(errors),
	)

	fmt.Printf("📊 %s\n", summary)

	return map[string]interface{}{
		"success":       len(errors) == 0,
		"created":       created,
		"alreadyExists": alreadyExists,
		"errors":        errors,
		"summary":       summary,
		"path":          clientPath,
	}
}

// CreateSubfolder crée un sous-dossier spécifique (SAFE - ne touche pas à l'existant)
func (a *App) CreateSubfolder(leopardNumber, subfolderName string) ClientFolderResult {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// Trouver le dossier du client
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return ClientFolderResult{Success: false, Error: "Impossible de lire le dossier Clients"}
	}

	var clientPath string
	targetLength := len(leopardNumber)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				clientPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if clientPath == "" {
		return ClientFolderResult{Success: false, Error: "Dossier client introuvable"}
	}

	// Vérifier le chemin du sous-dossier
	subPath := filepath.Join(clientPath, subfolderName)

	// SÉCURITÉ: Vérifier si existe déjà
	if info, err := os.Stat(subPath); err == nil {
		if info.IsDir() {
			// Compter les fichiers dedans
			items, _ := os.ReadDir(subPath)
			fileCount := 0
			for _, item := range items {
				if !item.IsDir() {
					fileCount++
				}
			}

			if fileCount > 0 {
				return ClientFolderResult{
					Success: false,
					Error:   fmt.Sprintf("⚠️ Le sous-dossier existe déjà avec %d fichier(s). Opération annulée pour éviter toute perte.", fileCount),
				}
			} else {
				return ClientFolderResult{
					Success: true,
					Path:    subPath,
					Error:   "ℹ️ Le sous-dossier existe déjà (vide).",
				}
			}
		}
	}

	// Créer le sous-dossier (il n'existe pas)
	if err := os.MkdirAll(subPath, 0755); err != nil {
		return ClientFolderResult{
			Success: false,
			Error:   fmt.Sprintf("Erreur création: %v", err),
		}
	}

	fmt.Printf("✅ Sous-dossier créé: %s/%s\n", clientPath, subfolderName)

	return ClientFolderResult{Success: true, Path: subPath}
}

// SubfolderDetail contient les détails enrichis d'un sous-dossier
type SubfolderDetail struct {
	Name           string            `json:"name"`
	Exists         bool              `json:"exists"`
	Path           string            `json:"path"`
	FileCount      int               `json:"fileCount"`
	SubfolderCount int               `json:"subfolderCount"`
	Files          []FileInfo        `json:"files,omitempty"`
	Children       []SubfolderDetail `json:"children,omitempty"`
}

// FileInfo contient les métadonnées d'un fichier
type FileInfo struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	ModifiedTime time.Time `json:"modifiedTime"`
	Extension    string    `json:"extension"`
}

// GetDetailedFolderStructure retourne la structure complète avec métadonnées
// ⚠️ NOUVELLE FONCTION - Ne remplace PAS GetClientFolderStructure
func (a *App) GetDetailedFolderStructure(leopardNumber string) map[string]interface{} {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	// Trouver le dossier du client
	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   "Impossible de lire le dossier Clients",
		}
	}

	var clientPath string
	targetLength := len(leopardNumber)

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= targetLength && name[:targetLength] == leopardNumber {
				clientPath = filepath.Join(clientsPath, name)
				break
			}
		}
	}

	if clientPath == "" {
		return map[string]interface{}{
			"success": false,
			"error":   "Dossier client introuvable",
		}
	}

	// Liste des sous-dossiers attendus (identique à votre liste existante)
	expectedSubfolders := []string{
		"Evaluations",
		"Notes",
		"Rapports",
		"Correspondance",
		"Documents",
		"Contrats",
	}

	var folders []SubfolderDetail

	for _, subfolderName := range expectedSubfolders {
		subPath := filepath.Join(clientPath, subfolderName)

		detail := SubfolderDetail{
			Name: subfolderName,
			Path: subPath,
		}

		// Vérifier si le dossier existe
		if info, err := os.Stat(subPath); err == nil && info.IsDir() {
			detail.Exists = true

			// Analyser le contenu (appel fonction ci-dessous)
			analyzeFolder(&detail, subPath)
		} else {
			detail.Exists = false
		}

		folders = append(folders, detail)
	}

	return map[string]interface{}{
		"success": true,
		"path":    clientPath,
		"folders": folders,
	}
}

// analyzeFolder analyse récursivement le contenu d'un dossier
// 🆕 FONCTION UTILITAIRE - Ne remplace rien
func analyzeFolder(detail *SubfolderDetail, folderPath string) {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return
	}

	var files []FileInfo
	var children []SubfolderDetail
	fileCount := 0
	subfolderCount := 0

	for _, entry := range entries {
		entryPath := filepath.Join(folderPath, entry.Name())

		if entry.IsDir() {
			// C'est un sous-dossier
			subfolderCount++

			childDetail := SubfolderDetail{
				Name:   entry.Name(),
				Exists: true,
				Path:   entryPath,
			}

			// Analyser ce sous-dossier (récursif, max 1 niveau)
			analyzeFolder(&childDetail, entryPath)
			children = append(children, childDetail)

		} else {
			// C'est un fichier
			fileCount++

			info, err := entry.Info()
			if err == nil {
				fileInfo := FileInfo{
					Name:         entry.Name(),
					Size:         info.Size(),
					ModifiedTime: info.ModTime(),
					Extension:    filepath.Ext(entry.Name()),
				}
				files = append(files, fileInfo)
			}
		}
	}

	// Trier les fichiers par date (plus récents en premier)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModifiedTime.After(files[j].ModifiedTime)
	})

	detail.FileCount = fileCount
	detail.SubfolderCount = subfolderCount
	detail.Files = files
	detail.Children = children
}

// OpenFolder ouvre un dossier spécifique dans l'explorateur
// 🆕 NOUVELLE FONCTION - Différente de OpenClientFolder (qui cherche par numéro)
func (a *App) OpenFolder(folderPath string) ClientFolderResult {
	// Vérifier que le dossier existe
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return ClientFolderResult{
			Success: false,
			Error:   "Le dossier n'existe pas",
		}
	}

	// Ouvrir dans l'explorateur selon l'OS
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", folderPath)
	case "darwin":
		cmd = exec.Command("open", folderPath)
	case "linux":
		cmd = exec.Command("xdg-open", folderPath)
	default:
		return ClientFolderResult{
			Success: false,
			Error:   "OS non supporté",
		}
	}

	if err := cmd.Start(); err != nil {
		return ClientFolderResult{
			Success: false,
			Error:   fmt.Sprintf("Erreur ouverture: %v", err),
		}
	}

	return ClientFolderResult{
		Success: true,
		Path:    folderPath,
	}
}
