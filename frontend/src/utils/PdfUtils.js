import { noteTypesList, noteTemplates } from './noteTemplates'

// 1. CONFIGURATION DES COULEURS (La palette de ta clinique)
export const PDF_THEME = {
    primary: [13, 148, 136],   // Turquoise Tailwind
    textMain: [0, 0, 0],       // Noir profond
    textLight: [100, 100, 100], // Gris pour libellés
    border: [220, 220, 220],    // Gris clair pour lignes
    background: [249, 250, 251], // Gris très clair pour zones
    badges: {
        green: [21, 128, 61],
        blue: [29, 78, 216],
        red: [185, 28, 28],
        purple: [126, 34, 206],
        orange: [194, 65, 12]
    }
}

/**
 * CUISINE : Prépare le terrain pour chaque section
 */
export const PDFUtils = {
    
    // Ajoute un entête avec Badge de type
    drawTypeHeader(doc, typeNote, y) {
        const typeConfig = noteTypesList.find(t => t.value === typeNote) || { color: 'blue' }
        const template = noteTemplates[typeNote] || {}
        const rgb = PDF_THEME.badges[typeConfig.color] || PDF_THEME.primary
        
        // Le Badge (Fond)
        doc.setFillColor(...rgb)
        doc.roundedRect(15, y, 45, 8, 1, 1, 'F')
        
        // Texte du badge (Blanc)
        doc.setFont("helvetica", "bold")
        doc.setFontSize(8)
        doc.setTextColor(255, 255, 255)
        doc.text(typeNote.toUpperCase(), 37.5, y + 5.5, { align: 'center' })
        
        // Le Sujet (Titre du template)
        doc.setFontSize(14)
        doc.setTextColor(...PDF_THEME.textMain)
        doc.text(template.sujet || "Note d'évolution", 65, y + 6)
        
        return y + 15 // Retourne la nouvelle position Y
    },

    // Dessine une ligne de séparation élégante
    drawDivider(doc, y) {
        doc.setDrawColor(...PDF_THEME.border)
        doc.setLineWidth(0.2)
        doc.line(15, y, 195, y)
        return y + 10
    },

    // Cuisine le texte long avec détection de titres
    drawNoteContent(doc, content, startY) {
        const margin = 15
        const width = 180
        let currentY = startY
        const lines = content.split('\n')
        
        lines.forEach(line => {
            // Détection des titres (Ligne qui finit par ":")
            if (line.trim().endsWith(':')) {
                currentY += 4
                doc.setFont("helvetica", "bold")
                doc.setFontSize(10)
                doc.setTextColor(...PDF_THEME.primary)
                doc.text(line.trim(), margin, currentY)
                currentY += 6
            } 
            // Texte normal
            else if (line.trim() !== "") {
                doc.setFont("helvetica", "normal")
                doc.setFontSize(10)
                doc.setTextColor(...PDF_THEME.textMain)
                
                const wrappedLines = doc.splitTextToSize(line, width)
                doc.text(wrappedLines, margin, currentY)
                currentY += (wrappedLines.length * 5.5)
            }
            
            // Gestion automatique du saut de page
            if (currentY > 275) {
                doc.addPage()
                currentY = 20
            }
        })
        
        return currentY
    },

    // Pied de page professionnel (Loi 25)
    drawFooter(doc, pageNum, pageCount, clientNom) {
        const y = 285
        doc.setFont("helvetica", "italic")
        doc.setFontSize(8)
        doc.setTextColor(...PDF_THEME.textLight)
        
        doc.text(`Document confidentiel - Dossier de ${clientNom}`, 15, y)
        doc.text(`Page ${pageNum} / ${pageCount}`, 195, y, { align: 'right' })
    },

    /**
     * 🖼️ AJOUTE UNE IMAGE dans le PDF
     * @param {Object} doc - Le document jsPDF
     * @param {string} imageData - Base64 de l'image (ex: depuis un input file)
     * @param {number} positionY - Où placer l'image verticalement
     * @param {Object} options - Personnalisation {largeur: 80, centrer: true, legende: "Photo avant"}
     * @returns {number} Nouvelle position Y après l'image
     */
    ajouterImage(doc, imageData, positionY, options = {}) {
        const config = {
            largeur: options.largeur || 80,
            centrer: options.centrer !== false, // Par défaut centré
            legende: options.legende || null,
            bordure: options.bordure || false
        }

        // Calcul de la hauteur proportionnelle
        const img = new Image()
        img.src = imageData
        const ratio = img.height / img.width
        const hauteur = config.largeur * ratio

        // Position X (centré ou marge gauche)
        const positionX = config.centrer 
            ? (210 - config.largeur) / 2  // Centré sur page A4
            : 15 // Marge gauche

        // Vérification saut de page
        if (positionY + hauteur > 270) {
            doc.addPage()
            positionY = 20
        }

        // Bordure optionnelle
        if (config.bordure) {
            doc.setDrawColor(...PDF_THEME.border)
            doc.rect(positionX - 1, positionY - 1, config.largeur + 2, hauteur + 2)
        }

        // Insertion de l'image
        doc.addImage(imageData, 'JPEG', positionX, positionY, config.largeur, hauteur)
        
        let nouvelleY = positionY + hauteur + 5

        // Légende optionnelle
        if (config.legende) {
            doc.setFont("helvetica", "italic")
            doc.setFontSize(9)
            doc.setTextColor(...PDF_THEME.textLight)
            const legendeX = config.centrer ? 105 : positionX + (config.largeur / 2)
            doc.text(config.legende, legendeX, nouvelleY, { align: 'center' })
            nouvelleY += 8
        }

        return nouvelleY
    },

    /**
     * 📊 CRÉER UN TABLEAU élégant
     * @param {Object} doc - Le document jsPDF
     * @param {Array} colonnes - Ex: ['Date', 'Symptôme', 'Intensité']
     * @param {Array} donnees - Ex: [['2024-01-15', 'Douleur', '7/10'], ...]
     * @param {number} positionY - Position de départ
     * @param {Object} options - {titre: "Historique", largeurColonnes: [40, 80, 40]}
     * @returns {number} Nouvelle position Y
     */
    creerTableau(doc, colonnes, donnees, positionY, options = {}) {
        const config = {
            titre: options.titre || null,
            largeurColonnes: options.largeurColonnes || null,
            couleurEntete: options.couleurEntete || PDF_THEME.primary
        }

        let currentY = positionY

        // Titre du tableau
        if (config.titre) {
            doc.setFont("helvetica", "bold")
            doc.setFontSize(11)
            doc.setTextColor(...PDF_THEME.textMain)
            doc.text(config.titre, 15, currentY)
            currentY += 8
        }

        // Calcul auto des largeurs si non spécifié
        const largeurs = config.largeurColonnes || 
            colonnes.map(() => 180 / colonnes.length)

        const margeGauche = 15
        const hauteurLigne = 8

        // Entête du tableau
        doc.setFillColor(...config.couleurEntete)
        doc.rect(margeGauche, currentY, 180, hauteurLigne, 'F')
        
        doc.setFont("helvetica", "bold")
        doc.setFontSize(9)
        doc.setTextColor(255, 255, 255)
        
        let xPosition = margeGauche + 2
        colonnes.forEach((colonne, index) => {
            doc.text(colonne, xPosition, currentY + 5.5)
            xPosition += largeurs[index]
        })
        
        currentY += hauteurLigne

        // Lignes de données
        doc.setFont("helvetica", "normal")
        doc.setTextColor(...PDF_THEME.textMain)
        
        donnees.forEach((ligne, indexLigne) => {
            // Alternance de fond
            if (indexLigne % 2 === 0) {
                doc.setFillColor(...PDF_THEME.background)
                doc.rect(margeGauche, currentY, 180, hauteurLigne, 'F')
            }

            xPosition = margeGauche + 2
            ligne.forEach((cellule, index) => {
                const texteCoupe = doc.splitTextToSize(cellule.toString(), largeurs[index] - 4)
                doc.text(texteCoupe[0], xPosition, currentY + 5.5)
                xPosition += largeurs[index]
            })
            
            currentY += hauteurLigne

            // Saut de page si nécessaire
            if (currentY > 270) {
                doc.addPage()
                currentY = 20
            }
        })

        // Bordure finale
        doc.setDrawColor(...PDF_THEME.border)
        doc.rect(margeGauche, positionY + (config.titre ? 8 : 0), 180, 
                 (donnees.length + 1) * hauteurLigne)

        return currentY + 10
    },

    /**
     * 💬 ENCADRÉ SPÉCIAL (Citation, alerte, conseil)
     * @param {Object} doc - Le document jsPDF
     * @param {string} texte - Contenu de l'encadré
     * @param {number} positionY - Position verticale
     * @param {Object} options - {type: 'info'|'alerte'|'conseil', icone: '⚠️'}
     * @returns {number} Nouvelle position Y
     */
    ajouterEncadre(doc, texte, positionY, options = {}) {
        const types = {
            info: { couleur: PDF_THEME.badges.blue, icone: 'ℹ️' },
            alerte: { couleur: PDF_THEME.badges.red, icone: '⚠️' },
            conseil: { couleur: PDF_THEME.badges.green, icone: '💡' },
            citation: { couleur: PDF_THEME.primary, icone: '📝' }
        }

        const config = types[options.type] || types.info
        const icone = options.icone || config.icone

        // Fond coloré léger
        const couleurFond = config.couleur.map(v => v + (255 - v) * 0.9)
        doc.setFillColor(...couleurFond)
        
        // Calcul hauteur nécessaire
        const lignes = doc.splitTextToSize(texte, 160)
        const hauteur = (lignes.length * 5) + 10

        // Vérification saut de page
        if (positionY + hauteur > 270) {
            doc.addPage()
            positionY = 20
        }

        // Dessin du fond
        doc.roundedRect(15, positionY, 180, hauteur, 2, 2, 'F')
        
        // Bordure gauche colorée
        doc.setFillColor(...config.couleur)
        doc.rect(15, positionY, 4, hauteur, 'F')

        // Icône et texte
        doc.setFontSize(12)
        doc.text(icone, 22, positionY + 7)
        
        doc.setFont("helvetica", "normal")
        doc.setFontSize(9)
        doc.setTextColor(...PDF_THEME.textMain)
        doc.text(lignes, 32, positionY + 7)

        return positionY + hauteur + 8
    },

    /**
     * 📋 LISTE À PUCES stylisée
     * @param {Object} doc - Le document jsPDF
     * @param {Array} elements - Ex: ['Item 1', 'Item 2', 'Item 3']
     * @param {number} positionY - Position de départ
     * @param {Object} options - {titre: "Objectifs", puce: '✓', indentation: 20}
     * @returns {number} Nouvelle position Y
     */
    ajouterListe(doc, elements, positionY, options = {}) {
        const config = {
            titre: options.titre || null,
            puce: options.puce || '•',
            indentation: options.indentation || 20,
            espacement: options.espacement || 6
        }

        let currentY = positionY

        // Titre de la liste
        if (config.titre) {
            doc.setFont("helvetica", "bold")
            doc.setFontSize(11)
            doc.setTextColor(...PDF_THEME.primary)
            doc.text(config.titre, 15, currentY)
            currentY += 8
        }

        // Éléments
        doc.setFont("helvetica", "normal")
        doc.setFontSize(10)
        doc.setTextColor(...PDF_THEME.textMain)

        elements.forEach(element => {
            // Puce
            doc.setTextColor(...PDF_THEME.primary)
            doc.text(config.puce, 15, currentY)
            
            // Texte avec wrapping
            doc.setTextColor(...PDF_THEME.textMain)
            const lignes = doc.splitTextToSize(element, 180 - config.indentation)
            doc.text(lignes, 15 + config.indentation, currentY)
            
            currentY += (lignes.length * 5) + config.espacement

            // Saut de page
            if (currentY > 270) {
                doc.addPage()
                currentY = 20
            }
        })

        return currentY + 5
    },

    /**
     * 🔤 SIGNATURE NUMÉRIQUE zone
     * @param {Object} doc - Le document jsPDF
     * @param {number} positionY - Position verticale
     * @param {Object} options - {nom: "Dr. Lavoie", titre: "Psychologue", date: true}
     * @returns {number} Nouvelle position Y
     */
    ajouterSignature(doc, positionY, options = {}) {
        const config = {
            nom: options.nom || "Professionnel de la santé",
            titre: options.titre || "",
            date: options.date !== false,
            imageSignature: options.imageSignature || null
        }

        let currentY = positionY + 5

        // Zone de signature
        doc.setDrawColor(...PDF_THEME.border)
        doc.line(15, currentY, 100, currentY)
        
        currentY += 5

        // Image de signature si fournie
        if (config.imageSignature) {
            doc.addImage(config.imageSignature, 'PNG', 15, currentY - 20, 40, 15)
        }

        // Nom du professionnel
        doc.setFont("helvetica", "bold")
        doc.setFontSize(10)
        doc.setTextColor(...PDF_THEME.textMain)
        doc.text(config.nom, 15, currentY)
        
        currentY += 5

        // Titre professionnel
        if (config.titre) {
            doc.setFont("helvetica", "italic")
            doc.setFontSize(9)
            doc.setTextColor(...PDF_THEME.textLight)
            doc.text(config.titre, 15, currentY)
            currentY += 5
        }

        // Date automatique
        if (config.date) {
            const dateActuelle = new Date().toLocaleDateString('fr-CA')
            doc.setFont("helvetica", "normal")
            doc.setFontSize(9)
            doc.text(`Signé le ${dateActuelle}`, 15, currentY)
        }

        return currentY + 10
    },

    /**
     * 📏 BARRE DE PROGRESSION (ex: pour échelles d'évaluation)
     * @param {Object} doc - Le document jsPDF
     * @param {number} valeur - Valeur actuelle (ex: 7)
     * @param {number} maximum - Valeur max (ex: 10)
     * @param {number} positionY - Position verticale
     * @param {Object} options - {label: "Douleur", couleur: 'red'}
     * @returns {number} Nouvelle position Y
     */
    ajouterBarreProgression(doc, valeur, maximum, positionY, options = {}) {
        const config = {
            label: options.label || "",
            largeur: options.largeur || 120,
            hauteur: options.hauteur || 6,
            couleur: options.couleur || 'blue'
        }

        const pourcentage = (valeur / maximum) * 100
        const couleurBarre = PDF_THEME.badges[config.couleur] || PDF_THEME.primary

        // Label
        if (config.label) {
            doc.setFont("helvetica", "normal")
            doc.setFontSize(9)
            doc.setTextColor(...PDF_THEME.textMain)
            doc.text(config.label, 15, positionY)
        }

        const yBarre = positionY + (config.label ? 5 : 0)

        // Fond de la barre
        doc.setFillColor(...PDF_THEME.background)
        doc.roundedRect(15, yBarre, config.largeur, config.hauteur, 1, 1, 'F')

        // Barre de progression
        const largeurRemplie = (config.largeur * pourcentage) / 100
        doc.setFillColor(...couleurBarre)
        doc.roundedRect(15, yBarre, largeurRemplie, config.hauteur, 1, 1, 'F')

        // Valeur affichée
        doc.setFont("helvetica", "bold")
        doc.setFontSize(9)
        doc.setTextColor(...couleurBarre)
        doc.text(`${valeur}/${maximum}`, 15 + config.largeur + 5, yBarre + 4.5)

        return yBarre + config.hauteur + 10
    }
}