package series

func All(n int, s string) (result []string) {
	if n == len(s) {
		return []string{s}
	}

	if n > len(s) {
		return
	}

	b := []byte(s)
	for i := 0; i < len(s)-n+1; i++ {
		if first, ok := First(n, string(b[i:])); ok {
			result = append(result, first)
		}
	}

	return
}

func UnsafeFirst(n int, s string) string {
	b := []byte(s)

	return string(b[:n])
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	b := []byte(s)

	return string(b[:n]), true
}
