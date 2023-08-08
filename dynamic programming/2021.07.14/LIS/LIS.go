package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	list := get_list()
	lis := find_LIS(list)

	fmt.Println(lis)
}

func get_list() []int {
	var num_amount uint8

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &num_amount)
	list := make([]int, num_amount)

	for i := range list {
		fmt.Fscanf(reader, "%d ", &list[i])
	}

	return list
}

func find_LIS(list []int) int {
	lis_with_len := make([]int, 1)
	lis_with_len[0] = math.MinInt64

	for _, num := range list {
		pos := get_pos_of_smaller_num(lis_with_len, num)

		if pos == len(lis_with_len)-1 {
			lis_with_len = append(lis_with_len, num)
		} else {
			lis_with_len[pos+1] = num
		}
	}

	return len(lis_with_len) - 1
}

func get_pos_of_smaller_num(lis_with_len []int, num int) int {
	var (
		pos         int
		left, right int = 0, len(lis_with_len) - 1
		center      int
	)

	if num > lis_with_len[right] {
		pos = right
	} else {
		for right-left > 1 {
			center = (left + right) / 2

			if lis_with_len[center] < num {
				left = center
			} else {
				right = center
			}
		}

		pos = left
	}

	return pos
}
