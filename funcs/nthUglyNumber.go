package funcs

import (
	"fmt"
	"sort"
)

func ThUglyNumber(n int) int {
	hash := make(map[int]int, n)
	nums := []int{2, 3, 5}
	dp := make([]int, 0, n)
	dp = append(dp, 0, 1)
	i := 1
	for {
		if len(dp) >= n*2 {
			sort.Ints(dp)
			fmt.Println(dp)
			return dp[n]
		}
		//dps := cale(dp[i])
		for _, num := range nums {
			dp2 := num * dp[i]
			//fmt.Println(hash)
			if _, ok := hash[dp2]; ok {
				continue
			}
			hash[dp2] = 1
			dp = append(dp, dp2)
		}
		sort.Ints(dp)
		i++
	}
}

func BestNthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	v2, v3, v5 := 1, 1, 1

	ugly := make([]int, n+1)
	p := 1

	for p <= n {
		m := min(min(v2, v3), v5)
		ugly[p] = m
		p++

		if m == v2 {
			v2 = ugly[p2] * 2
			p2++
		}
		if m == v3 {
			v3 = ugly[p3] * 3
			p3++
		}
		if m == v5 {
			v5 = ugly[p5] * 5
			p5++
		}
	}

	return ugly[n]
}
