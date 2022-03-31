package bit_operation

/**
判断一个数字是不是2的幂次方
*/

func isPowerOfTwo(n int) bool {

	if n&(n-1) == 0 {
		return true
	}
	return n == 0
}
