package array

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

分析：方法不可取，case 1 与 case 6 两个实例就说明了 先找行/列的方式是失败的；
另外， 在提取「列」的时候，对缓存「cache」非常的不友好，会造成缓存失效、频繁的读取主存。
*/

// 二分查找 在数组中返回k所在的下标，若不存在，它应该存在的位置
func binarySearch(nums []int, k int) (int, bool) {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == k {
			return mid, true
		}
		if nums[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -(right + 1), false
}

// 循环调用 binarySearch， 最大时间复杂度为 O(nlogn)
func findNumberIn2DArray(matrix [][]int, target int) bool {

	// 非空判断 -- 提高鲁棒性
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	var ok bool

	for _, row := range matrix {
		_, ok = binarySearch(row, target)
		if ok {
			return ok
		}
	}

	return ok
}

// 不断地在右上角向右下角检索
/*
首先定义右上角的位置坐标：
	1、如果右上角的数值 大于 target, 那么其所在的列都将大于 target。
	2、当右上角的数值小于 target，那么待查找的数值必然在该列。
	3、找到列后，开始从上到下排除行。
	4、直到找到或行结束为止。
同样的可以选择左下角为起始点；
*/
// 最大时间复杂度为 O(m+n)
// 缺点：因为涉及到矩阵的随机访存，对cache不友好。 但是内存优化方案没找出来；
func findNumberIn2DArray1(matrix [][]int, target int) bool {

	// 非空判断 -- 提高鲁棒性
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	top, right := 0, len(matrix[0])-1

	for top < len(matrix) && right >= 0 {
		if matrix[top][right] == target {
			return true
		} else if matrix[top][right] > target {
			right--
		} else {
			top++
		}
	}

	return false
}
