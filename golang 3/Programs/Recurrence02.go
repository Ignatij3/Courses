package main

import (
	"fmt"
	"math"
)

func Calculations() float64 {
	Check := 1 * math.Pow(10, -9)
	temp := 1.0
	Number := 0.1
	Total := 0.0
	change := -1.0
	for math.Abs(Number) > Check {
		change *= -1
		Number = ((change / (temp * math.Pow(2, temp)))) + ((change / (temp * math.Pow(3, temp))))
		temp += 2
		Total += Number
	}
	return Total
}

func main() {
	num := Calculations()
	num *= 4
	fmt.Println(num)
}
