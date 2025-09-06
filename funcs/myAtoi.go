package funcs

import (
	"strconv"
	"strings"
	"unicode"
)

func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	if n == 0 {
		return 0
	}
	var res string
	if s[0] == '+' || s[0] == '-' || unicode.IsDigit(rune(s[0])) {
		res += string(s[0])
	} else {
		return 0
	}
	for i := 1; i < n; i++ {
		if unicode.IsDigit(rune(s[i])) {
			res += string(s[i])
		} else {
			break
		}
	}
	final_res, _ := strconv.Atoi(res)

	if final_res < -(1 << 31) {
		return -(1 << 31)
	}
	if final_res > ((1 << 31) - 1) {
		return (1 << 31) - 1
	}

	return final_res
}
