package code

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		nums[abs(nums[i])-1] = -1 * abs(nums[abs(nums[i])-1])
	}
	result := make([]int, 0, len(nums)-1)
	for i := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
