package algo

import (
	"reflect"
	"testing"
)

func permutationsF(a []int) int {
	return a[0]*100 + a[1]*10 + a[2]
}

func permutationsT(in []int, expected [][]int, t *testing.T) {
	mt := make([]int, len(in))
	copy(mt, in)
	actual := Permutations(in)

	// inが書き換わっていないか確認
	for i := range in {
		if in[i] != mt[i] {
			t.Errorf("array in has been rewritten by Permutations")
		}
	}

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

func TestReverseCopy(t *testing.T) {
	in := []int{1, 2, 3}
	expected := []int{3, 2, 1}
	out := ReverseCopy(in)
	for i := range expected {
		if expected[i] != out[i] {
			t.Errorf("want [%v]=%v, but returned [%v]=%v", i, expected[i], i, out[i])
		}
	}
}

func TestContains(t *testing.T) {
	in := []int{4, 9, -6}
	testfalse := -16
	testtrue := 9
	if Contains(in, testfalse) == true {
		t.Errorf("want false, but returned true")
	}
	if Contains(in, testtrue) == false {
		t.Errorf("want true, but returned false")
	}
}

func TestValidateOrder(t *testing.T) {
	a := []int{3, 0, 1, 4, 2}
	vf1 := [][]int{[]int{0, 2}}
	vf2 := [][]int{[]int{4}}
	vf3 := [][]int{[]int{3, 4}}
	vt := [][]int{[]int{0, 3}, []int{4, 3}, []int{2, 0}, []int{1, 0}}
	if ValidateOrder(a, vf1) == true {
		t.Errorf("want false, but returned true")
	}
	if ValidateOrder(a, vf2) == true {
		t.Errorf("want false, but returned true")
	}
	if ValidateOrder(a, vf3) == true {
		t.Errorf("want false, but returned true")
	}
	if ValidateOrder(a, vt) == false {
		t.Errorf("want true, but returned false")
	}
}

func sliceAdderT(t *testing.T, s []int, n []int, expected []int, expectedf int) {
	actual, actualf := SliceAdder(s, n, len(s))
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("%v %v want slice %v, but returned %v", s, n, expected, actual)
	}
	if expectedf != actualf {
		t.Errorf("%v %v want flag %v, but returned %v", s, n, expectedf, actualf)
	}
}

func TestSliceAdder(t *testing.T) {
	sliceAdderT(t, []int{2, 4, 8, 1}, []int{4, 6, 12, 6}, []int{3, 4, 8, 1}, 0)
	sliceAdderT(t, []int{3, 4, 8, 1}, []int{4, 6, 12, 6}, []int{0, 5, 8, 1}, 1)
	sliceAdderT(t, []int{0, 5, 8, 1}, []int{4, 6, 12, 6}, []int{1, 5, 8, 1}, 0)
	sliceAdderT(t, []int{7, 5, 12, 5}, []int{4, 6, 12, 6}, []int{0, 0, 0, 0}, 3)
	sliceAdderT(t, []int{15, 1, 2, 3}, []int{4, 6, 12, 6}, []int{0, 2, 2, 3}, 1)
	sliceAdderT(t, []int{9, 5, 0, 3}, []int{10, 10, 10, 10}, []int{0, 6, 0, 3}, 1)
	sliceAdderT(t, []int{0, 5, 0, 3}, []int{0, 5, 2, 6}, []int{0, 0, 1, 3}, 1)
}

func sliceAdderRT(t *testing.T, s []int, n []int, expected []int, expectedf int) {
	actual, actualf := SliceAdderR(s, n, len(s))
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("%v %v want slice %v, but returned %v", s, n, expected, actual)
	}
	if expectedf != actualf {
		t.Errorf("%v %v want flag %v, but returned %v", s, n, expectedf, actualf)
	}
}

func TestSliceAdderR(t *testing.T) {
	sliceAdderRT(t, []int{2, 4, 7}, []int{4, 4, 4}, []int{3, 0, 0}, 1)
}
