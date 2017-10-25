# lgo

my Golang library

## functions list

### runes

control `[]rune` strings

    func Index(s []rune, str []rune, start int) int
    func Split(s []rune, sep []rune) [][]rune
    func Compare(a []rune, b []rune) bool
    func TrimSpace(s []rune) []rune
    func Join(s [][]rune, sep []rune) []rune

### algo

algorithm

    func Permutations(a []int) [][]int
    func Reverse(a []int)

