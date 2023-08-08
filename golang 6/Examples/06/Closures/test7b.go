package main

import "fmt"

func main() {
	counter := newCounter()
	fmt.Println(counter()) // 1
	// . . .
	fmt.Println(counter()) // 2
}

func newCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}
