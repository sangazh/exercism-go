package romannumerals

import (
	"errors"
	"math"
	"strings"
)

var m = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

var outOfRange = errors.New("out of range")

func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 3000 || arabic < 1 {
		return "", outOfRange
	}
	if v, ok := m[arabic]; ok {
		return v, nil
	}
	digits := DigitCnt(arabic)
	var result string
	for i := digits; i > 0; i-- {
		digit := CurDigit(arabic, i)
		switch i {
		case 4:
			result += lessThan10(digit, m[1000], "", "")
		case 3:
			result += lessThan10(digit, m[100], m[500], m[1000])
		case 2:
			result += lessThan10(digit, m[10], m[50], m[100])
		default:
			result += lessThan10(digit, m[1], m[5], m[10])
		}

	}
	return result, nil
}

func DigitCnt(num int) (cnt int) {
	if num < 10 {
		return 1
	}

	cnt = 1 + DigitCnt(num/10)
	return cnt
}

func CurDigit(num, position int) int {
	if position == 1 {
		return num % int(math.Pow10(position))
	}
	return (num / int(math.Pow10(position-1))) % 10
}

func lessThan10(digit int, one, five, ten string) string {
	switch {
	case digit <= 3:
		return strings.Repeat(one, digit)
	case digit == 4:
		return one + five
	case digit == 5:
		return five
	case digit > 5 && digit < 9:
		return five + strings.Repeat(one, digit-5)
	case digit == 9:
		return one + ten
	}
	return ""
}
