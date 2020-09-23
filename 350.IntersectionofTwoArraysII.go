package code

func intersect(nums1 []int, nums2 []int) []int {
	short, long := nums1, nums2
	if len(nums1) > len(nums2) {
		short, long = nums2, nums1
	}
	result := make([]int, 0, len(short))
	if len(short) < 1 {
		return result
	}
	check_map := make(map[int][]int, 0)
	for _, val := range short {
		if _, ok := check_map[val]; ok {
			check_map[val][0] = check_map[val][0] + 1
		} else {
			check_map[val] = []int{1, 0}
		}
	}
	for _, val := range long {
		if _, ok := check_map[val]; ok {
			check_map[val][1] = check_map[val][1] + 1
		}
	}
	for k, val := range check_map {
		for i := 0; i < minval(val); i++ {
			result = append(result, k)
		}
	}
	return result
}

func minval(set1 []int) int {
	min := 101
	for _, val := range set1 {
		if min < val {
			min = val
		}

	}
	return min
}

func Intersect(nums1 []int, nums2 []int) []int {
	return intersect(nums1, nums2)
}
