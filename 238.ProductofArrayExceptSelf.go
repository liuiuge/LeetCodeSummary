package code

// https://leetcode.com/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0], right[len(nums)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		right[len(nums)-1-i] = right[len(nums)-i] * nums[len(nums)-i]
	}
	for i := range result {
		result[i] = left[i] * right[i]
	}
	return result
}
