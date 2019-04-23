package say

import (
	"fmt"
	"strings"
)

var numberMap = map[int64]string{
	0:   "zero",
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	10:  "ten",
	11:  "eleven",
	12:  "twelve",
	13:  "thirteen",
	15:  "fifteen",
	18:  "eighteen",
	20:  "twenty",
	30:  "thirty",
	40:  "forty",
	50:  "fifty",
	80:  "eighty",
	100: "hundred",
}

var unitMap = []string{"thousand", "million", "billion", "trillion"}

const (
	Ten       = 10
	AHundred  = 100
	AThousand = 1000
	Max       = 999999999999
)

func Say(input int64) (string, bool) {
	if input < 0 || input > Max {
		return "", false
	}

	return say(input), true
}

func say(input int64) string {
	if input < AHundred {
		return twoDigits(input)
	}
	if input < AThousand {
		return threeDigits(input)
	}

	if input >= AThousand {
		last := input % AThousand
		var lastS, final string
		if last > 0 {
			lastS = say(last)
		}
		left := input / AThousand

		i := 0
		for left > 0 {
			cur := left % AThousand
			if cur > 0 {
				final = say(cur) + " " + unitMap[i] + " " + final
			}
			left = left / AThousand
			i++
		}
		return strings.TrimSpace(final + lastS)
	}

	return say(input)
}

// 100 - 999
func threeDigits(input int64) string {
	third := input / AHundred
	two := input % AHundred

	thirdS, twoDigitsS := numberMap[third], ""
	if two > 0 {
		twoDigitsS = " " + twoDigits(two)
	}

	return fmt.Sprintf("%s %s%s", thirdS, numberMap[AHundred], twoDigitsS)
}

// number 1-99
func twoDigits(input int64) string {
	if v, ok := numberMap[input]; ok {
		return v
	}

	first, second := input%Ten, input/Ten
	if input < 20 {
		return numberMap[first] + "teen"
	}

	//21 - 99
	var firstS, secondS string
	if v, ok := numberMap[second*Ten]; ok {
		secondS = v
	} else {
		secondS = numberMap[second] + "ty"
	}

	firstS = numberMap[first]
	return secondS + "-" + firstS
}
