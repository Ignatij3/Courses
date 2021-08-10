package main

import "fmt"

func IsA(n int) bool {
	var a, b bool
	if n == 1 {return true}
	if (n-1) % 2 == 0 {a = IsA((n-1)/2)} else {a = false}
	if (n-1) % 3 == 0 {b = IsA((n-1)/3)} else {b = false}
	if !a && !b {return false}
	return true
}

func main() {
	var x int
	for {
		fmt.Print("Введите число больше нуля, которое хотите проверить на принадлежность  A:")
		fmt.Scan(&x)
		switch IsA(x) {
			case true:
				fmt.Printf("Число %d принадлежит A\n\n", x)
			case false:
				fmt.Printf("Число %d не принадлежит A\n\n", x)
		}
	}
}
