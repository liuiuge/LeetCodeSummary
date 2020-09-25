package code

// https://leetcode.com/problems/koko-eating-bananas/

func minEatingSpeed(piles []int, H int) int {
	k, max, total := 1, 1, 0
	for _, x := range piles {
		if x > max {
			max = x
		}
		total += x
	}
	if len(piles) == H {
		k = max
	} else {
		if H < total {
			low, high := 1, max+1
			for low <= high {
				mid := (low + high) / 2
				if sumofhour(piles, mid, H) {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			k = low
		}
	}
	return k
}

func sumofhour(piles []int, k int, h int) bool {
	for _, x := range piles {
		h -= x / k
		if x%k > 0 {
			h -= 1
		}
		if h < 0 {
			return false
		}
	}
	return true
}
