package code

// https://leetcode.com/problems/maximum-product-of-word-lengths/

func maxProduct(words []string) int {
	max_cnt := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if cnt := hasSameLetter(words[i], words[j]); cnt > max_cnt {
				max_cnt = cnt
			}
		}
	}
	return max_cnt
}

func hasSameLetter(w1, w2 string) int {
	s, l := w1, w2
	if len(w1) > len(w2) {
		s, l = w2, w1
	}
	x := make(map[int]struct{}, len(s))
	var check_bit struct{}
	for _, b := range s {
		x[int(b-'a')] = check_bit
	}
	for _, b := range l {
		if _, ok := x[int(b-'a')]; ok {
			return 0
		}
	}
	return len(s) * len(l)
}
