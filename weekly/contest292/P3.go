package contest292

func countTexts(pressedKeys string) int {

	// 0. 首先打表
	const MOD int = 1e9 + 7
	const N int = 1e5 + 1

	dp3 := make([]int, N)
	dp3[1], dp3[2], dp3[3] = 1, 2, 4

	for i := 4; i < N; i++ {
		dp3[i] = (dp3[i-1]%MOD + dp3[i-2]%MOD + dp3[i-3]%MOD) % MOD
	}

	dp4 := make([]int, N)
	dp4[1], dp4[2], dp4[3], dp4[4] = 1, 2, 4, 8

	for i := 5; i < N; i++ {
		dp4[i] = (dp4[i-1]%MOD + dp4[i-2]%MOD + dp4[i-3]%MOD + dp4[i-4]%MOD) % MOD
	}

	ans := 1
	cont := 0
	// 1. 判断有多少个连续的数字及其个数k
	l, r := 0, 0
	for r < len(pressedKeys) {
		if pressedKeys[r] == pressedKeys[l] {
			cont++
		} else {
			if pressedKeys[l] == '7' || pressedKeys[l] == '9' {
				ans = (ans % MOD * dp4[cont] % MOD) % MOD
			} else {
				ans = (ans % MOD * dp3[cont] % MOD) % MOD
			}
			cont = 1
			l = r
		}
		r++
	}

	if pressedKeys[l] == '7' || pressedKeys[l] == '9' {
		ans = (ans % MOD * dp4[cont] % MOD) % MOD
	} else {
		ans = (ans % MOD * dp3[cont] % MOD) % MOD
	}

	// 2. 返回 ans 即可
	return ans
}
