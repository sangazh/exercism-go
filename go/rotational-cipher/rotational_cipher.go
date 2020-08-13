package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(input string, shift int) (output string) {
	for _, s := range input {
		if unicode.IsLower(s) {
			s = getLower(s + int32(shift))
		} else if unicode.IsUpper(s) {
			s = getUpper(s + int32(shift))
		}

		output += string(s)
	}

	return
}

func getLower(r rune) rune {
	if r > 'z' {
		r -= 26
	} else if r < 'a' {
		r += 26
	}
	return r
}

func getUpper(r rune) rune {
	if r > 'Z' {
		r -= 26
	} else if r < 'A' {
		r += 26
	}
	return r
}
