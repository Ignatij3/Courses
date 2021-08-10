package main

import (
	//"strings"
	"bufio"
	"fmt"
	"os"
)

func CompareLetters(a, b rune) bool {
	return (a == b) || (a >= 'a' && b < 'a' && a - 32 == b) || (a < 'a' && b >= 'a' && a + 32 == b)
}

func DisplaySlice(s []string) {
	if len(s) > 0 {
		fmt.Println("Here are all the words found:")
		for _, content := range s {fmt.Println(content)}
	} else {fmt.Println("No words found")}
	fmt.Println()
}

func CheckLetter(newstr []rune, letter rune) int {
	for n, p := range newstr {
		if letter == p {return n}
	}
	return -1
}
/*
func CompareToWord(origstr, newstr string) bool {
	a, b := []rune(origstr), []rune(newstr)
	for n, c := range a {
		if n == len(b) {return false}
		if c == '*' {
			if n == len(a) - 1 || (n == len(a) - 2 && b[len(b) - 1] == a[n + 1]) {return true}
			i := CheckLetter(b[n:], a[n + 1])
			if i == -1 {return false}
			if i + n == len(b) - 1 {
				return CompareToWord(strings.SplitAfter(origstr, "*")[1], string(b[i + n]))
			} else {
				return CompareToWord(strings.SplitAfter(origstr, "*")[1], strings.SplitAfterN(newstr, string(b[i + n]), 2)[1])
			}
		} else if c != '?' && !CompareLetters(c, b[n]) {return false}
	}
	return len(a) >= len(b)
}*/

func CompareToWord(mask, str []rune) bool {
	if len(str) == 0 {return false}
	if len(mask) == 1 {
		if mask[0] == '*' {return true}
		return (CompareLetters(mask[0], str[0]) || mask[0] == '?') && len(str) == 1
	}
	if mask[0] == '*' {
		return CompareToWord(mask[1:], str) || CompareToWord(mask, str[1:]) 
	} else if mask[0] == '?' {
		return CompareToWord(mask[1:], str[1:]) 
	}
	return CompareLetters(mask[0], str[0]) && CompareToWord(mask[1:], str[1:])
}

func SearchWords(s string) []string {
	var words []string
	file, err := os.Open("../Files/words.txt")
	if err != nil {fmt.Println(err)}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if CompareToWord([]rune(s), []rune(scanner.Text())) {
			words = append(words, scanner.Text())
		}
	}
	file.Close()
	return words
}

func main() {
	var (
		s string
		words []string
	)
	
	for {
		fmt.Print("Enter a word with mask:")
		fmt.Scan(&s)
		words = SearchWords(s)
		DisplaySlice(words)
		
		fmt.Print("Do you want to continue? (y/n)")
		fmt.Scan(&s)
		for s != "y" && s != "n" {
			fmt.Print("Error, try again (y/n)")
			fmt.Scan(&s)
		}
		if s == "n" {break}
	}
}
