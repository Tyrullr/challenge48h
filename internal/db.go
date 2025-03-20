package internal

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initialise la base de données
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// HashPassword hache un mot de passe avec SHA-256
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// RegisterUser ajoute un utilisateur à la BDD
func RegisterUser(email, password string) error {
	hashedPassword := HashPassword(password)

	_, err := db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hashedPassword)
	if err != nil {
		return errors.New("email déjà utilisé")
	}
	return nil
}

// Authenticate vérifie si l'utilisateur existe et si le mot de passe est correct
func Authenticate(email, password string) error {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		return errors.New("utilisateur non trouvé")
	}

	if storedPassword != HashPassword(password) {
		return errors.New("mot de passe incorrect")
	}
	return nil
}
