package code

// https://leetcode.com/problems/smallest-range-i/

func smallestRangeI(A []int, K int) int {
	if len(A) < 1 {
		return 0
	}
	smallest, biggest := A[0], A[0]
	for _, a := range A {
		if smallest > a {
			smallest = a
		} else if biggest < a {
			biggest = a
		}
	}
	value := biggest - smallest - 2*K
	if value > 0 {
		return value
	}
	return 0

}
