package handler

import (
	"SCORESABER--New/src/service"
	"html/template"
	"net/http"
	"strconv"
)

/* ajoute / retire un favoris */
func ToggleFavoriHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for _, m := range service.LastSearchResults {
		if m.ID == id {
			if service.IsFavori(id) {
				service.RemoveFavori(id)
			} else {
				service.AddFavori(m)
			}
			break
		}
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

/* affiche la page favoris */
func FavorisPageHandler(w http.ResponseWriter, r *http.Request) {
	favs := service.GetFavoris()
	t, _ := template.ParseFiles("html/Favorites.html")
	t.Execute(w, favs)
}
