package secret

import (
	"fmt"
)

var handshakeMap = map[int]string{
	1: "wink",
	2: "double blink",
	3: "close your eyes",
	4: "jump",
}

func Handshake(code uint) (result []string) {
	b := []byte(fmt.Sprintf("%b", code))

	j := 1
	isReverse := false
	for i := len(b) - 1; i > -1; i-- {
		if string(b[i]) == "1" {
			switch j {
			case 1, 2, 3, 4:
				result = append(result, handshakeMap[j])
			case 5:
				isReverse = true
			}
		}
		j++
	}

	if isReverse {
		result = reverse(result)
	}

	return result
}

func reverse(s []string) (result []string) {
	result = make([]string, 0, len(s))
	for i := len(s) - 1; i > -1; i-- {
		result = append(result, s[i])
	}

	return
}
