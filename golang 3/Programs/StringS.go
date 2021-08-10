package main

import (
	"fmt"
	"os"
)

func Calculating(str string) bool {
	File, err := os.Open("../Files/Numbers.txt")
	defer File.Close()
	if err != nil {fmt.Println(err)}
	var Text string
	fmt.Fscanln(File, &Text)
	i := 0
	strR := []rune(str)
	for _, symbol := range []rune(Text) {
		if symbol == strR[i] {
			i++
			if i >= len(strR) {
				return true
			}
		}
		
	}
	return false
}

func main() {
	var onceUsed string
	fmt.Print("Введите подстроку: ")
	fmt.Scan(&onceUsed)
	Calculating(onceUsed)
	/*if Calculating == true {
		fmt.Println("Subsequence found")
	} else {
		fmt.Println("Subsequence haven't been found")
	}*/
	//var temp string
	fmt.Println(Calculating(onceUsed))
}
