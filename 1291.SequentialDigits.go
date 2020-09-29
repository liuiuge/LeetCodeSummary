package code

import (
	"strconv"
)

// https://leetcode.com/problems/sequential-digits/

func sequentialDigits(low int, high int) []int {
	prepared := "123456789"
	result := []int{}
	for i := 2; i < 9; i++ {

		m := len(prepared) - i
		for j := 0; j < m; j++ {
			min_i, _ := strconv.Atoi(string(prepared[j : j+i]))
			if min_i < low {
				continue
			}
			if min_i > high {
				break
			}
			result = append(result, min_i)
		}

	}
	return result
}
