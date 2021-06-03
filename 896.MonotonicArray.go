package code

// https://leetcode.com/problems/monotonic-array/

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	up, start := true, false
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if start && up != (nums[i] > nums[i-1]) {
				return false
			} else {
				if !start {
					start = true
					up = (nums[i] > nums[i-1])
				}
			}
		}
	}

	return true
}
