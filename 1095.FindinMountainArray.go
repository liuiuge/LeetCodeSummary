package code

// https://leetcode.com/problems/find-in-mountain-array/

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	if index > len(this.arr)-1 {
		return 0
	}
	return this.arr[index]
}

func (this *MountainArray) length() int {
	return len(this.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	idx, length := -1, mountainArr.length()
	low, high := 0, length-1
	for low < high {
		mid := (high + low) / 2
		temp1, temp2 := mountainArr.get(mid+1), mountainArr.get(mid)
		if mountainArr.get(mid+1) < mountainArr.get(mid) {
			high = mid
		} else {
			low = mid + 1
		}
		if temp1 == target {
			idx = mid + 1
		}
		if temp2 == target {
			idx = mid
		}
	}
	if idx > high || idx < 0 {
		low = 0
		temp := high
		for low < temp {
			mid := (temp + low) / 2
			temp1 := mountainArr.get(mid)
			if temp1 < target {
				low = mid + 1
			} else if temp1 == target {
				idx = mid
				break
			} else {
				temp = mid
			}
		}
	}
	if idx < 0 {
		low = high
		temp := length - 1
		for low <= temp {
			mid := (temp + low) / 2
			temp1 := mountainArr.get(mid)
			if temp1 < target {
				temp = mid - 1
			} else if temp1 == target {
				idx = mid
				break
			} else {
				low = mid + 1
			}
		}
	}

	return idx
}
