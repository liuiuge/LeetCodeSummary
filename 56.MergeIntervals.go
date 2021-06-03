package code

// https://leetcode.com/problems/merge-intervals/

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0, 1)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if cur[1] >= next[0] {
			if next[1] > cur[1] {
				cur[1] = next[1]
			}
		} else {
			result = append(result, cur)
			cur = next
		}
	}
	result = append(result, cur)
	return result
}

func findEnd(start int, dict map[int]int) int {
	end, ok := dict[start]
	if ok {
		delete(dict, start)
		more := findEnd(end, dict)
		if more > end {
			end = more
		}
	} else {
		end = -1
	}
	return end
}
