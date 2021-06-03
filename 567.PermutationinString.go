package code

// https://leetcode.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	pattern_map := make(map[byte]int, 0)
	bs1 := []byte(s1)
	for i := range bs1 {
		if _, ok := pattern_map[bs1[i]]; ok {
			pattern_map[bs1[i]] += 1
		} else {
			pattern_map[bs1[i]] = 1
		}
	}
	bs2 := []byte(s2)
	l, r, count := 0, 0, 0
	for r < len(bs2) {
		if pattern_map[bs2[r]] > 0 {
			count++
		}
		pattern_map[bs2[r]]--
		if count == len(s1) {
			return true
		}
		r++
		if r-l > len(s1)-1 {
			if pattern_map[bs2[l]] >= 0 {
				count--
			}
			pattern_map[bs2[l]]++
			l++
		}
	}
	return false
}
