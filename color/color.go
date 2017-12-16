// Package color : 256 color CLI support
package color

import "fmt"

const FRed = "\x1b[31m"
const FGreen = "\x1b[32m"
const FYellow = "\x1b[33m"
const FBlue = "\x1b[34m"
const FPurple = "\x1b[35m"
const FLightBlue = "\x1b[36m"
const FWhite = "\x1b[37m"
const FGray = "\x1b[38m"

const BRed = "\x1b[41m"
const BGreen = "\x1b[42m"
const BYellow = "\x1b[43m"
const BBlue = "\x1b[44m"
const BPurple = "\x1b[45m"
const BLightBlue = "\x1b[46m"
const BWhite = "\x1b[47m"
const BGray = "\x1b[48m"

const Reset = "\x1b[m"

func F(n int) string {
	return fmt.Sprintf("\x1b[38;05;%vm", n)
}
func B(n int) string {
	return fmt.Sprintf("\x1b[48;05;%vm", n)
}
