package bit_operation

func singleNumber(nums []int) int {

	bitmap := make([]int, 32)

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		idx := 0
		for val > 0 {
			if val&1 == 1 {
				bitmap[idx]++
			}
			val >>= 1
			idx++
		}
	}

	ans := 0

	for i := len(bitmap) - 1; i >= 0; i-- {
		ans = ans<<1 + bitmap[i]%3
	}
	return ans
}
