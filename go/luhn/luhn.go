//Package luhn determines the given number whether it is valid per the Luhn formula
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

//Given a number determine whether or not it is valid per the Luhn formula
func Valid(code string) bool {
	code = strings.Replace(code, " ", "", -1)
	if len(code) < 2 {
		return false
	}
	slice := []rune(code)

	for _, c := range slice {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	j := 0
	doubling := ""
	for i := len(slice) - 1; i > -1; i-- {
		j++
		if j%2 == 0 {
			digit, _ := strconv.Atoi(string(slice[i]))
			doubling = strconv.Itoa(double(digit)) + doubling
		} else {
			doubling = string(slice[i]) + doubling
		}
	}
	if a := sum(doubling); a%10 == 0 {
		return true
	}

	return false
}
func sum(double string) (result int) {
	for _, d := range double {
		i, _ := strconv.Atoi(string(d))
		result += i
	}
	return result
}

func double(i int) int {
	i *= 2
	if i > 9 {
		i -= 9
	}
	return i
}
