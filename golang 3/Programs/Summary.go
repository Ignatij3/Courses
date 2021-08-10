package main
import (
		"fmt"
		)

func Adding(add []int) {
	for i := len(add) - 1; i < len(add); i-- {
		j := i - 1
		add[i] = add[i] + add[j]
		if j <= 0 {break}
	}
}

func main() {
	Array := []int {5, 12, 9, 2, 34, 22, 17, 8, 18, 48, 31, 20, 11, 19, 6, 51}
	Adding(Array)
	fmt.Println(Array)
}
