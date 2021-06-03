package code

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func numberOfSteps(num int) int {
	step := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
			step++
		} else {
			num--
			step++
		}
	}
	return step
}
