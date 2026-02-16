package route

import (
	"SCORESABER--New/src/handler"
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

/* chargement des template */
func LoadTemplates() {
	templates, _ = template.ParseGlob("html/*.html")
}

/* liste des routes */
func Routes(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("/player", handler.PlayerPageHandler)
	mux.HandleFunc("/search", handler.SearchHandler)
	mux.HandleFunc("/favorite/toggle", handler.ToggleFavoriHandler)
	mux.HandleFunc("/favorites", handler.FavorisPageHandler)
	mux.HandleFunc("/apropos", handler.AproposHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path == "/" || r.URL.Path == "/home" {
			templates.ExecuteTemplate(w, "Home.html", nil)
		} else {
			templates.ExecuteTemplate(w, "404.html", nil)
		}
	})
}
