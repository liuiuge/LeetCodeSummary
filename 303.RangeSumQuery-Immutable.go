package code

// https://leetcode.com/problems/range-sum-query-immutable

type NumArray struct {
	n []int
	s []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{
		n: nums,
	}
	tmp := make([]int, len(nums))
	tmp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		tmp[i] = nums[i] + tmp[i-1]
	}
	n.s = tmp
	return n

}

func (n *NumArray) SumRange(i int, j int) int {
	if i == j {
		return n.n[i]
	}

	return n.s[j] - n.s[i] + n.n[i]
}
