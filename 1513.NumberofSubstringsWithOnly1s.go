package code

// https://leetcode.com/problems/number-of-substrings-with-only-1s/

func numSub(s string) int {
	ret := make(map[int]int)
	tmp := 0
	for i := range []byte(s) {
		if s[i] == '1' {
			tmp += 1
		} else {
			if tmp > 0 {
				if _, ok := ret[tmp]; ok {
					ret[tmp] += 1
				} else {
					ret[tmp] = 1
				}
				tmp = 0
			}
		}
	}
	if tmp > 0 {
		if _, ok := ret[tmp]; ok {
			ret[tmp] += 1
		} else {
			ret[tmp] = 1
		}
		tmp = 0
	}
	retcnt := 0
	for k, v := range ret {
		retcnt += v * k * (k + 1) / 2
	}
	return retcnt % 1000000007
}
