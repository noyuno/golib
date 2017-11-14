// Package runes : control []rune strings
package runes

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
			p++
			if p >= len(str) {
				return i - len(str) + 1
			}
		} else {
			p = 0
		}
	}
	return -1
}

// Split : sを(sに含まれているであろう)sepで区切って部分文字列にスライスする
// s: 対象文字列
// sep: 区切る制御文字列（制御記号）
// return: sepで区切られた文字列
// sepがsに含まれていないときは大きさ1のスライスを返すが，その要素はsだけである．
// strings.Split()で行われる，sepが空の場合にUTF-8シーケンスでスライスするということはせず，上記に準ずる．
func Split(s []rune, sep []rune) [][]rune {
	ret := make([][]rune, 0, 10)
	a := 0
	b := Index(s, sep, a)
	slen := len(s)
	seplen := len(sep)
	for b != -1 && slen != b {
		ret = append(ret, s[a:b])
		a = b + seplen
		b = Index(s, sep, a)
	}
	last := s[a:]
	if len(last) != 0 {
		ret = append(ret, last)
	}
	return ret
}

// Compare : 文字列を比較する
// a: 比較文字列
// b: 比較文字列
// return: trueならばaとbの文字列は一致している．falseならばaとbの文字列は一致しない．
func Compare(a []rune, b []rune) bool {
	lena := len(a)
	if lena != len(b) {
		return false
	}
	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TrimSpace : 前後の空白文字を削除して残った文字列を返す
// s : 文字列
// return: トリムした文字列．ただし，トリムした結果文字列が残らなかったときは空文字列を返す．
// 空白文字とは，スペースおよびタブ文字をいう．
func TrimSpace(s []rune) []rune {
	lens := len(s)
	start := 0
	space := []rune(" ")[0]
	tab := []rune("\t")[0]
	spaceZenkaku := []rune("　")[0]
	for ; start < lens; start++ {
		if s[start] != space && s[start] != tab && s[start] != spaceZenkaku {
			break
		}
	}
	end := len(s) - 1
	for ; end >= start; end-- {
		if s[end] != space && s[end] != tab && s[end] != spaceZenkaku {
			break
		}
	}
	return s[start : end+1]
}

// Join : 文字列の配列の各要素をsepで連結します
// s : 連結する文字列の配列
// sep : 各要素を結びつけるための記号的役割を担う文字列
// return : 連結された文字列
// なお，sepが空文字列であっても正しく動作する．
func Join(s [][]rune, sep []rune) []rune {
	var ret []rune
	lens := len(s)
	var i int
	for i = 0; i < lens-1; i++ {
		ret = append(ret, s[i]...)
		ret = append(ret, sep...)
	}
	ret = append(ret, s[i]...)
	return ret
}
