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

// CreateClientFolderStructure crée la structure complète du dossier client
func (a *App) CreateClientFolderStructure(data map[string]interface{}) ClientFolderResult {
	leopardNumber, ok := data["leopardNumber"].(string)
	if !ok || leopardNumber == "" {
		return ClientFolderResult{Success: false, Error: "Numéro Leopard manquant"}
	}

	folderName, ok := data["folderName"].(string)
	if !ok || folderName == "" {
		folderName = leopardNumber
	}

	basePath := a.GetBasePath()
	clientPath := filepath.Join(basePath, "Clients", folderName)

	if _, err := os.Stat(clientPath); err == nil {
		return ClientFolderResult{Success: false, Error: "Le dossier existe déjà"}
	}

	if err := os.MkdirAll(clientPath, 0755); err != nil {
		return ClientFolderResult{Success: false, Error: fmt.Sprintf("Erreur création dossier: %v", err)}
	}

	subfolders := []string{"Evaluations", "Notes", "Rapports", "Correspondance", "Documents_medicaux", "Contrats"}

	for _, subfolder := range subfolders {
		subPath := filepath.Join(clientPath, subfolder)
		if err := os.MkdirAll(subPath, 0755); err != nil {
			fmt.Printf("⚠️ Erreur création sous-dossier %s: %v\n", subfolder, err)
		}
	}

	readmePath := filepath.Join(clientPath, "README.txt")
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

	fmt.Printf("✅ Dossier client créé: %s\n", clientPath)

	return ClientFolderResult{Success: true, Path: clientPath}
}

// OpenClientFolder ouvre le dossier d'un client dans l'explorateur
func (a *App) OpenClientFolder(folderName string) ClientFolderResult {
	basePath := a.GetBasePath()
	clientPath := filepath.Join(basePath, "Clients", folderName)

	if _, err := os.Stat(clientPath); os.IsNotExist(err) {
		return ClientFolderResult{Success: false, Error: "Le dossier n'existe pas"}
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", clientPath)
	case "darwin":
		cmd = exec.Command("open", clientPath)
	case "linux":
		cmd = exec.Command("xdg-open", clientPath)
	default:
		return ClientFolderResult{Success: false, Error: "OS non supporté"}
	}

	if err := cmd.Start(); err != nil {
		return ClientFolderResult{Success: false, Error: fmt.Sprintf("Erreur ouverture: %v", err)}
	}

	return ClientFolderResult{Success: true, Path: clientPath}
}

// ClientFolderExists vérifie si un dossier client existe
func (a *App) ClientFolderExists(leopardNumber string) bool {
	basePath := a.GetBasePath()
	clientsPath := filepath.Join(basePath, "Clients")

	entries, err := os.ReadDir(clientsPath)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= len(leopardNumber) && name[:len(leopardNumber)] == leopardNumber {
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
