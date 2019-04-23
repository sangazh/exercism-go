package pascal

func Triangle(size int) (result [][]int) {
	result = make([][]int, size)
	for i := 0; i < size; i++ {
		pascalMap[i+1] = single(i + 1)
		result[i] = pascalMap[i+1]
	}
	return
}

var pascalMap = map[int][]int{}

func single(n int) (result []int) {
	result = make([]int, n)
	result[0] = 1
	result[n-1] = 1
	if n == 1 {
		return
	}

	result[1] = n - 1
	result[n-2] = n - 1
	if n > 2 {
		last := pascalMap[n-1]
		for i := 2; i < n-2; i++ {
			result[i] = last[i-1] + last[i]
		}
	}

	return
}
