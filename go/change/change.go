package change

import (
	"errors"
	"math"
	"sort"
)

type coinChange struct {
	count float64
	coins []int
}

var temp = make(map[int]coinChange)

func Change(coins []int, target int) (change []int, err error) {
	if target == 0 {
		return []int{}, nil
	}
	temp = make(map[int]coinChange)
	cc, err := dp(target, coins)
	if err != nil {
		return nil, err
	}

	sort.Ints(cc.coins)
	return cc.coins, nil
}

func dp(n int, coins []int) (coinChange, error) {
	if v, ok := temp[n]; ok {
		return v, nil
	}

	if n == 0 {
		return coinChange{}, nil
	}

	if n < 0 {
		return coinChange{}, errors.New("invalid")
	}

	res := coinChange{
		count: math.Inf(64),
		coins: make([]int, 0),
	}

	for _, coin := range coins {
		if coin == 4 {

		}
		sub, err := dp(n-coin, coins)
		if err != nil {
			continue
		}

		if res.count > 1+sub.count {
			res.count = 1 + sub.count
			res.coins = append(sub.coins, coin)
		}
	}
	if res.count == math.Inf(64) {
		return coinChange{}, errors.New("no suitable")
	}

	temp[n] = res
	return temp[n], nil
}
