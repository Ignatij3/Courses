package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	top_num uint32
)

func main() {
	number, output_len := get_data()

	combs := make([]uint32, len(number))
	for i := 1; i < len(combs); i++ {
		combs[i] = 1
	}

	for end := 1; end < len(combs); end++ {
		if number_is_valid(number[:end]) {
			combs[end] = combs[end-1] + uint32(len(number[:end]))
		} else {
			combs[end] = combs[end-1]
		}
	}

	output := combs[len(combs)-1]
	divisor := 0
	for i := uint8(0); i < output_len; i++ {
		divisor += 10
	}
	fmt.Println(output % uint32(divisor))
}

func get_data() (string, uint8) {
	var (
		num_len    uint16
		output_len uint8
		number     string
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &num_len, &top_num, &output_len)

	fmt.Fscanf(reader, "%s\n", &number)

	return number, output_len
}

func number_is_valid(a string) bool {
	n, _ := strconv.Atoi(a)
	return n <= int(top_num) && ((len(a) > 1 && a[0] != '0') || (len(a) == 1))
}
