package brackets

func Bracket(input string) (bool, error) {
	brackets := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	stack := make([]string, 0)
	for _, s := range input {
		switch string(s) {
		case "{", "[", "(":
			stack = append(stack, string(s))
		case "}", "]", ")":
			if len(stack) == 0 {
				return false, nil
			}

			var left string
			left, stack = pop(stack)
			if right, ok := brackets[left]; ok {
				if right != string(s) {
					return false, nil
				}
			}
		}
	}

	if len(stack) > 0 {
		return false, nil
	}

	return true, nil
}

func pop(slice []string) (string, []string) {
	last := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return last, slice
}
