package sort

import "fmt"

// 将arr1与arr2合并成一个有序数组
func merge(arr1 []int, arr2 []int) []int {
	arr := []int{}

	len1, len2 := len(arr1), len(arr2)
	i, j := 0, 0
	for i < len1 && j < len2 {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}

	for i < len1 {
		arr = append(arr, arr1[i])
		i++
	}
	for j < len2 {
		arr = append(arr, arr2[j])
		j++
	}

	return arr
}

// MergeSort 归并排序
func MergeSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	arr1 := MergeSort(nums[:mid])
	arr2 := MergeSort(nums[mid:])

	return merge(arr1, arr2)
}

func Foo(arr1 []int, arr2 []int) {
	arr := merge(arr1, arr2)
	for _, elem := range arr {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}
