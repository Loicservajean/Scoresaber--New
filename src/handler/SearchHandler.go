package handler

import (
	"SCORESABER--New/src/service"
	"html/template"
	"net/http"
)

type SearchPageData struct {
	Query string
	Maps  []service.Leaderboard
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("q")

	maps, status, err := service.GetMaps(q)
	if err != nil || status != http.StatusOK {
		http.Error(w, "Impossible de récupérer les maps", status)
		return
	}

	for i := range maps {
		maps[i].Difficulty.DifficultyRaw = service.DifficultyName(maps[i].Difficulty.DifficultyRaw)
	}

	data := SearchPageData{
		Query: q,
		Maps:  maps,
	}

	t, _ := template.ParseFiles("html/Search.html")
	t.Execute(w, data)
}
