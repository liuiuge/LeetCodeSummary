package code

// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/

func kLengthApart(nums []int, k int) bool {
	if k < 1 {
		return true
	}
	min, cnt, start := -1, 0, false
	for i := range nums {
		if nums[i] == 0 && start {
			cnt++
		} else {
			if start && nums[i-1] == 1 {
				min = 0
			} else {
				start = true
			}
			if cnt > 0 {
				if cnt < min || min == -1 {
					min = cnt
				}
			}
			cnt = 0
		}
	}
	if min == -1 && cnt > 0 {
		return true
	}
	return min >= k
}
