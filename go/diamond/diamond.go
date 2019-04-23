package diamond

import (
	"errors"
)

func Gen(b byte) (string, error) {
	order := int(b - 'A' + 1)

	if order < 0 || order > 26 {
		return "", errors.New("out or range")
	}
	if order == 1 {
		return "A", nil
	}

	height := order*2 - 1
	result := make([][]byte, height)
	for i := 0; i < height; i++ {
		for j := 0; j < height; j++ {
			if j == order {

			}
			result[i][j] = '.'
		}
	}

	return "", nil
}
