package code

import "sort"

// https://leetcode.com/problems/sort-an-array/

func sortArray(nums []int) []int {
	// for i := 0; i < len(nums); i++ {
	// 	for j := i; j < len(nums); j++ {
	// 		if nums[j] < nums[i] {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 		}
	// 	}
	// }
	// quickSort(nums, 0, len(nums)-1)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func quickSort(nums []int, lo, high int) {
	if lo >= high {
		return
	}
	mid := parts(nums, lo, high)
	quickSort(nums, lo, mid-1)
	quickSort(nums, mid+1, high)
}

func parts(nums []int, lo, high int) int {
	pivot := nums[high]
	pidx := lo
	for i := lo; i < high; i++ {
		if nums[i] < pivot {
			nums[i], nums[pidx] = nums[pidx], nums[i]
			pidx++
		}
	}
	nums[high] = nums[pidx]
	nums[pidx] = pivot
	return pidx
}
