package code

// https://leetcode.com/problems/shifting-letters/

func shiftingLetters(S string, shifts []int) string {

	shiftList := make([]int, len(S))
	bs := []byte(S)
	base := 0
	for i := len(shifts) - 1; i >= 0; i-- {
		base += shifts[i]
		if len(shiftList) > i {
			shiftList[i] += base
			shiftList[i] += int(bs[i]) - 97
			shiftList[i] = shiftList[i] % 26
			bs[i] = byte(shiftList[i] + 97)
		}
	}
	return string(bs)
}
