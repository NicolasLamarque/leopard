// frontend/src/composables/useAppelReports.js
import { jsPDF } from "jspdf";
import autoTable from "jspdf-autotable"; // Importation nommée pour Vite

export function useAppelReports() {
  
  const generatePDF = (title, data) => {
    const doc = new jsPDF();
    
    // Header stylisé Rose (Ton interface)
    doc.setFillColor(236, 72, 153); // Rose 500
    doc.rect(0, 0, 210, 20, 'F');
    doc.setTextColor(255, 255, 255);
    doc.setFontSize(16);
    doc.text(title, 14, 13);
    
    // Date de génération en petit
    doc.setFontSize(9);
    doc.setTextColor(255, 255, 255);
    doc.text(`Généré le ${new Date().toLocaleDateString('fr-CA')}`, 160, 13);

    const tableRows = data.map(a => [
      a.date_appel ? new Date(a.date_appel).toLocaleDateString('fr-CA') : 'N/A',
      `${a.prospect_prenom || a.appelant_prenom || ''} ${a.prospect_nom || a.appelant_nom || ''}`,
      a.type_demande || 'Non spécifié',
      a.statut || 'Nouveau',
      a.priorite || 'Normale'
    ]);

    // Utilisation de la fonction autoTable directement (plus fiable)
    autoTable(doc, {
      startY: 25,
      head: [['Date', 'Nom / Prospect', 'Type de demande', 'Statut', 'Priorité']],
      body: tableRows,
      headStyles: { 
        fillColor: [120, 113, 108], // Stone 500 pour l'en-tête du tableau
        textColor: [255, 255, 255],
        fontSize: 10
      },
      alternateRowStyles: { 
        fillColor: [250, 250, 250] 
      },
      styles: {
        fontSize: 9,
        cellPadding: 3
      }
    });

    doc.save(`Rapport_${title.replace(/\s+/g, '_')}.pdf`);
  };

  const generateWeekly = (appels) => {
    const lastWeek = new Date();
    lastWeek.setDate(lastWeek.getDate() - 7);
    const data = appels.filter(a => new Date(a.date_appel) >= lastWeek);
    generatePDF("Rapport Hebdomadaire", data);
  };

  const generateMonthly = (appels) => {
    const lastMonth = new Date();
    lastMonth.setMonth(lastMonth.getMonth() - 1);
    const data = appels.filter(a => new Date(a.date_appel) >= lastMonth);
    generatePDF("Rapport Mensuel", data);
  };

  const generateByClient = (appels, clientId, clientName) => {
    // Filtrage strict par ID client
    const data = appels.filter(a => a.client_id === clientId);
    generatePDF(`Historique - ${clientName}`, data);
  };

  return { generateWeekly, generateMonthly, generateByClient };
}