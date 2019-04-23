package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	if len(divisors) == 0 || limit <= 1 {
		return 0
	}

	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}

	return
}
