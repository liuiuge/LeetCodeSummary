package code

// https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		n := nums[i]
		if n < 0 {
			n *= -1
		}
		nums[n] *= -1
		if nums[n] > 0 {
			return n
		}

	}
	return nums[0]
}

func findDuplicatefast(nums []int) int {
	hare := nums[0]     // fast pointer
	tortoise := nums[0] // slow pointer

	slowPtr := nums[tortoise]
	fastPtr := nums[nums[hare]]

	for slowPtr != fastPtr {
		slowPtr = nums[slowPtr]
		fastPtr = nums[nums[fastPtr]]
	}

	aPtr := nums[0]
	bPtr := slowPtr

	for aPtr != bPtr {
		aPtr = nums[aPtr]
		bPtr = nums[bPtr]
	}
	return aPtr
}
