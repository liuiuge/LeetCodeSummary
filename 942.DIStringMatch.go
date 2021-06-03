package code

// https://leetcode.com/problems/di-string-match/

func diStringMatch(s string) []int {
	if len(s) < 1 {
		return []int{0}
	}
	ret := make([]int, len(s)+1)
	left, right := 0, len(s)
	for i := range s {
		if s[i] == 'I' {
			ret[i] = left
			left++
		} else {
			ret[i] = right
			right--
		}
	}
	ret[len(s)+1] = left
	return ret
}
