package code

// https://leetcode.com/problems/single-number-ii/

func singleNumber(nums []int) int {
	for i := 0; i < len(nums)-3; {

		if nums[i+1] != nums[i] {
			tmp := i + 1
			for j := tmp; j < len(nums); j++ {
				if nums[j] == nums[i] {
					nums[i+1], nums[j] = nums[j], nums[i+1]
					tmp = j
					break
				}
			}
			if tmp == i+1 {
				return nums[i]
			}
		}
		if nums[i+2] != nums[i] {
			for j := i + 3; j < len(nums); j++ {
				if nums[j] == nums[i] {
					nums[i+2], nums[j] = nums[j], nums[i+2]
					break
				}
			}
		}
		i += 3
	}
	return nums[len(nums)-1]
}
