package code

// https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				if nums[i] == 0 {
					break
				}
			}
		}
	}
	return nums
}
