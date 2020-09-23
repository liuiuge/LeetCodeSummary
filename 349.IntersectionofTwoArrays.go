package code

func intersection(nums1 []int, nums2 []int) []int {
	short, long := nums1, nums2
	if len(nums1) > len(nums2) {
		short, long = nums2, nums1
	}
	result := make([]int, 0, len(short))
	if len(short) < 1 {
		return result
	}
	var check_bit struct{}
	check_map := make(map[int]struct{})
	for _, val := range short {
		check_map[val] = check_bit
	}
	for _, val := range long {
		if _, ok := check_map[val]; ok {
			delete(check_map, val)
			result = append(result, val)
		}
	}

	return result
}
