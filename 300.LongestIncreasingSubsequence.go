package code

// https://leetcode.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	ret := 1
	tmp := make([]int, len(nums))
	tmp[0] = 1
	for i := 1; i < len(nums); i++ {
		tmp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				if tmp[i] < tmp[j]+1 {
					tmp[i] = tmp[j] + 1
				}
			}
		}
		if ret < tmp[i] {
			ret = tmp[i]
		}
	}
	return ret
}

func maxInSlice(x []int) int {
	m := x[0]
	for i := range x {
		if x[i] > m {
			m = x[i]
		}
	}
	return m
}
