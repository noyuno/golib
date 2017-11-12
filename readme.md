# lgo

[![Build Status](https://travis-ci.org/noyuno/lgo.svg?branch=master)](https://travis-ci.org/noyuno/lgo)

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
    func ReverseCopy(a []int) []int 
    func ValidateOrder(a []int, v [][]int) bool 
    func Contains(s []int, t int) bool 

Do not edit this file, please edit description.md instead.
