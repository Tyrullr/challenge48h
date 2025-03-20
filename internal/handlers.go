package internal

import (
	"html/template"
	"net/http"
)

// renderTemplate rend un fichier HTML spécifique
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Page introuvable", http.StatusNotFound)
		return
	}
	t.Execute(w, nil)
}

// Handlers pour chaque page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Presentation")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Index")
}

func ActualiteHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Actualité")
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Admin")
}

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Connexion")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Contact")
}

func PresentationHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Presentation")
}
