// https://leetcode.com/problems/unique-paths/
package code

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		result[i][0] = 1
	}
	for i := range result[0] {
		result[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = result[i][j-1] + result[i-1][j]
		}
	}
	return result[m-1][n-1]
}
