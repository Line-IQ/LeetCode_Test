package funcs

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		maxA := (right - left) * min(height[left], height[right])
		if maxA > maxArea {
			maxArea = maxA
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
