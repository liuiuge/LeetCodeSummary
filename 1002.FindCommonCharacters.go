package code

func commonChars(A []string) []string {
	result := []string{}
	if len(A) < 1 {
		return result
	}
	tgt := make([][]int, 26)
	for i := range tgt {
		tgt[i] = make([]int, len(A))
	}
	for i, str := range A {
		for _, alph := range str {
			tgt[alph-'a'][i]++
		}
	}
	for i, list := range tgt {

		for j := 0; j < minval1002(list); j++ {
			result = append(result, string([]rune{rune('a' + i)}))
		}
	}
	return result
}

func minval1002(set1 []int) int {
	min := 101
	for _, val := range set1 {
		if min < val {
			min = val
		}

	}
	return min
}
