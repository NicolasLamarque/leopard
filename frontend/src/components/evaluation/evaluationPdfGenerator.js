/**
 * evaluationPdfGenerator.js
 * Générateur PDF professionnel pour évaluations du fonctionnement social
 * Utilise jsPDF (déjà dans package.json)
 * Les données arrivent DÉJÀ décryptées du backend Go
 */


import jsPDF from 'jspdf'
// ── Palette & constantes ───────────────────────────────────────────────────
const C = {
  // Grille
  pageW:      210,
  pageH:      297,
  marginL:    18,
  marginR:    18,
  marginT:    20,
  marginB:    20,
  contentW:   210 - 18 - 18,

  // Couleurs (RGB)
  primary:    [30,  41,  59],   // slate-800
  accent:     [13, 148, 136],   // teal-600
  accentLight:[204, 241, 239],  // teal-100
  muted:      [100, 116, 132],  // slate-500
  mutedLight: [241, 245, 249],  // slate-100
  white:      [255, 255, 255],
  border:     [226, 232, 240],  // slate-200
  danger:     [220,  38,  38],  // rouge signature
  success:    [16, 185, 129],   // emerald-500

  // Types
  tutelle:            { r: 37,  g: 99,  b: 235 },  // blue-600
  mandat:             { r: 147, g: 51,  b: 234 },  // purple-600
  conseil_tutelle:    { r: 217, g: 119, b: 6   },  // amber-600
  suivi_psychosocial: { r: 13,  g: 148, b: 136 },  // teal-600
}

// ── Helpers ────────────────────────────────────────────────────────────────
const rgb = (arr) => ({ r: arr[0], g: arr[1], b: arr[2] })

function typeColor(type) {
  return C[type] || { r: 30, g: 41, b: 59 }
}

function typeLabel(type) {
  const m = {
    tutelle:            'Régime de tutelle',
    mandat:             'Homologation de mandat',
    conseil_tutelle:    'Conseil de tutelle',
    suivi_psychosocial: 'Suivi psychosocial'
  }
  return m[type] || 'Évaluation du fonctionnement social'
}

function typeSubLabel(type) {
  const m = {
    tutelle:            'Évaluation du fonctionnement social — OTSTCFQ',
    mandat:             'Rapport pour homologation — OTSTCFQ',
    conseil_tutelle:    'Rapport d\'évaluation — OTSTCFQ',
    suivi_psychosocial: 'Note évolutive de suivi psychosocial'
  }
  return m[type] || 'Travail social — Secteur privé'
}

function formatDateLong(d) {
  if (!d) return '—'
  return new Date(d).toLocaleDateString('fr-CA', {
    year: 'numeric', month: 'long', day: 'numeric'
  })
}

function formatDateTimeLong(d) {
  if (!d) return '—'
  return new Date(d).toLocaleString('fr-CA', {
    year: 'numeric', month: 'long', day: 'numeric',
    hour: '2-digit', minute: '2-digit'
  })
}

// ── Classe Doc ─────────────────────────────────────────────────────────────
class EvalPDFDocument {
  constructor(evalData, client) {
    this.doc    = new jsPDF({ unit: 'mm', format: 'a4', orientation: 'portrait' })
    this.eval   = evalData
    this.client = client
    this.page   = 1
    this.y      = C.marginT
    this.type   = evalData.type_evaluation || 'tutelle'
    this.tc     = typeColor(this.type)
  }

  // ── Utilitaires PDF ──────────────────────────────────────────────────────

  rgb(...args) {
    if (args.length === 1) return args[0]
    return { r: args[0], g: args[1], b: args[2] }
  }

  setFont(size, style = 'normal', color = C.primary) {
    this.doc.setFontSize(size)
    this.doc.setFont('helvetica', style)
    this.doc.setTextColor(...color)
  }

  setFill(color) { this.doc.setFillColor(...color) }
  setDraw(color) { this.doc.setDrawColor(...color) }

  lineH(y, color = C.border, lw = 0.3) {
    this.doc.setLineWidth(lw)
    this.setDraw(color)
    this.doc.line(C.marginL, y, C.pageW - C.marginR, y)
  }

  rect(x, y, w, h, fill, draw = null, rounding = 0) {
    this.setFill(fill)
    if (draw) this.setDraw(draw)
    const style = draw ? 'FD' : 'F'
    if (rounding > 0) {
      this.doc.roundedRect(x, y, w, h, rounding, rounding, style)
    } else {
      this.doc.rect(x, y, w, h, style)
    }
  }

