package solutions

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	cur := intervals[0]
	result := [][]int{}

	for i := 1; i < len(intervals); i++ {
		nextStart := intervals[i][0]
		nextEnd := intervals[i][1]
		curEnd := cur[1]
		if curEnd >= nextStart {
			// continue searching for the end
			cur[1] = getMax(cur[1], nextEnd)
		} else {
			// append cur to answer and update cur to current position
			result = append(result, cur)
			cur = intervals[i]
		}
	}

	return append(result, cur)
}

func canMerge(a []int, b []int) bool {
	return !(a[1] < b[0])
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}