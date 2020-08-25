package code

// https://leetcode.com/problems/power-of-four

func isPowerOfFour(n int) bool {

	for n > 1 && n&(n-1) == 0 && n%4 == 0 {
		n /= 4
	}
	return n == 1
}
