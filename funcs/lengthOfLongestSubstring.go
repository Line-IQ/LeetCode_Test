package funcs

import "fmt"

func LengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	var str string
	var finalStr string
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; !ok {
			hash[s[i]] = i + 1
			str += string(s[i])
			continue
		}
		if len(str) > len(finalStr) {
			finalStr = str
			fmt.Println(finalStr)
		}
		str = s[hash[s[i]] : i+1]
		for k := range hash {
			if hash[k] < hash[s[i]] {
				delete(hash, k)
			}
		}
		hash[s[i]] = i + 1
	}
	if len(str) > len(finalStr) {
		fmt.Println(str)
		return len(str)
	}
	return len(finalStr)
}

func BestLengthOfLongestSubstring(s string) int {
	cacheMap := make(map[byte]int)
	l := 0
	ans := 0
	for i := range s {
		if index, ok := cacheMap[s[i]]; ok {
			if index+1 > l {
				l = index + 1
			}
		}
		cacheMap[s[i]] = i
		ans = max(ans, i-l+1)
	}
	return ans
}
