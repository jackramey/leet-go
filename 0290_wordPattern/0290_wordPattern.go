package wordPattern

import "strings"

func wordPattern(pattern string, s string) bool {
	elements := strings.Split(s, " ")
	if len(pattern) != len(elements) {
		return false
	}

	pMap := make(map[byte]string)
	wMap := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		word, ok := pMap[pattern[i]]
		if ok && word != elements[i] {
			return false
		} else {
			pMap[pattern[i]] = elements[i]
		}
		p, ok := wMap[elements[i]]
		if ok && p != pattern[i] {
			return false
		} else {
			wMap[elements[i]] = pattern[i]
		}
	}
	return true
}
