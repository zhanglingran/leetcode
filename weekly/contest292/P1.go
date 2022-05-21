package contest292

func largestGoodInteger(num string) string {

	// 双指针，fast,slow 两者之间差3，如果窗口满足条件则前后+=3，否则前后+=1.
	// 使用 ans = slow 来记录

	ans, slow, fast := -1, 0, 2

	for fast < len(num) {
		if num[slow] == num[slow+1] && num[slow] == num[fast] {
			if ans < 0 {
				ans = slow
			} else {
				if num[ans] < num[slow] {
					ans = slow
				}
			}
			slow += 3
			fast += 3
		} else {
			slow++
			fast++
		}
	}

	if ans < 0 {
		return ""
	} else {
		return num[ans : ans+3]
	}
}

// "6777133339"
