package code

func longestPalindrome(s string) int {
	if len(s) < 1 {
		return 0
	}
	cnt_map := make(map[rune]int, 0)
	for _, x := range []rune(s) {
		if _, ok := cnt_map[x]; !ok {
			cnt_map[x] = 1
		} else {
			cnt_map[x] += 1
		}
	}
	p1 := false
	max_len := 0
	for _, v := range cnt_map {
		if v&1 == 0 {
			max_len += v
		} else {
			if v > 2 {
				max_len += v - 1
			}
			p1 = true
		}
	}
	if p1 {
		max_len += 1
	}
	return max_len
}
