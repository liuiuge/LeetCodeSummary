package code

// https://leetcode.com/problems/h-index/

import (
	"sort"
)

func hIndex(citations []int) int {
	h := 0
	found := false
	for _, x := range citations {
		if x > 0 {
			found = true
			break
		}
	}
	if !found {
		return h
	}
	h = 1
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			h = len(citations) - i
		}
	}
	return h
}
