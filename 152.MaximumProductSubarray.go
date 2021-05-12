package code

// https://leetcode.com/problems/maximum-product-subarray/

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ret := nums[0]
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if cur > 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > 0 {
					cur *= nums[j]
				} else if j == 0 {
					break
				} else {
					step := explore(nums[j:])
					if step > 0 {
						for k := 0; k <= step; k++ {
							cur *= nums[j+k]
						}
					}
					break
				}
			}
		} else if cur < 0 {
			step := explore(nums[i:])
			if step > 0 {
				for k := 1; k <= step; k++ {
					cur *= nums[i+k]
				}
			}
		}
		if cur > ret {
			ret = cur
		}
	}
	return ret
}

func explore(x []int) int {
	cnt, step, ofs := 0, 0, 0
	for i := range x {

		if x[i] == 0 {
			break
		}
		ofs++
		if x[i] < 0 {
			cnt++
		}
		if cnt%2 == 0 {
			step = i
		}
	}
	if cnt%2 == 0 {
		step = ofs - 1
	}
	return step
}
