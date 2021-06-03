package code

// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/

func getNoZeroIntegers(n int) []int {
	a := 1
	b := n - 1
	for havingZero(a) || havingZero(b) {
		a++
		b--
	}
	return []int{a, b}
}

func havingZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}
