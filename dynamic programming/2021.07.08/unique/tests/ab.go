package main

import "fmt"

type stats struct {
	min int8
	max int8
	cur int8
}

func main() {
	word_pos := uint64(0)
	fmt.Scan(&word_pos)

	values := make([]map[stats]bool, 1)
	values[0] = make(map[stats]bool)
	values[0][stats{0, 0, 0}] = true

	total_words, words_with_len_n := uint64(0), uint64(0)
	for n := 1; total_words < word_pos; n++ {
		values = append(values, make(map[stats]bool))
		words_with_len_n, values = amount_words_with_len_n(values, n)
		total_words += words_with_len_n
	}

	sequence := make_sequence(values, words_with_len_n, word_pos-(total_words-words_with_len_n), int8(len(values)-1))
	fmt.Println(sequence)
}

func amount_words_with_len_n(values []map[stats]bool, n int) (uint64, []map[stats]bool) {
	var (
		amount_words uint64
		change       bool
	)

	for value := range values[n-1] {
		//test with added a
		value.cur++
		if value.cur > value.max {
			value.max++
			change = true
		}

		if value.max-value.min <= 3 {
			values[n][stats{min: value.min, max: value.max, cur: value.cur}] = true
			amount_words++
		}
		if change {
			value.max--
		}
		change = false

		//test with added b
		value.cur -= 2
		if value.cur < value.min {
			value.min--
			change = true
		}

		if value.max-value.min <= 3 {
			values[n][stats{min: value.min, max: value.max, cur: value.cur}] = true
			amount_words++
		}
		change = false
	}

	return amount_words, values
}

func make_sequence(values []map[stats]bool, words_with_len_n, word_pos uint64, n int8) string {
	var (
		sequence string
		cur_stat stats
	)
	fmt.Printf("words_with_len_n, word_pos: %d %d\n", words_with_len_n, word_pos)
	amnt := uint64(0)
	cur_stat = stats{0, 1, 1}

	{
		temp := make([]map[stats]bool, n+1)
		temp[0] = map[stats]bool{cur_stat: true}

		for j := int8(1); j <= n-1; j++ {
			temp[j] = make(map[stats]bool)
			amnt, temp = amount_words_with_len_n(temp, int(j))
		}
		fmt.Println("amnt, len(word)", amnt, n)
		for i := range temp {
			fmt.Println("temp", temp[i])
		}

		if amnt < word_pos {
			sequence = "b"
			cur_stat = stats{-1, 0, -1}
			fmt.Println("added b:", cur_stat)
			word_pos -= amnt
		} else {
			sequence = "a"
			cur_stat = stats{0, 1, 1}
			fmt.Println("added a:", cur_stat)
		}
		fmt.Printf("word_pos: %d\n", word_pos)
		fmt.Println()
	}

	for i := int8(2); i <= n; i++ {
		temp := make([]map[stats]bool, n+1)
		temp[0] = map[stats]bool{cur_stat: true}

		for j := int8(1); j <= n-i; j++ {
			temp[j] = make(map[stats]bool)
			amnt, temp = amount_words_with_len_n(temp, int(j))
		}

		fmt.Println("amnt, len(word)", amnt, n-i)
		for i := range temp {
			fmt.Println("temp", temp[i])
		}

		if amnt < word_pos {
			sequence += "b"
			cur_stat.cur--
			if cur_stat.cur < cur_stat.min {
				cur_stat.min--
			}
			fmt.Println("added b:", cur_stat)
			word_pos = (word_pos - amnt) + 1
		} else {
			sequence += "a"
			cur_stat.cur++
			if cur_stat.cur > cur_stat.max {
				cur_stat.max++
			}
			fmt.Println("added a:", cur_stat)
		}
		fmt.Printf("word_pos: %d\n", word_pos)
		fmt.Println()
	}

	return sequence
}
