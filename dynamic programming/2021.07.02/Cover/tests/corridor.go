package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	width, height := get_data()
	layout_num := calculate_layout_amount(width, height)
	fmt.Println(layout_num)
}

func get_data() (uint8, uint8) {
	var width, height uint8
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &width, &height)

	return width, height
}

func calculate_layout_amount(width, height uint8) uint64 {
	layouts := make([]uint64, height)
	for i := range layouts[:width-1] {
		layouts[i] = 1
	}
	layouts[width-1] = 2

	for i := width; i < height; i++ {
		layouts[i] = layouts[i-1] + layouts[uint8(i)-width]
	}

	return layouts[height-1]
}
