package code

// https://leetcode.com/problems/missing-number/

func missingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] != i {
			if nums[i] == len(nums) {
				nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
				nums = nums[:len(nums)-1]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		} else {
			i++
		}
	}
	return i
}

func missingNumber2(nums []int) int {
	c := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		c -= nums[i]
	}
	return c
}
