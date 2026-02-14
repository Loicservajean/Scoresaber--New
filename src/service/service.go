package service

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetUser(id string) (*User, int, error) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

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
