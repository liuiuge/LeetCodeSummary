package code

// https://leetcode.com/problems/count-numbers-with-unique-digits/

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	}
	result := 10
	start := 9
	for i := 1; i < n; i++ {
		start *= (10 - i)
		result += start
	}
	return result
}
