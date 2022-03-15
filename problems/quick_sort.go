package problems

import "math/rand"

func SortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	idx := partition(nums)
	quickSort(nums[:idx])
	quickSort(nums[idx+1:])
}

func partition(nums []int) int {
	i, head := 0, 0
	n := len(nums)
	pivot := rand.Intn(n)
	nums[n-1], nums[pivot] = nums[pivot], nums[n-1]
	for i < n-1 {
		if nums[i] <= nums[n-1] {
			nums[i], nums[head] = nums[head], nums[i]
			head++
		}
		i++
	}
	nums[head], nums[n-1] = nums[n-1], nums[head]
	return head
}
