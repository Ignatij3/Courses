package main

import (
	"fmt"
)

func main() {
	a := 5
	func(x int) {
		a += x
		fmt.Println("a +", x, "=", a) // a + 2 = 7
	}(2)
	fmt.Println(a) // 7
}
