package code

// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

func countSmaller(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		cnt := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				cnt += 1
			}
		}
		result[i] = cnt
	}

	return result
}

func CountSmaller(nums []int) []int {
	return countSmaller((nums))
}
