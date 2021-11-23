package code

// https://leetcode.com/problems/largest-divisible-subset/

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	type cval struct {
		cnt   int
		clist []int
	}
	result := make(map[int]cval, 0)
	tmax, tcnt := 0, 0
	for i := range nums {
		tmp := cval{
			cnt:   1,
			clist: []int{nums[i]},
		}
		maxi, cnt, found := 0, 0, false
		for j := i - 1; j > 0; j-- {
			if nums[i]%nums[j] == 0 {
				if cnt < result[j].cnt {
					maxi = j
					cnt = result[j].cnt
					found = true
				}
			}
		}
		if found {
			tmp.cnt += cnt
			tmp.clist = append(tmp.clist, result[maxi].clist...)
		}
		result[i] = tmp
		if tcnt < result[i].cnt {
			tmax = i
			tcnt = result[i].cnt
		}
	}
	return result[tmax].clist
}
