package strings

func Index(s []rune, str []rune) int {
	var i int = 0
	var p int = 0
	for ; i < len(s) && p < len(str); i++ {
		if s[i] == str[p] {
			p += 1
			if p >= len(str) {
				return i - len(str) + 1
			}
		}
	}
	return -1
}
