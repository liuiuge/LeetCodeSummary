package code

// https://leetcode.com/problems/house-robber/

func maxAB(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			a = maxAB(a+nums[i], b)
		} else {
			b = maxAB(a, b+nums[i])
		}
	}
	return maxAB(a, b)
}

func Rob(nums []int) int {
	return rob(nums)
}
