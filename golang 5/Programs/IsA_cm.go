package main

import "fmt"

func IsA(n int) bool {
	return (n == 1) || (n % 2 == 1 && IsA(n / 2)) || (n % 3 == 1 && IsA(n / 3))
}

func main() {
	for i := 1; i < 33; i++ {
		if IsA(i) {fmt.Println(i)}
	}
}
