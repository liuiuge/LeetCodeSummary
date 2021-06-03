package code

//  https://leetcode.com/problems/integer-replacement/

func integerReplacement(n int) int {
	cnt := 0
	for n > 1 {
		if n%2 == 1 {
			if (n-1)%4 == 0 && n != 3 {
				n = n - 1
			} else {
				n = n + 1
			}
			cnt++
		} else {
			n /= 2
			cnt++
		}
	}
	return cnt
}
