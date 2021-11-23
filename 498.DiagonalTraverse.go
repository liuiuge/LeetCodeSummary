package code

// "fmt"

// https://leetcode.com/problems/diagonal-traverse/

func FindDiagonalOrder(mat [][]int) []int {
	result := make([]int, 0, len(mat[0])*len(mat))
	if len(mat) < 1 {
		return result
	}
	x, y, up := 0, 0, true
	for x < len(mat[0]) && y < len(mat) {
		result = append(result, mat[y][x])
		if up {
			x++
			y--
			if x >= len(mat[0]) {
				x--
				y += 2
				up = false
			} else if y < 0 {
				y = 0
				up = false
			}
		} else {
			x--
			y++
			if y >= len(mat) {
				y--
				x += 2
				up = true
			} else if x < 0 {
				x = 0
				up = true
			}
		}
	}
	return result
}
