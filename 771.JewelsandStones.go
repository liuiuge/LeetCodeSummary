package code

// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	cnt := 0
	if len(J) < 1 {
		return cnt
	}
	x := make(map[rune]struct{}, len(J))
	var check_bit struct{}
	for _, cx := range []rune(J) {
		x[cx] = check_bit
	}
	for _, cx := range []rune(S) {
		if _, ok := x[cx]; ok {
			cnt += 1
		}
	}
	return cnt
}
