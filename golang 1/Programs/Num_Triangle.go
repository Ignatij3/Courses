package main
import (
	"fmt"
)

type Square [50][50]int

func main() {
	var (
		Box	Square
		height, width int = 10, 10
	)
	for y := 0 ; y < height; y++ {
		Box[y][0] = 1
	}
	for x := 0 ; x < width; x++ {
		Box[0][x] = 1
	}
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			Box[y][x] = Box[y - 1][x] + Box[y][x - 1]
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%10d", Box[y][x])
		}
		fmt.Println()
	}
	
	fmt.Println()
}
