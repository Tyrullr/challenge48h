package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3" // Import pour activer le driver SQLite
)

func initDB(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (l
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table:", err)
	}
}

func EmailExists(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		log.Fatal("Erreur lors de la vérification de l'email:", err)
	}
	return count > 0
}

func AddUser(db *sql.DB, email, name string) {
	if EmailExists(db, email) {
		fmt.Println("L'utilisateur avec cet email existe déjà :", email)
		return
	}

	_, err := db.Exec("INSERT INTO users (email, name) VALUES (?, ?)", email, name)
	if err != nil {
		log.Fatal("Erreur lors de l'ajout de l'utilisateur:", err)
	}
	fmt.Println("Utilisateur ajouté:", name, "(", email, ")")
}

func RemoveUser(db *sql.DB, email string) {
	_, err := db.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal("Erreur lors de la suppression de l'utilisateur:", err)
	}
	fmt.Println("Utilisateur supprimé:", email)
}

func GetUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, email, name FROM users")
	if err != nil {
		log.Fatal("Erreur lors de la récupération des utilisateurs:", err)
	}
	defer rows.Close()

	fmt.Println("\nListe des utilisateurs :")
	for rows.Next() {
		var id int
		var email, name string
		if err := rows.Scan(&id, &email, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Email: %s | Nom: %s\n", id, email, name)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal("Erreur d'ouverture de la base de données:", err)
	}
	defer db.Close()

	initDB(db)

	AddUser(db, "alice@example.com", "Alice")
	AddUser(db, "bob@example.com", "Bob")
	AddUser(db, "alice@example.com", "Alice (duplicate)") // Ce test ne doit pas fonctionner

	GetUsers(db)

	RemoveUser(db, "alice@example.com")

	GetUsers(db)
}
