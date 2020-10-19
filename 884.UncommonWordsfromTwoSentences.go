package code

import (
	"strings"
)

// https://leetcode.com/problems/uncommon-words-from-two-sentences/

func uncommonFromSentences(A string, B string) []string {
	listA := strings.Split(A, " ")
	listB := strings.Split(B, " ")
	dict := make(map[string]int, len(listA))
	for _, x := range listA {
		ret, ok := dict[x]
		if !ok {
			dict[x] = 1
		} else {
			dict[x] = ret + 1
		}
	}
	for _, x := range listB {
		ret, ok := dict[x]
		if !ok {
			dict[x] = 1
		} else {
			dict[x] = ret + 1
		}
	}
	result := []string{}
	for k, v := range dict {
		if v > 1 {
			continue
		}
		result = append(result, k)
	}
	return result
}
