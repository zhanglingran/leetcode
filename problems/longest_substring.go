package problems

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LengthOfLongestSubstring(s string) int {

	set := make(map[byte]bool)
	right := 0
	ret := 0

	for left := 0; left < len(s); left++ {

		// 将right向后移动，判断set中是不是有重复的
		for right < len(s) {
			if set[s[right]] {
				break
			}
			set[s[right]] = true
			right++
		}

		ret = Max(ret, right-left)

		delete(set, s[left])
	}
	return ret
}
