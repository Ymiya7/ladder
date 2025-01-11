package kama

func generateMatrix(n int) [][]int {

	rowStart, colStart := 0, 0
	rowEnd, colEnd := n-1, n-1

	// init
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	k := 1
	for {
		// 从左到右
		for i := colStart; i <= colEnd; i++ {
			matrix[rowStart][i] = k
			k++
		}
		rowStart++
		if rowStart > rowEnd {
			break
		}
		// 从上到下
		for i := rowStart; i <= rowEnd; i++ {
			matrix[i][colEnd] = k
			k++
		}
		colEnd--
		if colEnd < colStart {
			break
		}
		// 从右到左
		for i := colEnd; i >= colStart; i-- {
			matrix[rowEnd][i] = k
			k++
		}
		rowEnd--
		if rowEnd < rowStart {
			break
		}
		//从下到上
		for i := rowEnd; i >= rowStart; i-- {
			matrix[i][colStart] = k
			k++
		}
		colStart++
		if colStart > colEnd {
			break
		}
	}
	return matrix
}
