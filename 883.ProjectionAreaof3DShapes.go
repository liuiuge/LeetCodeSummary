package code

// https://leetcode.com/problems/projection-area-of-3d-shapes/

func projectionArea(grid [][]int) int {
	result := 0
	max_array := make([][]int, 0, 2)
	max_array = append(max_array, make([]int, len(grid)))
	max_array = append(max_array, make([]int, len(grid[0])))
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				result++
			}
			if grid[i][j] > max_array[0][i] {
				max_array[0][i] = grid[i][j]
			}
			if j == len(max_array[1]) {
				max_array[1] = append(max_array[1], grid[i][j])
			} else if max_array[1][j] < grid[i][j] {
				max_array[1][j] = grid[i][j]
			}

		}
	}
	for i := range max_array {
		for j := range max_array[i] {
			result += max_array[i][j]
		}
	}
	return result
}
