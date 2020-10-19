package code

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/

func luckyNumbers(matrix [][]int) []int {
	x := make([]int, 0, len(matrix))
	if len(matrix) < 1 {
		return x
	}
	skip_map := make([]int, len(matrix[0]))
	for i := range matrix {
		min_i, min_idx := matrix[i][0], 0
		for j := range matrix[i] {
			if min_i > matrix[i][j] {
				min_i = matrix[i][j]
				min_idx = j
			}
		}
		if skip_map[min_idx] > 0 {
			continue
		}
		is_lucky := true
		k := 0
		for ; k < len(matrix); k++ {
			if matrix[k][min_idx] > min_i {
				is_lucky = false
				break
			}
		}
		if is_lucky {
			x = append(x, min_i)
			skip_map[min_i] = 1
		}
	}
	return x
}
