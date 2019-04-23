package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)

	var parsed string
	for _, r := range phrase {
		if unicode.IsLetter(r) || unicode.IsSpace(r) || unicode.IsDigit(r) || string(r) == "'" {
			parsed += string(r)
			continue
		}
		if unicode.IsPunct(r) {
			parsed += " "
			continue
		}
	}

	f := make(map[string]int)
	for _, word := range strings.Fields(parsed) {
		word = strings.Trim(word, "'")
		f[word] += 1
	}

	return f
}
