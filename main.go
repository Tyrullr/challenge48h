package main

import (
	"log"
	"net/http"

	"challenge48h/internal" // Utilise le nom du module défini dans go.mod
)

func main() {
	log.Println("Serveur démarré sur http://localhost:5050")
	log.Fatal(http.ListenAndServe(":5050", internal.NewServer()))
}
