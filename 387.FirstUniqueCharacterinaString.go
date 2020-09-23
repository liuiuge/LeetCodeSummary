package code

// https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	uniq_dict := make(map[rune]int, len(s))
	r_s := []rune(s)
	for i := range r_s {
		_, ok := uniq_dict[r_s[i]]
		if ok {
			uniq_dict[r_s[i]]++
		} else {
			uniq_dict[r_s[i]] = 1
		}
	}
	min_v := -1
	for i := range r_s {
		if uniq_dict[r_s[i]] == 1 {
			return i
		}
	}
	return min_v
}
