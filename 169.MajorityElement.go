package code

func majorityElement(nums []int) int {
	x, cnt := 0, 0
	for _, elem := range nums {
		if cnt == 0 {
			x = elem
			cnt++
		} else if x != elem {
			cnt--
		} else {
			cnt++
		}
	}
	return x
}

func MajorityElement(nums []int) int {
	return majorityElement(nums)
}
