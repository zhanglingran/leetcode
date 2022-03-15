package problems

import "math"

// 首先是暴力解法
// 以其中一个为起始数，向后遍历找到最大的值记录下来 [超时]
func maxSubArray1(nums []int) int {
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// 思路：从左向右遍历，如果加上当前元素后，sum比当前元素大 sum+=当前元素；如果加上当前元素后 其和比当前元素小，那么sum=当前元素
// 为什么？因为要得到最大的子数组的和，且是向右的，是没有后效行的。
func maxSubArray2(nums []int) int {
	max := math.MinInt32
	sum := 0

	for _, val := range nums {
		sum += val
		if sum < val {
			sum = val
		}
		if sum > max {
			max = sum
		}
	}

	return max
}

// 思路：dp
// 如果要选择当前元素，那么其和为sum + nums[i]; 如果不选，那么其和为nums[i]
// dp[i] 表示以第i个元素结尾的时候，其子序列的和的最大值是多少
func maxSubArray(nums []int) int {
	max := nums[0]
	dp := 0

	for _, val := range nums {
		dp = Max(dp+val, val)
		max = Max(max, dp)
	}

	return max
}
