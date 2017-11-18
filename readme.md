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
    func ValidateOrder(a []int, v [][]int) bool 

### runes

control []rune strings

    func Compare(a []rune, b []rune) bool 
    func Index(s []rune, str []rune, start int) int 
    func Join(s [][]rune, sep []rune) []rune 
    func Split(s []rune, sep []rune) [][]rune 
    func TrimSpace(s []rune) []rune 


## Make readme

    ./make-readme.sh

