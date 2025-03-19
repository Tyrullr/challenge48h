package main

import (
	"log"
	"net/http"

	"challenge48h/internal" // Utilise le nom du module défini dans go.mod
)

func main() {
	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", internal.NewServer()))
}