  text(str, x, y, opts = {}) {
    this.doc.text(str, x, y, opts)
  }

  /**
   * Wrap text et retourne le nb de lignes utilisées
   */
  wrappedText(str, x, y, maxW, lineH, size, style = 'normal', color = C.primary) {
    if (!str?.trim()) return 0
    this.setFont(size, style, color)
    const lines = this.doc.splitTextToSize(str, maxW)
    lines.forEach((line, i) => {
      this.checkPageBreak(lineH + 2)
      this.doc.text(line, x, this.y + (i === 0 ? 0 : lineH))
      if (i < lines.length - 1) this.y += lineH
    })
    return lines.length
  }

  checkPageBreak(needed = 20) {
    if (this.y + needed > C.pageH - C.marginB) {
      this.addPage()
    }
  }

  addPage() {
    this.doc.addPage()
    this.page++
    this.y = C.marginT
    this.drawRunningHeader()
    this.drawPageNumber()
  }

  // ── Éléments récurrents ──────────────────────────────────────────────────

  drawRunningHeader() {
    // Bande fine en haut de chaque page suivante
    this.rect(0, 0, C.pageW, 10, [this.tc.r, this.tc.g, this.tc.b])
    this.setFont(6, 'bold', C.white)
    this.text(typeLabel(this.type).toUpperCase(), C.marginL, 6.5)

    const clientStr = `${this.client?.prenom || ''} ${this.client?.nom || ''}`.trim()
    this.setFont(6, 'normal', C.white)
    this.doc.text(clientStr, C.pageW - C.marginR, 6.5, { align: 'right' })

    this.y = 14
  }

  drawPageNumber() {
    this.setFont(7, 'normal', C.muted)
    this.doc.text(`Page ${this.page}`, C.pageW / 2, C.pageH - 8, { align: 'center' })

    // Bas de page trait
    this.doc.setLineWidth(0.2)
    this.setDraw(C.border)
    this.doc.line(C.marginL, C.pageH - 12, C.pageW - C.marginR, C.pageH - 12)
  }

  // ── COUVERTURE ───────────────────────────────────────────────────────────

