package string

import (
	"strings"
)

func reverseWords(s string) string {

	str := strings.Split(strings.TrimSpace(s), "\\s+")

	k := 0
	for i := 0; i < len(str); i++ {
		if str[i] != "" {
			str[k] = str[i]
			k++
		}
	}

	str = str[:k]

	reverse(str)

	return strings.Join(str, " ")
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
