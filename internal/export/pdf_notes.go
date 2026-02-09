package export

import (
	"fmt"
	models "leopard/internal/model"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// GenerateNotesPDF génère un PDF simple avec les notes
func GenerateNotesPDF(notes []models.Note, leopardNumber string, outputDir string) (string, error) {
	if len(notes) == 0 {
		return "", fmt.Errorf("aucune note a exporter")
	}

	// Créer le dossier
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("impossible de creer le dossier: %w", err)
	}

	// Nom du fichier
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("notes_export_%s_%s.pdf", leopardNumber, timestamp)
	outputPath := filepath.Join(outputDir, filename)

	// Créer le PDF - config simple
	m := maroto.New(config.NewBuilder().Build())

	// Pour chaque note, ajouter du contenu
	for i, note := range notes {
		if i > 0 {
			// Saut de page entre notes
			m.AddRows(text.NewRow(50, "", props.Text{}))
		}
		addNoteRows(m, note, i+1, len(notes))
	}

	// Générer et sauvegarder
	doc, err := m.Generate()
	if err != nil {
		return "", fmt.Errorf("erreur generation: %w", err)
	}

	err = doc.Save(outputPath)
	if err != nil {
		return "", fmt.Errorf("erreur sauvegarde: %w", err)
	}

	return outputPath, nil
}

// addNoteRows ajoute les rows d'une note avec layout professionnel
func addNoteRows(m core.Maroto, note models.Note, pageNum, totalPages int) {
	// EN-TETE: LOGO + INFO CLIENT
	m.AddRows(
		row.New(35).Add(
			// LOGO à gauche
			image.NewFromFileCol(3, "frontend/src/assets/images/logoBlanc.png", props.Rect{
				Center:  false,
				Percent: 90,
				Top:     2,
			}),
			// Espace
			col.New(1),
			// INFO CLIENT à droite
			text.NewCol(8, buildClientInfoBox(note), props.Text{
				Size: 9,
				Top:  2,
				Left: 2,
			}),
		),
	)

	// Ligne de séparation
	m.AddRows(row.New(1))
	m.AddRows(line.NewRow(1, props.Line{
		Color:     &props.Color{Red: 100, Green: 100, Blue: 100},
		Thickness: 0.8,
	}))
	m.AddRows(row.New(3))

	// TITRE
	titre := safeString(note.Titre)
	if titre == "" {
		titre = "[Sans titre]"
	}

	m.AddRows(text.NewRow(12, titre, props.Text{
		Size:  16,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))

	// Ligne sous le titre
	m.AddRows(line.NewRow(1, props.Line{
		Color:     &props.Color{Red: 200, Green: 200, Blue: 200},
		Thickness: 0.5,
	}))
	m.AddRows(row.New(5))

	// CONTENU
	m.AddRows(text.NewRow(6, "CONTENU:", props.Text{
		Size:  10,
		Style: fontstyle.Bold,
	}))

	contenu := safeString(note.Contenu)
	if contenu == "" {
		contenu = "[Aucun contenu]"
	}

	// Paragraphes
	paragraphs := strings.Split(contenu, "\n")
	for _, p := range paragraphs {
		p = strings.TrimSpace(p)
		if p == "" {
			m.AddRows(row.New(2))
			continue
		}
		m.AddRows(text.NewRow(0, p, props.Text{Size: 10}))
	}

	m.AddRows(row.New(10))

	// ═══════════ SIGNATURE ET VALIDATION ═══════════
	if note.Verrouille == 1 {
		// Ligne de séparation
		m.AddRows(line.NewRow(1, props.Line{
			Color:     &props.Color{Red: 100, Green: 100, Blue: 100},
			Thickness: 0.5,
		}))
		m.AddRows(row.New(2))

		// Titre section
		m.AddRows(text.NewRow(6, "SIGNATURE ET VALIDATION", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
		}))

		m.AddRows(row.New(2))

		// Info signature
		sigNom := safeString(note.SignatureNom)
		if sigNom == "" {
			sigNom = "[Non signe]"
		}

		m.AddRows(text.NewRow(5, fmt.Sprintf("Note signee par: %s", sigNom), props.Text{
			Size: 9,
			Left: 2,
		}))

		sigDate := formatDatePtr(note.SignatureDate)
		if sigDate != "" {
			m.AddRows(text.NewRow(5, fmt.Sprintf("Date de signature: %s", sigDate), props.Text{
				Size: 9,
				Left: 2,
			}))
		}

		m.AddRows(text.NewRow(5, "Statut: [Note verrouillee et signee]", props.Text{
			Size: 9,
			Left: 2,
		}))

		m.AddRows(row.New(3))
	}

	// ═══════════ INFORMATIONS DE LA NOTE ═══════════
	// Ligne de séparation
	m.AddRows(line.NewRow(1, props.Line{
		Color:     &props.Color{Red: 100, Green: 100, Blue: 100},
		Thickness: 0.5,
	}))
	m.AddRows(row.New(2))

	// Titre section
	m.AddRows(text.NewRow(6, "INFORMATIONS DE LA NOTE", props.Text{
		Size:  10,
		Style: fontstyle.Bold,
	}))

	m.AddRows(row.New(2))

	// Métadonnées (une par ligne)
	metaLines := buildNoteMetadataLines(note)
	for _, line := range metaLines {
		m.AddRows(text.NewRow(4, line, props.Text{
			Size: 8,
			Left: 2,
		}))
	}

	m.AddRows(row.New(3))

	// ═══════════ FOOTER ═══════════
	// Ligne de séparation
	m.AddRows(line.NewRow(1, props.Line{
		Color:     &props.Color{Red: 100, Green: 100, Blue: 100},
		Thickness: 0.5,
	}))
	m.AddRows(row.New(2))

	// Date génération
	m.AddRows(text.NewRow(4, fmt.Sprintf("Document genere le %s", time.Now().Format("2006-01-02 a 15:04")), props.Text{
		Size:  7,
		Align: align.Center,
	}))

	// Pagination
	m.AddRows(text.NewRow(4, fmt.Sprintf("Page %d sur %d", pageNum, totalPages), props.Text{
		Size:  7,
		Align: align.Center,
	}))
}

