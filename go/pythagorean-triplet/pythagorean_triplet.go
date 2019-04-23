package pythagorean

import "math"

type Triplet [3]int

func (t Triplet) isValid() bool {
	if len(t) < 3 {
		return false
	}
	return t[0]*t[0]+t[1]*t[1] == t[2]*t[2]
}

func Range(min, max int) []Triplet {
	ts := make([]Triplet, 0)

	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			a := int(math.Sqrt(float64(max*max - i*i - j*j)))
			for k := a; k <= max; k++ {
				t := Triplet{i, j, k}
				if t.isValid() {
					ts = append(ts, t)
				}
			}
		}
	}
	return ts
}

func Sum(p int) []Triplet {
	ts := make([]Triplet, 0)
	for n := 2; n <= p; n++ {
		for m := n; m <= p-n; m++ {
			t := Triplet{n, m, p - n - m}
			if t[2] < t[1] {
				break
			}

			if t.isValid() {
				ts = append(ts, t)
			}
		}
	}

	return ts
}
