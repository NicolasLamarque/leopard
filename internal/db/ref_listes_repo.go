package database

import (
	"fmt"

	models "leopard/internal/model"
	"leopard/internal/services/importer"
	"strings"
)

// ─────────────────────────────────────────────────────────────
// LECTURE
// ─────────────────────────────────────────────────────────────

// GetRefListeByCategorie retourne toutes les entrées actives d'une catégorie,
// triées par ordre. C'est la fonction principale pour alimenter les dropdowns Vue.
// Exemple : db.GetRefListeByCategorie("type_intervenant")
func (db *Database) GetRefListeByCategorie(categorie string) ([]models.RefListe, error) {
	var items []models.RefListe
	err := db.Select(&items, `
		SELECT id, categorie, libelle, couleur, icone, is_systeme, ordre, actif
		FROM ref_listes
		WHERE categorie = ? AND actif = 1
		ORDER BY ordre ASC, libelle ASC
	`, categorie)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération liste '%s': %w", categorie, err)
	}
	return items, nil
}

// GetAllCategories retourne la liste unique de toutes les catégories existantes.
// Utilisé pour la page d'administration des listes.
func (db *Database) GetAllRefCategories() ([]string, error) {
	var categories []string
	err := db.Select(&categories, `
		SELECT DISTINCT categorie
		FROM ref_listes
		WHERE actif = 1
		ORDER BY categorie ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération catégories: %w", err)
	}
	return categories, nil
}

// GetRefListeByID retourne une entrée spécifique par son ID.
func (db *Database) GetRefListeByID(id int) (*models.RefListe, error) {
	var item models.RefListe
	err := db.Get(&item, `
		SELECT id, categorie, libelle, couleur, icone, is_systeme, ordre, actif
		FROM ref_listes
		WHERE id = ?
	`, id)
	if err != nil {
		return nil, fmt.Errorf("ref_liste ID %d non trouvée: %w", id, err)
	}
	return &item, nil
}

// ─────────────────────────────────────────────────────────────
// CRÉATION
// ─────────────────────────────────────────────────────────────

// CreateRefListe ajoute une nouvelle entrée utilisateur (is_systeme = 0).
// Les entrées système sont insérées uniquement via le seed SQL.
func (db *Database) CreateRefListe(req models.CreateRefListeRequest) (int64, error) {
	// Calcul automatique de l'ordre si non fourni
	if req.Ordre == 0 {
		var maxOrdre int
		err := db.Get(&maxOrdre, `
			SELECT COALESCE(MAX(ordre), 0) FROM ref_listes WHERE categorie = ?
		`, req.Categorie)
		if err != nil {
			return 0, fmt.Errorf("erreur calcul ordre: %w", err)
		}
		req.Ordre = maxOrdre + 1
	}

	// Valeurs par défaut
	if req.Couleur == "" {
		req.Couleur = "slate"
	}

	resultat, err := db.NamedExec(`
		INSERT INTO ref_listes (categorie, libelle, couleur, icone, is_systeme, ordre, actif)
		VALUES (:categorie, :libelle, :couleur, :icone, 0, :ordre, 1)
	`, map[string]interface{}{
		"categorie": req.Categorie,
		"libelle":   req.Libelle,
		"couleur":   req.Couleur,
		"icone":     req.Icone,
		"ordre":     req.Ordre,
	})
	if err != nil {
		return 0, fmt.Errorf("erreur création ref_liste: %w", err)
	}
	return resultat.LastInsertId()
}

// ─────────────────────────────────────────────────────────────
// MODIFICATION
// ─────────────────────────────────────────────────────────────

// UpdateRefListe modifie le libellé, la couleur et l'icône d'une entrée.
// Protège les entrées système (is_systeme = 1).
func (db *Database) UpdateRefListe(id int, req models.CreateRefListeRequest) error {
	// Vérification : on ne touche pas aux entrées système
	var isSysteme bool
	err := db.Get(&isSysteme, `SELECT is_systeme FROM ref_listes WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("ref_liste ID %d non trouvée: %w", id, err)
	}
	if isSysteme {
		return fmt.Errorf("impossible de modifier une entrée système (id: %d)", id)
	}

	_, err = db.Exec(`
		UPDATE ref_listes
		SET libelle = ?, couleur = ?, icone = ?
		WHERE id = ?
	`, req.Libelle, req.Couleur, req.Icone, id)
	if err != nil {
		return fmt.Errorf("erreur mise à jour ref_liste: %w", err)
	}
	return nil
}

// UpdateRefListeOrdre met à jour l'ordre d'affichage d'une entrée.
func (db *Database) UpdateRefListeOrdre(id int, ordre int) error {
	_, err := db.Exec(`
		UPDATE ref_listes SET ordre = ? WHERE id = ?
	`, ordre, id)
	if err != nil {
		return fmt.Errorf("erreur mise à jour ordre ref_liste: %w", err)
	}
	return nil
}

// ─────────────────────────────────────────────────────────────
// SUPPRESSION
// ─────────────────────────────────────────────────────────────

// DeleteRefListe supprime une entrée utilisateur.
// Les entrées système (is_systeme = 1) sont protégées.
func (db *Database) DeleteRefListe(id int) error {
	// Vérification : entrée système intouchable
	var isSysteme bool
	err := db.Get(&isSysteme, `SELECT is_systeme FROM ref_listes WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("ref_liste ID %d non trouvée: %w", id, err)
	}
	if isSysteme {
		return fmt.Errorf("impossible de supprimer une entrée système (id: %d)", id)
	}

	_, err = db.Exec(`DELETE FROM ref_listes WHERE id = ? AND is_systeme = 0`, id)
	if err != nil {
		return fmt.Errorf("erreur suppression ref_liste: %w", err)
	}
	return nil
}

// SaveRefListeList insère une liste de RefListeRow issus de l'import CSV.
// Retourne le nombre d'entrées importées et le nombre de doublons ignorés.
func (db *Database) SaveRefListeList(items []importer.RefListeRow, createdBy int) (int, int, error) {
	imported := 0
	ignored := 0

	for _, item := range items {
		req := models.CreateRefListeRequest{
			Categorie: item.Categorie,
			Libelle:   item.Libelle,
			Couleur:   item.Couleur,
			Icone:     item.Icone,
			Ordre:     item.Ordre,
		}

		_, err := db.CreateRefListe(req)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				ignored++
			} else {
				return imported, ignored, fmt.Errorf("erreur sauvegarde '%s/%s': %w", item.Categorie, item.Libelle, err)
			}
			continue
		}
		imported++
	}

	return imported, ignored, nil
}
