package main

import (
	"fmt"
)

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	a = append(a, 6)
	a = append(a, 7)
	a = append(a, 8)
	a = append(a, 9)
	a = append(a, 10)
	println("Slice:", a)
	fmt.Println("Slice:", a)
}
