package allyourbase

import (
	"errors"
	"math"
)

var ErrInvalidInputBase = errors.New("input base must be >= 2")
var ErrInvalidOutputBase = errors.New("output base must be >= 2")
var ErrInvalidDigit = errors.New("all digits must satisfy 0 <= d < input base")

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (output []int, err error) {
	if inputBase < 2 {
		return nil, ErrInvalidInputBase
	}
	if outputBase < 2 {
		return nil, ErrInvalidOutputBase
	}

	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	//check if digits are valid
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, ErrInvalidDigit
		}
	}

	var t int
	for i, j := len(inputDigits)-1, 0; i >= 0; i-- {
		t += inputDigits[i] * int(math.Pow(float64(inputBase), float64(j)))
		j++
	}

	for {
		o := t / outputBase
		output = append([]int{t % outputBase}, output...)
		if o <= 0 {
			break
		}

		t = o
	}

	return
}
