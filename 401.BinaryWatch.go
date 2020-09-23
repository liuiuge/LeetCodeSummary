package code

import (
	"strconv"
)

// https://leetcode.com/problems/binary-watch/

func readBinaryWatch(num int) []string {
	hour := make([][]string, 4)
	minute := make([][]string, 6)
	for i := 0; i < 60; i++ {
		ret := strconv.Itoa(i)
		if i < 10 {
			ret = "0" + ret
		}
		cnt := count1(i)

		minute[cnt] = append(minute[cnt], ret)
		if i > 11 {
			continue
		}
		hour[cnt] = append(hour[cnt], ret)
	}

	x_ret := []string{}
	for j := 0; j <= num && j < 4; j++ {
		for _, h := range hour[j] {
			for k, m := range minute[num-j] {
				if k > 5 {
					continue
				}
				x_ret = append(x_ret, h+":"+m)
			}
		}
	}
	return x_ret
}

func count1(n int) int {
	cnt := 0
	for n > 0 {
		cnt += 1
		n = n & (n - 1)
	}
	return n
}
