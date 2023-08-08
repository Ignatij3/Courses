package main

import "fmt"

func main() {
	nextFibo := makeGenerator()
	for i := 0; i < 10; i++ {
		fmt.Printf("f(%d) = %d\n", i, nextFibo())
	}
}

//	f(0) = 1
//	f(1) = 1
//	f(2) = 2
//	f(3) = 3
//	f(4) = 5
//	f(5) = 8
//	f(6) = 13
//	f(7) = 21
//	f(8) = 34
//	f(9) = 55

func makeGenerator() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}
