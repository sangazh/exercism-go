package railfence

import (
	"strings"
)

func Encode(message string, rails int) string {
	transform := make([]string, rails)
	r, direction := 1, 1

	for _, letter := range message {
		transform[r-1] += string(letter)
		r, direction = zipZap(r, rails, direction)
	}

	return strings.Join(transform, "")
}

func Decode(message string, rails int) string {
	//make an empty zig-zag shape
	empty := make([][]string, rails)
	r, direction := 1, 1

	for i := 0; i < len(message); i++ {
		empty = fillEmpty(empty, r-1)
		r, direction = zipZap(r, rails, direction)
	}

	//fill the cipher text
	var i int
	for _, line := range empty {
		for l, v := range line {
			if v == "?" {
				line[l] = string(message[i])
				i++
			}
		}
	}

	// read original message
	r, direction = 1, 1
	var decoded string
	for i := 0; i < len(message); i++ {
		decoded += string(empty[r-1][i])
		r, direction = zipZap(r, rails, direction)
	}
	return decoded
}

func fillEmpty(empty [][]string, n int) [][]string {
	for i := range empty {
		if i == n {
			empty[i] = append(empty[i], "?")
		} else {
			empty[i] = append(empty[i], ".")
		}
	}
	return empty
}

func zipZap(r, rails, direction int) (newR, newD int) {
	r += direction
	switch {
	case r == rails:
		direction = -1
	case r == 1:
		direction = 1
	}
	return r, direction
}
