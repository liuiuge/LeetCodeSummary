package code

import (
	"strings"
)

// https://leetcode.com/problems/last-substring-in-lexicographical-order/

func lastSubstring(s string) string {
	subId := []int{}
	rs := []rune(s)
	max_rn := rs[0]
	for _, r := range rs {
		if r > max_rn {
			max_rn = r
		}
	}
	for i, r := range rs {
		if r == max_rn {
			subId = append(subId, i)
		}
	}
	if len(subId) == len(s) {
		return s
	}
	current := string(s[subId[0]:])
	for i := 1; i < len(subId); i++ {
		if strings.Compare(current, string(s[subId[i]:])) < 0 {
			current = string(s[subId[i]:])
		}
	}

	return current
}
