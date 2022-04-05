package string

import "strings"

/**
判断是不是数值: A[.[B]][e|E C]
其中 A和C是整数，其前边可以包含正负号
*/
func isNumber(s string) bool {

	s = strings.Trim(s, " ")

	if len(s) == 0 {
		return false
	}
	k := 0

	if s[k] == '+' || s[k] == '-' {
		k++
	}

	// 接下来就是判断整数部分, 并用isInteger标记是不是包含了数字，如果不是整数，直接返回false
	isInteger := false
	for k < len(s) {

		// 如果遇到了 . 或者 e 进入到下个判断的环节
		if s[k] == '.' || s[k] == 'e' || s[k] == 'E' {
			break
		}
		// 如果除了非 . 或 e 之外的其他数字，直接返回false
		if s[k] < '0' || s[k] > '9' {
			return false
		}
		// 标记有数字
		isInteger = true
		k++
	}

	// 如果 遇到了., 后边就要判断小数部分
	isDecimal := false
	if k < len(s) && s[k] == '.' {
		k++
		for k < len(s) {
			// 判断后边是不是 e
			if s[k] == 'e' || s[k] == 'E' {
				break
			}
			if s[k] < '0' || s[k] > '9' {
				return false
			}
			k++
			isDecimal = true
		}
	}
	// 如果前边没有数字，那么直接返回false
	if !isInteger && !isDecimal {
		return false
	}

	// 最后就是了 e, 如果有e 但是后边没有数字 那么返回false
	if k < len(s) && (s[k] == 'e' || s[k] == 'E') {
		k++
		if k < len(s) && s[k] == '+' || s[k] == '-' {
			k++
		}
		flag := false
		for k < len(s) {
			if s[k] < '0' || s[k] > '9' {
				return false
			}
			flag = true
			k++
		}
		if !flag {
			return false
		}
	}

	return true
}

// 参考剑指offer中的解题思路
func isNumber1(s string) bool {

	// 将 s 前后的空格裁掉
	s = strings.Trim(s, " ")

	// 1 首先扫描整数
	isInteger, s := scanNumber(s)

	// 2 扫面小数部分
	var ok bool
	if len(s) > 0 && s[0] == '.' {
		s = s[1:]
		ok, s = scanUnsignedNumber(s)
		isInteger = isInteger || ok
	}

	// 3 后边是 e
	if len(s) > 0 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		ok, s = scanNumber(s)
		isInteger = isInteger && ok
	}

	return isInteger && len(s) == 0
}

func scanNumber(s string) (bool, string) {
	if len(s) == 0 {
		return false, s
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	return scanUnsignedNumber(s)
}

// 扫描字符串 s, 如果包含了正整数 返回true，否则返回 false
func scanUnsignedNumber(s string) (bool, string) {
	if len(s) == 0 {
		return false, s
	}
	idx := 0
	for idx < len(s) && s[idx] <= '9' && s[idx] >= '0' {
		idx++
	}
	return idx > 0, s[idx:]
}
