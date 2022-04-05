package contest287

/**
给你一个 下标从 0 开始 的整数数组 candies 。数组中的每个元素表示大小为 candies[i] 的一堆糖果。
你可以将每堆糖果分成任意数量的 子堆 ，但 无法 再将两堆合并到一起。
另给你一个整数 k 。你需要将这些糖果分配给 k 个小孩，使每个小孩分到 相同 数量的糖果。
每个小孩可以拿走 至多一堆 糖果，有些糖果可能会不被分配。
返回每个小孩可以拿走的 最大糖果数目。
*/

// 效率最高的代码：使用均值和最大值进行优化
func maximumCandies(candies []int, k int64) int {

	// 1、找到最大值
	var sum int64
	var max int
	for i := 0; i < len(candies); i++ {
		sum += int64(candies[i])
		if max < candies[i] {
			max = candies[i]
		}
	}

	if sum <= k {
		return int(sum / k)
	}

	// 使用均值与最大值对比后，对右端点进行优化
	mean := sum / k
	if mean > int64(max) {
		mean = int64(max)
	}

	left, right := int64(0), mean
	// 对 mean 使用 二分优化
	for left < right {

		// +1的目的是为了防止程序陷入死循环，如出现 5+6/2=5, 而5又是最优解
		m := (left + right + 1) >> 1

		var count int64
		for i := 0; i < len(candies); i++ {
			count += int64(candies[i]) / m
			// 判断放入循环内，节省不必要的后续计算
			if count >= k {
				left = m
				break
			}
		}

		if count < k {
			right = m - 1
		}
	}

	return int(left)
}

// 最简洁的写法； 省去不必要的判断，直接将最后边的值设置为 candies[i]的最大值
func maximumCandies1(candies []int, k int64) int {

	left, right := 0, int(1e7)
	for left < right {

		// +1的目的是为了防止程序陷入死循环，如出现 5+6/2=5, 而5又是最优解
		m := (left + right + 1) >> 1

		var count int
		for i := 0; i < len(candies); i++ {
			count += candies[i] / m
		}

		if int64(count) >= k {
			left = m
		} else {
			right = m - 1
		}
	}

	return left
}
