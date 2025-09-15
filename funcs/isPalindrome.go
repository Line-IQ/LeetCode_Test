package funcs

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	n := len(str)
	l := n / 2

	for i := 0; i < l; i++ {
		if str[i] == str[n-i-1] {
			continue
		}
		return false
	}
	return true
}

func BetterIsPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == x || x == reversed/10
}
