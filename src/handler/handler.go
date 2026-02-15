package handler

import (
	"SCORESABER--New/src/service"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

type PlayerPageData struct {
	User   *service.User
	Scores []service.PlayerScore
}

func PlayerPageHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Aucun nom fourni", http.StatusBadRequest)
		return
	}

	if !isNum(id) {
		n, err := searchName(id)
		if err != nil {
			http.Error(w, "Joueur introuvable", http.StatusNotFound)
			return
		}
		id = n
	}

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

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func searchName(name string) (string, error) {
	u := "https://scoresaber.com/api/players?search=" + url.QueryEscape(name)
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var d struct {
		Players []struct {
			ID string `json:"id"`
		} `json:"players"`
	}

	if json.NewDecoder(resp.Body).Decode(&d) != nil {
		return "", errors.New("decode error")
	}

	if len(d.Players) == 0 {
		return "", errors.New("not found")
	}

	return d.Players[0].ID, nil
}

func AproposHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/Apropos.html")
	t.Execute(w, nil)
}
