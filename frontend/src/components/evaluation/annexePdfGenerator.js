/**
 * annexePdfGenerator.js
 * Générateur PDF pour l'Annexe — Personnes consultées
 * Approche : reconstruction jsPDF programmatique (ZERO html2canvas)
 * Compatible Wails : retourne base64 au lieu de doc.save()
 */

import jsPDF from 'jspdf'

// ── Constantes mise en page ────────────────────────────────────────────────
const PAGE = {
  w: 215.9,   // mm — format lettre
  h: 279.4,
  mL: 18,     // marge gauche
  mR: 18,     // marge droite
  mT: 18,     // marge haut
  mB: 16,     // marge bas
  get cW() { return this.w - this.mL - this.mR }, // largeur contenu
}

// ── Palette — look "formulaire officiel" ──────────────────────────────────
const COL = {
  black:      [0,   0,   0  ],
  white:      [255, 255, 255],
  headerBg:   [216, 216, 216], // gris clair — les en-têtes de section
  borderGray: [136, 136, 136],
  numBg:      [232, 232, 232], // badge numéro personne
  textDark:   [0,   0,   0  ],
  textGray:   [85,  85,  85 ],
  textLight:  [119, 119, 119],
  accentBlue: [26,  107, 181], // couleur accent (facultatif)
}

// ── Classe générateur ──────────────────────────────────────────────────────
class AnnexePDFGenerator {
  constructor(data) {
    this.doc      = new jsPDF({ unit: 'mm', format: 'letter', orientation: 'portrait' })
    this.data     = data         // { client, personnes, dateDoc }
    this.y        = PAGE.mT
    this.page     = 1
  }

  // ── Utilitaires ───────────────────────────────────────────────────────────

  setFont(size, style = 'normal', color = COL.textDark) {
    this.doc.setFontSize(size)
    this.doc.setFont('helvetica', style)
    this.doc.setTextColor(...color)
  }

  setFill(color) { this.doc.setFillColor(...color) }
  setDraw(color) { this.doc.setDrawColor(...color) }
  setLW(lw)      { this.doc.setLineWidth(lw) }

  rect(x, y, w, h, fill, draw = null, lw = 0.3, radius = 0) {
    this.setFill(fill)
    if (draw) { this.setDraw(draw); this.setLW(lw) }
    const style = draw ? 'FD' : 'F'
    if (radius > 0) {
      this.doc.roundedRect(x, y, w, h, radius, radius, style)
    } else {
      this.doc.rect(x, y, w, h, style)
    }
  }

  border(x, y, w, h, color = COL.borderGray, lw = 0.5) {
    this.setDraw(color)
    this.setLW(lw)
    this.doc.setFillColor(255, 255, 255)
    this.doc.rect(x, y, w, h, 'FD')
  }

  line(x1, y1, x2, y2, color = COL.borderGray, lw = 0.5) {
    this.setDraw(color)
    this.setLW(lw)
    this.doc.line(x1, y1, x2, y2)
  }

  text(str, x, y, opts = {}) {
    this.doc.text(str || '', x, y, opts)
  }

  // Texte avec coupure automatique — retourne la hauteur utilisée
  multiText(str, x, y, maxW, lineH, size, style = 'normal', color = COL.textDark) {
    if (!str?.trim()) return 0
    this.setFont(size, style, color)
    const lines = this.doc.splitTextToSize(str, maxW)
    lines.forEach((line, i) => {
      this.checkBreak(lineH)
      this.doc.text(line, x, this.y)
      if (i < lines.length - 1) this.y += lineH
    })
    return lines.length * lineH
  }

  checkBreak(needed = 20) {
    if (this.y + needed > PAGE.h - PAGE.mB) {
      this.newPage()
    }
  }

  newPage() {
    this.doc.addPage()
    this.page++
    this.y = PAGE.mT
    this.drawRunningHeader()
    this.drawFooter()
  }

  // ── En-tête répété (pages suivantes) ──────────────────────────────────────

  drawRunningHeader() {
    const client = this.data.client
    const name   = client?.nom ? `${client.nom}` : '—'

    this.setFont(8, 'bold', COL.textGray)
    this.text('Annexe — Personnes consultées (suite)', PAGE.mL, this.y)

    this.setFont(8, 'normal', COL.textGray)
    this.doc.text(name, PAGE.w - PAGE.mR, this.y, { align: 'right' })

    this.line(PAGE.mL, this.y + 2, PAGE.w - PAGE.mR, this.y + 2, COL.borderGray, 0.3)
    this.y += 8
  }

  // ── Pied de page ──────────────────────────────────────────────────────────

