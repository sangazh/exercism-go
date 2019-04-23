package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("two sequences should have equal length")
	}

	result := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result += 1
		}
	}
	return result, nil
}
