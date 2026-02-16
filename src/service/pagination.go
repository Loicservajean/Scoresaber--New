package service

/* pagination des résultat */
func Pagination(maps []Leaderboard, page int, limit int) []Leaderboard {
	start := (page - 1) * limit
	end := start + limit

	if start > len(maps) {
		return []Leaderboard{}
	}
	if end > len(maps) {
		end = len(maps)
	}

	return maps[start:end]
}
