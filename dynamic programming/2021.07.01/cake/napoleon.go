package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start1 := time.Now()
	cakes := get_data()
	end1 := time.Now()

	start2 := time.Now()
	cakes_creme := make([][]bool, len(cakes))
	for pos, cake := range cakes {
		cakes_creme[pos] = make([]bool, len(cakes[pos]))
		cakes_creme[pos] = make_cake(cake)
	}
	end2 := time.Now()

	start3 := time.Now()
	write_output(cakes_creme)
	end3 := time.Now()

	fmt.Printf("input took %v\n", end1.Sub(start1))
	fmt.Printf("processing took %v\n", end2.Sub(start2))
	fmt.Printf("output took %v\n", end3.Sub(start3))
}

func get_data() [][]uint32 {
	var (
		data_sets   uint32
		cake_layers uint32
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &data_sets)
	cakes := make([][]uint32, data_sets)

	for i := uint32(0); i < data_sets; i++ {
		fmt.Fscanf(reader, "%d\n", &cake_layers)
		cakes[i] = make([]uint32, cake_layers)
		for j := uint32(0); j < cake_layers; j++ {
			fmt.Fscanf(reader, "%d ", &cakes[i][j])
		}
		fmt.Fscanf(reader, "\n")
	}

	return cakes
}

func make_cake(cake []uint32) []bool {
	cake_cream := make([]bool, len(cake))
	start_end_pos := make([]int, 0)

	start := len(cake) - 1
	stop := start + 1
	start_unassigned := true

	for i := start; i >= 0; i-- {
		if cake[i] > 0 {
			if start_unassigned {
				start = i
				start_unassigned = false
			}
			if i-int(cake[i])+1 < stop {
				stop = i - int(cake[i]) + 1
			}
		}
		if (stop != len(cake) && cake[i] == 0) || i == 0 {
			start_end_pos = append(start_end_pos, start)
			start_end_pos = append(start_end_pos, stop) //or 0

			start_unassigned = true
			stop = len(cake)
		}
	}

	for i := len(start_end_pos) - 1; i >= 0; i -= 2 {
		for start, end := start_end_pos[i], start_end_pos[i-1]; start >= end; start-- {
			cake_cream[start] = true
		}
	}

	return cake_cream
}

func write_output(cakes_cream [][]bool) {
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < len(cakes_cream); i++ {
		for j := 0; j < len(cakes_cream[i]); j++ {
			if cakes_cream[i][j] {
				writer.WriteByte('1')
				writer.WriteByte(' ')
			} else {
				writer.WriteByte('0')
				writer.WriteByte(' ')
			}
		}
		writer.WriteByte('\n')
	}
	writer.Flush()
}
