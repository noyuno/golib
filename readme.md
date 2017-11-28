# lgo

[![Build Status](https://travis-ci.org/noyuno/lgo.svg?branch=master)](https://travis-ci.org/noyuno/lgo)

My Golang library

## Functions list


### algo

my algorithm library

    const Carry = 1
    const Overflow = 3
    func Contains(s []int, t int) bool 
    func Permutations(a []int) [][]int 
    func Reverse(a []int) 
    func ReverseCopy(a []int) []int 
    func SliceAdder(s []int, n []int, l int) ([]int, int) 
    func SliceAdderR(s []int, n []int, l int) ([]int, int) 
    func ValidateOrder(a []int, v [][]int) bool 

### color

256 color CLI support

    const BBlue = "\x1b[44m"
    const BGray = "\x1b[48m"
    const BGreen = "\x1b[42m"
    const BLightBlue = "\x1b[46m"
    const BPurple = "\x1b[45m"
    const BRed = "\x1b[41m"
    const BWhite = "\x1b[47m"
    const BYellow = "\x1b[43m"
    const FBlue = "\x1b[34m"
    const FGray = "\x1b[38m"
    const FGreen = "\x1b[32m"
    const FLightBlue = "\x1b[36m"
    const FPurple = "\x1b[35m"
    const FRed = "\x1b[31m"
    const FWhite = "\x1b[37m"
    const FYellow = "\x1b[33m"
    const Reset = "\x1b[m"
    func B(n int) string 
    func F(n int) string 

### runes

control []rune strings

    func Compare(a []rune, b []rune) bool 
    func Copy(s []rune) []rune 
    func Index(s []rune, str []rune, start int) int 
    func Join(s [][]rune, sep []rune) []rune 
    func Split(s []rune, sep []rune) [][]rune 
    func TrimSpace(s []rune) []rune 


## Make readme

    ./make-readme.sh

