package math

import "strconv"

/**
该问题已经有了大概的思路，但是并未深入思考解决问题，看到题解后，发现自己的解题思路是正确的，然后才能攻克难题
*/
func findNthDigit(n int) int {

	// 如果n是个位数，直接返回即可
	if n < 10 {
		return n
	}

	// 将数字0剔除：此时，满足一定的规律性
	// 即个位数是9个、10位数90个、100位数900个
	n -= 1

	// 根据上边的 9、90、900 可以定义一个
	const baseMulti int = 9

	// 位数, 1表示1位数、2表示两位数、3表示三位数
	digits := 1

	// 当前n位数的最小值：10^n
	start := 1

	for true {
		count := baseMulti * start * digits
		if n > count {
			n -= count
		} else {
			break
		}
		digits++

		// 每次都将 power*10
		start *= 10
	}

	// 循环结束后，得到的数字k表示：最终结果是存在于 digits 位数中的

	// 得到的 num 就是 要计算的第n个数位所在的数字
	num := start + n/digits

	// 最终结果在 num 数字中的偏移量是
	offset := n % digits

	// 将 num 转成string后直接通过下标获取结果比较方便，不用写循环
	ans := strconv.Itoa(num)[offset]

	return int(ans - '0')
}
