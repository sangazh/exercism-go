package allergies

var scoresMap = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(score uint) []string {
	result := make([]string, 0)
	if score == 0 {
		return result
	}

	for i := 0; i < len(scoresMap); i++ {
		var s uint = 1 << uint(i)
		substance := scoresMap[s]
		if score&s == s {
			result = append(result, substance)
		}
	}

	return result
}

func AllergicTo(score uint, substance string) bool {
	for i := 0; i < len(scoresMap); i++ {
		var s uint = 1 << uint(i)
		if s > score {
			break
		}

		if v, ok := scoresMap[s]; ok {
			if v == substance && score&s == s {
				return true
			}
		}
	}

	return false
}
