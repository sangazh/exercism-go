package piglatin

import (
	"strings"
)

func Sentence(input string) string {
	words := strings.Split(input, " ")
	result := make([]string, len(words))
	for i, word := range words {
		result[i] = handle(word)
	}

	return strings.Join(result, " ")
}

func handle(input string) string {
	if isVowel(input[:2]) {
		return input + "ay"
	}

	if beforeY, other, yes := hasYInMiddle(input); yes {
		return other + beforeY + "ay"
	}

	if consonant, other, yes := isConsonant(input); yes {
		return other + consonant + "ay"
	}
	return ""
}

var vowels = "aoieu"
var consonants = "bcdfghjklmnpqrstvxyz"

// if words begins with vowel
func isVowel(input string) bool {
	for _, v := range vowels {
		if string(v) == string(input[0]) {
			return true
		}
	}

	if input == "xr" || input == "yt" {
		return true
	}

	return false
}

// if word begins with consonant
func isConsonant(input string) (consonant, other string, yes bool) {
	for i, letter := range input {
		// if word begins with consonant sound followed by "qu
		if i < len(input)-1 && input[i:i+2] == "qu" {
			consonant += "qu"
			yes = true
			break
		}
		index := strings.IndexRune(consonants, letter)
		if index == -1 {
			break
		}
		consonant += string(letter)
		yes = true
	}

	if yes {
		return consonant, strings.TrimLeft(input, consonant), yes
	}

	return "", "", false

}

// if word contains a "y"
func hasYInMiddle(input string) (beforeY, other string, yes bool) {
	index := strings.IndexRune(input, 'y')
	if index < 1 || (len(input) > 2 && index == len(input)-1) {
		return "", "", false
	}

	return input[:index], input[index:], true
}
