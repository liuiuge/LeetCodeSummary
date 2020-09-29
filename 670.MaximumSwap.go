package code

import (
	"sort"
)

// https://leetcode.com/problems/maximum-swap/

func maximumSwap(num int) int {
	list := []int{}
	temp := num
	for num > 0 {
		list = append(list, num%10)
		num /= 10
	}
	sorted_list := make([]int, len(list))
	copy(sorted_list, list)
	sort.Slice(sorted_list, func(i, j int) bool {
		return sorted_list[i] > sorted_list[j]
	})

	max, max_index, front_index := 0, 0, len(list)-1
	for i := 0; i < len(list); i++ {
		if list[len(list)-1-i] == sorted_list[i] {
			continue
		} else {
			front_index = len(list) - 1 - i
			max = sorted_list[i]
			break
		}
	}

	if max == 0 {
		return temp
	} else {
		for i := 0; i < front_index; i++ {
			if list[i] == max {
				max_index = i
				break
			}
		}
		list[max_index], list[front_index] = list[front_index], list[max_index]
	}

	new_num := 0
	for i := len(list) - 1; i >= 0; i-- {
		new_num = 10*new_num + list[i]
	}

	return new_num
}
