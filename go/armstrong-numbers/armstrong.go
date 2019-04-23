package armstrong

import (
	"math"
)

func IsNumber(input int) bool {
	if input < 10 {
		return true
	}

	power := float64(DigitCnt(input))
	var armstrong float64
	var num = input
	for {
		d := float64(num % 10)
		armstrong += math.Pow(d, power)
		num /= 10
		if num == 0 {
			break
		}

	}

	return int(armstrong) == input
}

func DigitCnt(num int) (cnt int) {
	if num < 10 {
		return 1
	}

	cnt = 1 + DigitCnt(num/10)
	return cnt
}
