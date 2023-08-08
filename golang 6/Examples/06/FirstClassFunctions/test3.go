package main

import (
	"fmt"
)

func call(p func(a *int, b int) int, divisible *int, divisor int) {
	fmt.Println(p(divisible, divisor))
}

func main() {
	f := func(a *int, b int) (power int) {
		for *a%b == 0 {
			*a /= b
			power++
		}
		return
	}
	x, y := 2000, 5
	call(f, &x, y)    // 3
	fmt.Println(x, y) // 16 5
}
