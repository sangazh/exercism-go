package palindrome

import (
	"errors"
	"fmt"
)

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmax < fmin {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			np := NewProduct(i, j)
			if np.isPalindrome() {
				if np.gt(pmax) {
					pmax = *np
				}
				if np.eq(pmax) {
					pmax.addFactors(i, j)
				}

				if pmin.isEmpty() {
					pmin = *np
					continue
				}

				if np.lt(pmin) {
					pmin = *np
				}
				if np.eq(pmin) {
					pmin.addFactors(i, j)
				}

			}
		}
	}
	if pmin.isEmpty() && pmax.isEmpty() {
		return Product{}, Product{}, errors.New("no palindromes...")
	}

	return
}

func (p Product) isPalindrome() bool {
	s := string([]rune(fmt.Sprint(p.Product)))
	if len(s) == 1 {
		return true
	}

	for i := range s[:(len(s))/2] {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func (p *Product) addFactors(a, b int) {
	var f = [2]int{}
	if a <= b {
		f[0], f[1] = a, b
	} else {
		f[0], f[1] = b, a
	}
	for _, exist := range p.Factorizations {
		if exist[0] == f[0] {
			return
		}
	}

	p.Factorizations = append(p.Factorizations, f)
}

func (p Product) isEmpty() bool {
	return p.Product == 0
}

func (p Product) gt(q Product) bool {
	return p.Product > q.Product
}

func (p Product) lt(q Product) bool {
	return p.Product < q.Product
}

func (p Product) eq(q Product) bool {
	return p.Product == q.Product
}

func NewProduct(a, b int) *Product {
	var f = [2]int{}
	if a <= b {
		f[0], f[1] = a, b
	} else {
		f[0], f[1] = b, a
	}

	return &Product{
		Product:        a * b,
		Factorizations: [][2]int{f},
	}
}
