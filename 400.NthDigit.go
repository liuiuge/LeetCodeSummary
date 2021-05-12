package code

// https://leetcode.com/problems/nth-digit/

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	bitcount, bitbase, cnt := 1, 1, 0
	temp := n
	for {

		if temp < bitcount*9*bitbase {
			break
		} else {
			cnt += bitcount * 9 * bitbase
			temp -= bitcount * 9 * bitbase
			bitcount++
			bitbase *= 10
		}
	}
	start := bitbase
	tgt, tgtbit := (n-cnt)/(bitcount)+start, (n-cnt)%(bitcount)
	if tgtbit == 0 {
		return (tgt - 1) % 10
	}
	for i := 0; i <= bitcount-tgtbit-1; i++ {
		tgt /= 10
	}
	x := tgt % 10

	return x
}
