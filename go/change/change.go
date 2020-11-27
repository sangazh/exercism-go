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

// BenchmarkChange-8            804           1474806 ns/op          626932 B/op      23074 allocs/op
func Change1(coins []int, target int) (change []int, err error) {
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

// BenchmarkChange-8           4923            232675 ns/op          379201 B/op       4016 allocs/op
func Change(coins []int, target int) (change []int, err error) {
	if target < 0 {
		return nil, errors.New("invalid target")
	}
	dp := make([]coinChange, target+1, target+1)
	dp[0] = coinChange{
		count: 0,
		coins: []int{},
	}
	for i := 1; i < len(dp); i++ {
		dp[i].count = float64(target + 1)
	}

	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			if i == coin {
				dp[i] = coinChange{
					count: 1,
					coins: []int{coin},
				}
				continue
			}

			if dp[i-coin].count+1 < dp[i].count {
				dp[i].count = dp[i-coin].count + 1
				cc := make([]int, len(dp[i-coin].coins))
				copy(cc, dp[i-coin].coins)
				dp[i].coins = append(cc, coin)
			}
		}
	}

	if int(dp[target].count) == target+1 {
		return nil, errors.New("invalid")
	}
	sort.Ints(dp[target].coins)
	return dp[target].coins, nil
}
