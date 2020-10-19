package code

// https://leetcode.com/problems/h-index-ii/

func hIndexII(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}
	lo, hi := 0, n-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if citations[mid] >= n-mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if n-lo <= citations[lo] {
		return n - lo
	}
	return 0
}
