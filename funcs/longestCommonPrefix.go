package funcs

func LongestCommonPrefix(strs []string) string {
	n := len(strs)
	res := strs[0]
	minNum := len(res)
	for i := 1; i < n; i++ {
		num := match(res, strs[i])
		if num < minNum {
			res = res[:num]
			minNum = num
		}
	}
	if minNum == 0 {
		return ""
	}
	return res[:minNum]
}

func match(str1, str2 string) int {
	num := 0
	for num < min(len(str1), len(str2)) {
		if str1[num] == str2[num] {
			num++
			continue
		}
		break
	}
	return num
}