  drawFooter() {
    const client  = this.data.client
    const dossier = client?.no_dossier_leopard || ''
    const footerY = PAGE.h - PAGE.mB + 2

    this.line(PAGE.mL, footerY - 2, PAGE.w - PAGE.mR, footerY - 2, COL.borderGray, 0.3)

    this.setFont(7.5, 'normal', COL.textLight)
    this.text("Annexe à l'évaluation psychosociale — Homologation du mandat de protection", PAGE.mL, footerY + 2)

    if (dossier) {
      this.doc.text(`Dossier : ${dossier}`, PAGE.w - PAGE.mR, footerY + 2, { align: 'right' })
    }

    // Numéro de page
    this.setFont(7.5, 'normal', COL.textLight)
    this.doc.text(`Page ${this.page}`, PAGE.w / 2, footerY + 2, { align: 'center' })
  }

  // ── EN-TÊTE DU DOCUMENT ───────────────────────────────────────────────────

  drawDocHeader() {
    const client = this.data.client

    // Ligne de titre avec bordure bas
    this.setFont(12, 'bold', COL.textDark)
    this.text("Annexe à l'évaluation psychosociale", PAGE.mL, this.y)

    this.setFont(10, 'normal', COL.textDark)
    this.y += 5
    this.text("dans le cadre de l'homologation d'un mandat de protection", PAGE.mL, this.y)

    // Référence doc (droite)
    this.setFont(8, 'normal', COL.textGray)
    this.doc.text('096-DGSP-2025-07', PAGE.w - PAGE.mR, this.y - 5, { align: 'right' })
    this.doc.text('Annexe — Personnes consultées', PAGE.w - PAGE.mR, this.y, { align: 'right' })

    // Ligne séparatrice double
    this.line(PAGE.mL, this.y + 3, PAGE.w - PAGE.mR, this.y + 3, COL.textDark, 1.5)
    this.y += 8

    // ── ENCADRÉ CLIENT ──────────────────────────────────────────────────────
    const boxH = 22
    this.border(PAGE.mL, this.y, PAGE.cW, boxH, COL.textDark, 1.2)

    this.setFont(7.5, 'bold', COL.textGray)
    this.text('PERSONNE VISÉE PAR L\'ÉVALUATION', PAGE.mL + 3, this.y + 5)

    const clientName = client?.nom
      ? `${client.prenom || ''} ${client.nom}`.trim()
      : '—'

    this.setFont(12, 'bold', COL.textDark)
    this.text(clientName, PAGE.mL + 3, this.y + 12)

    const ddn = client?.date_naissance
      ? `Date de naissance : ${client.date_naissance}`
      : ''

    const dossier = client?.no_dossier_leopard
      ? `No de dossier : ${client.no_dossier_leopard}`
      : ''

    this.setFont(9, 'normal', COL.textDark)
    if (ddn) this.text(ddn, PAGE.mL + 3, this.y + 17)
    if (dossier) this.text(dossier, PAGE.mL + 3 + (ddn ? 70 : 0), this.y + 17)

    this.y += boxH + 5
  }

  // ── EN-TÊTE DE SECTION ────────────────────────────────────────────────────

  drawSectionHeader(pageLabel = 'Annexe — page 1') {
    const hH = 9

    // Fond gris
    this.rect(PAGE.mL, this.y, PAGE.cW, hH, COL.headerBg, COL.borderGray, 0.5)

    this.setFont(9.5, 'bold', COL.textDark)
    this.text(
      "C. Personnes consultées dans le cadre de l'évaluation",
      PAGE.mL + 3,
      this.y + 6
    )

    // Texte italique (suite)
    this.setFont(8, 'normal', COL.textGray)
    this.text('(suite du formulaire principal)', PAGE.mL + 120, this.y + 6)

    // Label page (droite)
    this.setFont(7.5, 'normal', COL.textGray)
    this.doc.text(pageLabel, PAGE.w - PAGE.mR - 2, this.y + 6, { align: 'right' })

    this.y += hH

    // Corps de section : début du border bas
    this.border(PAGE.mL, this.y, PAGE.cW, 4, COL.borderGray, 0.5)
    this.y += 3
  }

  // ── BLOC PERSONNE ─────────────────────────────────────────────────────────

