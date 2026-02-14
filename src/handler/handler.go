package handler

import (
	"SCORESABER--New/src/service"
	"html/template"
	"net/http"
)

func PlayerPageHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	user, status, err := service.GetUser(id)
	if err != nil || status != http.StatusOK {
		http.Error(w, "Impossible de récupérer le joueur", status)
		return
	}

	t, _ := template.ParseFiles("html/Player.html")
	t.Execute(w, user)
}
