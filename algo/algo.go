// Package algo : my algorithm library
package algo

// Permutations : 入力配列の順列（配列を並び替えてできる配列の群）をすべて求める
// a : 配列
// return : 順列
func Permutations(data []int) <-chan []int {
	nd := make([]int, len(data))
	copy(nd, data)
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, nd)
	}(c)
	return c
}

// http://www.quickperm.org/ の実装
func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func Combinations(data [][]int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		pos := make([]int, len(data))
		max := make([]int, len(data))
		f := 0
		for i := range data {
			max[i] = len(data[i])
		}
		for f != 3 {
			out := make([]int, len(data))
			for k := range out {
				out[k] = data[k][pos[k]]
			}
			c <- out
			pos, f = SliceAdderR(pos, max, len(out))
		}
	}(c)
	return c
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

// Overflow : オーバーフロー
const Overflow = 3

// Carry : 桁上がり
const Carry = 1

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
			return []int{0}, Overflow
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
		f = Carry
	}
	return ret, f
}

func SliceAdderR(s []int, n []int, l int) ([]int, int) {
	pos := len(s) - 1
	if s[pos] >= n[pos]-1 {
		// 桁上がり
		if len(s) == 1 {
			// 最上位桁であって，これ以上桁上がりできない
			return []int{0}, Overflow
		}
		r, f := SliceAdderR(s[:pos], n[:pos], l)
		r = append(r, 0)
		return r, f
	}
	ret := make([]int, len(s))
	copy(ret, s)
	ret[pos]++
	var f int
	if l == len(s) {
		f = 0
	} else {
		f = Carry
	}
	return ret, f
}
