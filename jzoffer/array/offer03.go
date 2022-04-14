package array

// 获取数组中任意一个重复的数字, n >= 2
// 时间/空间复杂度均为 O(n)
func findRepeatNumber(nums []int) int {

	maps := make(map[int]int)

	for _, val := range nums {
		maps[val]++
		if maps[val] > 1 {
			return val
		}
	}

	return -1
}

// 空间复杂度为O(1)的方法
// 首先捋一下思路
// 从下标为0的数字开始，如果a[i]!=i,那么将这个数字与 下标为 a[i] 的数字进行比较，如果两个数值相同，返回true。
// 否则，两者位置交换；之后再次比较a[i]直达a[i] = i的时候，才再次向后移动i
// 如果到了某个位置，a[i]的数值与 下标为a[i]上的数值相同的时候，那么将其进行交换即可
func findRepeatNumber1(nums []int) int {

	// 首先从第0个下标开始循环
	for idx := 0; idx < len(nums); idx++ {

		// 如果当前下标的元素不等于当前下标，就进入循环；
		for nums[idx] != idx {

			// 直到当前下标的数值与以当前数值为下标的数值相等的时候，说明出现了重复的元素，直接返回即可
			if nums[idx] == nums[nums[idx]] {
				return nums[idx]
			}

			// 如果上述判断为false，那么将当前数值放到以其为下标的位置，并将nums[nums[idx]] 的数值放到nums[idx]处
			// swap nums[idx] & nums[nums[idx]]
			nums[idx], nums[nums[idx]] = nums[nums[idx]], nums[idx]

		}
	}

	return -1
}

/*
以下边的数组为例：
[2, 3, 1, 0, 2, 5, 3]

第一轮循环
	循环判断：
		因为 nums[0]=2 != 1
			将 nums[0] 与 nums[2] 比较，发现不相等；交换位置， 此时 nums[0]=1
		因为 nums[0]=1 != 0
			将 nums[0] 与 nums[1] 比较，发现不相等；交换位置， 此时 nums[0]=3
		因为 nums[0]=3 != 0
			将 nums[0] 与 nums[3] 比较，发现不相等；交换位置， 此时 nums[0]=0
		因为 nums[0]=0 == 0
			内部循环结束
第二轮循环
	循环判断：
		因为 nums[1] = 1
			内部循环结束

...

直到，
	因为 nums[4]=2 != 2
		将nums[4] 与 nums[2] 比较，发现相等，返回结果即可







*/
