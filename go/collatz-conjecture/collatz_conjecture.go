package collatzconjecture

import (
	"errors"
)

var ErrInvalidInput = errors.New("input must be positive")

func CollatzConjecture(input int) (steps int, err error) {
	if input <= 0 {
		return 0, ErrInvalidInput
	}
	for {
		if input == 1 {
			break
		}
		if isEven(input) {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		steps++
	}
	return
}

func isEven(n int) bool {
	return n%2 == 0
}
