package yacht

func Score(dice []int, category string) int {
	diceMap := handleDict(dice)
	switch category {
	case "yacht":
		if len(diceMap) == 1 {
			return 50
		}
	case "ones":
		if v, ok := diceMap[1]; ok {
			return v
		}
	case "twos":
		if v, ok := diceMap[2]; ok {
			return v * 2
		}
	case "threes":
		if v, ok := diceMap[3]; ok {
			return v * 3
		}
	case "fours":
		if v, ok := diceMap[4]; ok {
			return v * 4
		}
	case "fives":
		if v, ok := diceMap[5]; ok {
			return v * 5
		}
	case "sixes":
		if v, ok := diceMap[6]; ok {
			return v * 6
		}
	case "full house":
		if len(diceMap) != 2 {
			break
		}
		for _, count := range diceMap {
			if count == 1 || count == 4 {
				return 0
			}
		}

		return sum(dice)
	case "four of a kind":
		if len(diceMap) > 2 {
			break
		}
		for n, count := range diceMap {
			if count >= 4 {
				return n * 4
			}
		}
	case "little straight":
		if len(diceMap) == 5 {
			if _, ok := diceMap[6]; !ok {
				return 30
			}
		}
	case "big straight":
		if len(diceMap) == 5 {
			if _, ok := diceMap[1]; !ok {
				return 30
			}
		}
	case "choice":
		return sum(dice)
	}

	return 0
}

func sum(nums []int) int {
	n := 0
	for _, num := range nums {
		n += num
	}
	return n
}

func handleDict(nums []int) map[int]int {
	target := make(map[int]int)
	for _, num := range nums {
		if _, ok := target[num]; ok {
			target[num] += 1
		} else {
			target[num] = 1
		}
	}
	return target
}
