package main

import (
	"fmt"
)

func ClosestSquare(n uint64) uint64 {
	p := uint64(0)
	for p * p <= n {p++}
	return p - 1
}

func FindSquares(n uint64, amount int) {
	var sq uint64
	if n > 0 {
		sq = ClosestSquare(n)
		if amount == 3 {
			fmt.Printf("%d^2", sq)
		} else {
			fmt.Printf("%d^2 + ", sq)
			FindSquares(n - sq * sq, amount + 1)
		}
	} else {
		for ; amount < 4; amount++ {
			if amount == 3 {
				fmt.Printf("%d^2", sq)
			} else {
				fmt.Printf("%d^2 + ", sq)
			}
		}
	}
}

func main() {
	var n uint64
	fmt.Println("Enter numbers, 4 squares of which you want to find")
	for {
		fmt.Print("Enter number:")
		fmt.Scan(&n)
		fmt.Printf("%d = ", n)
		FindSquares(n, 0)
		fmt.Println("\n")
	}
}
