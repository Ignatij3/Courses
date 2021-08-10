package main

import (
	"fmt"
	"math"
)

func main() {
	var Input1 string
	var Input2 string
	var Num1 int
	var Num2 int
	var Ctrl int
	var Check bool = true
	fmt.Scan(&Input1)
	fmt.Scan(&Input2)
	Compare1 := []rune(Input1)
	Compare2 := []rune(Input2)
	for index1, compare1 := range Compare1 {
		if !Check {
			break
		}
		if compare1 != '0' && compare1 != '1' {
			fmt.Println("ERROR! \n Вы ввели небинарное число, попробуйте снова!")
			Check = false
		}
		if compare1 == '1' {
			Ctrl = int(math.Pow(2, float64(index1)))
			Num1 = Ctrl + Num1
		}
	}
	for index2, compare2 := range Compare2 {
		if !Check {
			break
		}
		if compare2 != '0' && compare2 != '1' {
			fmt.Println("ERROR! \n Вы ввели небинарное число, попробуйте снова!")
			Check = false
		}
		if compare2 == '1' {
			Ctrl = int(math.Pow(2, float64(index2)))
			Num2 = Ctrl + Num2
		}
	}
	if Num1 < Num2 {
		fmt.Println(Input1, " < ", Input2)
		fmt.Println(Num1, " < ", Num2)
	}
	if Num1 > Num2 {
		fmt.Println(Input1, " > ", Input2)
		fmt.Println(Num1, " > ", Num2)
	}
	if Num1 == Num2 {
		fmt.Println(Input1, " = ", Input2)
		fmt.Println(Num1, " = ", Num2)
	}
}
