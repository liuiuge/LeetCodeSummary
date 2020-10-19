package code

// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/

func sumOddLengthSubarrays(arr []int) int {

	if len(arr) < 1 {
		return 0
	}
	sum := 0
	if len(arr) < 3 {
		for i := range arr {
			sum += arr[i]
		}
		return sum
	}

	for i := 1; i <= len(arr); i += 2 {
		index := 0
		for j := range arr {
			if j < i && j < len(arr)-i+1 {
				index++
			} else if j >= i && j+i-1 >= len(arr) {
				index--
			}

			sum += arr[j] * index
		}
	}
	return sum

}
