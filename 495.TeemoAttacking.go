package code

// https://leetcode.com/problems/teemo-attacking/

func findPoisonedDuration(timeSeries []int, duration int) int {
	cnt := 0
	if duration == 0 || len(timeSeries) == 0 {
		return cnt
	}
	start, end := timeSeries[0], timeSeries[0]+duration
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] <= end {
			end = timeSeries[i] + duration
		} else {
			cnt += end - start
			start = timeSeries[i]
			end = start + duration
		}
	}
	cnt += end - start
	return cnt
}
