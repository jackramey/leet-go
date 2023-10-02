package isPalindrome

import (
	"regexp"
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all
//   non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//

const alphaNumPattern = `[^a-zA-Z0-9]+`

// isPalindrome first formats the string by removing all whitespace and non-alphanumeric characters and then makes sure
// we're using a consistent case for non-numeric characters by casting them toLower. Once normalized we keep to cursors,
// i and j, that start at the beginning and end of the string respectively. If at any point the characters at those
// positions are not equivalent then we return false. If we've traversed the full string without any early exit that
// means the string is a valid palindrome.
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
