package Amazon

func subString(search, result string) int32 {
	k := 0
	for i := 0; i < len(search); i++ {
		if search[i] == result[k] {
			k++
		}
	}
	return int32(len(result) - k)
}
