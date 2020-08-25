package code

// https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) < 1 {
			continue
		}
		j := len(matrix[i]) - 1
		if target > (matrix[i][j]) {
			continue
		}
		for ; j >= 0; j-- {
			if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}
