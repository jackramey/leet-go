package isomorphicStrings

func isIsomorphic(s string, t string) bool {
	sMap := make(map[uint8]uint8, 26)
	tMap := make(map[uint8]uint8, 26)
	for i := 0; i < len(t); i++ {
		c, ok := sMap[s[i]]
		if ok && c != t[i] {
			return false
		} else {
			sMap[s[i]] = t[i]
		}

		d, ok := tMap[t[i]]
		if ok && d != s[i] {
			return false
		} else {
			tMap[t[i]] = s[i]
		}
	}
	return true
}
