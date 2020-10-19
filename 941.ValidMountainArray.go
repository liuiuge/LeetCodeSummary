package code

// https://leetcode.com/problems/valid-mountain-array/

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	up, peak := true, false
	for i := 0; i < len(A); i++ {
		if i != len(A)-1 && A[i+1] > A[i] && up || i != len(A)-1 && A[i+1] < A[i] && !up {
			continue
		} else if i == len(A)-1 {
			if A[i] >= A[i-1] {
				return false
			}
		} else {
			if peak {
				return false
			} else {
				if i != 0 {
					peak = true
					up = false
				} else {
					return false
				}
			}
		}
	}
	return true
}
