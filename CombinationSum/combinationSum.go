package solutions

import "sort"

var result [][]int
var candis []int
var tar int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = nil
	candis = candidates
	tar = target

	dfs(0, []int{}, 0)

	return result
}

func dfs(index int, curArr []int, total int) {
	if total == tar {
		result = append(result, append([]int{}, curArr...)) // create copy of curArr
		return
	}
	if index >= len(candis) || total > tar {
		return
	}
	dfs(index, append(curArr, candis[index]), total+candis[index])
	dfs(index+1, curArr, total)
	return
}