// algo : my algorithm library
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

// Reverse : 配列を逆順に並び替えて，そのコピーを返す．
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
