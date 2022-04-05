package bit_operation

func myPow(x float64, n int) float64 {

	absN := n
	if n < 0 {
		absN = -n
	}

	var ans float64 = 1.0

	for absN > 0 {
		if absN&1 == 1 {
			ans = ans * x
		}
		x = x * x
		absN >>= 1
	}

	if n > 0 {
		return ans
	}
	return 1 / ans
}
