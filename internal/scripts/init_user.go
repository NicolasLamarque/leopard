package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

	// Vérifier si un user existe déjà
	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = ?", "admin")
	if err != nil {
		log.Fatalf("❌ Erreur vérification users: %v", err)
	}

	if count > 0 {
		fmt.Println("ℹ️  L'utilisateur 'admin' existe déjà")
		fmt.Println("🔄 Mise à jour du mot de passe...")

		// Données
		password := "admin123" // Change ce mot de passe!

		// Générer hash avec bcrypt (IMPORTANT: cost 14 comme dans app.go)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			log.Fatalf("❌ Erreur génération hash: %v", err)
		}

		// Mettre à jour le mot de passe
		_, err = db.Exec("UPDATE users SET password = ?, salt = NULL WHERE username = ?", string(hashedPassword), "admin")
		if err != nil {
			log.Fatalf("❌ Erreur mise à jour: %v", err)
		}

		fmt.Println("✅ Mot de passe mis à jour!")
		fmt.Printf("   Username: admin\n")
		fmt.Printf("   Password: %s\n", password)
		return
	}

	// Données du premier utilisateur
	username := "admin"
	password := "admin123" // Change ce mot de passe!
	fullName := "Administrateur"
	role := "admin"

	// Générer hash avec bcrypt (IMPORTANT: cost 14 comme dans app.go)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("❌ Erreur génération hash: %v", err)
	}

	// Insertion dans la DB (sans salt car on utilise bcrypt)
	query := `
		INSERT INTO users (username, password, salt, full_name, role)
		VALUES (?, ?, NULL, ?, ?)
	`

	result, err := db.Exec(query, username, string(hashedPassword), fullName, role)
	if err != nil {
		log.Fatalf("❌ Erreur insertion: %v", err)
	}

	id, _ := result.LastInsertId()
	fmt.Println("\n✅ Utilisateur créé avec succès!")
	fmt.Printf("   ID: %d\n", id)
	fmt.Printf("   Username: %s\n", username)
	fmt.Printf("   Password: %s\n", password)
	fmt.Printf("   Role: %s\n", role)
	fmt.Println("\n⚠️  N'oublie pas de changer le mot de passe après la première connexion!")
	fmt.Println("\n🚀 Tu peux maintenant te connecter à l'application!")
}
