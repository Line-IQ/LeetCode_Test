package funcs

func TwoSum(nums []int, target int) []int {
	for index, value := range nums {
		for j := index + 1; j < len(nums); j++ {
			if value+nums[j] == target {
				return []int{index, j}
			}
		}
	}
	return nil
}

func BestTwoSum(nums []int, target int) []int {
	// 直接用 map 去重
	hash := make(map[int]int)
	// 只用到了 一次循环, 这种循环 比 for range 快
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// 只需判断 是否存在等于 target-num 的 key
		if _, ok := hash[target-num]; ok {
			return []int{i, hash[target-num]}
		}
		// key 为 num, value 为 索引
		hash[num] = i
	}
	return []int{}
}
