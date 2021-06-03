package code

// https://leetcode.com/problems/unique-number-of-occurrences/

func uniqueOccurrences(arr []int) bool {
	result := make(map[int]int, len(arr))
	for i := range arr {
		ret, ok := result[arr[i]]
		if ok {
			result[arr[i]] = ret + 1
		} else {
			result[arr[i]] = 1
		}
	}
	var bit struct{}
	tmp := make(map[int]struct{}, len(arr))
	for _, v := range result {
		if _, ok := tmp[v]; !ok {
			tmp[v] = bit
		} else {
			return false
		}
	}
	return true
}
