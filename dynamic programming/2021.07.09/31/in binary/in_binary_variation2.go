package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		num_pos     int32
		zero_amount uint8
	)

	fmt.Scanf("%d %d\n", &num_pos, &zero_amount)
	print_all_binary_with_n_zeros(zero_amount)

	num := find_nth_num(num_pos, zero_amount)
	fmt.Printf("on %d place: %s\n", num_pos, num)
}

func print_all_binary_with_n_zeros(zero_amount uint8) {
	if zero_amount < 64 {
		fmt.Print(1)
		for i := 0; i < int(zero_amount); i++ {
			fmt.Print(0)
		}
		fmt.Print("\n")
	}

	for bin_num_len := zero_amount + 2; bin_num_len <= 64; bin_num_len++ {
		bin_repr := make_bin_repr_with_len(zero_amount, bin_num_len)
		fmt.Printf("%s\n", bin_repr)

		for i := len(bin_repr) - 1; i > 1; i-- {
			if bin_repr[i-1] == '0' {
				bin_repr[i] = '0'
				bin_repr[i-1] = '1'

			} else {
				j := i
				for end_offset := 1; j > 0 && bin_repr[j-1] == '1'; end_offset++ {
					bin_repr[j] = '0'
					bin_repr[len(bin_repr)-end_offset] = '1'
					j--
				}

				if j <= 1 {
					break
				}
				bin_repr[j] = '0'
				bin_repr[j-1] = '1'
				i = len(bin_repr)
			}

			fmt.Printf("%s\n", bin_repr)
		}
	}
}

func make_bin_repr_with_len(zero_amount, bin_num_len uint8) []byte {
	bin_repr := make([]byte, 0)

	bin_repr = append(bin_repr, '1')
	bin_num_len--

	for i := uint8(0); i < zero_amount; i++ {
		bin_repr = append(bin_repr, '0')
		bin_num_len--
	}
	for i := uint8(0); i < bin_num_len; i++ {
		bin_repr = append(bin_repr, '1')
	}

	return bin_repr
}

func find_nth_num(n int32, zero_amount uint8) string {
	var (
		total_amount_nums int64
		bin_num_len       uint8
	)

	for bin_num_len = zero_amount + 1; ; bin_num_len++ {
		curr_len_quantity := c(int64(bin_num_len)-1, int64(zero_amount))
		total_amount_nums += curr_len_quantity
		if total_amount_nums >= int64(n) {
			n -= int32(total_amount_nums - curr_len_quantity)
			total_amount_nums = curr_len_quantity
			break
		}
	}

	var result string
	if zero_amount < 64 {
		result = "1"
		for length := bin_num_len - 1; length > 0; length-- {
			quantity_with_zero := c(int64(length)-1, int64(zero_amount)-1)

			if quantity_with_zero < int64(n) {
				result += "1"
				n -= int32(quantity_with_zero)

			} else if quantity_with_zero >= int64(n) {
				result += "0"
				zero_amount--

				if zero_amount == 0 {
					for i := length - 1; i > 0; i-- {
						result += "1"
					}
					break
				}
			}
		}
	}

	return result
}

func c(n, k int64) int64 {
	if k > n {
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
