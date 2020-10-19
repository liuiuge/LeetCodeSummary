package code

// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/

func numPairsDivisibleBy60(time []int) int {
	cnt := 0
	for i := range time {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				cnt += 1
			}
		}
	}
	return cnt
}
