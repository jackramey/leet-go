package ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	library := map[rune]int{}
	for _, r := range magazine {
		library[r]++
	}

	for _, r := range ransomNote {
		library[r]--
		if library[r] < 0 {
			return false
		}
	}
	return true
}
