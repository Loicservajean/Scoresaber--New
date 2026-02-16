package handler

import (
	"SCORESABER--New/src/service"
	"html/template"
	"net/http"
	"strconv"
)

type SearchPageData struct {
	Query      string
	Maps       []service.Leaderboard
	Difficulty string
	Mode       string
	Status     string
	Page       int
	PrevPage   int
	NextPage   int
	Limit      int
}

/* gestion de la recherche */
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		p, _ := strconv.Atoi(pageStr)
		if p > 0 {
			page = p
		}
	}

	maps, status, err := service.GetMaps(q, page)
	if err != nil || status != http.StatusOK {
		http.Error(w, "Impossible de récupérer les maps", status)
		return
	}

	maps = service.FilterMapsBy(maps, q)

	difficulty := r.URL.Query().Get("difficulty")
	mode := r.URL.Query().Get("mode")
	statusFilter := r.URL.Query().Get("status")
	maps = service.FilterBy(maps, difficulty, mode, statusFilter)

	for i := range maps {
		maps[i].Difficulty.DifficultyRaw = service.DifficultyName(maps[i].Difficulty.DifficultyRaw)
		maps[i].IsFav = service.IsFavori(maps[i].ID)
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 30
	if limitStr != "" {
		l, _ := strconv.Atoi(limitStr)
		if l == 10 || l == 20 || l == 30 {
			limit = l
		}
	}

	prev := page - 1
	if prev < 1 {
		prev = 1
	}

	data := SearchPageData{
		Query:      q,
		Maps:       maps,
		Difficulty: difficulty,
		Mode:       mode,
		Status:     statusFilter,
		Page:       page,
		PrevPage:   prev,
		NextPage:   page + 1,
		Limit:      limit,
	}

	t, _ := template.ParseFiles("html/Search.html")
	t.Execute(w, data)
}
