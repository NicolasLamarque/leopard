package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	// Connexion à la base de données
	dbPath := "../app.db" // Ajuste le chemin si nécessaire
	db, err := sqlx.Connect("sqlite", dbPath)
	if err != nil {
		log.Fatalf("❌ Erreur de connexion: %v", err)
	}
	defer db.Close()

	fmt.Println("📦 Connexion à la base de données OK")

	// Activation des foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatalf("❌ Erreur PRAGMA: %v", err)
	}

	// Récupérer l'ID d'un user existant pour created_by
	var userID int
	err = db.Get(&userID, "SELECT id FROM users LIMIT 1")
	if err != nil {
		log.Fatalf("❌ Aucun utilisateur trouvé. Crée d'abord un user avec init_user.go")
	}

	fmt.Printf("👤 Utilisation de l'user ID %d pour created_by\n", userID)

	// Liste de clients fictifs
	clients := []map[string]interface{}{
		{
			"nom":                        "Tremblay",
			"prenom":                     "Jean",
			"date_naissance":             "1945-03-15",
			"telephone":                  "819-555-0101",
			"cellulaire":                 "819-555-0102",
			"email":                      "jean.tremblay@example.com",
			"adresse":                    "123 Rue Principale",
			"code_postal":                "J8V 1A1",
			"ville":                      "Gatineau",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "TREJ 4503 1501",
			"numero_securite_sociale":    "123-456-789",
			"no_hcm":                     "HCM001",
			"no_chaur":                   "CH001",
			"no_dossier_leopard":         "LEO-2024-001",
			"medecin_famille_No_Licence": "12345",
			"note_fixe":                  "Client prioritaire - Suivi hebdomadaire",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Gagnon",
			"prenom":                     "Marie",
			"date_naissance":             "1952-07-22",
			"telephone":                  "819-555-0201",
			"cellulaire":                 "819-555-0202",
			"email":                      "marie.gagnon@example.com",
			"adresse":                    "456 Avenue des Érables",
			"code_postal":                "J8V 2B2",
			"ville":                      "Gatineau",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "GAGM 5207 2201",
			"numero_securite_sociale":    "234-567-890",
			"no_hcm":                     "HCM002",
			"no_chaur":                   "CH002",
			"no_dossier_leopard":         "LEO-2024-002",
			"medecin_famille_No_Licence": "12346",
			"note_fixe":                  "Allergie aux arachides",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Roy",
			"prenom":                     "Pierre",
			"date_naissance":             "1938-11-10",
			"telephone":                  "819-555-0301",
			"cellulaire":                 "",
			"email":                      "",
			"adresse":                    "789 Boulevard Saint-Joseph",
			"code_postal":                "J8V 3C3",
			"ville":                      "Gatineau",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "ROYP 3811 1001",
			"numero_securite_sociale":    "345-678-901",
			"no_hcm":                     "HCM003",
			"no_chaur":                   "CH003",
			"no_dossier_leopard":         "LEO-2024-003",
			"medecin_famille_No_Licence": "12347",
			"note_fixe":                  "Suivi mensuel - Mobilité réduite",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Côté",
			"prenom":                     "Lucie",
			"date_naissance":             "1950-04-28",
			"telephone":                  "819-555-0401",
			"cellulaire":                 "819-555-0402",
			"email":                      "lucie.cote@example.com",
			"adresse":                    "321 Rue du Parc",
			"code_postal":                "J8V 4D4",
			"ville":                      "Cantley",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "COTL 5004 2801",
			"numero_securite_sociale":    "456-789-012",
			"no_hcm":                     "HCM004",
			"no_chaur":                   "CH004",
			"no_dossier_leopard":         "LEO-2024-004",
			"medecin_famille_No_Licence": "12348",
			"note_fixe":                  "Famille très impliquée",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Bouchard",
			"prenom":                     "Robert",
			"date_naissance":             "1942-09-05",
			"telephone":                  "819-555-0501",
			"cellulaire":                 "",
			"email":                      "",
			"adresse":                    "654 Chemin de la Montagne",
			"code_postal":                "J8V 5E5",
			"ville":                      "Chelsea",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "BOUR 4209 0501",
			"numero_securite_sociale":    "567-890-123",
			"no_hcm":                     "HCM005",
			"no_chaur":                   "CH005",
			"no_dossier_leopard":         "LEO-2024-005",
			"medecin_famille_No_Licence": "12349",
			"note_fixe":                  "Rendez-vous bilingue (FR/EN)",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Lavoie",
			"prenom":                     "Suzanne",
			"date_naissance":             "1955-12-18",
			"telephone":                  "819-555-0601",
			"cellulaire":                 "819-555-0602",
			"email":                      "suzanne.lavoie@example.com",
			"adresse":                    "987 Rue des Pins",
			"code_postal":                "J8V 6F6",
			"ville":                      "Gatineau",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "LAVS 5512 1801",
			"numero_securite_sociale":    "678-901-234",
			"no_hcm":                     "HCM006",
			"no_chaur":                   "CH006",
			"no_dossier_leopard":         "LEO-2024-006",
			"medecin_famille_No_Licence": "12350",
			"note_fixe":                  "Préfère les rendez-vous matinaux",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Bergeron",
			"prenom":                     "Claude",
			"date_naissance":             "1948-06-30",
			"telephone":                  "819-555-0701",
			"cellulaire":                 "",
			"email":                      "",
			"adresse":                    "147 Avenue Laurier",
			"code_postal":                "J8V 7G7",
			"ville":                      "Gatineau",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "BERC 4806 3001",
			"numero_securite_sociale":    "789-012-345",
			"no_hcm":                     "HCM007",
			"no_chaur":                   "CH007",
			"no_dossier_leopard":         "LEO-2024-007",
			"medecin_famille_No_Licence": "12351",
			"note_fixe":                  "Ancien militaire - Prothèse auditive",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
		{
			"nom":                        "Morin",
			"prenom":                     "Denise",
			"date_naissance":             "1953-02-14",
			"telephone":                  "819-555-0801",
			"cellulaire":                 "819-555-0802",
			"email":                      "denise.morin@example.com",
			"adresse":                    "258 Rue Wellington",
			"code_postal":                "J8V 8H8",
			"ville":                      "Gatineau",
			"pays":                       "Canada",
			"numero_assurance_maladie":   "MORD 5302 1401",
			"numero_securite_sociale":    "890-123-456",
			"no_hcm":                     "HCM008",
			"no_chaur":                   "CH008",
			"no_dossier_leopard":         "LEO-2024-008",
			"medecin_famille_No_Licence": "12352",
			"note_fixe":                  "Diabète type 2 - Suivi nutritionnel",
			"Actif":                      1,
			"dcd":                        0,
			"created_by":                 userID,
		},
	}

	// Query d'insertion
	query := `
		INSERT INTO clients (
			nom, prenom, date_naissance, telephone, cellulaire, email, adresse,
			code_postal, ville, pays, numero_assurance_maladie, numero_securite_sociale,
			no_hcm, no_chaur, no_dossier_leopard, medecin_famille_No_Licence,
			note_fixe, Actif, dcd, created_by
		) VALUES (
			:nom, :prenom, :date_naissance, :telephone, :cellulaire, :email, :adresse,
			:code_postal, :ville, :pays, :numero_assurance_maladie, :numero_securite_sociale,
			:no_hcm, :no_chaur, :no_dossier_leopard, :medecin_famille_No_Licence,
			:note_fixe, :Actif, :dcd, :created_by
		)
	`

	// Insertion de chaque client
	insertedCount := 0
	for _, client := range clients {
		result, err := db.NamedExec(query, client)
		if err != nil {
			log.Printf("⚠️  Erreur insertion %s %s: %v", client["prenom"], client["nom"], err)
			continue
		}

		id, _ := result.LastInsertId()
		fmt.Printf("✅ Client créé: %s %s (ID: %d)\n", client["prenom"], client["nom"], id)
		insertedCount++
	}

	fmt.Printf("\n🎉 %d clients sur %d insérés avec succès!\n", insertedCount, len(clients))
	fmt.Println("🚀 Ta table clients est maintenant peuplée!")
}
