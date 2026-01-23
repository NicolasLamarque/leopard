/**
 * clientFolderManager.js
 * Système de gestion des dossiers clients Leopard
 */

/**
 * Génère le numéro de dossier Leopard
 * Format: NOM(3 lettres) + PRENOM(1 lettre) + _ + DATE(YYYYMMDD) + _XXX
 * Exemple: DUPA_20231215_001 pour Dupont Alice créé le 15/12/2023
 */
export function generateLeopardNumber(nom, prenom, createdAt = null) {
  try {
    // Nettoyer et normaliser les noms
    const cleanNom = normalizeString(nom).toUpperCase();
    const cleanPrenom = normalizeString(prenom).toUpperCase();
    
    // Prendre les 3 premières lettres du nom
    const nomPart = cleanNom.substring(0, 3).padEnd(3, 'X');
    
    // Prendre la première lettre du prénom
    const prenomPart = cleanPrenom.substring(0, 1) || 'X';
    
    // Formater la date (YYYYMMDD)
    const date = createdAt ? new Date(createdAt) : new Date();
    const datePart = formatDateForFolder(date);
    
    // Format final: NNNP_YYYYMMDD_XXX (le compteur sera ajouté par le backend)
    return `${nomPart}${prenomPart}_${datePart}`;
  } catch (error) {
    console.error('❌ Erreur génération numéro Leopard:', error);
    return null;
  }
}

/**
 * Normalise une chaîne (enlève accents, caractères spéciaux)
 */
function normalizeString(str) {
  if (!str) return '';
  
  return str
    .normalize('NFD')
    .replace(/[\u0300-\u036f]/g, '') // Enlever les accents
    .replace(/[^a-zA-Z]/g, '') // Garder SEULEMENT les lettres (pas de chiffres)
    .trim();
}

/**
 * Formate une date en YYYYMMDD
 */
