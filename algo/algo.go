// Package algo : my algorithm library
package algo

// Permutations : 入力配列の順列（配列を並び替えてできる配列の群）をすべて求める
// a : 配列
// return : 順列
func Permutations(a []int) [][]int {
	var recurse func([]int, int)
	ret := [][]int{}

	recurse = func(a []int, n int) {
		if n == 1 {
			t := make([]int, len(a))
			copy(t, a)
			ret = append(ret, t)
		} else {
			for i := 0; i < n; i++ {
				recurse(a, n-1)
				if n%2 == 1 {
					a[i], a[n-1] = a[n-1], a[i]
				} else {
					a[0], a[n-1] = a[n-1], a[0]
				}
			}
		}
	}
	recurse(a, len(a))
	return ret
}

// Reverse : 配列を逆順に並び替える
// a : 配列（スライス）．変更される．
func Reverse(a []int) {
	lena := len(a)
	for i, j := 0, lena-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// ReverseCopy : 配列を逆順に並び替えて，そのコピーを返す．
// a : 配列（スライス）．変更されない
// return : 逆順になった配列
func ReverseCopy(a []int) []int {
	lena := len(a)
	ret := make([]int, lena)
	for i, j := 0, lena-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = a[j], a[i]
	}
	if lena%2 == 1 {
		ret[(lena-1)/2] = a[(lena-1)/2]
	}
	return ret
}

// ValidateOrder : 順序が正しいかどうかテストする
// a : テストする配列
// v : 現れるべきでない部分パターンの集合
// return : true : 正しい, false : 正しくない
func ValidateOrder(a []int, v [][]int) bool {
	e := make([]int, len(v))
	for i := 0; i < len(a); i++ {
		for p := 0; p < len(v); p++ {
			if v[p][e[p]] == a[i] {
				// step
				e[p]++
				if len(v[p]) <= e[p] {
					// accepted by v[p]
					return false
				}
			}
		}
	}
	return true
}

// Contains : テストする値が配列中に含まれているかどうかテストする
// s : 配列
// t : 配列sに含まれていることを期待する値
// return : true: 含まれている, false: 含まれていない
func Contains(s []int, t int) bool {
	for _, v := range s {
		if t == v {
			return true
		}
	}
	return false
}

// SliceAdder : スライスをn進数として桁上がり可能な加算器（nは要素ごとに異なる）
// s: 配列
// n: 要素の進数を示す値が含まれた配列
// l: 配列の要素数
// return: 1つ加算したs,
// 検知ビット(1: 桁上がり, 3: オーバーフロー)
// s = [2 4 8 1], n =  [4 6 12 6]のとき，
// SliceAdder(s, n)とすると[3 4 8 1]が返るが，
// SliceAdder([3 4 8 1], [4 6 12 6])とすると[0 5 8 1]が返る
func SliceAdder(s []int, n []int, l int) ([]int, int) {
	if s[0] >= n[0]-1 {
		// 桁上がり
		if len(s) == 1 {
			// 最上位桁であって，これ以上桁上がりできない
			return []int{0}, 3
		}
		r, f := SliceAdder(s[1:], n[1:], l)
		ret := make([]int, 1, l)
		ret[0] = 0
		ret = append(ret, r...)
		return ret, f
	}
	ret := make([]int, len(s))
	copy(ret, s)
	ret[0]++
	var f int
	if l == len(s) {
		f = 0
	} else {
		f = 1
	}
	return ret, f
}
