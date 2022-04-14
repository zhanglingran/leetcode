package dp

/**
在上题的基础上将 n 的上限设置为 1000， 并对最终的结果进行取模运算
通过数学证明可以知道，当每段的长度为 e 的时候，可以截取得到最长；其实可以得到一个数学公式
假设每段的长度为 x，那么截成 a 段一共的长度为 x^a, 其中 a = n/x
然后得到一个公式 y = x^(1/x) 求导计算极大值得到 x = e 的时候取最大值；
然后比较 x = 2 或 3 的时候，其最终的结果是多少？可知道 x = 3 的时候，取得max
*/
func cuttingRopeIII(n int) int {

	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	const MOD int = 1e9 + 7

	ans := 1
	for n > 4 {
		ans = (ans * 3) % MOD
		n -= 3
	}

	return (ans * n) % MOD
}
