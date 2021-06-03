package code

// https://leetcode.com/problems/course-schedule/

func canFinish(n int, pre [][]int) bool {
	if len(pre) == 0 || n <= 0 {
		return true
	}
	map_ret := make(map[int][]int, 0)
	for i := range pre {
		if _, ok := map_ret[pre[i][0]]; !ok {
			map_ret[pre[i][0]] = []int{}
		}
		map_ret[pre[i][0]] = append(map_ret[pre[i][0]], pre[i][1])
	}
	visited := make([]bool, n)
	rec := make([]bool, n)
	for i := range map_ret {
		if isCyclic(visited, rec, i, map_ret) {
			return false
		}
	}
	return true
}

func isCyclic(visited, rec []bool, key int, dict map[int][]int) bool {
	if rec[key] {
		return true
	}
	if visited[key] {
		return false
	}
	rec[key] = true
	adjust := []int{}
	if _, ok := dict[key]; ok {
		adjust = dict[key]
	}
	for i := range adjust {
		if isCyclic(visited, rec, i, dict) {
			return true
		}
	}
	rec[key] = false
	visited[key] = true
	return false
}
