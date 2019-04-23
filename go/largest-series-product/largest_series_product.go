package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return -1, errors.New("span should be greater or equal than 0")
	}
	if len(digits) == 0 && span == 0 {
		return 1, nil
	}
	if len(digits) < span {
		return -1, errors.New("span shouldn't be larger than the length")
	}

	var largest int64
	for i := 0; i <= (len(digits) - span); i++ {
		product, err := cal(digits[i : span+i])
		if err != nil {
			return -1, err
		}
		if product > largest {
			largest = product
		}
	}

	return largest, nil
}

func cal(digits string) (int64, error) {
	var p int64 = 1
	for _, ds := range digits {
		if d, err := strconv.Atoi(string(ds)); err != nil {
			return 0, errors.New("digits contain non-digits")
		} else {
			p *= int64(d)
		}
	}
	return p, nil
}
