package isPalindrome

import (
	"regexp"
	"strings"
)

const alphaNumPattern = `[^a-zA-Z0-9]+`

func isPalindrome(s string) bool {
	s = regexp.MustCompile(alphaNumPattern).ReplaceAllString(s, "")
	if s == "" {
		return true
	}

	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
