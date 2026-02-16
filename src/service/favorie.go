package service

import (
	"encoding/json"
	"os"
)

type Favori struct {
	ID             int        `json:"id"`
	SongName       string     `json:"songName"`
	SongAuthorName string     `json:"songAuthorName"`
	CoverImage     string     `json:"coverImage"`
	Difficulty     Difficulty `json:"difficulty"`
	Plays          int        `json:"plays"`
	Ranked         bool       `json:"ranked"`
	Qualified      bool       `json:"qualified"`
	Loved          bool       `json:"loved"`
}

var Favoris []Favori
var favFile = "data/favoris.json"

/* charge les favoris depuis le json */
func init() {
	data, err := os.ReadFile(favFile)
	if err == nil {
		json.Unmarshal(data, &Favoris)
	}
}

/* ecrit les favoris dans le json */
func save() {
	data, _ := json.MarshalIndent(Favoris, "", "  ")
	os.WriteFile(favFile, data, 0644)
}

/* ajout de favori si non présent */
func AddFavori(lb Leaderboard) {
	for _, f := range Favoris {
		if f.ID == lb.ID {
			return
		}
	}

	Favoris = append(Favoris, Favori{
		ID:             lb.ID,
		SongName:       lb.SongName,
		SongAuthorName: lb.SongAuthorName,
		CoverImage:     lb.CoverImage,
		Difficulty:     lb.Difficulty,
		Plays:          lb.Plays,
		Ranked:         lb.Ranked,
		Qualified:      lb.Qualified,
		Loved:          lb.Loved,
	})

	save()
}

/* retirer un favoris */
func RemoveFavori(id int) {
	newList := []Favori{}
	for _, f := range Favoris {
		if f.ID != id {
			newList = append(newList, f)
		}
	}
	Favoris = newList
	save()
}

/* check si favoris ou pas */
func IsFavori(id int) bool {
	for _, f := range Favoris {
		if f.ID == id {
			return true
		}
	}
	return false
}

/* retourne tous les favoris */
func GetFavoris() []Favori {
	return Favoris
}
