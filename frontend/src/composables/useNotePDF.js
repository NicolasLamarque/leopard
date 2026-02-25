// composables/useNotesPDF.js
import { jsPDF } from 'jspdf'
import logoBlanc from '../assets/images/logoBlanc.png'
export function useNotesPDF() {
  // Couleurs corporatives
  const colors = {
    teal: [20, 184, 166],
    tealDark: [13, 148, 136],
    slate: [51, 65, 85],
    slateDark: [30, 41, 59],
    gray: [107, 114, 128],
    grayLight: [229, 231, 235]
  }

  // Fonction pour ajouter l'en-tête avec logo
  const addHeader = (doc, titre, logoBase64) => {
    const pageWidth = doc.internal.pageSize.getWidth()
    
    // Logo à gauche (si disponible)
    if (logoBase64) {
      try {
        doc.addImage(logoBase64, 'PNG', 15, 15, 30, 15)
      } catch (e) {
        console.warn('Impossible de charger le logo:', e)
      }
    }
    
    // Titre à droite
    doc.setFontSize(11)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.slateDark)
    doc.text(titre, pageWidth - 15, 22, { align: 'right' })
    
    // Ligne de séparation
    doc.setDrawColor(...colors.teal)
    doc.setLineWidth(1)
    doc.line(15, 35, pageWidth - 15, 35)
    
    return 40 // Position Y après l'en-tête
  }

  // Fonction pour ajouter le pied de page confidentiel
  const addFooter = (doc, pageNumber, totalPages) => {
    const pageWidth = doc.internal.pageSize.getWidth()
    const pageHeight = doc.internal.pageSize.getHeight()
    
    // Ligne de séparation
    doc.setDrawColor(...colors.grayLight)
    doc.setLineWidth(0.5)
    doc.line(15, pageHeight - 25, pageWidth - 15, pageHeight - 25)
    
    // Texte confidentiel
    doc.setFontSize(6)
    doc.setFont('helvetica', 'normal')
    doc.setTextColor(...colors.gray)
    
    const confidentialText = "JS dit : DOCUMENT CONFIDENTIEL - Les renseignements contenus dans ce document sont de nature confidentielle et privilégiée. " +
      "Ils sont destinés uniquement à l'usage de la personne ou de l'organisme autorisé. Toute divulgation, distribution ou " +
      "reproduction non autorisée est strictement interdite et peut être illégale."
    
    const splitText = doc.splitTextToSize(confidentialText, pageWidth - 30)
    doc.text(splitText, pageWidth / 2, pageHeight - 20, { align: 'center' })
    
    // Numéro de page
    doc.setFontSize(8)
    doc.text(`Page ${pageNumber} / ${totalPages}`, pageWidth / 2, pageHeight - 10, { align: 'center' })
  }

  // Fonction pour ajouter les informations client
  const addClientInfo = (doc, note, startY) => {
    let y = startY
    const pageWidth = doc.internal.pageSize.getWidth()
    
    // Titre de section
    doc.setFontSize(9)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.tealDark)
    doc.text('INFORMATIONS CLIENT', 15, y)
    y += 7
    
    // Contenu
    doc.setFontSize(8)
    doc.setFont('helvetica', 'normal')
    doc.setTextColor(...colors.slate)
    
    // Nom complet
    const nomComplet = `${note.client_nom || ''}, ${note.client_prenom || ''} ${note.client_no_leopard ? `(No: ${note.client_no_leopard})` : ''}`
    doc.text(nomComplet.trim(), 17, y)
    y += 5
    
    // RAMQ
    if (note.client_no_ramq) {
      doc.text(`RAMQ: ${note.client_no_ramq}`, 17, y)
      y += 5
    }
    
    // Adresse
    const adresseParts = []
    if (note.client_adresse) {
      let adr = note.client_adresse
      if (note.client_appartement) adr += `, app. ${note.client_appartement}`
      adresseParts.push(adr)
    }
    
    const cityParts = []
    if (note.client_ville) cityParts.push(note.client_ville)
    if (note.client_province) cityParts.push(note.client_province)
    if (note.client_code_postal) cityParts.push(note.client_code_postal)
    if (cityParts.length) adresseParts.push(cityParts.join(', '))
    
    if (adresseParts.length) {
      doc.text(`Adresse: ${adresseParts.join(', ')}`, 17, y)
      y += 5
    }
    
    // Téléphones
    const tels = []
    if (note.client_telephone) tels.push(`Tél: ${note.client_telephone}`)
    if (note.client_cellulaire) tels.push(`Cell: ${note.client_cellulaire}`)
    if (tels.length) {
      doc.text(tels.join(' | '), 17, y)
      y += 5
    }
    
    // Ligne de séparation
    doc.setDrawColor(...colors.grayLight)
    doc.setLineWidth(0.5)
    doc.line(15, y + 2, pageWidth - 15, y + 2)
    
    return y + 7
  }

  // Fonction pour ajouter les détails de l'intervention
  const addInterventionDetails = (doc, note, startY) => {
    let y = startY
    const pageWidth = doc.internal.pageSize.getWidth()
    
    // Titre de section
    doc.setFontSize(9)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.tealDark)
    doc.text('DÉTAILS DE L\'INTERVENTION', 15, y)
    y += 7
    
    doc.setFontSize(8)
    doc.setFont('helvetica', 'normal')
    doc.setTextColor(...colors.slate)
    
    // Ligne 1: Type, Date, Heure
    const type = note.type_note || 'Non spécifié'
    const dateInt = note.date_intervention ? new Date(note.date_intervention).toLocaleDateString('fr-CA') : 
                    new Date(note.date_note).toLocaleDateString('fr-CA')
    const heure = note.heure_intervention || '-'
    
    doc.text(`Type: ${type}`, 15, y)
    doc.text(`Date: ${dateInt}`, 80, y)
    doc.text(`Heure: ${heure}`, 145, y)
    y += 5
    
    // Ligne 2: Mode, Type intervention, Durée
    const mode = note.mode_intervention || '-'
    const typeInt = note.type_intervention || '-'
    const duree = note.duree_intervention || '-'
    
    doc.text(`Mode: ${mode}`, 15, y)
    doc.text(`Type intervention: ${typeInt}`, 80, y)
    doc.text(`Durée: ${duree}`, 145, y)
    y += 5
    
    // Indicateurs spéciaux
    const indicators = []
    if (note.note_tardive === 1) indicators.push('NOTE TARDIVE')
    if (note.note_de_tier === 1) indicators.push('NOTE DE TIERS')
    if (note.note_liee_id) {
      const typeLien = note.type_lien || 'Liée'
      indicators.push(`${typeLien} à note #${note.note_liee_id}`)
    }
    
    if (indicators.length) {
      doc.setFont('helvetica', 'bold')
      doc.setTextColor(...colors.tealDark)
      doc.text(indicators.join(' • '), 15, y)
      y += 6
    }
    
    // Ligne de séparation
    doc.setDrawColor(...colors.grayLight)
    doc.setLineWidth(0.5)
    doc.line(15, y + 2, pageWidth - 15, y + 2)
    
    return y + 7
  }

  // Fonction pour ajouter le titre de la note
  const addNoteTitle = (doc, note, startY) => {
    let y = startY
    const pageWidth = doc.internal.pageSize.getWidth()
    
    const titre = note.titre || 'Note sans titre'
    
    doc.setFontSize(14)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.slateDark)
    doc.text(titre, 15, y)
    y += 7
    
    // Ligne sous le titre
    doc.setDrawColor(...colors.teal)
    doc.setLineWidth(1.5)
    doc.line(15, y, pageWidth - 15, y)
    
    return y + 7
  }

  // Fonction pour ajouter le contenu de la note
  const addNoteContent = (doc, note, startY) => {
    let y = startY
    const pageWidth = doc.internal.pageSize.getWidth()
    const pageHeight = doc.internal.pageSize.getHeight()
    const maxY = pageHeight - 35 // Marge pour le footer
    
    const contenu = note.contenu || '[Aucun contenu]'
    const paragraphs = contenu.split('\n')
    
    for (const para of paragraphs) {
      const trimmed = para.trim()
      
      if (!trimmed) {
        y += 3
        continue
      }
      
      // Détecter les sections (MAJUSCULES + ':')
      if (trimmed.endsWith(':') && trimmed === trimmed.toUpperCase()) {
        // Section
        if (y > maxY - 10) {
          doc.addPage()
          y = 40
        }
        
        doc.setFontSize(10)
        doc.setFont('helvetica', 'bold')
        doc.setTextColor(...colors.tealDark)
        const lines = doc.splitTextToSize(trimmed, pageWidth - 30)
        doc.text(lines, 15, y)
        y += lines.length * 5 + 2
        
      } else if (trimmed.startsWith('•') || trimmed.startsWith('-') || trimmed.startsWith('*')) {
        // Élément de liste
        if (y > maxY - 5) {
          doc.addPage()
          y = 40
        }
        
        doc.setFontSize(9)
        doc.setFont('helvetica', 'normal')
        doc.setTextColor(...colors.slate)
        const lines = doc.splitTextToSize(trimmed, pageWidth - 35)
        doc.text(lines, 20, y)
        y += lines.length * 5
        
      } else {
        // Paragraphe normal
        if (y > maxY - 10) {
          doc.addPage()
          y = 40
        }
        
        doc.setFontSize(9)
        doc.setFont('helvetica', 'normal')
        doc.setTextColor(...colors.slate)
        const lines = doc.splitTextToSize(trimmed, pageWidth - 30)
        doc.text(lines, 15, y)
        y += lines.length * 4.5 + 2
      }
    }
    
    return y
  }

  // Fonction pour ajouter la signature
  const addSignature = (doc, note, startY) => {
    let y = startY + 10
    const pageWidth = doc.internal.pageSize.getWidth()
    
    // Ligne de séparation
    doc.setDrawColor(...colors.grayLight)
    doc.setLineWidth(1)
    doc.line(pageWidth / 2 + 10, y, pageWidth - 15, y)
    y += 7
    
    // Texte de signature
    doc.setFontSize(8)
    doc.setFont('helvetica', 'italic')
    doc.setTextColor(...colors.gray)
    doc.text('Note signée électroniquement par:', pageWidth / 2 + 10, y)
    y += 5
    
    doc.setFontSize(10)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.slateDark)
    doc.text(note.signature_nom || 'Non signé', pageWidth / 2 + 10, y)
    y += 5
    
    const sigDate = note.signature_date ? 
      new Date(note.signature_date).toLocaleString('fr-CA', { 
        year: 'numeric', 
        month: '2-digit', 
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
      }) : '-'
    
    doc.setFontSize(8)
    doc.setFont('helvetica', 'normal')
    doc.setTextColor(...colors.gray)
    doc.text(`Le ${sigDate}`, pageWidth / 2 + 10, y)
    
    return y + 10
  }

  // Fonction pour générer le PDF d'une seule note
  const generateSingleNotePDF = async (note, logoBase64 = null) => {
    const doc = new jsPDF()
    
    // Page de la note
    let y = addHeader(doc, 'Note Psychosociale', logoBase64)
    y = addClientInfo(doc, note, y)
    y = addInterventionDetails(doc, note, y)
    y = addNoteTitle(doc, note, y)
    y = addNoteContent(doc, note, y)
    
    if (note.verrouille === 1) {
      addSignature(doc, note, y)
    }
    
    // Ajouter le footer
    const totalPages = doc.internal.pages.length - 1
    for (let i = 1; i <= totalPages; i++) {
      doc.setPage(i)
      addFooter(doc, i, totalPages)
    }
    
    return doc
  }

  // Fonction pour générer le PDF avec sommaire (export multiple)
  const generateNotesWithSummaryPDF = async (notes, leopardNumber, logoBase64 = null) => {
    const doc = new jsPDF()
    
    // PAGE 1: SOMMAIRE
    let y = addHeader(doc, 'RAPPORT SOMMAIRE DES INTERVENTIONS', logoBase64)
    
    if (notes.length > 0) {
      y = addClientInfo(doc, notes[0], y)
    }
    
    y += 5
    
    // Titre du tableau
    doc.setFontSize(12)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.tealDark)
    doc.text('RÉSUMÉ DES INTERVENTIONS', doc.internal.pageSize.getWidth() / 2, y, { align: 'center' })
    y += 7
    
    // Ligne de séparation
    const pageWidth = doc.internal.pageSize.getWidth()
    doc.setDrawColor(...colors.teal)
    doc.setLineWidth(2)
    doc.line(15, y, pageWidth - 15, y)
    y += 7
    
    // En-têtes du tableau
    doc.setFontSize(9)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.slate)
    doc.text('DATE', 25, y, { align: 'center' })
    doc.text('TITRE', 95, y, { align: 'center' })
    doc.text('MODE', 165, y, { align: 'center' })
    y += 5
    
    // Ligne de séparation
    doc.setDrawColor(...colors.grayLight)
    doc.setLineWidth(0.5)
    doc.line(15, y, pageWidth - 15, y)
    y += 5
    
    // Lignes du tableau
    doc.setFontSize(8)
    doc.setFont('helvetica', 'normal')
    
    for (const note of notes) {
      const dateStr = note.date_intervention ? 
        new Date(note.date_intervention).toLocaleDateString('fr-CA') :
        new Date(note.date_note).toLocaleDateString('fr-CA')
      const titre = note.titre || 'Sans titre'
      const mode = note.mode_intervention || '-'
      
      doc.text(dateStr, 25, y, { align: 'center' })
      
      // Titre tronqué si nécessaire
      const titreTronque = doc.splitTextToSize(titre, 100)
      doc.text(titreTronque[0], 95, y, { align: 'center' })
      doc.text(mode, 165, y, { align: 'center' })
      y += 5
      
      // Ligne légère
      doc.setDrawColor(...colors.grayLight)
      doc.setLineWidth(0.3)
      doc.line(15, y, pageWidth - 15, y)
      y += 4
    }
    
    // Total des interventions
    y += 5
    doc.setFontSize(9)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(...colors.tealDark)
    doc.text(`Total des interventions: ${notes.length}`, pageWidth - 15, y, { align: 'right' })
    
    // PAGES SUIVANTES: NOTES DÉTAILLÉES
    for (const note of notes) {
      doc.addPage()
      let y = addHeader(doc, 'Note Psychosociale', logoBase64)
      y = addClientInfo(doc, note, y)
      y = addInterventionDetails(doc, note, y)
      y = addNoteTitle(doc, note, y)
      y = addNoteContent(doc, note, y)
      
      if (note.verrouille === 1) {
        addSignature(doc, note, y)
      }
    }
    
    // Ajouter les footers sur toutes les pages
    const totalPages = doc.internal.pages.length - 1
    for (let i = 1; i <= totalPages; i++) {
      doc.setPage(i)
      addFooter(doc, i, totalPages)
    }
    
    return doc
  }

  // Fonction pour charger le logo en base64
  const loadLogoAsBase64 = async () => {
    try {
      const response = await fetch(logoBlanc)
      const blob = await response.blob()
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.onloadend = () => resolve(reader.result)
        reader.onerror = reject
        reader.readAsDataURL(blob)
      })
    } catch (error) {
      console.warn('Impossible de charger le logo:', error)
      return null
    }
  }

  return {
    generateSingleNotePDF,
    generateNotesWithSummaryPDF,
    loadLogoAsBase64
  }
}