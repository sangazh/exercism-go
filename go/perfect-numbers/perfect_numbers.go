package perfect

import (
	"errors"
	"math"
)

type Classification int

const (
	ClassificationDeficient Classification = iota + 1
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("not a natural number")

func Classify(input int64) (Classification, error) {
	if input <= 0 {
		return 0, ErrOnlyPositive
	}

	if input == 1 {
		return ClassificationDeficient, nil
	}

	factors := map[int64]int{1: 1}
	if input%2 == 0 {
		factors[input/2] = 1
	}

	for i := int64(2); i < int64(math.Sqrt(float64(input))+1); i++ {
		if input%int64(i) == 0 {
			factors[i] += 1
			factors[input/i] += 1
		}
	}

	var sum int64
	for f := range factors {
		sum += f
	}

	switch {
	case input == sum:
		return ClassificationPerfect, nil
	case sum > input:
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
