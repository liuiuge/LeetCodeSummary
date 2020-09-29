package code

// https://leetcode.com/problems/number-of-matching-subsequences/

func numMatchingSubseq(S string, words []string) int {
	result := 0
	for _, t := range words {
		if isSubsequence(t, S) {
			result += 1
		}
	}
	return result
}
