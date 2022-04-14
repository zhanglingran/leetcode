package divide_and_conquer

import "fmt"

// 剑指offer 51题 : 数组中的逆序对

func reversePairs(nums []int) int {

	ans := 0
	_, ans = mergeSort(nums)
	return ans
}

func mergeSort(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}

	mid := len(nums) / 2
	arr1, c1 := mergeSort(nums[:mid])
	arr2, c2 := mergeSort(nums[mid:])
	arr, c := merge(arr1, arr2)
	return arr, c + c1 + c2
}

func merge(arr1, arr2 []int) ([]int, int) {

	count := 0

	ans := []int{}

	i, j := len(arr1)-1, len(arr2)-1
	for {
		if arr1[i] > arr2[j] {
			ans = append(ans, arr1[i])
			i--
			count = count + j + 1
		} else {
			ans = append(ans, arr2[j])
			j--
		}
		if i < 0 || j < 0 {
			break
		}
	}

	for i >= 0 {
		ans = append(ans, arr1[i])
		i--
	}

	for j >= 0 {
		ans = append(ans, arr2[j])
		j--
	}

	reverse(ans)

	return ans, count
}

// 数组逆转
func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func Foo(arr1 []int, arr2 []int) {
	arr, c := merge(arr1, arr2)
	fmt.Println("c is ", c)
	for _, elem := range arr {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}
