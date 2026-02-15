package route

import (
	"SCORESABER--New/src/handler"
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates, _ = template.ParseGlob("html/*.html")
}

func Routes(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/player", handler.PlayerPageHandler)
	mux.HandleFunc("/search", handler.SearchHandler)

	mux.HandleFunc("/favorite/toggle", handler.ToggleFavoriHandler)
	mux.HandleFunc("/favorites", handler.FavorisPageHandler)

	mux.HandleFunc("/apropos", handler.AproposHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "Home.html", nil)
	})
}
