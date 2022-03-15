package problems

// FindMedianSortedArraysFun1 求两个有序数组的中位数
// 方法一：归并后计算得到返回值
// 归并后计算，时间复杂度是 O(n+m)
func FindMedianSortedArraysFun1(nums1 []int, nums2 []int) float64 {

	n := len(nums1) + len(nums2)

	var ret [2000]int

	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			ret[k] = nums1[i]
			i++
		} else {
			ret[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		ret[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		ret[k] = nums2[j]
		k++
		j++
	}

	if n%2 == 1 {
		return float64(ret[n/2])
	} else {
		return (float64(ret[n/2]) + float64(ret[n/2-1])) / 2
	}
}

// FindMedianSortedArrays 求两个有序数组的中位数
// 方法二：以O(log (m+n))的时间复杂度进行计算
// 这个问题转换，在两个有序数组中，找到第k个数，这个k即为中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	n := len(nums1) + len(nums2)

	if n%2 == 1 {
		return getKthInTwoArr(nums1, nums2, n/2+1)
	} else {
		return (getKthInTwoArr(nums1, nums2, n/2) + getKthInTwoArr(nums1, nums2, n/2+1)) / 2.0
	}
}

// 获取两个有序数组中第k小的元素
func getKthInTwoArr(nums1 []int, nums2 []int, k int) float64 {

	len1, len2 := len(nums1), len(nums2)
	idx1, idx2 := 0, 0
	for true {
		if idx1 == len1 {
			return float64(nums2[idx2+k-1])
		}
		if idx2 == len2 {
			return float64(nums1[idx1+k-1])
		}
		if k == 1 {
			return float64(min(nums1[idx1], nums2[idx2]))
		}

		half := k / 2
		newIdx1 := min(idx1+half, len1) - 1
		newIdx2 := min(idx2+half, len2) - 1

		pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]

		if pivot1 < pivot2 {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}

	}
	return 0
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
