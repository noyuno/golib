package strings

// Index : 最初に現れる文字の位置を返す
// s: 対象文字列
// str: 検索文字列
// start: sの先頭を0としたstrの開始位置
// return: sの先頭を0としたstrの開始位置.
// len(s) < len(str)であれば-1を返す．
// strが空文字であればmin(len(s), start)を返す．
// strを見つけることができなければ，-1を返す．
func Index(s []rune, str []rune, start int) int {
	i := 0
	if start > 0 {
		i = start
	}
	p := 0
	if len(s) < len(str) {
		return -1
	}
	if len(str) == 0 {
		if start < len(s) {
			return start
		}
		return len(s)
	}
	for ; i < len(s) && p < len(str); i++ {
		if s[i] == str[p] {
			p += 1
			if p >= len(str) {
				return i - len(str) + 1
			}
		} else {
			p = 0
		}
	}
	return -1
}
