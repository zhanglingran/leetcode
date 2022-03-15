package dp

// 最长回文子串

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}

	dp := [1000][1000]bool{}

	maxlen, left := 0, 0

	// 长度为1的时候，全部都是回文串
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	// 枚举长度
	for L := 2; L <= len(s); L++ {
		// 枚举左边的下标
		for i := 0; i < len(s); i++ {
			// 得到右边的下标
			j := i + L - 1
			if j >= len(s) {
				break
			}

			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]

			// 如果最边上的两个不相同，那么该子串必不是回文串
			if s[i] != s[j] {
				dp[i][j] = false

			} else {
				// 若最边上的两个字符相同
				// 如果这两个字符是相邻的，那么必为回文串
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 如果这两个字符是不相邻的，那么其是不是回文串取决于其内部的串是不是回文串
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxlen {
				maxlen = j - i + 1
				left = i
			}

		}
	}

	return s[left : left+maxlen]
}
