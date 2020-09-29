package code

func combine(n int, k int) [][]int {
	temp := make([]int, 0, k)
	result := make([][]int, 0)
	dfs(n, k, 1, &temp, &result)
	return result
}

func dfs(n, k, idx int, temp *[]int, result *[][]int) {
	if len(*temp) == k {
		n_temp := make([]int, k)
		for i, x := range *temp {
			n_temp[i] = x
		}
		*result = append(*result, n_temp)
		return
	}
	for i := idx; i < n; i++ {
		*temp = append(*temp, i)
		dfs(n, k, i+1, temp, result)
		*temp = (*temp)[:len(*temp)-1]
	}
}