function formatDateForFolder(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}${month}${day}`;
}

/**
 * Structure du dossier client
 */
const FOLDER_STRUCTURE = {
  root: 'Documents', // Documents/Leopard/Clients/
  subfolders: [
    'Evaluations',
    'Notes',
    'Rapports',
    'Correspondance',
    'Documents_medicaux',
    'Contrats'
  ]
};

/**
 * Crée le dossier complet du client sur le disque
 */
/**
 * Crée le dossier complet du client sur le disque
 */
export async function createClientFolder(clientData) {
  try {
    console.log('📨 Données envoyées au backend:', clientData);
    
    // Vérifier que l'API Wails est disponible
    if (!window.go || !window.go.main || !window.go.main.App) {
      throw new Error('API Wails non disponible');
    }
    
    // Préparer les données pour Go
    const dataForGo = {
      leopardNumber: clientData.no_dossier_leopard || clientData.leopardNumber,
      folderName: clientData.folderName || clientData.no_dossier_leopard
    };
    
    console.log('🔧 Données formatées pour Go:', dataForGo);
    
    // Appel au backend
    const result = await window.go.main.App.CreateClientFolderStructure(dataForGo);
    
    console.log('📥 Résultat du backend:', result);
    
    if (result.success) {
      return {
        success: true,
        leopardNumber: dataForGo.leopardNumber,
        path: result.path
      };
    }
    
    // Si le dossier existe déjà, on considère ça comme un succès
    if (result.error && result.error.includes('existe déjà')) {
      return {
        success: true,
        leopardNumber: dataForGo.leopardNumber,
        path: result.path,
        alreadyExists: true
      };
    }
    
    throw new Error(result.error || 'Erreur inconnue');
    
  } catch (error) {
    console.error('❌ Erreur création dossier client:', error);
    return {
      success: false,
      error: error.message
    };
  }
}

/**
 * Ouvre le dossier d'un client spécifique
 */
export async function openClientFolder(leopardNumber) {
  try {
    console.log('📂 Ouverture du dossier:', leopardNumber);
    
    // Vérifier que l'API Wails est disponible
    if (!window.go || !window.go.main || !window.go.main.App) {
      throw new Error('API Wails non disponible');
    }
    
    // Appel Wails - le backend cherchera le dossier automatiquement
    const result = await window.go.main.App.OpenClientFolder(leopardNumber);
    
    console.log('📥 Résultat ouverture:', result);
    
    if (result.success) {
      return {
        success: true,
        path: result.path
      };
    }
    
    throw new Error(result.error || 'Erreur inconnue');
    
  } catch (error) {
    console.error('❌ Erreur ouverture dossier:', error);
    throw error;
  }
}

/**
 * Vérifie si le dossier client existe
 */
export async function clientFolderExists(leopardNumber) {
  try {
    // Vérifier que l'API Wails est disponible
    if (!window.go || !window.go.main || !window.go.main.App) {
      console.warn('⚠️ API Wails non disponible');
      return false;
    }
    
    const result = await window.go.main.App.ClientFolderExists(leopardNumber);
    console.log(`🔍 Vérification dossier ${leopardNumber}:`, result);
    return result;
  } catch (error) {
    console.error('❌ Erreur vérification dossier:', error);
    return false;
  }
}

/**
 * Renomme un dossier client (si nom changé)
 */
export async function renameClientFolder(oldLeopardNumber, newClientData) {
  try {
    const newLeopardNumber = generateLeopardNumber(
      newClientData.nom, 
      newClientData.prenom, 
      newClientData.created_at
    );
    
    const oldFolderName = oldLeopardNumber;
    const newFolderName = `${newLeopardNumber} - ${newClientData.nom} ${newClientData.prenom}`;
    
    const result = await window.go.main.App.RenameClientFolder(oldFolderName, newFolderName);
    
    if (result.success) {
      console.log(`✅ Dossier renommé: ${result.newPath}`);
      return {
        success: true,
        newLeopardNumber,
        newPath: result.newPath
      };
    }
    
    throw new Error(result.error);
  } catch (error) {
    console.error('❌ Erreur renommage dossier:', error);
    return {
      success: false,
      error: error.message
    };
  }
}

/**
 * Liste tous les dossiers clients
 */
export async function listClientFolders() {
  try {
    const folders = await window.go.main.App.ListClientFolders();
    return folders || [];
  } catch (error) {
    console.error('❌ Erreur listage dossiers:', error);
    return [];
  }
}

/**
 * Exporte un document dans le dossier client
 */
export async function exportToClientFolder(leopardNumber, subfolder, filename, data) {
  try {
    const result = await window.go.main.App.ExportToClientFolder({
      leopardNumber,
      subfolder,
      filename,
      data
    });
    
    if (result.success) {
      console.log(`✅ Document exporté: ${result.path}`);
      return result.path;
    }
    
    throw new Error(result.error);
  } catch (error) {
    console.error('❌ Erreur export document:', error);
    return null;
  }
}

// Tests unitaires pour la génération du numéro
export function testLeopardNumberGeneration() {
  console.log('🧪 Tests de génération du numéro Leopard:');
  
  const tests = [
    { nom: 'Dupont', prenom: 'Alice', date: new Date('2023-12-15'), expected: 'DUPA_20231215' },
    { nom: 'Martin', prenom: 'Bob', date: new Date('2024-01-01'), expected: 'MARB_20240101' },
    { nom: 'Lefebvre', prenom: 'Catherine', date: new Date('2024-06-30'), expected: 'LEFC_20240630' },
    { nom: 'Li', prenom: 'Jean', date: new Date('2024-12-21'), expected: 'LIXXJ_20241221' },
    
    { nom: 'José-María', prenom: 'François', date: new Date('2024-03-15'), expected: 'JOSF_20240315' },
  ];
  
  tests.forEach(test => {
    const result = generateLeopardNumber(test.nom, test.prenom, test.date);
    const status = result === test.expected ? '✅' : '❌';
    console.log(`${status} ${test.nom} ${test.prenom} → ${result} (attendu: ${test.expected})`);
  });
}