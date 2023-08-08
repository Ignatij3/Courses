package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y int
}

func main() {
	coords, bridge_amount, bridge_len := get_coordinates()
	length := calculate_length(coords, bridge_amount, bridge_len)
	fmt.Printf("%.5f", length[len(length)-1][bridge_amount])
}

func get_coordinates() ([]point, int, int) {
	var (
		points_amount int
		bridge_amount int
		bridge_len    int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &points_amount, &bridge_amount, &bridge_len)
	coordinates := make([]point, points_amount)

	for i := 0; i < points_amount; i++ {
		fmt.Fscanf(reader, "%d %d\n", &coordinates[i].x, &coordinates[i].y)
	}

	return coordinates, bridge_amount, bridge_len
}

func calculate_length(coords []point, bridge_amount, bridge_len int) [][]float64 {
	list := make([][]float64, len(coords))
	list[0] = make([]float64, bridge_amount+1)
	for i := 1; i < len(list); i++ {
		list[i] = make([]float64, bridge_amount+1)
		list[i][0] = list[i-1][0] + distance(coords[i-1], coords[i])
	}

	for ppos := 1; ppos < len(list); ppos++ {
		for br_quantity := 1; br_quantity <= bridge_amount; br_quantity++ {
			list[ppos][br_quantity] = smallest_distance(list, coords, ppos, br_quantity, bridge_len)
		}
	}

	return list
}

func smallest_distance(list [][]float64, coords []point, ppos, br_quantity, bridge_len int) float64 {
	min := list[ppos][br_quantity-1]

	temp := list[ppos-1][br_quantity] + distance(coords[ppos-1], coords[ppos])
	if temp < min {
		min = temp
	}

	for start := 0; start < ppos-1; start++ {
		if bridge_is_valid(coords[start:ppos+1], bridge_len) {
			temp := list[start][br_quantity-1] + distance(coords[start], coords[ppos])
			if temp < min {
				min = temp
			}
		}
	}

	return min
}

func bridge_is_valid(coords []point, bridge_len int) bool {
	a := coords[0]
	b := coords[len(coords)-1]
	for _, c := range coords {
		if c_lies_higher_ab_line(a, b, c) {
			return false
		}
	}

	xdiff := a.x - b.x
	ydiff := a.y - b.y
	return xdiff*xdiff+ydiff*ydiff <= bridge_len*bridge_len
}

func c_lies_higher_ab_line(a, b, c point) bool {
	return (c.y-a.y)*(b.x-a.x) > (c.x-a.x)*(b.y-a.y)
}

func distance(a, b point) float64 {
	xdiff := float64(a.x - b.x)
	ydiff := float64(a.y - b.y)
	return math.Sqrt(xdiff*xdiff + ydiff*ydiff)
}
