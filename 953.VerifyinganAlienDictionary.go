package code

// https://leetcode.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {

	dict := make(map[byte]int, len(order))
	for i, b := range []byte(order) {
		dict[b] = i
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			b := false
			for k := 0; k < len(words[i]) && k < len(words[j]); k++ {
				if dict[words[i][k]] < dict[words[j][k]] {
					b = true
					break
				} else if dict[words[i][k]] == dict[words[j][k]] {
					continue
				} else {
					return false
				}

			}
			if !b {
				if len(words[i]) > len(words[j]) {
					return false
				}
			}
		}
	}
	return true

}
