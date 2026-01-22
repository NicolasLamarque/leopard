/**
 * WailsFolder.js
 * WailsIntegration.js
 * Fonctions pour interagir avec le système de fichiers via Wails
 */

import { OpenDirectoryDialog, BrowserOpenURL } from '@wailsapp/runtime';

/**
 * Configuration des chemins de dossiers
 */
const FOLDER_PATHS = {
  clients: 'Leopard/Clients',
  dossiers: 'Leopard/Dossiers',
  notes: 'Leopard/Notes',
  exports: 'Leopard/Exports'
};

/**
 * Ouvre un dossier système spécifique
 * @param {string} folderType - Type de dossier (clients, dossiers, notes)
 * @param {string} clientId - ID du client (optionnel)
 */
export async function openSystemFolder(folderType, clientId = null) {
  try {
    let folderPath = FOLDER_PATHS[folderType];
    
    if (clientId) {
      folderPath = `${folderPath}/${clientId}`;
    }
    
    // Ouvrir le dossier dans l'explorateur de fichiers
    // Note: Utilise file:// protocole pour Windows, macOS, Linux
    const platform = await getPlatform();
    let url;
    
    if (platform === 'windows') {
      url = `file:///${folderPath.replace(/\//g, '\\')}`;
    } else {
      url = `file://${folderPath}`;
    }
    
    await BrowserOpenURL(url);
    console.log(`✅ Dossier ouvert: ${folderPath}`);
    return true;
  } catch (error) {
    console.error(`❌ Erreur ouverture dossier ${folderType}:`, error);
    return false;
  }
}

/**
 * Crée un nouveau dossier pour un client
 * @param {string} clientId - ID du client
 * @param {string} clientName - Nom du client
 */
export async function createClientFolder(clientId, clientName) {
  try {
    // Appel vers le backend Wails Go
    const result = await window.go.main.App.CreateClientFolder(clientId, clientName);
    
    if (result.success) {
      console.log(`✅ Dossier client créé: ${result.path}`);
      return result.path;
    }
    
    throw new Error(result.error);
  } catch (error) {
    console.error('❌ Erreur création dossier client:', error);
    return null;
  }
}

/**
 * Sélectionne un dossier via dialogue système
 */
export async function selectFolder() {
  try {
    const result = await OpenDirectoryDialog({
      title: 'Sélectionner un dossier',
      defaultDirectory: FOLDER_PATHS.clients
    });
    
    if (result) {
      console.log(`✅ Dossier sélectionné: ${result}`);
      return result;
    }
    
    return null;
  } catch (error) {
    console.error('❌ Erreur sélection dossier:', error);
    return null;
  }
}

/**
 * Obtient la plateforme du système
 */
async function getPlatform() {
  try {
    const { Environment } = await import('@wailsapp/runtime');
    const info = await Environment();
    return info.platform;
  } catch (error) {
    console.error('❌ Erreur détection plateforme:', error);
    return 'unknown';
  }
}

/**
 * Exporte un fichier vers le dossier d'exports
 * @param {string} filename - Nom du fichier
 * @param {Blob} data - Données du fichier
 */
export async function exportFile(filename, data) {
  try {
    // Appel vers le backend Wails pour écrire le fichier
    const result = await window.go.main.App.ExportFile(filename, data);
    
    if (result.success) {
      console.log(`✅ Fichier exporté: ${result.path}`);
      
      // Ouvrir le dossier d'exports
      await openSystemFolder('exports');
      
      return result.path;
    }
    
    throw new Error(result.error);
  } catch (error) {
    console.error('❌ Erreur export fichier:', error);
    return null;
  }
}

/**
 * Vérifie si un dossier existe
 * @param {string} folderPath - Chemin du dossier
 */
export async function folderExists(folderPath) {
  try {
    const result = await window.go.main.App.FolderExists(folderPath);
    return result;
  } catch (error) {
    console.error('❌ Erreur vérification dossier:', error);
    return false;
  }
}

/**
 * Obtient la liste des fichiers dans un dossier
 * @param {string} folderPath - Chemin du dossier
 */
export async function listFiles(folderPath) {
  try {
    const files = await window.go.main.App.ListFiles(folderPath);
    return files || [];
  } catch (error) {
    console.error('❌ Erreur listage fichiers:', error);
    return [];
  }
}