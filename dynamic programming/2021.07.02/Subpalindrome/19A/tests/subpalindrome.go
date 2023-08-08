package main

import (
	"bufio"
	"fmt"
	"os"
)

type window struct {
	is_palindrome     bool
	palindrome_length uint8
}

func main() {
	str := get_data()
	length := find_max_subpalindrome_length(str)
	fmt.Println(length)
}

func get_data() (str string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%s", &str)
	return
}

func find_max_subpalindrome_length(str string) uint8 {
	list := init_window_list(str)

	for window_size := 2; window_size < len(list); window_size++ {
		for left, right := 0, window_size; right < len(list); {
			if str[left] == str[right] {
				length := return_biggest_length(list, left, right)
				list[left][right] = window{is_palindrome: true, palindrome_length: length + 2}
			}
			left++
			right++
		}
	}

	return return_biggest_length(list, -1, len(list))
}

func init_window_list(str string) [][]window {
	list := make([][]window, len(str))
	list[0] = make([]window, len(str))
	list[0][0] = window{is_palindrome: true, palindrome_length: 1}

	for i := 1; i < len(list); i++ {
		list[i] = make([]window, len(str))
		list[i][i] = window{is_palindrome: true, palindrome_length: 1}
		if str[i-1] == str[i] {
			list[i-1][i] = window{is_palindrome: true, palindrome_length: 2}
		}
	}

	return list
}

func return_biggest_length(list [][]window, left, right int) uint8 {
	max := uint8(1)
	min_pos, max_pos := left+1, right-1
	delta := max_pos - min_pos

	for start, end := min_pos, max_pos; start != end; {
		if end > max_pos {
			start = min_pos
			delta--
			end = start + delta
		} else if list[start][end].is_palindrome {
			max = list[start][end].palindrome_length
			break
		} else {
			start++
			end++
		}
	}

	return max
}
