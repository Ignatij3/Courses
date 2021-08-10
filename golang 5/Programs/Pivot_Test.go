package main

import (
	"fmt"
)

func main() {
	for pivot := 1; pivot < 21; pivot++ {
		fmt.Printf("(%d - 1) / 2 - %d\n", pivot, (pivot - 1) / 2)
	}
}
