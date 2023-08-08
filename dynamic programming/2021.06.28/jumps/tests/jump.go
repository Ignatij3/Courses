package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	list, step_length := getList()
	fmt.Println(list)
	total_value, max_step := find_greatest_value_and_step(list, step_length)
	fmt.Println(total_value, max_step)
}

func getList() ([]int, int) {
	/*var (
		len_list    int
		step_length int
	)

	fmt.Scanf("%d %d\n", &len_list, &step_length)
	list := make([]int, len_list)*/

	list := make([]int, 20)
	add_data_to_list(list)
	//return list, step_length
	return list, 10
}

func add_data_to_list(list []int) {
	var data int
	/*reader := bufio.NewReader(os.Stdin)
	byte_num_slice, _ := reader.ReadBytes('\n')*/
	file, _ := os.Open("jump.01.in")
	reader := bufio.NewReader(file)
	reader.ReadLine()
	byte_num_slice, _ := reader.ReadBytes('\n')
	byte_num_slice = byte_num_slice[:len(byte_num_slice)-2]

	num_slice := strings.Split(string(byte_num_slice), " ")
	for pos, number := range num_slice {
		data, _ = strconv.Atoi(number)
		list[pos] = data
	}
}

func find_greatest_value_and_step(list []int, step_length int) (int, int) {
	var (
		max_step    int
		total_value int
	)

	for start_pos := 0; start_pos < len(list); {
		max_num, step := return_highest_number_and_step(list, step_length, &start_pos)
		if step > max_step {
			max_step = step
		}
		total_value += max_num
	}

	return total_value, max_step
}

func return_highest_number_and_step(list []int, step_length int, start_pos *int) (int, int) {
	var (
		max_num int = math.MinInt64
		step    int
		pos     int
		end_pos int = *start_pos + step_length
	)

	if end_pos >= len(list) {
		end_pos = len(list) - 1
	}

	for pos = *start_pos; pos < end_pos && max_num < 0; pos++ {
		max_num = list[pos]
	}
	fmt.Println("\nafter for", max_num, pos)

	step = (pos - *start_pos) + 1
	if pos == len(list)-1 && max_num < 0 {
		set_new_start_pos(list, start_pos, end_pos)
		//find greatest here
		//max_num = 0
	}

	if *start_pos >= len(list) {
		fmt.Printf("%v chosen %d end %d(%d)\n", list[:*start_pos], max_num, pos, list[pos])
	} else {
		fmt.Printf("%v chosen %d end %d(%d)\n", list[*start_pos:end_pos], max_num, pos, list[pos])
	}

	*start_pos = pos + 1

	return max_num, step
}

func set_new_start_pos(list []int, start_pos *int, end_pos int) {
	var (
		num     int = math.MinInt64
		num_pos int
	)

	for i := *start_pos; num < 0 && i < len(list); i++ {
		num = list[i]
		num_pos = i
	}

	difference := num_pos - end_pos

}
