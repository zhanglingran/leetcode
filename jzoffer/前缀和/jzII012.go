package 前缀和

func pivotIndex(nums []int) int {

	leftSum := make([]int, len(nums))
	rightSum := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			leftSum[i] = nums[i]
		} else {
			leftSum[i] = leftSum[i-1] + nums[i]
		}
	}

	ans := -1

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightSum = nums[i]
		} else {
			rightSum += nums[i]
		}
		if rightSum == leftSum[i] {
			ans = i
		}
	}

	return ans
}

func pivotIndex1(nums []int) int {

	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	ans := -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum*2+nums[i] == total {
			ans = i
			break
		}
		sum += nums[i]
	}

	return ans
}
