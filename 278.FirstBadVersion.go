package code

// https://leetcode.com/problems/first-bad-version/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */

func isBadVersion(version int) bool {
	if version > 3 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	x := n
	if n < 2 {
		return x
	}
	low, high := 1, n
	for low < high {
		mid := (high-low)/2 + low
		if isBadVersion(mid) {
			high = mid
			if !isBadVersion(high - 1) {
				return high
			}
		} else {
			low = mid + 1
			if isBadVersion(low) {
				return low
			}
		}
	}
	return high
}
