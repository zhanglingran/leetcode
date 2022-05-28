package main

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	// write code here

	K -= 1
	for true {
		m, val := partition(a)
		if m == K {
			return val
		} else if m > K {
			a = a[:m]
		} else {
			a = a[m+1:]
			K -= (m + 1)
		}
	}
	return -1
}

func partition(a []int) (int, int) {
	n := len(a)
	head := 0
	pivot := a[0]
	for i := 0; i < n-1; i++ {
		if a[i] > pivot {
			a[head] = a[i]
			head++
		}
	}

	a[n-1], a[head] = a[head], a[n-1]
	return head, a[head]
}
