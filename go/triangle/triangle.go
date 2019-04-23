package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
	Deg             // degenerate
)

// KindFromSides Determine if a triangle is equilateral, isosceles,
// scalene, or degenerate.
func KindFromSides(a, b, c float64) Kind {
	for _, i := range []float64{a, b, c} {
		if !check(i) {
			return NaT
		}
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a+b == c || a+c == b || b+c == a {
		return Deg
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}

func check(f float64) bool {
	if f > 0 && f < math.MaxFloat64 {
		return true
	}
	return false
}
