package handler

import (
	"SCORESABER--New/src/service"
	"html/template"
	"net/http"
)

type PlayerPageData struct {
	User   *service.User
	Scores []service.PlayerScore
}

func PlayerPageHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	user, status, err := service.GetUser(id)
	if err != nil || status != http.StatusOK {
		http.Error(w, "Impossible de récupérer le joueur", status)
		return
	}

	scores, status2, err2 := service.GetUserScores(id)
	if err2 != nil || status2 != http.StatusOK {
		http.Error(w, "Impossible de récupérer les scores", status2)
		return
	}

	data := PlayerPageData{
		User:   user,
		Scores: scores,
	}

	t, _ := template.ParseFiles("html/Player.html")
	t.Execute(w, data)
}
