package code

import (
	"sort"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	// sort.Slice(func(i, j int) bool {
	// 	return nums[i] > nums[j]
	// })
	HeapSort(&nums, k)
	return nums[len(nums)-k]
}

func HeapSort(nums *[]int, k int) {
	for i := len(*nums)/2 - 1; i >= 0; i-- {
		AdjustDown(nums, i, len(*nums))
	}
	for i := len(*nums) - 1; i >= len(*nums)-k; i-- {
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		AdjustDown(nums, 0, i)
	}
}

func AdjustDown(nums *[]int, i, j int) {
	temp := (*nums)[i]
	for k := 2*i + 1; k < j; k = 2*k + 1 {
		if k+1 < j && (*nums)[k] < (*nums)[k+1] {
			k++
		}
		if temp < (*nums)[k] {
			(*nums)[i] = (*nums)[k]
			i = k
		} else {
			break
		}
	}
	(*nums)[i] = temp
}
