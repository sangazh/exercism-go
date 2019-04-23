package etl

import "strings"

func Transform(old map[int][]string) map[string]int {
	result := make(map[string]int)
	if len(old) == 0 {
		return result
	}
	for score, letters := range old {
		for _, l := range letters {
			result[strings.ToLower(l)] = score
		}
	}

	return result
}
