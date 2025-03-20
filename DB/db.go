package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal("Erreur d'ouverture de la base de données:", err)
	}

	query := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table:", err)
	}

	return db
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func RegisterUser(db *sql.DB, username, password string) {
	hashedPassword := hashPassword(password)

	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		log.Fatal("Erreur lors de l'enregistrement de l'utilisateur:", err)
	}

	fmt.Println("Utilisateur enregistré avec succès!")
}

func LoginUser(db *sql.DB, username, password string) bool {
	var storedHashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Utilisateur non trouvé.")
			return false
		}
		log.Fatal("Erreur lors de la récupération de l'utilisateur:", err)
	}

	if storedHashedPassword != hashPassword(password) {
		fmt.Println("Mot de passe incorrect.")
		return false
	}

	fmt.Println("Connexion réussie!")
	return true
}

func main() {
	db := initDB()
	defer db.Close()

	// Exemples d'utilisation
	RegisterUser(db, "Alice", "password123")
	RegisterUser(db, "Bob", "securePass")

	LoginUser(db, "Alice", "password123") // Succès
	LoginUser(db, "Bob", "wrongpass")     // Échec
}
