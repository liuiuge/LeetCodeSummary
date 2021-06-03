package code

func findSpecialInteger(arr []int) int {
	ret, tmp, max, cnt := arr[0], arr[0], 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == tmp {
			cnt++
		} else {
			if cnt > max {
				max = cnt
				ret = arr[i-1]
			}
			tmp = arr[i]
			cnt = 1
		}
	}
	if cnt > max {
		ret = arr[len(arr)-1]
	}
	return ret
}
