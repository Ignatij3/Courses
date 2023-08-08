package main

import (
	"fmt"
)

func main() {
	a := func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println("[non] the first first class function :)")
		}
	}
	a(2)
	fmt.Printf("%T", a)
}

// [non] the first first class function :)
// [non] the first first class function :)
// func(int)
