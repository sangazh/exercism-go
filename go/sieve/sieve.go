package sieve

func Sieve(limit int) []int {
	result := make([]int, 0)
	if limit < 2 {
		return result
	}
	primes := make([]int, 0)
	lists := makeList(limit)

	prime := 2
	list := genList(prime, lists)

	for i := 0; i <= limit-1; i++ {
		primes = append(primes, prime)
		if len(list) == 0 {
			break
		}

		prime, list = list[0], list[1:]
		list = genList(prime, list)
	}

	return primes
}

func genList(prime int, lists []int) []int {
	list := make([]int, 0)
	for _, n := range lists {
		if n%prime == 0 {
			continue
		}
		list = append(list, n)
	}
	return list
}

func makeList(limit int) []int {
	lists := make([]int, 0)
	for i := 2; i <= limit; i++ {
		lists = append(lists, i)
	}

	return lists
}
