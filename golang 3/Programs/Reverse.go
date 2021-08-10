package main

import (
	"fmt"
	"os"
)

func main() {
	var slice []int
	var temp int
	var x int
	var i int
	var check int
	array, err := os.Open("../Files/Numb.dat")
	if err != nil {fmt.Println(err)}
	defer array.Close()
	for {
		if _, err := fmt.Fscanln(array, &temp); err == nil {
			slice = append(slice, temp)
		}  else  {
			break
		}
	}
	fmt.Print("Введите \"x\":")
	fmt.Scan(&x)
	for i = 0; i < len(slice) - 1; {
		check++
		if slice[i] <= x {i++}
		if slice[i] > x {
			break
			}
	}
	fmt.Println("\"x\" найден на позиции", check)
}
