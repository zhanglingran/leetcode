package array

import "math/rand"

func getLeastNumbers(arr []int, k int) []int {
	// 基于快速排序的思想，找到前k个数

	if k == 0 || len(arr) == 0 {
		return []int{}
	}

	left, right := 0, len(arr)

	tmp := k

	for left < right {
		mid := partition(arr[left:right])
		if mid+1 == k {
			break
		} else if mid+1 < k {
			left = mid + 1
			k -= (mid + 1)
		} else {
			right = mid
		}
	}

	return arr[:tmp]
}

func partition(arr []int) int {
	n := len(arr)
	pivot := rand.Intn(n)
	arr[n-1], arr[pivot] = arr[pivot], arr[n-1]
	head := 0

	for i := 0; i < n-1; i++ {
		if arr[i] <= arr[n-1] {
			arr[head], arr[i] = arr[i], arr[head]
			head++
		}
	}

	arr[head], arr[n-1] = arr[n-1], arr[head]

	return head
}
