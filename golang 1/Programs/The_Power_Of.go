package main

import (
	"fmt"
)

func power2(n int) int {
	var f int
	if n%2 == 1 {
		fmt.Println("Error #104")
	}
	for ; n != 1; n /= 2 {
		f++
	}
	return f
}
func main() {
	var n, f int
	fmt.Scan(&n)
	f = power2(n)
	fmt.Println(f)
}
