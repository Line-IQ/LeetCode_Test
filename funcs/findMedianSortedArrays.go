package funcs

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	n := len(nums1) + len(nums2)
	nums3 := make([]int, 0, n)
	for {
		if i == len(nums1) {
			nums3 = append(nums3, nums2[j:]...)
			break
		}
		if j == len(nums2) {
			nums3 = append(nums3, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
			continue
		} else {
			nums3 = append(nums3, nums2[j])
			j++
			continue
		}
	}
	if n%2 == 0 {
		half := float64(nums3[n/2] + nums3[n/2-1])
		return half / 2.0
	}
	half := float64(nums3[n/2])
	return half
}
