package array

func constructArr(a []int) []int {

	n := len(a)
	// 从前向后
	L := make([]int, n-1)
	// 从后向前
	R := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		if i == 0 {
			L[i] = a[i]
		} else {
			L[i] = L[i-1] * a[i]
		}
	}

	pos := 0

	for i := n - 1; i > 0; i-- {
		if pos == 0 {
			R[pos] = a[i]
		} else {
			R[pos] = R[pos-1] * a[i]
		}
		pos++
	}

	B := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			B[i] = R[len(R)-1]
		} else if i == n-1 {
			B[i] = L[len(L)-1]
		} else {
			B[i] = R[len(R)-i-1] * L[i-1]
		}
	}

	return B
}

func constructArr2(a []int) []int {

	n := len(a)
	B := make([]int, n)

	preFix := make([]int, n)
	subFix := make([]int, n)

	if n == 1 {
		return []int{1}
	}
	if n == 0 {
		return []int{}
	}

	preFix[0] = 1
	subFix[n-1] = 1
	for i := 1; i < n; i++ {
		preFix[i] = preFix[i-1] * a[i]
		subFix[n-i-1] = subFix[n-i] * a[n-i]
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			B[i] = subFix[i]
		} else if i == n-1 {
			B[i] = preFix[i-1]
		} else {
			B[i] = subFix[i] * preFix[i-1]
		}
	}

	return B
}
