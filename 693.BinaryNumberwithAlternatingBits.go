package code

// https://leetcode.com/problems/binary-number-with-alternating-bits/

func HasAlternatingBits(n int) bool {
	if n == 1 {
		return true
	}
	x := n % 2
	for n > 1 {
		n /= 2
		d := n % 2
		if x == d {
			return false
		} else {
			x = d
		}
	}
	return x == n
}

func HasAlternatingBits2(n int) bool {
	prev := n & 1

	for n > 1 {
		n = n >> 1
		d := n & 1
		if n&1 == prev {
			return false
		}
		prev = d
	}
	return prev == n
}
