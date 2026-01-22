// src/templates/pdf/NoteTemplate.js
import { jsPDF } from 'jspdf'
import logoPath from '@/assets/images/LeopardLogo.png' // Ton logo

export const genererNotePDF = async (client, fullNotes) => {
  const doc = new jsPDF()
  let y = 20

  // --- 1. EN-TÊTE (LOGO + IDENTIFICATION) ---
  const img = new Image()
  img.src = logoPath
  await new Promise((resolve) => {
    img.onload = () => {
      doc.addImage(img, 'PNG', 15, 10, 25, 25)
      resolve()
    }
  })

  doc.setFontSize(10)
  doc.setFont('helvetica', 'bold')
  doc.text('IDENTIFICATION DU CLIENT', 50, 15)
  
  doc.setFont('helvetica', 'normal')
  doc.setFontSize(9)
  const infos = [
    `Nom : ${client.nom}, ${client.prenom}`,
    `Dossier : ${client.no_dossier_leopard || 'N/A'}`,
    `Date de naissance : ${client.date_naissance || 'N/A'}`,
    `RAMQ : ${client.numero_assurance_maladie || 'N/A'}`
  ]
  
  let infoY = 20
  infos.forEach(line => {
    doc.text(line, 50, infoY)
    infoY += 5
  })

  y = 45
  doc.setDrawColor(200)
  doc.line(15, y, 195, y) // Ligne de séparation
  y += 10

  // --- 2. TITRE DU DOCUMENT ---
  doc.setFontSize(14)
  doc.setFont('helvetica', 'bold')
  doc.text('NOTES ÉVOLUTIVES DU DOSSIER', 105, y, { align: 'center' })
  y += 15

  // --- 3. BOUCLE SUR LES NOTES ---
  fullNotes.forEach((note, index) => {
    // Vérifier s'il reste de la place sur la page
    if (y > 250) {
      doc.addPage()
      y = 20
    }

    // Encadré de la note
    doc.setFillColor(245, 245, 245)
    doc.rect(15, y, 180, 8, 'F')
    
    doc.setFontSize(10)
    doc.setFont('helvetica', 'bold')
    const dateNote = new Date(note.date_note).toLocaleDateString('fr-CA')
    doc.text(`NOTE DU ${dateNote} - ${note.type_note}`, 18, y + 5)
    y += 13

    // Contenu de la note
    doc.setFont('helvetica', 'normal')
    doc.setFontSize(10)
    const lignes = doc.splitTextToSize(note.contenu, 170)
    
    lignes.forEach(ligne => {
      if (y > 275) {
        doc.addPage()
        y = 20
      }
      doc.text(ligne, 17, y)
      y += 5
    })

    // Signature
    y += 5
    doc.setFont('helvetica', 'italic')
    doc.setFontSize(9)
    doc.text(`Signé électroniquement par : ${note.signature_nom || 'Intervenant'}`, 17, y)
    
    y += 15 // Espace avant la prochaine note
  })

  return doc
}