package main
import (
	"fmt"
)

type Square [10][10]int

func main() {
	var (
		Box	Square
		height, width int = 10, 10
	)
	for y := 0 ; y < height; y++ {
		Box[y][0] = 1
	}
	for x, y := 0, 0 ; x < width && y < height; {
		Box[y][x] = 1
		y++
		x++
	}
	fmt.Println(Box)
}
