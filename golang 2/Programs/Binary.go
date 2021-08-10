package main
import "fmt"

func main() {
	var Input string
	var Equation string
	var first bool = true
	fmt.Scan(&Input)
	Count := [] rune (Input)
	for index, contains := range Count {
		if contains == '1' {
			if !first {
				fmt.Print("| ")
			} else {
				first = false
			}
			if index == 0 {
				Equation = "(!A & !B & !C) ";
			} else if index == 1 {
				Equation = "(!A & !B & C) ";
			} else if index == 2 {
				Equation = "(!A & B & !C) ";
			} else if index == 3 {
				Equation = "(!A & B & C) ";
			} else if index == 4 {
				Equation = "(A & !B & !C) ";
			} else if index == 5 {
				Equation = "(A & !B & C) ";
			} else if index == 6 {
				Equation = "(A & B & !C) ";
			} else if index == 7 {
				Equation = "(A & B & C) ";
			}
			fmt.Print(Equation)
		}
		index++
	}
	
}
