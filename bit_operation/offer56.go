package bit_operation

func singleNumbers(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum ^ nums[i]
	}

	// 此时若sum = 0，说明不存在两个只出现了一次的数
	if sum == 0 {
		return nil
	}

	// 否则按照 sum 中的 某个二进制位，将数组分成两部分
	mode := 1
	for mode&sum == 0 {
		mode <<= 1
	}

	ans1, ans2 := 0, 0

	for i := 0; i < len(nums); i++ {
		// 这里一定要注意：判断是不是mode位是1的时候，与运算后的结果必然是mode本身
		if nums[i]&mode == mode {
			ans2 = ans2 ^ nums[i]
		} else {
			ans1 = ans1 ^ nums[i]
		}
	}

	return []int{ans2, ans1}
}
