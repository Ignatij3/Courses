package main

import (
	"fmt"
	"math"
)

func Calculations() float64 {
	Check := 1 * math.Pow(10, -12)
	Number := 1.0
	temp01 := 1.0
	temp02 := 1.0
	temp03 := 1.0
	Adding := 1.0
	Total := 0.0
	x := 1.0
	y := 0
	for Number > Check {
		Adding *= (temp01 / temp02)
		Number = x + (1 / 1 * Adding) / math.Pow(4, temp03)
		temp01 += 2
		temp02 += 1
		temp03 += 1
		Total += Number
		y = 1
		if y == 1 {x = 0}
	}
	return Total
}

func main() {
	num := Calculations()
	num = math.Pow(num, 2)
	fmt.Println(num)
}
