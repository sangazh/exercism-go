package encode

import (
	"fmt"
)

func RunLengthEncode(input string) (output string) {
	if len(input) == 0 {
		return input
	}

	type charEncode struct {
		char  int32
		count int
	}

	chars := make([]charEncode, 0)

	for i, c := range input {
		if i == 0 || input[i] != input[i-1] {
			chars = append(chars, charEncode{char: c, count: 1})
			continue
		}

		chars[len(chars)-1].count++
	}

	for _, c := range chars {
		n := c.count
		if n > 1 {
			output += fmt.Sprintf("%d%s", n, string(c.char))
		} else {
			output += string(c.char)
		}
	}

	return output

}

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return input
	}

	return ""

}
