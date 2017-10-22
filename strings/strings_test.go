package strings

import "testing"

func tindex(t *testing.T, s string, str string, start int, expected int) {
	actual := Index([]rune(s), []rune(str), start)
	if actual != expected {
		t.Errorf("Index(%s, %s, %d) want %v, but %v", s, str, start, expected, actual)
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
