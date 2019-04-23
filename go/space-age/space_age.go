package space

import (
	"math"
)

// Age accept an input of age in seconds, calculate how old someone would be on
func Age(seconds float64, p Planet) (result float64) {
	earthYear := seconds / (60 * 60 * 24 * 365.25)

	formula := map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	result = earthYear / formula[p]
	return math.Round(result*100) / 100
}

type Planet string
