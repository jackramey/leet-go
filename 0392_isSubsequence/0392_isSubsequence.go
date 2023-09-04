package isSubsequence

import "strings"

func isSubsequence(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		at := strings.IndexRune(t, rune(s[i]))
		if at == -1 {
			return false
		}

		t = t[at+1:]
	}

	return true
}
