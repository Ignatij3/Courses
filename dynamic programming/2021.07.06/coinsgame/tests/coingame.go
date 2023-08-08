package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	coin_slice, max_move_len := get_data()
	max_difference := calc_max_difference_matrix(coin_slice, max_move_len)
	num_sum := sum(coin_slice)
	max_win := (max_difference[0][max_move_len-1] + num_sum) / 2
	fmt.Println(max_win)
}

func get_data() ([]int64, uint8) {
	var (
		stack_amount uint8
		max_move_len uint8
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d ", &stack_amount)
	coin_slice := make([]int64, stack_amount)

	for i := range coin_slice {
		fmt.Fscanf(reader, "%d ", &coin_slice[i])
	}
	fmt.Fscanf(reader, "%d ", &max_move_len)

	return coin_slice, max_move_len
}

func calc_max_difference_matrix(coin_slice []int64, max_move_len uint8) [][]int64 {
	diff_len := len(coin_slice)
	max_difference := make([][]int64, diff_len)

	max_difference[diff_len-1] = make([]int64, max_move_len)
	for i := range max_difference[diff_len-1] {
		max_difference[diff_len-1][i] = coin_slice[diff_len-1]
	}

	for i := diff_len - 2; i >= 0; i-- {
		max_difference[i] = make([]int64, max_move_len)
		for j := range max_difference[i] {
			sum_obtained_nums := int64(0)
			for stack_pos := i; stack_pos <= i+j; stack_pos++ {
				if stack_pos < diff_len {
					sum_obtained_nums += coin_slice[stack_pos]
				}
			}
			if j == 0 {
				max_difference[i][j] = int64(math.MinInt64)
			} else {
				max_difference[i][j] = max_difference[i][j-1]
			}
			var tmp int64
			if i+j+1 < diff_len {
				tmp = sum_obtained_nums - max_difference[i+j+1][j]
			} else {
				tmp = sum_obtained_nums
			}
			if tmp > max_difference[i][j] {
				max_difference[i][j] = tmp
			}
		}
	}

	return max_difference
}

func sum(a []int64) int64 {
	sum := int64(0)
	for _, num := range a {
		sum += num
	}
	return sum
}
