package code

// https://leetcode.com/problems/reverse-bits

func reverseBits(num uint32) uint32 {
	num2, mask := 0, 1<<31
	for i := 0; i < 32; i++ {
		idx := uint(i)
		mid := ((int(num) & mask) >> (31 - idx)) << idx
		num2 |= mid
		mask >>= 1
	}
	return uint32(num2)

}

func ReverseBits(num uint32) uint32 {
	return reverseBits(num)
}
