package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structure utilisateur
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler gère l'inscription
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err = RegisterUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Inscription réussie !")
}

// LoginHandler gère la connexion
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err = Authenticate(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Connexion réussie !")
}

// NewServer configure les routes et retourne le serveur
func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes existantes
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/actualite", ActualiteHandler)
	mux.HandleFunc("/admin", AdminHandler)
	mux.HandleFunc("/connection", ConnectionHandler)
	mux.HandleFunc("/contact", ContactHandler)
	mux.HandleFunc("/presentation", PresentationHandler)
	mux.HandleFunc("/index", IndexHandler)

	// Nouvelles routes (authentification)
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/login", LoginHandler)

	return mux
}
