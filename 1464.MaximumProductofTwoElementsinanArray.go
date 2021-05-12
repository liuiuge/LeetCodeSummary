package code

// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

func maxProduct1(nums []int) int {
	st, nd := 0, 1
	if nums[st] < nums[nd] {
		st, nd = 1, 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] >= nums[st] {
			nd = st
			st = i

		} else if nums[i] > nums[nd] {
			nd = i
		}
	}
	return (nums[st] - 1) * (nums[nd] - 1)
}
