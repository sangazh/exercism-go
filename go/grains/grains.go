package grains

import (
	"github.com/pkg/errors"
	"math"
)

//calculate squares between 1 and 64
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("input should be between 1 and 64")
	}

	s := math.Pow(2, float64(input-1))
	return uint64(s), nil
}

// Calculate the number of grains of wheat on a chessboard given that the number
//on each square doubles.
func Total() (result uint64) {
	for i := 1; i < 65; i++ {
		square, _ := Square(i)
		result += square
	}

	return result
}