  drawCover() {
    const tc = this.tc

    // Fond haut page
    this.rect(0, 0, C.pageW, 80, [tc.r, tc.g, tc.b])

    // Texture diagonale subtile (lignes fines)
    this.doc.setLineWidth(0.1)
    this.doc.setDrawColor(255, 255, 255)
    this.doc.setGState(this.doc.GState({ opacity: 0.05 }))
    for (let i = -40; i < 250; i += 8) {
      this.doc.line(i, 0, i + 80, 80)
    }
    this.doc.setGState(this.doc.GState({ opacity: 1 }))

    // Logo texte en haut
    this.setFont(8, 'bold', C.white)
    this.text('TRAVAIL SOCIAL — SECTEUR PRIVÉ', C.marginL, 14)
    this.doc.setLineWidth(0.5)
    this.doc.setDrawColor(255, 255, 255)
    this.doc.setGState(this.doc.GState({ opacity: 0.4 }))
    this.doc.line(C.marginL, 16, C.pageW - C.marginR, 16)
    this.doc.setGState(this.doc.GState({ opacity: 1 }))

    // Titre principal
    this.setFont(22, 'bold', C.white)
    this.text(typeLabel(this.type).toUpperCase(), C.marginL, 38)

    // Sous-titre
    this.setFont(10, 'normal', [200, 230, 230])
    this.text(typeSubLabel(this.type), C.marginL, 48)

    // Nom Léopard
    if (this.eval.nom_leopard) {
      this.rect(C.marginL, 56, 80, 10, [0, 0, 0], null, 2)
      this.doc.setGState(this.doc.GState({ opacity: 0.4 }))
      this.rect(C.marginL, 56, 80, 10, [0, 0, 0], null, 2)
      this.doc.setGState(this.doc.GState({ opacity: 1 }))
      this.setFont(7, 'normal', [150, 220, 210])
      this.text('IDENTIFIANT', C.marginL + 4, 62)
      this.setFont(8, 'bold', [180, 240, 230])
      this.text(this.eval.nom_leopard, C.marginL + 26, 62)
    }

    // Zone CLIENT ─────────────────────────────────────────────
    this.y = 92
    const clientName = `${this.client?.prenom || ''} ${this.client?.nom || ''}`.trim()
    const dossier    = this.client?.no_dossier_leopard || '—'
    const dateNaiss  = this.client?.date_naissance ? formatDateLong(this.client.date_naissance) : '—'

    // Cadre client
    this.rect(C.marginL - 2, this.y - 6, C.contentW + 4, 44, C.mutedLight, C.border, 3)
    this.lineH(this.y - 6 + 14, [tc.r, tc.g, tc.b], 0.8)

    this.setFont(7, 'bold', [tc.r, tc.g, tc.b])
    this.text('DOSSIER CLIENT', C.marginL + 2, this.y + 2)

    this.setFont(16, 'bold', C.primary)
    this.text(clientName || '—', C.marginL + 2, this.y + 13)

    const col2 = C.pageW / 2 + 5
    this.setFont(7, 'bold', C.muted)
    this.text('N° LÉOPARD', C.marginL + 2, this.y + 22)
    this.text('DATE DE NAISSANCE', col2, this.y + 22)

    this.setFont(9, 'bold', C.primary)
    this.text(dossier, C.marginL + 2, this.y + 29)
    this.text(dateNaiss, col2, this.y + 29)

    this.y += 50

    // Zone ÉVALUATION ────────────────────────────────────────
    this.rect(C.marginL - 2, this.y - 6, C.contentW + 4, 36, C.mutedLight, C.border, 3)
    this.lineH(this.y - 6 + 14, [tc.r, tc.g, tc.b], 0.8)

    this.setFont(7, 'bold', [tc.r, tc.g, tc.b])
    this.text('ÉVALUATION', C.marginL + 2, this.y + 2)

    const createdAt   = formatDateLong(this.eval.created_at)
    const signedAt    = this.eval.date_signature ? formatDateTimeLong(this.eval.date_signature) : '—'
    const auteur      = this.eval.auteur_nom || this.eval.signature_nom || '—'

    this.setFont(7, 'bold', C.muted)
    this.text('DATE DE CRÉATION', C.marginL + 2, this.y + 14)
    this.text('ÉVALUATEUR', col2, this.y + 14)
    this.setFont(9, 'bold', C.primary)
    this.text(createdAt, C.marginL + 2, this.y + 20)
    this.text(auteur, col2, this.y + 20)

    this.y += 44

    // Objet / motif
    if (this.eval.objet_evaluation || this.eval.motif_reference) {
      const titre = this.eval.objet_evaluation || this.eval.motif_reference
      this.rect(C.marginL - 2, this.y - 4, C.contentW + 4, 22, [tc.r, tc.g, tc.b], null, 3)
      this.setFont(7, 'bold', [200, 230, 230])
      this.text('OBJET DE L\'ÉVALUATION', C.marginL + 2, this.y + 3)
      this.setFont(9, 'normal', C.white)
      const lines = this.doc.splitTextToSize(titre, C.contentW - 4)
      const display = lines.slice(0, 2)
      display.forEach((l, i) => this.text(l, C.marginL + 2, this.y + 9 + i * 5))
      this.y += 28
    }

    // Avertissement confidentialité
    this.y = C.pageH - 40
    this.rect(C.marginL - 2, this.y, C.contentW + 4, 20, [254, 243, 199], [253, 224, 71], 2)
    this.setFont(7, 'bold', [146, 64, 14])
    this.text('⚠  DOCUMENT CONFIDENTIEL — LOI 25 (QUÉBEC)', C.marginL + 2, this.y + 6)
    this.setFont(6.5, 'normal', [146, 64, 14])
    const avert = 'Ce document contient des renseignements personnels protégés. Sa reproduction, divulgation ou utilisation est strictement limitée aux personnes autorisées.'
    const avertLines = this.doc.splitTextToSize(avert, C.contentW - 4)
    avertLines.forEach((l, i) => this.text(l, C.marginL + 2, this.y + 11 + i * 4))

    // Signature sceau
    if (this.eval.verrouille && this.eval.signature_nom) {
      this.y = C.pageH - 16
      this.setFont(7, 'bold', C.success)
      this.text(`✓ Scellé par : ${this.eval.signature_nom}`, C.marginL, this.y)
      this.setFont(6.5, 'normal', C.muted)
      this.doc.text(`Le ${signedAt}`, C.pageW - C.marginR, this.y, { align: 'right' })
    }

    // Numéro page
    this.setFont(7, 'normal', C.muted)
    this.doc.text('1', C.pageW / 2, C.pageH - 5, { align: 'center' })

    // Page suivante
    this.addPage()
  }

