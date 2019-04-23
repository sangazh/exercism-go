package diffsquares

import "math"

func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
	}
	result := math.Pow(float64(sum), 2)
	return int(result)
}

func SumOfSquares(n int) int {
	result := 0
	for i := 1; i < n+1; i++ {
		result += i * i
	}
	return int(result)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
