package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	input = normalize(input)

	r := int(math.Round(math.Sqrt(float64(len(input)))))
	c := len(input) / r
	if len(input)%r != 0 {
		c += 1
	}

	input = padding(input, r, c)

	output := make([]string, c)
	for i, s := range input {
		output[i%c] += string(s)
	}

	return strings.Join(output, " ")
}

func normalize(input string) (after string) {
	before := strings.ToLower(input)
	for _, s := range before {
		if unicode.IsLetter(s) || unicode.IsDigit(s) {
			after += string(s)
		}
	}
	return after
}

func padding(input string, r, c int) string {
	if c*r == len(input) || r*r == len(input) {
		return input
	}
	n := c*r - len(input)
	for i := 0; i < n; i++ {
		input += " "
	}
	return input
}
