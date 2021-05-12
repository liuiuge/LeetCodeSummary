package code

// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/

func maximumSum(arr []int) int {
	if len(arr) < 2 {
		return arr[0]
	}
	mx := arr[len(arr)-1]
	sec := mx
	fst := arr[len(arr)-2] + max([]int{0, arr[len(arr)-1]})
	for i := len(arr) - 3; i >= 0; i-- {
		fst = arr[i] + max([]int{0, fst, sec})
		sec = arr[i+1] + max([]int{0, sec})
		mx = max([]int{mx, fst})
	}
	return mx
}

func max(s1 []int) int {
	x := s1[0]
	for i := range s1 {
		if s1[i] > x {
			x = s1[i]
		}
	}
	return x
}
