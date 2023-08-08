package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var num_length uint64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d", &num_length)

	number_amount := find_phone_number_amount(num_length)
	fmt.Println(number_amount)
}

func find_phone_number_amount(num_length uint64) *big.Int {
	if num_length == 1 {
		return big.NewInt(8)
	}

	//5 is nonexistent, therefore dial_num_sum[7], for example, is 8
	var dial_num_sum [9]*big.Int = [9]*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(0),
		big.NewInt(1),
	}
	var temp_num_sum [9]*big.Int = return_copy(dial_num_sum)

	for i := uint64(1); i < num_length; i++ {
		temp_num_sum[0].Add(dial_num_sum[4], dial_num_sum[5])
		temp_num_sum[1].Add(dial_num_sum[5], dial_num_sum[7])
		temp_num_sum[2].Add(dial_num_sum[6], dial_num_sum[8])
		temp_num_sum[3].Add(dial_num_sum[4], dial_num_sum[7])

		temp_num_sum[4].Add(dial_num_sum[0], dial_num_sum[3])
		temp_num_sum[4].Add(temp_num_sum[4], dial_num_sum[8])

		temp_num_sum[5].Add(dial_num_sum[0], dial_num_sum[1]) //6
		temp_num_sum[5].Add(temp_num_sum[5], dial_num_sum[6])

		temp_num_sum[6].Add(dial_num_sum[2], dial_num_sum[5]) //7
		temp_num_sum[7].Add(dial_num_sum[1], dial_num_sum[3]) //8
		temp_num_sum[8].Add(dial_num_sum[2], dial_num_sum[4]) //9

		dial_num_sum = return_copy(temp_num_sum)
	}

	return sum(dial_num_sum)
}

func sum(a [9]*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, num := range a {
		sum.Add(sum, num)
	}
	return sum
}

func return_copy(a [9]*big.Int) [9]*big.Int {
	var b [9]*big.Int = [9]*big.Int{
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
	}

	b[0].Add(b[0], a[0])
	b[1].Add(b[1], a[1])
	b[2].Add(b[2], a[2])
	b[3].Add(b[3], a[3])
	b[4].Add(b[4], a[4])
	b[5].Add(b[5], a[5])
	b[6].Add(b[6], a[6])
	b[7].Add(b[7], a[7])
	b[8].Add(b[8], a[8])

	return b
}
