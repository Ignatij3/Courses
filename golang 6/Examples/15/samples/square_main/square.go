package main

import "fmt"

func Square(a int) int {
	return a*a
}

func main() {
	n:= 32
    fmt.Printf("%d^2 = %d", n, Square(n))
}

