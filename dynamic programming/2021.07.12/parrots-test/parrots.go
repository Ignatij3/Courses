package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	K = 100000000 //10*N, where N=8
	R = 65535
)

func main() {
	message_len, message := get_data()
	encoded_message := encode(message_len, message)
	shuffle(encoded_message)
	decode(encoded_message)
}

func get_data() (int, []uint) {
	var message_len int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &message_len)
	message := make([]uint, message_len)

	for i := range message {
		fmt.Fscanf(reader, "%d ", &message[i])
	}

	return message_len, message
}

func encode(message_len int, message []uint) []string {
	max_num := get_max(message)
	encoding_table := get_encoding_table(max_num)

	var encoded_message []string
	for _, num := range message {
		encoded_message = make([]string, 0)
		for cipher := len(encoding_table) - 1; cipher >= 0; cipher-- {
			encoded_message = append(encoded_message, encode_number(encoding_table, num))
		}
	}

	return encoded_message
}

func encode_number(encoding_table [][]int64, num uint) string {
	var (
		encoded   string
		start_pos int
	)

	for length := len(encoding_table) - 1; length >= 0; length-- {
		sum := int64(0)
		for single_int := start_pos; single_int <= R; single_int++ {
			sum += encoding_table[length][single_int]
			if sum >= int64(num) {
				encoded += strconv.Itoa(single_int)

				if single_int > start_pos {
					num -= uint(sum - encoding_table[length][single_int-1])
				}
				break
			}
		}
	}

	return encoded
}

func get_encoding_table(max_num int) [][]int64 {
	encoding_table := make([][]int64, 1)

	for i := range encoding_table[0] {
		encoding_table[0][i] = 1
	}

	for length := 1; length <= K; length++ {
		encoding_table = append(encoding_table, make([]int64, R+1))
		encoding_table[length][R] = 1

		total := int64(1)
		for first := R - 1; first >= 0; first-- {
			encoding_table[length][first] = encoding_table[length-1][first] + encoding_table[length][first+1]
			total += encoding_table[length][first]
		}

		if total >= int64(max_num) {
			break
		}
	}

	return encoding_table
}

func get_max(message []uint) int {
	max := uint(0)
	for _, num := range message {
		if num > max {
			max = num
		}
	}
	return int(max)
}

func shuffle(encoded_message []string) {
	rand.Seed(time.Now().UnixNano())
	var left, right int
	for i := 0; i < 10000; i++ {
		left = rand.Intn(len(encoded_message))
		right = rand.Intn(len(encoded_message))
		encoded_message[left], encoded_message[right] = encoded_message[right], encoded_message[left]
	}
}

func decode(encoded_message []string) {

}
