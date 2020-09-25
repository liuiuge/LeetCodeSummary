package code

// https://leetcode.com/problems/find-peak-element/

func findPeakElement(nums []int) int {

	if len(nums) < 1 {
		return -1
	}
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			idx = i
		} else {
			idx = i - 1
			break
		}
	}
	return idx
}
