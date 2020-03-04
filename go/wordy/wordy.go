package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if len(question) == 0 {
		return 0, false
	}
	s := strings.TrimLeft(question, "What is")
	s = strings.TrimRight(s, "?")
	words := strings.Fields(s)
	numbers := make([]int, 0)
	operations := make([]string, 0)

	var flag bool // default false is number, true is operations
	for _, k := range words {
		switch k {
		case "plus", "minus", "multiplied", "divided":
			if !flag {
				return 0, false
			}
			operations = append(operations, k)
			flag = false
		case "by":
			continue
		default:
			if flag {
				return 0, false
			}
			n, err := strconv.Atoi(k)
			if err != nil {
				return 0, false
			}
			numbers = append(numbers, n)
			flag = true
		}
	}
	if len(numbers) == 0 {
		return 0, false
	}

	if len(numbers) == len(operations) {
		return 0, false
	}

	if len(numbers) == 1 {
		return numbers[0], true
	}

	for _, o := range operations {
		n1 := numbers[0]
		numbers = numbers[1:]

		switch o {
		case "plus":
			numbers[0] += n1
		case "minus":
			numbers[0] = n1 - numbers[0]
		case "multiplied":
			numbers[0] *= n1
		case "divided":
			numbers[0] = n1 / numbers[0]
		}
	}

	return numbers[0], true
}
