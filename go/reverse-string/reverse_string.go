package reverse

func String(s string) (result string) {
	if len(s) == 0 {
		return s
	}

	for _, char := range s {
		result = string(char) + result
	}
	return result
}
