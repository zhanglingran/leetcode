package array

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
*/

/**
使用两个指针，right 向右移动找到一个奇数，然后与left的偶数进行交换即可
快慢指针不好
*/
func exchange(nums []int) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	left, right := 0, 1

	for right < len(nums) && left < right {

		// left 向右移动找到第一个偶数
		for nums[left]%2 == 1 {
			left++
		}

		// right 向右移动找到第一个奇数
		for nums[right]%2 == 0 {
			right++
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right++

	}
	return nums
}

func exchange2(nums []int) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	left, right := 0, len(nums)-1

	for left < right {

		// left 向右移动找到第一个偶数
		for left < len(nums) && nums[left]%2 == 1 {
			left++
		}

		// right 向左移动找到第一个奇数
		for right < len(nums) && nums[right]%2 == 0 {
			right--
		}
		// 此处的判断，可以防止，111222 这种情况的数据出现时，将其转化为 112122
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

/**
进阶3 可扩展的解法
	将 满足xx 的放到左边，满足yy 的放到右边，何解？
	将判断抽象出来一个函数，该函数可以处理以上问题
*/
func exchange3(nums []int, f func(int) bool) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	left, right := 0, len(nums)-1

	for left < right {

		// left 向右移动找到第一个偶数
		for left < len(nums) && f(nums[left]) {
			left++
		}

		// right 向左移动找到第一个奇数
		for right >= 0 && !f(nums[right]) {
			right--
		}
		// 此处的判断，可以防止，111222 这种情况的数据出现时，将其转化为 112122
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

func isEven(a int) bool {
	return a%2 == 1
}

func fakeMain() []int {
	nums := []int{1, 4, 61}
	return exchange3(nums, isEven)
}
