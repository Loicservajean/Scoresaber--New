package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ScoreStats struct {
	TotalScore            int     `json:"totalScore"`
	TotalRankedScore      int     `json:"totalRankedScore"`
	AverageRankedAccuracy float64 `json:"averageRankedAccuracy"`
	TotalPlayCount        int     `json:"totalPlayCount"`
	RankedPlayCount       int     `json:"rankedPlayCount"`
	ReplaysWatched        int     `json:"replaysWatched"`
}

type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Profile     string     `json:"profilePicture"`
	Country     string     `json:"country"`
	PP          float64    `json:"pp"`
	Rank        int        `json:"rank"`
	CountryRank int        `json:"countryRank"`
	ScoreStats  ScoreStats `json:"scoreStats"`
}

type PlayerScoreResponse struct {
	PlayerScores []PlayerScore `json:"playerScores"`
}

type PlayerScore struct {
	Leaderboard struct {
		ID             int    `json:"id"`
		SongName       string `json:"songName"`
		SongAuthorName string `json:"songAuthorName"`
		DifficultyRaw  string `json:"difficultyRaw"`
		CoverImage     string `json:"coverImage"`
		MaxScore       int    `json:"maxScore"`
	} `json:"leaderboard"`

	Score struct {
		Rank      int     `json:"rank"`
		BaseScore int     `json:"baseScore"`
		Accuracy  float64 `json:"accuracy"`
		PP        float64 `json:"pp"`
	} `json:"score"`
}

func GetUser(id string) (*User, int, error) {
	client := http.Client{Timeout: 5 * time.Second}

	url := "https://scoresaber.com/api/player/" + id + "/full"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur requête : %s", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur réseau : %s", err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("erreur API : %s", res.Status)
	}

	var data User
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur JSON : %s", decodeErr.Error())
	}

	return &data, res.StatusCode, nil
}

func GetUserScores(id string) ([]PlayerScore, int, error) {
	client := http.Client{Timeout: 5 * time.Second}

	url := "https://scoresaber.com/api/player/" + id + "/scores?limit=10&sort=recent"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur requête : %s", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur réseau : %s", err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("erreur API : %s", res.Status)
	}

	var data PlayerScoreResponse
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur JSON : %s", decodeErr.Error())
	}

	for i := range data.PlayerScores {
		if data.PlayerScores[i].Score.Accuracy == 0 {
			bs := float64(data.PlayerScores[i].Score.BaseScore)
			ms := float64(data.PlayerScores[i].Leaderboard.MaxScore)
			if ms > 0 {
				data.PlayerScores[i].Score.Accuracy = (bs / ms) * 100
			}
		}
	}

	return data.PlayerScores, res.StatusCode, nil
}

type LeaderboardSearch struct {
	Leaderboards []Leaderboard `json:"leaderboards"`
	Metadata     Metadata      `json:"metadata"`
}

type Metadata struct {
	Total        int `json:"total"`
	Page         int `json:"page"`
	ItemsPerPage int `json:"itemsPerPage"`
}

type Leaderboard struct {
	ID              int        `json:"id"`
	SongHash        string     `json:"songHash"`
	SongName        string     `json:"songName"`
	SongSubName     string     `json:"songSubName"`
	SongAuthorName  string     `json:"songAuthorName"`
	LevelAuthorName string     `json:"levelAuthorName"`
	Difficulty      Difficulty `json:"difficulty"`
	MaxScore        int        `json:"maxScore"`
	Ranked          bool       `json:"ranked"`
	Qualified       bool       `json:"qualified"`
	Loved           bool       `json:"loved"`
	Plays           int        `json:"plays"`
	DailyPlays      int        `json:"dailyPlays"`
	CoverImage      string     `json:"coverImage"`
}

type Difficulty struct {
	LeaderboardId int    `json:"leaderboardId"`
	Difficulty    int    `json:"difficulty"`
	GameMode      string `json:"gameMode"`
	DifficultyRaw string `json:"difficultyRaw"`
}

func GetMaps(name string, page int) ([]Leaderboard, int, error) {
	client := http.Client{Timeout: 5 * time.Second}

	encoded := url.QueryEscape(name)
	url := "https://scoresaber.com/api/leaderboards?search=" + encoded + "&page=" + strconv.Itoa(page)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur requête : %s", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur réseau : %s", err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("erreur API : %s", res.Status)
	}

	var data LeaderboardSearch
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("erreur JSON : %s", decodeErr.Error())
	}

	return data.Leaderboards, res.StatusCode, nil
}

func DifficultyName(raw string) string {
	switch {
	case strings.Contains(raw, "Easy"):
		return "Easy"
	case strings.Contains(raw, "Normal"):
		return "Normal"
	case strings.Contains(raw, "Hard"):
		return "Hard"
	case strings.Contains(raw, "ExpertPlus"):
		return "Expert+"
	case strings.Contains(raw, "Expert"):
		return "Expert"
	}
	return raw
}

func FilterMaps(maps []Leaderboard, query string) []Leaderboard {
	if query == "" {
		return maps
	}

	query = strings.ToLower(query)
	var result []Leaderboard

	for _, m := range maps {
		name := strings.ToLower(m.SongName)
		author := strings.ToLower(m.SongAuthorName)

		if strings.Contains(name, query) || strings.Contains(author, query) {
			result = append(result, m)
		}
	}

	return result
}
