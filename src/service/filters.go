package service

import "strings"

func FilterMapsBy(maps []Leaderboard, query string) []Leaderboard {
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

func FilterBy(maps []Leaderboard, difficulty, mode, status string) []Leaderboard {
	var result []Leaderboard

	for _, m := range maps {
		if difficulty != "" && m.Difficulty.DifficultyRaw != difficulty {
			continue
		}

		if mode != "" && m.Difficulty.GameMode != mode {
			continue
		}

		if status != "" {
			if status == "ranked" && !m.Ranked {
				continue
			}
			if status == "qualified" && !m.Qualified {
				continue
			}
			if status == "loved" && !m.Loved {
				continue
			}
		}

		result = append(result, m)
	}

	return result
}
