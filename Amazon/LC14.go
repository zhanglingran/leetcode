package Amazon

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	// 优化点 1：将 strs[0] 提取出来作为变量L（存放在栈中，此时可以防止多次访存）
	pattern := strs[0]

	// 优化点 2： minLen <= 0 直接返回
	for i := 1; i < len(strs) && minLen > 0; i++ {
		ans := 0
		for k := 0; k < min(minLen, len(strs[i])); k++ {
			if strs[i][k] == pattern[k] {
				ans++
			} else {
				break
			}
		}
		if ans < minLen {
			minLen = ans
		}
	}

	return pattern[:minLen]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
