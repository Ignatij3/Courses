package main

import "fmt"

func Solve(n int, dir bool) {
	if n == 1 {
		if dir {
			fmt.Printf("  Put %d right down\n", n)
		} else if !dir {
			fmt.Printf("  Put %d left  up\n", n)
		}
	}
	if n > 1 {
		if dir {
			Solve(n-1, true)
			fmt.Printf("  Put %d right up\n", n)
			Solve(n-1, false)
			fmt.Printf("Shift %d down\n", n)
			Solve(n-1, true)
		} else if !dir {
			Solve(n-1, false)
			fmt.Printf("Shift %d up\n", n)
			Solve(n-1, true)
			fmt.Printf("  Put %d left  up\n", n)
			Solve(n-1, false)
		}
	}
}

func main() {
	Solve(5, true)
}
