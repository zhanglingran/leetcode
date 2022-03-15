package problems

import "math/rand"

// 找到一个数组中第k大的元素
func findKthLargest(nums []int, k int) int {
	k -= 1
	for len(nums) >= 1 {
		idx := getPivotIdx(nums)
		if idx == k {
			return nums[idx]
		} else if idx > k {
			nums = nums[:idx]
		} else if idx < k {
			nums = nums[idx:]
			k -= idx
		}
	}
	return -1
}

func getPivotIdx(nums []int) int {
	n := len(nums)
	pivot := rand.Intn(n)
	nums[n-1], nums[pivot] = nums[pivot], nums[n-1]
	i, head := 0, 0
	for i < n-1 {
		if nums[i] >= nums[n-1] {
			nums[i], nums[head] = nums[head], nums[i]
			head++
		}
		i++
	}
	nums[head], nums[n-1] = nums[n-1], nums[head]
	return head
}
