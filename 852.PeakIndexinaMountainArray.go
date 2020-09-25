package code

// https://leetcode.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		if arr[mid+1] < arr[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
