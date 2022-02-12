package solutions

func maxSubArray(nums []int) int {
	max := -9223372036854775808 // minimun vakue if int64
	sum := 0
	for _, val := range nums {
		sum = sum + val
		max = getMax(max, sum)

		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}