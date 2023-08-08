package main

import (
	"fmt"
)

//interface definition
type DigitsFinder interface {
	FindDigits() string
}

type MyString string

//MyString implements DigitsFinder
func (ms MyString) FindDigits() string {
	var digits []rune
	for _, r := range []rune(ms) {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}
	return string(digits)
}

func main() {
	var d DigitsFinder
	date := MyString("Friday, 30 July 2021. 12:05:35")
	d = date                                      // possible since MyString implements DigitsFinder
	fmt.Printf("Digits are %s\n", d.FindDigits()) // Digits are 302021120535
}
