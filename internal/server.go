package internal

import (
	"net/http"
)

// NewServer crée un serveur HTTP et configure les routes
func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Servir les fichiers statiques (CSS, JS, images)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Définition des routes
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/actualite", ActualiteHandler)
	mux.HandleFunc("/admin", AdminHandler)
	mux.HandleFunc("/connection", ConnectionHandler)
	mux.HandleFunc("/contact", ContactHandler)
	mux.HandleFunc("/presentation", PresentationHandler)
	mux.HandleFunc("/index", IndexHandler)

	return mux
}
