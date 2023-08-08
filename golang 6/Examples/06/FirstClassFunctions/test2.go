package main

import (
	"fmt"
)

type precede func(a int, b int) bool

func before(a int, b int) bool {
	return a < b
}

func main() {
	var compare1 precede
	var compare2 = func(a int, b int) bool {
		return a < b
	}
	x, y, compare1 := 7, 5, before
	s := compare1(x, y)
	if s {
		fmt.Println(x, "precedes", y)
	} else if compare2(y, x) {
		fmt.Println(y, "precedes", x)
	} else {
		fmt.Println(x, "and", y, "is equal")
	}
	fmt.Printf("compare1 type is %T\n", compare1)
	fmt.Printf("compare2 type is %T\n", compare2)
	fmt.Printf("compare1 = %v, compare2 = %v\n", compare1, compare2)
	compare1 = compare2
	fmt.Printf("compare1 = %v, compare2 = %v\n", compare1, compare2)
	compare2 = before
	fmt.Printf("compare2 = %v\n", compare2)
}

// 5 precedes 7
// compare1 type is func(int, int) bool
// compare2 type is func(int, int) bool
// compare1 = 0x45dd40, compare2 = 0x45e1c0
// compare1 = 0x45e1c0, compare2 = 0x45e1c0
// compare2 = 0x45dd40
