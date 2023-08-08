package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type contestant struct {
	beauty int32
	charm  int32
}

func main() {
	girls := get_data()
	thrown_out := find_minimal_sacrifices(girls)
	fmt.Println(thrown_out)
}

func get_data() []contestant {
	var girl_amount uint32

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &girl_amount)
	girls := make([]contestant, girl_amount)

	for i := range girls {
		fmt.Fscanf(reader, "%d %d\n", &girls[i].beauty, &girls[i].charm)
	}

	return girls
}

func find_minimal_sacrifices(girls []contestant) int {
	sort.Slice(girls, func(i, j int) bool { return girls[i].beauty > girls[j].beauty })
	lnds := find_LNDS(girls)

	return len(girls) - lnds
}

func find_LNDS(girls []contestant) int {
	lis_with_len := make([]int32, 1)
	lis_with_len[0] = math.MinInt32

	for _, girl := range girls {
		pos := get_pos_of_smaller_num(lis_with_len, girl.charm)

		if pos == len(lis_with_len)-1 {
			lis_with_len = append(lis_with_len, girl.charm)
		} else {
			lis_with_len[pos+1] = girl.charm
		}
	}

	return len(lis_with_len) - 1
}

func get_pos_of_smaller_num(lis_with_len []int32, charm int32) int {
	var (
		pos         int
		left, right int = 0, len(lis_with_len) - 1
		center      int
	)

	if charm > lis_with_len[right] {
		pos = right
	} else {
		for right-left > 1 {
			center = (left + right) / 2

			if lis_with_len[center] <= charm {
				left = center
			} else {
				right = center
			}
		}

		pos = left
	}

	return pos
}
