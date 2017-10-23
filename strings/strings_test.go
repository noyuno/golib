package strings

import "testing"

func tindex(t *testing.T, s string, str string, start int, expected int) {
	actual := Index([]rune(s), []rune(str), start)
	if actual != expected {
		t.Errorf("Index(%s, %s, %d): want %v, but returned %v", s, str, start, expected, actual)
	}
}

func TestIndex(t *testing.T) {
	tindex(t, "abcde", "d", 0, 3)
	tindex(t, "abcde", "o", 2, -1)
	tindex(t, "あえいうえお", "えお", 0, 4)
	tindex(t, "あい", "あいうえお", 0, -1)
	tindex(t, "カキクケキクケキクケコ", "キクケ", 3, 4)
	tindex(t, "カキクケキクケキクケコ", "", 3, 3)
}

func TestSplit(t *testing.T) {
	in := []rune("abcde-fghij")
	expected := [][]rune{[]rune("abcde"), []rune("fghij")}
	actual := Split(in, []rune("-"))
	for i := 0; i < len(expected); i++ {
		if !Compare(expected[i], actual[i]) {
			t.Errorf("Split(%s, %s): want %v, but %v", string(in), "-", expected, actual)
		}
	}
}

func tTrim(t *testing.T, test string, expected string) {
	actual := TrimSpace([]rune(test))
	if !Compare(actual, []rune(expected)) {
		t.Errorf("TrimSpace(%s): want %s, but returned %s", test, expected, string(actual))
	}
}

func TestTrimSpace(t *testing.T) {
	a := "abc"
	at := a
	b := " 　d e  	f	"
	bt := "d e  	f"
	tTrim(t, a, at)
	tTrim(t, b, bt)
}
