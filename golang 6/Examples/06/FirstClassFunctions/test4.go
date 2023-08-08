package main

import (
	"fmt"
)

func get() (p func(a *int, b int) int) {
	return func(a *int, b int) (power int) {
		for *a%b == 0 {
			*a /= b
			power++
		}
		return
	}
}

func main() {
	x, y := 2000, 5
	fmt.Println(get()(&x, y)) // 3
	fmt.Println(x, y)         // 16 5
}
