package string

func firstUniqChar(s string) byte {

	var ans byte = ' '
	left, right := 0, 0

	for right < len(s) {
		if left != right && s[right] == s[left] {
			left++
			right = left
		} else {
			right++
		}
	}

	ans = s[left]

	return ans

}
