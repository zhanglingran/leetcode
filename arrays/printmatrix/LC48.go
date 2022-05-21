package printmatrix

import "fmt"

/**
顺时针旋转图像
*/
func rotate(matrix [][]int) {
	// 首先对矩阵转置
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 其次对矩阵做镜像（按列）
	for i := 0; i < n; i++ {
		l, r := 0, n-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
	//fmt.Println(matrix)
}

/**
举一反三：逆时针针旋转图像
	转置+按行做镜像
	反转置+按列做镜像
*/
func reRotate(matrix [][]int) {
	// 首先对矩阵反转置（按照副对角线进行对称交换）
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}

	// 其次对矩阵做镜像（按列）
	for i := 0; i < n; i++ {
		l, r := 0, n-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
	fmt.Println(matrix)
}
