package solution

func rotate(matrix [][]int) {
	var tm int
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		for m, n := i, j; m < j; m, n = m+1, n-1 {
			tm = matrix[i][m]
			matrix[i][m] = matrix[n][i]
			matrix[n][i] = matrix[j][n]
			matrix[j][n] = matrix[m][j]
			matrix[m][j] = tm
		}
	}
}