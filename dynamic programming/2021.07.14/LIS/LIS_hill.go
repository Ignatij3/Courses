package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	list := get_list()
	lis, lds := make_subseq_lists(list)

	max_hill := find_max_hill(lis, lds)
	fmt.Println(max_hill)
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

func make_subseq_lists(list []int) ([]int, []int) {
	lis := make([]int, len(list))
	lds := make([]int, len(list))

	for lis_pos, lds_pos := 0, len(lds)-1; lds_pos >= 0; lis_pos, lds_pos = lis_pos+1, lds_pos-1 {
		lis[lis_pos] = find_top_lis(list, lis, lis_pos) + 1
		lds[lds_pos] = find_top_lds(list, lds, lds_pos) + 1
	}

	return lis, lds
}

func find_top_lis(list, lis []int, lis_pos int) int {
	top_lis := 0
	for i := lis_pos - 1; i >= 0; i-- {
		if top_lis < lis[i] && list[i] < list[lis_pos] {
			top_lis = lis[i]
		}
	}
	return top_lis
}

func find_top_lds(list, lds []int, lds_pos int) int {
	top_lds := 0
	for i := lds_pos + 1; i < len(lds); i++ {
		if top_lds < lds[i] && list[i] < list[lds_pos] {
			top_lds = lds[i]
		}
	}
	return top_lds
}

func find_max_hill(lis, lds []int) int {
	var max_hill int

	temp := 0
	for pos := range lis {
		temp = lis[pos] + lds[pos]
		if max_hill < temp {
			max_hill = temp
		}
	}

	return max_hill - 1
}
