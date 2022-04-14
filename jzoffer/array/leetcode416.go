package array

/**
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

思路：
	计算和，如果之和为奇数那么肯定不可分
如果之和为偶数，那么就是计算数组中是不是存在一个子集和，使子集的和为 sum/2

nums 的长度不大于 200，可以使用深度优先+剪枝的方法进行检索
也可以使用dp中背包问题的变形（是否可以装满的情况）
*/
func canPartition(nums []int) bool {

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	// 在 nums 找到是不是满足子集之和为 target 的（背包装满的问题）
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}

	dp[0][0] = true
	for i := 1; i < len(nums); i++ {
		w := nums[i]
		for j := 1; j <= target; j++ {
			if j >= w {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-w]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
