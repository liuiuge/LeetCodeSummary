package code

// https://leetcode.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i, j, idx := 0, len(nums)-1, len(nums)-1
	for i <= j {
		if nums[i]*-1 < nums[j] {
			result[idx] = nums[j] * nums[j]
			j--
		} else {
			result[idx] = nums[i] * nums[i]
			i++
		}
		idx--
	}
	return result
}
