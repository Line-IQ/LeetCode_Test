package funcs

import (
	"math"
	"strconv"
)

func Reverse(x int) int {

	s := strconv.Itoa(x)
	var res string
	if s[0] == 45 {
		res += "-"
		for i := len(s) - 1; i > 0; i-- {
			res += string(s[i])
		}
		result, _ := strconv.Atoi(res)
		if result < math.MinInt32 {
			return 0
		}
		return result
	}
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	result, _ := strconv.Atoi(res)
	if result > math.MaxInt32 {
		return 0
	}
	return result
}
