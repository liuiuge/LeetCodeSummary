package code

// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

func repeatedNTimes(nums []int) int {
	if len(nums) < 4 {
		return 0
	}
	ret := nums[0]
	x := make(map[int]struct{}, 0)
	for i := range nums {
		if _, ok := x[nums[i]]; ok {
			return nums[i]
		}
		x[nums[i]] = struct{}{}
	}
	return ret
}
