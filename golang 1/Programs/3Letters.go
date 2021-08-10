package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
		var (
			text, res string
		)
		file, err := os.Open("..\\Files\\Homework.txt")
		defer file.Close()
		if err != nil {fmt.Println(err)}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {text += scanner.Text()}
		text = strings.ReplaceAll(text, "	", "")
		text = strings.ReplaceAll(text, " ", "")
		for i, r := range text {
			res = res + string(r)
			if i > 0 && (i+1)%3 == 0 {
				fmt.Printf("'%v'     ", res)
				res = ""
			}
		}
}