  // ── SECTION HEADER ───────────────────────────────────────────────────────

  drawSectionHeader(number, label) {
    this.checkPageBreak(22)

    const tc = this.tc
    // Fond section
    this.rect(C.marginL - 2, this.y, C.contentW + 4, 14, [tc.r, tc.g, tc.b], null, 2)

    // Numéro
    this.rect(C.marginL - 2, this.y, 14, 14, [0, 0, 0], null, 2)
    this.doc.setGState(this.doc.GState({ opacity: 0.3 }))
    this.rect(C.marginL - 2, this.y, 14, 14, [0, 0, 0], null, 2)
    this.doc.setGState(this.doc.GState({ opacity: 1 }))
    this.setFont(9, 'bold', C.white)
    this.text(String(number), C.marginL + 4.5, this.y + 9.5, { align: 'center' })

    this.setFont(10, 'bold', C.white)
    this.text(label.toUpperCase(), C.marginL + 16, this.y + 9.5)

    this.y += 18
  }

  // ── BLOC CONTENU ─────────────────────────────────────────────────────────

  drawContentBlock(text) {
    if (!text?.trim()) return

    const lines   = this.doc.splitTextToSize(text, C.contentW - 6)
    const lineH   = 5
    const blockH  = lines.length * lineH + 10

    this.checkPageBreak(Math.min(blockH, 40))

    this.rect(C.marginL - 2, this.y, C.contentW + 4, blockH, C.mutedLight, C.border, 2)

    let localY = this.y + 7
    this.setFont(9, 'normal', C.primary)

    lines.forEach((line) => {
      if (localY + lineH > C.pageH - C.marginB - 6) {
        this.doc.addPage()
        this.page++
        localY = C.marginT
        this.drawRunningHeader()
        this.drawPageNumber()
        localY = this.y
      }
      this.doc.text(line, C.marginL + 2, localY)
      localY += lineH
    })

    this.y = localY + 5
  }

  // ── SIGNATURE FINALE ─────────────────────────────────────────────────────

  drawSignaturePage() {
    this.checkPageBreak(80)

    const tc = this.tc

    this.y += 10
    this.lineH(this.y, C.border, 0.5)
    this.y += 10

    // Titre
    this.setFont(11, 'bold', C.primary)
    this.text('ATTESTATION PROFESSIONNELLE', C.pageW / 2, this.y, { align: 'center' })
    this.y += 8

    // Texte attestation
    const attest = [
      'Je soussigné(e), travailleur(euse) social(e) membre de l\'Ordre des travailleurs sociaux et',
      'des thérapeutes conjugaux et familiaux du Québec (OTSTCFQ), atteste avoir procédé à',
      'l\'évaluation du fonctionnement social de la personne ci-haut mentionnée conformément',
      'aux standards de pratique professionnelle en vigueur.'
    ]
    this.setFont(9, 'normal', C.muted)
    attest.forEach(line => {
      this.text(line, C.pageW / 2, this.y, { align: 'center' })
      this.y += 5
    })

    this.y += 12

    // Bloc signature
    const sigW = (C.contentW - 10) / 2
    const col1 = C.marginL
    const col2 = C.marginL + sigW + 10

    // Signature gauche
    this.rect(col1, this.y, sigW, 28, C.mutedLight, C.border, 2)
    this.setFont(7, 'bold', C.muted)
    this.text('SIGNATURE DU PROFESSIONNEL', col1 + 4, this.y + 6)
    this.doc.setLineWidth(0.5)
    this.setDraw([tc.r, tc.g, tc.b])
    this.doc.line(col1 + 4, this.y + 22, col1 + sigW - 4, this.y + 22)

    if (this.eval.signature_nom) {
      this.setFont(9, 'bold', C.primary)
      this.text(this.eval.signature_nom.split(',')[0], col1 + 4, this.y + 20)
      this.setFont(7.5, 'normal', C.muted)
      const role = this.eval.signature_nom.split(',')[1]?.trim() || 'T.S.'
      this.text(role, col1 + 4, this.y + 26)
    }

    // Date droite
    this.rect(col2, this.y, sigW, 28, C.mutedLight, C.border, 2)
    this.setFont(7, 'bold', C.muted)
    this.text('DATE DE SIGNATURE', col2 + 4, this.y + 6)
    this.doc.line(col2 + 4, this.y + 22, col2 + sigW - 4, this.y + 22)

    const signedAt = this.eval.date_signature ? formatDateTimeLong(this.eval.date_signature) : formatDateLong(new Date())
    this.setFont(9, 'bold', C.primary)
    this.text(signedAt, col2 + 4, this.y + 20)

    this.y += 36

    // Badge scellé
    if (this.eval.verrouille) {
      this.rect(C.marginL, this.y, C.contentW, 14, [209, 250, 229], [167, 243, 208], 3)
      this.setFont(8, 'bold', [6, 95, 70])
      this.text('✓  Ce document est scellé numériquement et ne peut être modifié', C.pageW / 2, this.y + 9, { align: 'center' })
      this.y += 20
    }

    // Mention légale
    this.y += 5
    this.setFont(6.5, 'normal', C.muted)
    const legal = [
      'Ce document a été produit dans le cadre de l\'exercice de la profession de travailleur social régie par la Loi sur les travailleurs sociaux',
      'et les thérapeutes conjugaux et familiaux du Québec (RLRQ c T-11.1). Il est protégé par la Loi sur la protection des renseignements',
      'personnels dans le secteur privé (Loi 25). Toute reproduction, divulgation ou utilisation non autorisée est interdite.'
    ]
    legal.forEach(line => {
      this.doc.text(line, C.pageW / 2, this.y, { align: 'center' })
      this.y += 4
    })
  }

