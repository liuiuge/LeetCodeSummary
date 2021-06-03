package code

// https://leetcode.com/problems/maximum-of-absolute-value-expression/

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	res := 0
	signs := [][]int{[]int{1, 1}, []int{-1, 1}, []int{-1, -1}, []int{1, -1}}
	for i := range signs {
		ret := check(signs[i][0], signs[i][1], arr1, arr2)
		if ret > res {
			res = ret
		}
	}
	return res
}

func check(sign1, sign2 int, arr1, arr2 []int) int {
	mi, ma := 1000000, -10000000
	for i := range arr1 {
		cur := arr1[i]*sign1 + arr2[i]*sign2 + i
		if cur > ma {
			ma = cur
		}
		if cur < mi {
			mi = cur
		}
	}
	return ma - mi
}
