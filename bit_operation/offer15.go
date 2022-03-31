package bit_operation

/**
计算 二进制中 1的个数
数字 num 与 1 进行与运算；然后将 num 右移一位
*/
func hammingWeight(num uint32) int {

	var ans uint32
	ans = 0
	for num != 0 {
		ans += num & 1
		num >>= 1
	}
	return int(ans)
}

/**
方法二：
	首先我们可知道的是：一个数字减去1后，其最后一个1必然会变成0，而其后边的位都是原来的取反操作
如： 1100 - 1 = 1011; 1001 - 1 = 1000
	那么，每次将 num - 1 并与自身与运算，如果不是0则1的个数多一个
*/
func hammingWeightII(num uint32) int {

	ans := 0
	for num != 0 {
		ans++
		num = num & (num - 1)
	}
	return ans
}
