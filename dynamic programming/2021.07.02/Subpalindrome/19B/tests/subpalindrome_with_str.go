package main

import (
	"bufio"
	"fmt"
	"os"
)

type window struct {
	is_palindrome bool
	palindrome    string
}

func main() {
	str := get_data()
	subpalindrome := find_max_subpalindrome(str)
	fmt.Printf("%d\n%s\n", len(subpalindrome), subpalindrome)
}

func get_data() (str string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%s", &str)
	return
}

func find_max_subpalindrome(str string) string {
	list := init_window_list(str)

	for window_size := 2; window_size < len(list); window_size++ {
		for left, right := 0, window_size; right < len(list); {
			if str[left] == str[right] {
				pal := return_biggest_palindrome_between(list, left, right)
				list[left][right] = window{is_palindrome: true, palindrome: string(str[left]) + pal + string(str[right])}
			}
			left++
			right++
		}
	}

	return return_biggest_palindrome_between(list, -1, len(list))
}

func init_window_list(str string) [][]window {
	list := make([][]window, len(str))
	list[0] = make([]window, len(str))
	list[0][0] = window{is_palindrome: true, palindrome: string(str[0])}

	for i := 1; i < len(list); i++ {
		list[i] = make([]window, len(str))
		list[i][i] = window{is_palindrome: true, palindrome: string(str[i])}
		if str[i-1] == str[i] {
			list[i-1][i] = window{is_palindrome: true, palindrome: string(str[i-1 : i+1])}
		}
	}

	return list
}

func return_biggest_palindrome_between(list [][]window, left, right int) string {
	max_palindrome := list[left+1][left+1].palindrome
	min_pos, max_pos := left+1, right-1
	delta := max_pos - min_pos

	for start, end := min_pos, max_pos; start != end; {
		if end > max_pos {
			start = min_pos
			delta--
			end = start + delta
		} else if list[start][end].is_palindrome {
			max_palindrome = list[start][end].palindrome
			break
		} else {
			start++
			end++
		}
	}

	return max_palindrome
}
