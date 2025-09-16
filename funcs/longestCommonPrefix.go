package funcs

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minNum := len(strs[0])
	for i := 1; i < len(strs); i++ {
		num := match(strs[0], strs[i])
		if num < minNum {
			minNum = num
		}
		if minNum == 0 {
			return ""
		}
	}

	return strs[0][:minNum]
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
