package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println("ab" + BRed + "c" + FBlue + "def" + Reset + "gh")
	fmt.Println("a" + F(5) + "b" + B(208) + "c" + F(103) + "d" + B(126) + Reset + "e")
}
