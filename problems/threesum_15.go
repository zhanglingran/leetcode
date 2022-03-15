package problems

import (
	"math/rand"
	"sort"
)

/**
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := partitionFun(nums)
	QuickSort(nums[:mid])
	QuickSort(nums[mid+1:])
}

// 返回 pivot 的位置
func partitionFun(nums []int) int {
	i, head := 0, 0
	n := len(nums)
	// 找到一个随机数字
	pivot := rand.Intn(n)
	nums[n-1], nums[pivot] = nums[pivot], nums[n-1]

	// nums[n-1]作为一个pivot
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

// 使用slice操作远比数组要快得多

func ThreeSum(nums []int) [][]int {

	ret := [][]int{}

	sort.Ints(nums)

	for idx, val := range nums {
		if val > 0 {
			return ret
		}
		// 去除重复
		if idx > 0 && val == nums[idx-1] {
			continue
		}

		L := idx + 1
		R := len(nums) - 1
		for L < R {
			if val+nums[L]+nums[R] == 0 {
				ret = append(ret, []int{val, nums[L], nums[R]})

				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if val+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}

	return ret
}
