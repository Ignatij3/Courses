package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 11, -5, 8, 2, 0, 12}
	for _, x := range numbers {
		go func() {
			process(x)
		}()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	//	12 12 12 12 12 12 12

	for _, x := range numbers {
		tmp := x
		go func() {
			process(tmp)
		}()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	//	8 1 12 -5 11 0 2

	for _, x := range numbers {
		go func(input int) {
			process(input)
		}(x)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	//	2 12 0 1 11 -5 8
}

func process(n int) {
	time.Sleep(50 * time.Millisecond)
	fmt.Print(n, " ")
}
