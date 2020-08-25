package code

// https://leetcode.com/problems/range-sum-query-2d-immutable

type NumMatrix struct {
	nums  []int
	snums []int
	row   int
	col   int
}

func TDConstructor(matrix [][]int) NumMatrix {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return NumMatrix{}
	}
	row := len(matrix)
	col := len(matrix[0])
	nums := make([]int, 0, row*col)
	snums := make([]int, row*col)
	snums[0] = matrix[0][0]
	x := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			nums = append(nums, matrix[i][j])
			if x > 0 {
				snums[x] = snums[x-1] + matrix[i][j]
			}
			x++
		}
	}
	return NumMatrix{
		nums:  nums,
		snums: snums,
		row:   row,
		col:   col,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.row == 0 || this.col == 0 {
		return 0
	}
	if row1 == row2 && col1 == col2 {
		return this.nums[row1*this.col+col1]
	}
	cnt := 0
	for i := row1; i <= row2; i++ {
		cnt += this.snums[i*this.col+col2] - this.snums[i*this.col+col1] + this.nums[i*this.col+col1]
	}
	return cnt
}
