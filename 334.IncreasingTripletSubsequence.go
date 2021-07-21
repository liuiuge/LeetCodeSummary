package code

// https://leetcode.com/problems/increasing-triplet-subsequence/

func increasingTriplet(nums []int) bool {
	s, m := int(^uint(0)>>1), int(^uint(0)>>1)
	for i := range nums {
		if nums[i] <= s {
			s = nums[i]
		} else if nums[i] <= m {
			m = nums[i]
		} else {
			return true
		}
	}
	return false
}