  drawPersonBlock(personne, num) {
    const p  = PAGE
    const mI = p.mL + 2  // marge intérieure
    const wI = p.cW - 4  // largeur intérieure

    // Calcul hauteur estimée du bloc
    const noteLines = personne.notes
      ? this.doc.splitTextToSize(personne.notes, wI - 4).length
      : 0
    const estimH = 72 + (noteLines > 0 ? noteLines * 4 + 10 : 0)

    this.checkBreak(estimH + 6)

    const startY = this.y

    // Badge numéro
    const badgeW = 28
    this.rect(mI, this.y, badgeW, 6, COL.numBg, COL.borderGray, 0.3)
    this.setFont(7.5, 'bold', COL.textGray)
    this.text(`Personne ${num}`, mI + 2, this.y + 4.5)
    this.y += 8

    // ── Rangée 1 : Nom / Prénom ─────────────────────────────────────────────
    const col1W = wI / 2
    this.drawField('Nom',    personne.nom || '',    mI,           this.y, col1W - 3)
    this.drawField('Prénom', personne.prenom || '', mI + col1W,   this.y, col1W - 3)
    this.y += 10

    // ── Rangée 2 : Lien / Date / Titre ──────────────────────────────────────
    const c3w = [wI * 0.40, wI * 0.32, wI * 0.28]
    this.drawField('Lien avec la personne visée',      personne.lien || '',    mI,                        this.y, c3w[0] - 3)
    this.drawField('Date de consultation (aaaa-mm-jj)', personne.dateConsult || '', mI + c3w[0],          this.y, c3w[1] - 3)
    this.drawField('Titre / fonction',                  personne.titre || '',   mI + c3w[0] + c3w[1],     this.y, c3w[2] - 3)
    this.y += 10

    // ── Rangée 3 : Adresse / CP / Tél / Poste ───────────────────────────────
    const c4w = [wI * 0.40, wI * 0.15, wI * 0.27, wI * 0.18]
    this.drawField('Adresse (numéro, rue, ville)', personne.adresse || '',  mI,                         this.y, c4w[0] - 3)
    this.drawField('Code postal',                  personne.codePostal || '', mI + c4w[0],              this.y, c4w[1] - 3)
    this.drawField('No tél. (travail)',             personne.tel || '',      mI + c4w[0] + c4w[1],      this.y, c4w[2] - 3)
    this.drawField('Poste',                         personne.poste || '',    mI + c4w[0] + c4w[1] + c4w[2], this.y, c4w[3] - 3)
    this.y += 10

    // ── Rangée 4 : Courriel ──────────────────────────────────────────────────
    this.drawField('Courriel', personne.courriel || '', mI, this.y, col1W - 3)
    this.y += 10

    // ── Cases à cocher — Type de consultation ───────────────────────────────
    this.setFont(7.5, 'bold', COL.textGray)
    this.text('TYPE DE CONSULTATION', mI, this.y)
    this.y += 4

    const types = [
      { label: 'Entretien téléphonique', key: 'tel' },
      { label: 'Rencontre',              key: 'rencontre' },
      { label: 'Téléconsultation',       key: 'video' },
      { label: 'Courriel et/ou courrier', key: 'courriel' },
    ]

    let cbX = mI
    types.forEach(type => {
      const checked = personne.typesConsultation?.includes(type.key) || false
      this.drawCheckbox(cbX, this.y, type.label, checked)
      cbX += 48
    })
    this.y += 8

    // ── Organisme / employeur (si renseigné) ────────────────────────────────
    if (personne.organisme || personne.adresseOrganisme || personne.telOrganisme) {
      this.setFont(7.5, 'bold', COL.textGray)
      // Ligne pointillée
      this.setDraw([187, 187, 187])
      this.setLW(0.3)
      this.doc.setLineDashPattern([1, 2], 0)
      this.line(mI, this.y - 1, mI + wI, this.y - 1)
      this.doc.setLineDashPattern([], 0)

      this.text('Contact au travail / Organisme', mI, this.y + 3)
      this.y += 6

      const co2W = wI / 2
      this.drawField('Organisme / employeur',      personne.organisme || '',        mI,          this.y, co2W - 3)
      this.drawField('Adresse professionnelle',    personne.adresseOrganisme || '', mI + co2W,   this.y, co2W - 3)
      this.y += 10

      this.drawField('No tél. (organisme)', personne.telOrganisme || '', mI, this.y, wI * 0.33 - 3)
      this.drawField('Courriel (organisme)', personne.courrielOrganisme || '', mI + wI * 0.33, this.y, wI * 0.33 - 3)
      this.y += 10
    }

    // ── Notes libres ─────────────────────────────────────────────────────────
    if (personne.notes?.trim()) {
      this.setFont(7.5, 'bold', COL.textGray)
      this.text('NOTES LIBRES', mI, this.y)
      this.y += 4

      this.setFont(9, 'normal', COL.textDark)
      const notesLines = this.doc.splitTextToSize(personne.notes, wI - 4)
      const notesH     = notesLines.length * 4 + 6

      this.border(mI, this.y, wI, notesH, [204, 204, 204], 0.3)
      let nY = this.y + 4
      notesLines.forEach(line => {
        this.doc.text(line, mI + 2, nY)
        nY += 4
      })
      this.y = nY + 4
    }

    // Bordure autour de tout le bloc personne
    const blockH = this.y - startY + 4
    this.setDraw([187, 187, 187])
    this.setLW(0.4)
    this.doc.rect(p.mL, startY - 2, p.cW, blockH, 'S')

    this.y += 6
  }

