package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bracket_num, depth := get_data()
	comb := find_bracket_combinations(bracket_num, depth) - find_bracket_combinations(bracket_num, depth-1)
	fmt.Println(comb)
}

func get_data() (uint8, uint8) {
	var bracket_num, depth uint8
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &bracket_num, &depth)

	return bracket_num, depth
}

func find_bracket_combinations(bracket_num, depth uint8) uint64 {
	combs := make([][]uint64, (bracket_num*2)+1)
	for i := range combs {
		combs[i] = make([]uint64, depth+1)
	}
	combs[0][0] = 1

	for num := uint8(1); num <= bracket_num*2; num++ {
		combs[num][0] = combs[num-1][1]
		for dep := uint8(1); dep < depth; dep++ {
			combs[num][dep] = combs[num-1][dep-1] + combs[num-1][dep+1]
		}
		combs[num][depth] = combs[num-1][depth-1]
	}

	return combs[bracket_num*2][0]
}
