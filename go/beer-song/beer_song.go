package beer

import (
	"errors"
	"fmt"
)

var verseA = "%v %s of beer on the wall, %v %s of beer.\n"
var verseB = "Take %s down and pass it around, %v %s of beer on the wall.\n"
var verseC = "Go to the store and buy some more, 99 bottles of beer on the wall.\n"

func Verse(n int) (string, error) {
	bottles := "bottles"
	bottle := "bottle"
	switch {
	case n == 0:
		return fmt.Sprintf(verseA+verseC, "No more", bottles, "no more", bottles), nil
	case n == 2:
		return fmt.Sprintf(verseA+verseB, n, bottles, n, bottles, "one", n-1, bottle), nil
	case n == 1:
		return fmt.Sprintf(verseA+verseB, n, bottle, n, bottle, "it", "no more", bottles), nil
	case n > 0 && n <= 99:
		return fmt.Sprintf(verseA+verseB, n, bottles, n, bottles, "one", n-1, bottles), nil
	}
	return "", errors.New("out of range")
}

func Verses(max, min int) (s string, err error) {
	if max < min {
		return "", errors.New("upper bound should greater than lower bound")
	}
	for i := max; i >= min; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}
		s += v + "\n"
	}
	return s, nil
}

func Song() string {
	song, err := Verses(99, 0)
	if err != nil {
		return ""
	}

	return song
}
