package code

// https://leetcode.com/problems/insert-interval/

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	if len(intervals) == 0 {
		result = append(result, newInterval)
	} else {
		cur := newInterval
		if cur[1] < intervals[0][0] {
			result = append(result, cur)
			result = append(result, intervals...)
		} else if cur[0] > intervals[len(intervals)-1][1] {
			result = append(intervals, cur)
		} else {
			finish := false
			for i := 0; i < len(intervals); i++ {
				if cur[0] > intervals[i][1] {
					result = append(result, intervals[i])
				} else if cur[1] < intervals[i][0] {
					if !finish {
						result = append(result, cur)
						finish = true
					}
					result = append(result, intervals[i])
				} else {
					if cur[0] > intervals[i][0] {
						cur[0] = intervals[i][0]
					}
					if cur[1] < intervals[i][1] {
						cur[1] = intervals[i][1]
					}
				}
			}
			if !finish {
				result = append(result, cur)
			}
		}
	}
	return result
}
