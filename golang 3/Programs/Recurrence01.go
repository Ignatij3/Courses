package main

import (
	"fmt"
	"math"
)

func Calculations() float64 {
	Check := 1 - 1 * math.Pow(10, -9)
	Number := math.Sqrt(0.5)
	Total := Number
	for Number <= Check {
		Number = math.Sqrt(0.5 + 0.5 * Number)
		Total *= Number
	}
	return Total
}

func main() {
	num := Calculations()
	fmt.Println(2/num)
}
