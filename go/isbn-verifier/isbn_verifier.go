package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	if len(isbn) == 0 {
		return false
	}

	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}

	var check int
	for i, d := range isbn {
		var n int
		if d == 'X' {
			if i != len(isbn)-1 {
				// "X" at wrong position
				return false
			}
			n = 10
		} else {
			num, err := strconv.Atoi(string(d))
			if err != nil {
				//invalid digit
				return false
			}
			n = num
		}

		check += (len(isbn) - i) * n
	}

	return (check % 11) == 0
}
