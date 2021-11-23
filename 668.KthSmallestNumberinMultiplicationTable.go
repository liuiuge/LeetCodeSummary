package code

// https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/

func countinrow(n, m, num, k int) bool {
	c := 0
	for i := 1; i <= m; i++ {
		tmp := n
		if tmp > num/i {
			tmp = num / i
		}
		c += tmp
	}
	return c >= k
}

func findKthNumber(m int, n int, k int) int {
	if k == 1 {
		return 1
	}
	if k == m*n {
		return m * n
	}
	lo, hi := 0, m*n
	for lo < hi {
		mid := (hi-lo)/2 + lo
		if countinrow(n, m, mid, k) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}
