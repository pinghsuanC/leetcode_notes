package solutions

func insert(intervals [][]int, newInterval []int) [][]int {
	length := len(intervals)
	// insert interval
	if length == 0 {
		return append(intervals, newInterval)
	}

	for i := 0; i < length; i++ {
		if intervals[i][0] >= newInterval[0] {
			// insert in the front
			tmp := intervals[:i]
			tmp1 := intervals[i:]
			intervals = append(append(append([][]int{}, tmp...), newInterval), tmp1...)
			break
		}
		if i == length-1 {
			intervals = append(intervals, newInterval)
		}
	}

	return mergeIntervals(intervals)
}

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	cur := intervals[0]
	result := [][]int{}

	for i := 1; i < len(intervals); i++ {
		first := intervals[i][0]
		last := intervals[i][1]
		if cur[1] >= first {
			// merge
			cur[1] = getMax(cur[1], last)
		} else {
			// can't merge
			result = append(result, cur)
			cur = intervals[i]
		}
	}

	return append(result, cur)
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}