package binarysearch

import "math/rand"

type Solution struct {
	w    []int
	sums []int
}

func Constructor(w []int) Solution {

	// 计算前缀和
	sums := make([]int, len(w)+1)
	for i := 1; i <= len(w); i++ {
		sums[i] = sums[i-1] + w[i-1]
	}
	return Solution{
		w:    w,
		sums: sums,
	}
}

func (this *Solution) PickIndex() int {

	n := len(this.sums)
	// 生成一个pick元素, 其取值范围在 1-sum 之间
	pick := rand.Intn(this.sums[n-1]) + 1

	// 然后在 sums 中找到 第一个比pick大的元素，并将其位置返回
	left, right := 0, len(this.sums)-1
	for left <= right {
		mid := (left + right) >> 1
		if this.sums[mid] < pick {
			left = mid + 1
		} else if this.sums[mid] == pick {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}

	if left >= n {
		return 0
	} else {
		return left - 1
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
