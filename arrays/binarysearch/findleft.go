package binarysearch

// 二分查找：查找最左侧的边界
func binarySearch_leftbound(nums []int, target int) int {

	// 因为 检索区间为 [left, right]
	left, right := 0, len(nums)-1

	// 那么检索暂停条件为 左边大于右边
	for left <= right {

		// 找到中间元素
		mid := (left + right) >> 1
		if nums[mid] < target {
			// 如果中间元素小于target，说明target在右半截，
			left = mid + 1
		} else if nums[mid] == target {
			// 如果 target 与 中间元素相同，那么需要在 left, mid 中间继续检索
			right = mid - 1
		} else {
			right = mid - 1
		}
	}

	// 如果 中间元素等于 target的时候还要将区间向左边移动，那么最终结束的时候，检索区间必然为 [left > right] 那么最终位置必然是 left.
	// 但是，如果target比全部元素还要大的话，那么left必然>right = len(nums) , 即left=len(nums)
	// 如果 target 比全部元素还要下的话，那么返回值将会是 left=0
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 二分查找：查找最右侧的边界
func binarySearch_rightbound(nums []int, target int) int {

	// 因为 检索区间为 [left, right]
	left, right := 0, len(nums)-1

	// 那么检索暂停条件为 左边大于右边
	for left <= right {

		// 找到中间元素
		mid := (left + right) >> 1
		if nums[mid] < target {
			// 如果中间元素小于target，说明target在右半截，
			left = mid + 1
		} else if nums[mid] == target {
			// 如果 target 与 中间元素相同，那么需要在 mid,right 中间继续检索
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 如果 中间元素等于 target的时候还要将区间向右边移动，那么最终结束的时候，检索区间必然为 [left > right] 那么最终位置必然是 right.
	// 但是，如果target比全部元素还要大的话，那么必然会出现left>right = len(nums)的情况 , 此时right=len(nums)-1
	// 如果 target 比全部元素还要下的话，那么right会一直向左边移动，直到 right=-1
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
