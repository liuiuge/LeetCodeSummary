package code

import "sort"

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	pattern := []byte(p)
	sort.Slice(pattern, func(i, j int) bool {
		return pattern[i] < pattern[j]
	})
	result := make([]int, 0, len(s)-len(p))
	for i := 0; i < len(s)-len(p); i++ {
		x := string(s[i : i+len(p)])
		if isAnagramSub(x, string(pattern)) {
			result = append(result, i)
		}
	}
	return result
}

func isAnagramSub(s string, p string) bool {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	return string(bs) == p
}
