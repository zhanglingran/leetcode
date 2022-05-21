package printmatrix

// 顺时针遍历矩阵
func spiralOrder(matrix [][]int) []int {

	// 就是按照 右、下、左、上四个顺序遍历数组
	rows, cols := len(matrix), len(matrix[0])

	// 定义四个变量
	left, up, right, down := 0, 0, cols-1, rows-1

	lens := rows * cols
	ans := make([]int, lens)
	cur := 0

	// 一个死循环，里边分别包含了四个循环遍历
	for cur < lens {
		// 从 left -> right; 遍历完成后，up 向下移动一个位置
		if up <= down {
			for i := left; i <= right; i++ {
				ans[cur] = matrix[up][i]
				cur++
			}
			up++
		}

		// 从 up -> down; 遍历完成后，right 向左边移动一个位置
		if left <= right {
			for i := up; i <= down; i++ {
				ans[cur] = matrix[i][right]
				cur++
			}
			right--
		}

		// 从 right -> left; 遍历完成后，down 向上移动一个位置
		if up <= down {
			for i := right; i >= left; i-- {
				ans[cur] = matrix[down][i]
				cur++
			}
			down--
		}

		// 从 down -> up; 遍历完成后，left 向右移动一个位置
		if left <= right {
			for i := down; i >= up; i-- {
				ans[cur] = matrix[i][left]
				cur++
			}
			left++
		}
	}
	return ans
}
