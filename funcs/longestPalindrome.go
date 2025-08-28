package funcs

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	t := "^"
	for _, ch := range s {
		t += "#" + string(ch)
	}
	t += "#$"

	n := len(t)
	p := make([]int, n)
	center, right := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*center - i // i 关于 center的对称点

		if i < right {
			p[i] = min(right-i, p[mirror])
		}

		for t[i+(1+p[i])] == t[i-(1+p[i])] {
			p[i]++
		}

		if i+p[i] > right {
			center, right = i, i+p[i]
		}
	}

	maxLen, centerIndex := 0, 0
	for i := 1; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	start := (centerIndex - maxLen) / 2

	return s[start : start+maxLen]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BestLongestPalindrome(s string) string {
	var res string
	var length = len(s)
	for i := 0; i < length; i++ {
		s1 := findLongest(s, i, i)
		s2 := findLongest(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func findLongest(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
