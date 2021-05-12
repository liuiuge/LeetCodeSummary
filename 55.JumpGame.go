package code

// https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	target := len(nums) - 1
	i := target - 1
	for i >= 0 {
		if nums[i] >= target-i {
			target = i
			i = target - 1
		} else {
			i--
		}
	}
	if target == 0 {
		return true
	}
	return false
}
