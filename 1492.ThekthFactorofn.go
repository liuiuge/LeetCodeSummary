package code

// https://leetcode.com/problems/the-kth-factor-of-n/

func kthFactor(n int, k int) int {
	ret, cnt := -1, 0
	fs := make([]int, 0, n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			cnt++
			fs = append(fs, i)
			if cnt == k {
				return i
			}
		}
	}
	total := 2 * cnt
	if fs[cnt-1]*fs[cnt-1] == n {
		total--
	}
	if k <= total {
		return n / fs[total-k]
	}
	return ret
}
