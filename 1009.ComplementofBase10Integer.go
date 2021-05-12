package code

// https://leetcode.com/problems/complement-of-base-10-integer/

import "fmt"

func bitwiseComplement(N int) int {
	var start uint = 0
	for N > 1<<(start+1) {
		start++
	}

	fmt.Println(start)
	n := 0
	for {
		bitval := (1 << start) & N
		fmt.Println(bitval)
		n |= (bitval ^ (1 << start))

		if start == 0 {
			break
		}
		start--
	}
	return n
}

func BitwiseComplement(N int) int {
	return bitwiseComplement(N)
}
