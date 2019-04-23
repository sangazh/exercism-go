package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	if len(input) == 0 {
		return false
	}
	input = strings.ToLower(input)
	alpha := make(map[string]int, 0)
	for _, l := range input {
		if unicode.IsLetter(l) {
			alpha[string(l)]++
		}
	}
	if len(alpha) < 26 {
		return false
	}

	return true
}
