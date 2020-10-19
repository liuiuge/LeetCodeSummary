package code

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

func average(salary []int) float64 {
	var sum float64 = 0

	min, max := salary[0], salary[0]
	for _, x := range salary {
		if min > x {
			min = x
		} else if max < x {
			max = x
		}
		sum += float64(x)
	}
	return (sum - float64(min) - float64(max)) / float64(len(salary)-2)
}
