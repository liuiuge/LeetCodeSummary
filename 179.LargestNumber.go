package code

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/largest-number/

func largestNumber(nums []int) string {

	retList := make([]string, len(nums))
	for i, num := range nums {
		retList[i] = strconv.Itoa(num)
	}
	sort.Slice(retList, func(i, j int) bool {
		return retList[i]+retList[j] > retList[j]+retList[i]
	})
	fmt.Println(retList)
	var zero string = "0"
	if string(retList[0][0]) == zero {
		return zero
	}
	return strings.Join(retList, "")
}

func LargestNumber(nums []int) string {
	return largestNumber(nums)
}
