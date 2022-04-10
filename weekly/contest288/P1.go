package contest288

import "sort"

//
func largestInteger(num int) int {

	// 将数字拆分成了两个数组 一个奇数、一个偶数
	odd, even, all := split(num)

	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})

	sort.Slice(even, func(i, j int) bool {
		return even[i] > even[j]
	})

	res := 0
	k1, k2 := 0, 0
	for i := len(all) - 1; i >= 0; i-- {
		if all[i]%2 == 0 {
			res = res*10 + even[k1]
			k1++
		} else {
			res = res*10 + odd[k2]
			k2++
		}
	}

	return res
}

func split(num int) ([]int, []int, []int) {
	odd := make([]int, 0, 0)
	even := make([]int, 0, 0)
	all := make([]int, 0, 0)
	for num > 0 {
		tmp := num % 10
		if tmp%2 == 0 {
			even = append(even, tmp)
		} else {
			odd = append(odd, tmp)
		}
		all = append(all, tmp)
		num /= 10
	}

	return odd, even, all
}
