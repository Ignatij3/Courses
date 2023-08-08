package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	width  int16
	height int16
}

func main() {
	coordinates := get_data()
	length := calc_max_envelope_length(coordinates)
	fmt.Println(length)
}

func get_data() []point {
	var coordinates_amount uint16
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &coordinates_amount)
	coordinates := make([]point, coordinates_amount)

	for i := uint16(0); i < coordinates_amount; i++ {
		fmt.Fscanf(reader, "%d %d\n", &coordinates[i].width, &coordinates[i].height)
	}

	return coordinates
}

func calc_max_envelope_length(coords []point) uint32 {
	coords[0] = point{width: coords[0].height, height: coords[0].width}
	temp_coords := make([]point, 0)
	temp_coords = append(temp_coords, coords...)

	for pos := 1; pos < len(coords); pos++ {
		max := coords[pos].height + int16(math.Abs(float64(coords[pos-1].height-coords[pos].width)))
		temp := coords[pos-1].height + int16(math.Abs(float64(coords[pos].height-coords[pos-1].width)))

		if max > temp {
			temp_coords[pos] = point{width: coords[pos].height, height: coords[pos].width}
		}
	}
	coords = temp_coords

	total := uint32(coords[0].width)
	for pos := 1; pos < len(coords); pos++ {
		total += uint32(int16(math.Abs(float64(coords[pos].height-coords[pos-1].height))) + coords[pos].width)
	}

	return total
}