  // ── Champ avec label + ligne ───────────────────────────────────────────────

  drawField(label, value, x, y, maxW) {
    // Label
    this.setFont(7.5, 'bold', COL.textGray)
    this.text(label.toUpperCase(), x, y)

    // Valeur
    this.setFont(9.5, 'normal', COL.textDark)
    const displayVal = value || ''
    const truncated  = displayVal.length > 0
      ? this.doc.splitTextToSize(displayVal, maxW)[0]
      : ''
    this.text(truncated, x, y + 5)

    // Ligne de saisie
    this.line(x, y + 6.5, x + maxW, y + 6.5, COL.borderGray, 0.4)
  }

  // ── Case à cocher ────────────────────────────────────────────────────────

  drawCheckbox(x, y, label, checked = false) {
    const size = 3.5

    // Carré
    this.setDraw(COL.borderGray)
    this.setLW(0.4)
    this.doc.setFillColor(...COL.white)
    this.doc.rect(x, y - size, size, size, 'FD')

    // Coche si coché
    if (checked) {
      this.setFont(7, 'bold', COL.accentBlue)
      this.text('✓', x + 0.3, y - 0.5)
    }

    // Label
    this.setFont(8.5, 'normal', COL.textDark)
    this.text(label, x + size + 1.5, y - 0.5)
  }

  // ── GÉNÉRATION COMPLÈTE ──────────────────────────────────────────────────

  generate() {
    // Page 1 : en-tête + section header
    this.drawDocHeader()
    this.drawSectionHeader('Annexe — page 1')

    // Blocs personnes
    const personnes = this.data.personnes || []
    personnes.forEach((p, i) => {
      this.drawPersonBlock(p, i + 5) // commence à 5 (suite du formulaire officiel)
    })

    // Pied de page sur toutes les pages
    const totalPages = this.doc.getNumberOfPages()
    for (let i = 1; i <= totalPages; i++) {
      this.doc.setPage(i)
      this.drawFooter()
    }

    return this.doc
  }
}

// ── Fonction principale exportée ───────────────────────────────────────────

/**
 * Génère le PDF de l'annexe et le retourne en base64
 * Compatible Wails : pas de doc.save(), retourne la donnée brute
 *
 * @param {Object} data
 *   - client    : { nom, prenom, date_naissance, no_dossier_leopard }
 *   - personnes : Array de personnes consultées
 *   - nomLeopard : string — identifiant Léopard de l'évaluation associée (pour nommage)
 *
 * @returns {{ base64: string, filename: string }}
 */
export async function generateAnnexePDF(data) {
  if (!data) throw new Error('Données manquantes')

  const generator = new AnnexePDFGenerator(data)
  const doc       = generator.generate()

  // Générer en base64 (pas de download navigateur)
  const base64    = doc.output('datauristring') // "data:application/pdf;base64,..."

  // Nommage
  const clientNom  = (data.client?.nom || 'client').replace(/\s+/g, '_')
  const leopard    = data.nomLeopard ? `_${data.nomLeopard}` : ''
  const dateStr    = new Date().toISOString().slice(0, 10).replace(/-/g, '')
  const filename   = `Annexe_PersonnesConsultees${leopard}_${clientNom}_${dateStr}.pdf`

  return { base64, filename }
}

/**
 * Sauvegarde le PDF via la fonction Go de Wails
 * Doit être appelé dans un contexte où window.go est disponible
 *
 * @param {Object} data       — données pour generateAnnexePDF
 * @param {string} leopardNumber — numéro Léopard du CLIENT (pour trouver son dossier)
 * @returns {string} path du fichier sauvegardé
 */
export async function saveAnnexePDFToFolder(data, leopardNumber) {
  const { base64, filename } = await generateAnnexePDF(data)

  const result = await window.go.main.App.SavePDFToClientFolder({
    leopardNumber,
    subfolder:  'Evaluations',
    filename,
    pdfBase64:  base64,
  })

  if (!result.Success) {
    throw new Error(result.Error || 'Erreur sauvegarde PDF')
  }

  return result.Path
}