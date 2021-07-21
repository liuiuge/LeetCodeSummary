package code

// https://leetcode.com/problems/shuffle-string/

func restoreString(s string, indices []int) string {
	result := []byte(s)
	for i := range s {
		result[indices[i]] = s[i]
	}
	return string(result)
}
