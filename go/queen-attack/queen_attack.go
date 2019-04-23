package queenattack

import (
	"github.com/pkg/errors"
	"math"
)

func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errors.New("two chess cannot at same position")
	}

	white, ok := parse(w)
	if !ok {
		return false, errors.New("white invalid")
	}

	black, ok := parse(b)
	if !ok {
		return false, errors.New("black invalid")
	}

	if white[0] == black[0] || white[1] == black[1] {
		return true, nil
	}

	if math.Abs(float64(white[0]-black[0])) == math.Abs(float64(white[1]-black[1])) {
		return true, nil
	}

	return false, nil
}

func parse(chess string) ([]int, bool) {
	if len(chess) != 2 {
		return nil, false
	}

	result := make([]int, 2)
	result[0] = int(chess[0]-'a') + 1
	result[1] = int(chess[1] - '0')

	if result[0] < 1 || result[0] > 8 || result[1] < 1 || result[1] > 8 {
		return nil, false
	}

	return result, true
}
