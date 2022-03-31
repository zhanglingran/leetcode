package problem

/**
首先使用暴力算法解题
因为 n 最大取值为 10^9 不可以使用记忆化搜索
*/
func numberOf2sInRange(n int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		tmp := 0
		x := i
		for x != 0 {
			if x%10 == 2 {
				tmp++
			}
			x = x / 10
		}
	}
	return ans
}
