package array

/*
剑指 Offer 29. 顺时针打印矩阵
	规定好行进顺序后，直接使用深搜？
*/

func spiralOrder(matrix [][]int) []int {

	res := make([]int, 0, 0)
	if len(matrix) == 0 {
		return res
	}
	rows, cols := len(matrix), len(matrix[0])

	startX, endX := 0, rows-1
	startY, endY := 0, cols-1

	for true {

		for i := startY; i <= endY; i++ {
			res = append(res, matrix[startX][i])
		}
		startX++
		if startX > endX {
			break
		}

		for i := startX; i <= endX; i++ {
			res = append(res, matrix[i][endY])
		}
		endY--
		if startY > endY {
			break
		}

		for i := endY; i >= startY; i-- {
			res = append(res, matrix[endX][i])
		}
		endX--
		if startX > endX {
			break
		}

		for i := endX; i >= startX; i-- {
			res = append(res, matrix[i][startY])
		}
		startY++
		if startY > endY {
			break
		}

	}

	return res
}
