package code

// https://leetcode.com/problems/count-square-submatrices-with-all-ones/

func countSquares(matrix [][]int) int {
	result := 0
	x, y := len(matrix[0]), len(matrix)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if j == 0 {
				continue
			}
			result++
			for length := 1; i+length < y && j+length < x; length++ {
				zero := false
				for a := i; a < i+length; a++ {
					for b := j; b < i+length; b++ {
						if matrix[a][b] == 0 {
							zero = true
							break
						}
					}
				}
				if zero {
					break
				}
				result++
			}
		}
	}
	return result
}
