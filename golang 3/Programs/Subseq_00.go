package main

import (
	"fmt"
	"os"
)

func main() {
	File, err := os.Open("../Files/Numbers.dat")
	if err != nil {fmt.Println(err)}
	defer File.Close()
	var (
		var slice []int
		temp, scan, check int
	)
	for {
		if _, err := fmt.Fscanln(File, &temp); err == nil {
			slice = append(slice, temp)
		}  else  {
			break
		}
	}
	fmt.Println(slice)
	fmt.Print("Введите число, которое хотите проверить (до 100000): ")
	fmt.Scan(&scan)
	for i := 0; i < len(slice); i++ {
		if slice[i] == scan {check++}
		if slice[i] > scan {break}
	}
	fmt.Println("Вот число вхождений этого числа: ", check)
}
