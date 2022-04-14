package math

func nthUglyNumber(n int) int {
	// dp[i] 表示第i个丑数

	pointer2, pointer3, pointer5 := 1, 1, 1

	dp := make([]int, 0, n+1)
	dp = append(dp, 1)

	for i := 2; i <= n; i++ {

		val2, val3, val5 := dp[pointer2]*2, dp[pointer3]*3, dp[pointer5]*5

		dp[i] = min(val2, min(val3, val5))
		if dp[i] == val2 {
			pointer2++
		} else if dp[i] == val3 {
			pointer3++
		} else {
			pointer5++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
