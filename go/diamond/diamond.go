package diamond

import (
	"errors"
)

func Gen(b byte) (s string, err error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("out of range")
	}
	n := b - 'A' + 1

	if n == 1 {
		return "A\n", nil
	}

	for i := uint8(1); i <= n; i++ {
		s += render(i, n)
	}

	for i := n - 1; i > 0; i-- {
		s += render(i, n)
	}
	return s, nil
}

func render(i, n uint8) (s string) {
	if i == 1 {
		return renderRowContainsOneA(n)
	}
	return renderOtherRows(i, n)
}

func renderRowContainsOneA(n uint8) (s string) {
	for i := uint8(0); i < 2*n-1; i++ {
		if i == n-1 {
			s += "A"
		} else {
			s += " "

		}
	}
	return s + "\n"
}

func renderOtherRows(i, n uint8) (s string) {
	char := string('A' + i - 1)
	for j := n; j > 0; j-- {
		if j == i {
			s += char
		} else {
			s += " "
		}
	}

	for j := uint8(2); j <= n; j++ {
		if j == i {
			s += char
		} else {
			s += " "
		}
	}

	return s + "\n"
}
