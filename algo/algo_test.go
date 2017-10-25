package algo

import "testing"

func permutationsF(a []int) int {
	return a[0]*100 + a[1]*10 + a[2]
}

func permutationsT(in []int, expected [][]int, t *testing.T) {
	actual := Permutations(in)

	if len(expected) != len(actual) {
		t.Errorf("array size mismatch. want %v, but returned %v",
			len(expected), len(actual))
	}
	m := map[int][]int{}
	for i := 0; i < len(expected); i++ {
		m[permutationsF(expected[i])] = expected[i]
	}
	for i := 0; i < len(actual); i++ {
		_, ok := m[permutationsF(actual[i])]
		if !ok {
			t.Errorf("unknown array")
		}
	}
}

func TestPermutations(t *testing.T) {
	in1 := []int{1, 2, 3}
	expected1 := [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1}}
	in2 := []int{1, 2, 2}
	expected2 := [][]int{
		[]int{1, 2, 2},
		[]int{1, 2, 2},
		[]int{2, 1, 2},
		[]int{2, 2, 1},
		[]int{2, 1, 2},
		[]int{2, 2, 1}}
	permutationsT(in1, expected1, t)
	permutationsT(in2, expected2, t)
}

func TestReverse(t *testing.T) {
	in := []int{1, 2, 3}
	expected := []int{3, 2, 1}
	Reverse(in)
	for i := range expected {
		if expected[i] != in[i] {
			t.Errorf("want [%v]=%v, but returned [%v]=%v", i, expected[i], i, in[i])
		}
	}
}
