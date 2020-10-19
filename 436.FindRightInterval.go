package code

import (
	"sort"
)

// https://leetcode.com/problems/find-right-interval/

func findRightInterval(intervals [][]int) []int {
	index := make([]int, len(intervals))
	if len(index) < 1 {
		return index
	}
	leftpoint := make([][]int, len(intervals))
	for i, inter := range intervals {
		leftpoint[i] = []int{inter[0], i}
	}
	sort.Slice(leftpoint, func(i, j int) bool {
		return leftpoint[i][0] < leftpoint[j][0]
	})
	for i, inter := range intervals {
		idx := -1
		for _, point := range leftpoint {
			if point[0] > inter[1] {
				idx = point[1]
				break
			}
		}
		index[i] = idx
	}
	return index
}

func findRightIntervalBS(intervals [][]int) []int {
	rst := make([]int, len(intervals))
	a := make([][]int, len(intervals))

	for i, interval := range intervals {
		a[i] = []int{interval[0], i}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	for i, interval := range intervals {
		// bin-search for lower bound
		t := interval[1]
		l, r := 0, len(a)-1
		for l <= r && l < len(a) {
			m := l + (r-l)/2
			mt := a[m][0]
			if mt < t {
				l = m + 1
			} else {
				// mt >= t
				r = m - 1
			}
		}
		if l >= len(a) {
			rst[i] = -1
		} else {
			rst[i] = a[l][1]
		}
	}

	return rst
}
