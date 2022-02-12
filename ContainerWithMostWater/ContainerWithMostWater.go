package solutions

func maxArea(height []int) int {
	maxVol := 0
	lIndex := 0
	rIndex := len(height) - 1

	for lIndex != rIndex {
		vol := getMin(height[lIndex], height[rIndex]) * (rIndex - lIndex)
		if vol > maxVol {
			maxVol = vol
		}
		if height[lIndex] > height[rIndex] {
			rIndex--
		} else {
			lIndex++
		}
	}
	return maxVol
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}