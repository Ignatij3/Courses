package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var first, second, zero_amount int32
	fmt.Scanf("%d %d %d\n", &first, &second, &zero_amount)
	first_zeros := calculate_zeros(first, zero_amount)
	second_zeros := calculate_zeros(second, zero_amount)
	fmt.Println(second_zeros - first_zeros)
}

func calculate_zeros(top_num, zero_amount int32) int64 {
	top_num_binary := strconv.FormatInt(int64(top_num), 2)
	num_len := int64(len(top_num_binary))

	result := int64(0)
	for bin_num_len := zero_amount + 1; bin_num_len < int32(num_len); bin_num_len++ {
		result += c(int64(bin_num_len)-1, int64(zero_amount))
	}

	zeros_left := zero_amount
	reduced_num_len := num_len - 1
	for _, bit := range top_num_binary {
		reduced_num_len--
		if bit == '1' {
			result += c(reduced_num_len, int64(zeros_left)-1)
		} else {
			zeros_left--
		}
	}

	return result
}

func c(n, k int64) int64 {
	if n <= 0 || k <= 0 || k > n {
		return 0
	}

	n_fact := fact(n)
	k_fact := fact(k)
	k_n_fact := fact(n - k)

	total := big.NewInt(0)
	k_fact.Mul(k_fact, k_n_fact)
	total.Div(n_fact, k_fact)

	return total.Int64()
}

func fact(n int64) *big.Int {
	total := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		total.Mul(total, big.NewInt(i))
	}
	return total
}
