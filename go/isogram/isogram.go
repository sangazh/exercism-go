package isogram

import "strings"

//Determine if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}

	word = strings.ToLower(word)
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)

	repeats := make(map[string]int, 0)
	for _, letter := range word {
		if _, ok := repeats[string(letter)]; !ok {
			repeats[string(letter)] = 1
		} else {
			repeats[string(letter)] += 1
		}
	}

	for _, count := range repeats {
		if count > 1 {
			return false
		}
	}
	return true
}