// buildClientInfoBox construit l'encadré d'info client
func buildClientInfoBox(note models.Note) string {
	lines := []string{}

	// No Léopard
	leopard := safeString(note.ClientNoLeopard)
	if leopard != "" {
		lines = append(lines, fmt.Sprintf("No. Dossier: %s", leopard))
	}

	// Nom, Prénom
	nom := safeString(note.ClientNom)
	prenom := safeString(note.ClientPrenom)
	if nom != "" && prenom != "" {
		lines = append(lines, fmt.Sprintf("Nom: %s, %s", nom, prenom))
	} else if nom != "" || prenom != "" {
		lines = append(lines, fmt.Sprintf("Nom: %s%s", nom, prenom))
	}

	// RAMQ
	ramq := safeString(note.ClientNoRAMQ)
	if ramq != "" {
		lines = append(lines, fmt.Sprintf("No RAMQ: %s", ramq))
	}

	// Adresse complète
	adresse := buildAdresseComplete(note)
	if adresse != "" {
		lines = append(lines, fmt.Sprintf("Adresse: %s", adresse))
	}

	// Téléphones
	tel := safeString(note.ClientTelephone)
	cell := safeString(note.ClientCellulaire)
	if tel != "" && cell != "" {
		lines = append(lines, fmt.Sprintf("Tel: %s | Cell: %s", tel, cell))
	} else if tel != "" {
		lines = append(lines, fmt.Sprintf("Tel: %s", tel))
	} else if cell != "" {
		lines = append(lines, fmt.Sprintf("Cell: %s", cell))
	}

	if len(lines) == 0 {
		return "[Client non identifie]"
	}

	return strings.Join(lines, "\n")
}

// buildAdresseComplete construit l'adresse complète
func buildAdresseComplete(note models.Note) string {
	parts := []string{}

	addr := safeString(note.ClientAdresse)
	app := safeString(note.ClientAppartement)

	if addr != "" {
		if app != "" {
			parts = append(parts, fmt.Sprintf("%s, app. %s", addr, app))
		} else {
			parts = append(parts, addr)
		}
	}

	ville := safeString(note.ClientVille)
	province := safeString(note.ClientProvince)
	cp := safeString(note.ClientCodePostal)

	if ville != "" {
		cityPart := ville
		if province != "" {
			cityPart += ", " + province
		}
		if cp != "" {
			cityPart += " " + cp
		}
		parts = append(parts, cityPart)
	}

	return strings.Join(parts, ", ")
}

// buildNoteMetadataLines construit les lignes de métadonnées (une par ligne)
func buildNoteMetadataLines(note models.Note) []string {
	lines := []string{}

	// Type de note
	typeNote := safeString(note.TypeNote)
	if typeNote != "" {
		lines = append(lines, fmt.Sprintf("Type: %s", typeNote))
	}

	// Date intervention
	dateIntervention := formatDatePtr(note.DateIntervention)
	if dateIntervention != "" {
		lines = append(lines, fmt.Sprintf("Date intervention: %s", dateIntervention))
	}

	// Heure et durée
	heure := safeString(note.HeureIntervention)
	duree := safeString(note.DureeIntervention)
	if heure != "" && duree != "" {
		lines = append(lines, fmt.Sprintf("Heure: %s (Duree: %s)", heure, duree))
	} else if heure != "" {
		lines = append(lines, fmt.Sprintf("Heure: %s", heure))
	}

	// Mode
	mode := safeString(note.ModeIntervention)
	if mode != "" {
		lines = append(lines, fmt.Sprintf("Mode: %s", mode))
	}

	// Type intervention
	typeIntervention := safeString(note.TypeIntervention)
	if typeIntervention != "" {
		lines = append(lines, fmt.Sprintf("Type intervention: %s", typeIntervention))
	}

	// Indicateurs
	if note.NoteTardive == 1 {
		lines = append(lines, "[Note tardive]")
	}
	if note.NoteDeTier == 1 {
		lines = append(lines, "[Note de tiers]")
	}

	typeLien := safeString(note.TypeLien)
	if typeLien != "" {
		lines = append(lines, fmt.Sprintf("[%s]", typeLien))
	}

	if len(lines) == 0 {
		lines = append(lines, fmt.Sprintf("Note creee le %s", formatDate(note.DateNote)))
	}

	return lines
}

// Helpers

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

func formatDatePtr(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}
