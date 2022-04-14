package dp

func maxSubArray(nums []int) int {
	// 双指针向右移动，如果加上right后，sum增加，那么right继续后移；否则，left右移，直到sum出现新增；

	left, right := 0, 1
	sum := nums[0]
	ans := sum
	for right < len(nums) {
		sum += nums[right]
		if nums[right] < 0 {
			for left < right && sum-nums[left] <= sum {
				sum -= nums[left]
				left++
			}
			sum -= nums[left]
		}
		ans = max(ans, sum)
		right++
	}

	return ans
}
