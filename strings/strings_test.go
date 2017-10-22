package strings

import "testing"

func tindex(t *testing.T, s string, str string, expected int) {
	actual := Index([]rune(s), []rune(str))
	if actual != expected {
		t.Errorf("want %v, but %v", expected, actual)
	}
}

func TestIndex(t *testing.T) {
	tindex(t, "abcde", "d", 3)
	tindex(t, "abcde", "o", -1)
	tindex(t, "あえいうえお", "えお", 4)
	tindex(t, "カキクケキクケキクケコ", "キクケコ", 7)
}
