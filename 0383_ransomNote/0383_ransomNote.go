package ransomNote

// canConstruct takes both the ransomNote and magazine as inputs and then catalogues the characters available
// from the magazine and then removes them from the catalogue when iterating through the ransomNote. If at any point
// the catalogue contains a negative number for a character in the library then we will return false.
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
