package code

// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	z, cnt := x^y, 0
	var i uint = 0
	for ; i < 32; i++ {
		cnt += (z >> i) & 1
	}
	return cnt
}
