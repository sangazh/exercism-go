package darts

import (
	"math"
)

func Score(x, y float64) int {
	radius := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case radius <= 10 && radius > 5:
		return 1
	case radius <= 5 && radius > 1:
		return 5
	case radius <= 1:
		return 10
	default:
		return 0
	}

}
