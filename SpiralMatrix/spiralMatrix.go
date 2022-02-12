package solutions

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	result := make([]int, 0, m*n)

	top, bottom, left, right := 0, m-1, 0, n-1

	for len(result) < m*n {
		// go right
		for j := left; j <= right && len(result) < n*m; j++ {
			result = append(result, matrix[top][j])
		}
		// go down
		for j := top + 1; j <= bottom && len(result) < n*m; j++ {
			result = append(result, matrix[j][right])
		}
		// go left
		for j := right - 1; j >= left && len(result) < n*m; j-- {
			result = append(result, matrix[bottom][j])
		}
		// go up
		for j := bottom - 1; j >= top+1 && len(result) < n*m; j-- {
			result = append(result, matrix[j][left])
		}

		left++
		right--
		top++
		bottom--
	}

	return result

}