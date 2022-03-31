package dp

/**
剪绳子问题：给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1].
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

思路：
	dp[n] = max( max(dp[i], i) * max(dp[n-i], n-i)) i = 1->n/2
*/
func cuttingRope(n int) int {

	dp := make([]int, 0, n+1)
	dp = append(dp, 0, 0, 1)

	for i := 3; i <= n; i++ {
		ans := 0
		for j := 1; j <= i/2; j++ {
			ans = max(max(dp[j], j)*max(dp[n-j], n-j), ans)
		}
		dp = append(dp, ans)
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
参考剑指offer题解
*/
func cuttingRopeII(n int) int {

	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, 0, n+1)
	dp = append(dp, 0, 1, 2, 3)

	for i := 4; i <= n; i++ {
		tmp := 0
		for j := 1; j <= i/2; j++ {
			tmp = maxII(dp[i]*dp[i-j], tmp)
		}
		dp = append(dp, tmp)
	}
	return dp[n]
}

func maxII(a, b int) int {
	if a > b {
		return a
	}
	return b
}
