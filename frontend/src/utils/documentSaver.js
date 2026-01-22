/ src/utils/documentSaver.js

/**
 * Utilitaire centralisé pour sauvegarder des documents dans les dossiers Leopard
 */

/**
 * Sauvegarde un PDF (déjà généré avec jsPDF) dans un dossier client
 * 
 * @param {jsPDF} doc - Instance jsPDF
 * @param {string} leopardNumber - Numéro Leopard du client (ex: "DUPA_20231215_001")
 * @param {string} subfolder - Sous-dossier ("Notes", "Rapports", "Documents/Identification", etc.)
 * @param {string} filename - Nom du fichier (ex: "Note_2024-01-17.pdf")
 * @returns {Promise<{success: boolean, path?: string, error?: string}>}
 */
export async function savePDFToClientFolder(doc, leopardNumber, subfolder, filename) {
  try {
    // 1. Récupérer le chemin du dossier client
    const folderInfo = await window.go.main.App.GetClientFolderInfo(leopardNumber);
    
    if (!folderInfo || !folderInfo.path) {
      throw new Error(`Dossier client introuvable pour ${leopardNumber}`);
    }

    // 2. Construire le chemin complet du sous-dossier
    const targetFolder = `${folderInfo.path}/${subfolder}`;

    // 3. Convertir le PDF en base64
    const pdfBase64 = doc.output('datauristring').split(',')[1];

    // 4. Sauvegarder via le backend
    const result = await window.go.main.App.SavePDFToFolder(
      targetFolder,
      filename,
      pdfBase64
    );

    if (!result.success) {
      throw new Error(result.error || 'Erreur inconnue lors de la sauvegarde');
    }

    console.log(`✅ PDF sauvegardé: ${result.path}`);
    return result;

  } catch (error) {
    console.error('❌ Erreur sauvegarde PDF:', error);
    return {
      success: false,
      error: error.message
    };
  }
}

/**
 * Génère un nom de fichier standardisé avec horodatage
 * 
 * @param {string} prefix - Préfixe (ex: "Note", "Rapport", "Fiche")
 * @param {string} clientNom - Nom du client
 * @param {string} extension - Extension (défaut: "pdf")
 * @returns {string} Ex: "Note_DUPONT_2024-01-17_143052.pdf"
 */
export function generateFileName(prefix, clientNom, extension = 'pdf') {
  const now = new Date();
  const dateStr = now.toISOString().split('T')[0]; // 2024-01-17
  const timeStr = now.toTimeString().split(' ')[0].replace(/:/g, ''); // 143052
  
  const safeName = clientNom
    .toUpperCase()
    .normalize('NFD')
    .replace(/[\u0300-\u036f]/g, '') // Enlever accents
    .replace(/[^A-Z]/g, ''); // Garder que les lettres
  
  return `${prefix}_${safeName}_${dateStr}_${timeStr}.${extension}`;
}

/**
 * Mapping des types de documents vers leurs sous-dossiers
 */
export const DOCUMENT_FOLDERS = {
  NOTE_CLINIQUE: 'Notes',
  RAPPORT: 'Rapports',
  EVALUATION: 'Evaluations',
  CORRESPONDANCE: 'Correspondance',
  FICHE_IDENTIFICATION: 'Documents/Identification',
  CONSENTEMENT: 'Documents/Consentements',
  DOCUMENT_MEDICAL: 'Documents/Medicaux',
  DOCUMENT_LEGAL: 'Documents/Legaux',
  CONTRAT: 'Contrats',
  PLAN_INTERVENTION: 'Rapports' // ou créer un dossier "Plans" si tu préfères
};