  // ── GÉNÉRATION COMPLÈTE ──────────────────────────────────────────────────

  generate() {
    // 1. Couverture
    this.drawCover()

    // 2. Corps du document — sections dans l'ordre
    const allSections = [
      { key: 'contexte_evaluation',       label: 'Contexte de l\'évaluation' },
      { key: 'motif_reference',           label: 'Motif de référence' },
      { key: 'objet_evaluation',          label: 'Objet de l\'évaluation' },
      { key: 'capacites_cognitives',      label: 'Capacités cognitives' },
      { key: 'inaptitude_constatee',      label: 'Inaptitude constatée' },
      { key: 'etat_sante_physique',       label: 'État de santé physique' },
      { key: 'dimensions_psycho_sociales',label: 'Dimensions psycho-sociales' },
      { key: 'roles_sociaux',             label: 'Rôles sociaux' },
      { key: 'evaluation_mandataire',     label: 'Évaluation du mandataire' },
      { key: 'reseau_social_soutien',     label: 'Réseau social et soutien' },
      { key: 'evolution_situation',       label: 'Évolution de la situation' },
      { key: 'objectifs_intervention',    label: 'Interventions réalisées' },
      { key: 'situation_globale',         label: 'Situation globale' },
      { key: 'analyse_clinique',          label: 'Analyse clinique' },
      { key: 'opinion_professionnelle',   label: 'Opinion professionnelle' },
      { key: 'recommandations',           label: 'Recommandations' },
      { key: 'plan_action',              label: 'Plan d\'action' }
    ]

    let sectionNum = 1
    allSections.forEach(s => {
      if (this.eval[s.key]?.trim()) {
        this.drawSectionHeader(sectionNum++, s.label)
        this.drawContentBlock(this.eval[s.key])
        this.y += 4
      }
    })

    // 3. Page de signature
    this.drawSignaturePage()

    // Numéro de page sur toutes les pages
    const totalPages = this.doc.getNumberOfPages()
    for (let i = 2; i <= totalPages; i++) {
      this.doc.setPage(i)
      this.setFont(7, 'normal', C.muted)
      this.doc.text(`${i} / ${totalPages}`, C.pageW / 2, C.pageH - 5, { align: 'center' })
    }

    return this.doc
  }
}

// ── Export principal ───────────────────────────────────────────────────────
/**
 * Génère et télécharge le PDF de l'évaluation
 * @param {Object} evalData  — données DÉCRYPTÉES retournées par le backend
 * @param {Object} client    — données du client
 */
export async function generateEvaluationPDF(evalData, client) {
  if (!evalData) throw new Error('Données d\'évaluation manquantes')

  const generator = new EvalPDFDocument(evalData, client)
  const doc       = generator.generate()

  // Nom de fichier
  const clientName = `${client?.nom || 'Client'}`.replace(/\s+/g, '_')
  const dateStr    = new Date().toISOString().slice(0, 10)
  const type       = (evalData.type_evaluation || 'evaluation').replace(/_/g, '-')
  const leopard    = evalData.nom_leopard ? `_${evalData.nom_leopard}` : ''
  const fileName   = `EFS_${clientName}_${type}${leopard}_${dateStr}.pdf`

  doc.save(fileName)
  return fileName
}