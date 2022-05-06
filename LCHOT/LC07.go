package LCHOT

import "math"

func reverse(x int) int {

	ans := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}

	ans *= sign
	if ans < math.MinInt32 {
		return 0
	}
	if ans > math.MaxInt32 {
		return 0
	}

	return ans
}
