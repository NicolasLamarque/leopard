package database

import (
	"fmt"
	schema "leopard/internal/db/schemas" // Import de ton sous-dossier

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Database struct {
	*sqlx.DB
}

func New(path string) (*Database, error) {
	db, err := sqlx.Connect("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("impossible de se connecter à SQLite: %w", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("impossible d'activer les foreign_keys: %w", err)
	}

	// On assemble les morceaux de tes 13 fichiers ici.
	// L'ordre est important pour les clés étrangères (Users et Clients en premier).
	// 1. Les fondations (Celles qui ne dépendent de personne)
	fullSchema := schema.System + // Municipalités, pays
		schema.Users + // Utilisateurs
		schema.Logs + // Journal d'audit (Loi 25) - Nécessite Users
		schema.TableMedecins + // Ressources
		schema.TableNotaires +
		schema.TableResidence +
		schema.TableCHSLD +
		schema.TableRi +
		schema.TableClients + // Pivot - Nécessite Users, Medecins, Résidences
		schema.TableContacts + // Nécessite Clients
		schema.TableNotes + // Nécessite Clients, Users
		schema.TableEval + // Nécessite Clients (Vue incluse)
		schema.TablePI + // Nécessite Clients (Vue incluse)
		schema.TableAppels + // Nécessite Clients, Users
		schema.TableIntervenants // Nécessite Clients (Vue incluse)

	_, err = db.Exec(fullSchema)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du schéma: %w", err)
	}

	return &Database{db}, nil
}
