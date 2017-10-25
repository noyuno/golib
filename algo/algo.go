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

// Revrse : 配列を逆順に並び替える
// a : 配列（スライス）．変更される．
func Reverse(a []int) {
	lena := len(a)
	for i, j := 0, lena-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
