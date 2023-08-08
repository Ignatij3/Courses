package main

import (
	"fmt"
	"math/big"
)

func getInput() uint16 {
	var n uint16
	fmt.Scan(&n)
	return n
}

func calculate(n uint16) *big.Int {
	a := make([]*big.Int, (9*n)+1)
	for i := 0; i <= 9*int(n); i++ {
		a[i] = big.NewInt(1)
	}

	var start uint16
	for length := uint16(2); length <= n; length++ {
		for totalsum := 9 * uint16(length); totalsum > 0; totalsum-- {
			if totalsum < 9*(length-1) {
				start = 0
			} else {
				start = totalsum - 9*(length-1)
			}

			total := big.NewInt(0)
			for sum := start; sum < 10 && sum <= totalsum; sum++ {
				total = total.Add(total, a[totalsum-sum])
			}
			a[totalsum] = a[totalsum].Set(total)
		}
	}

	return sumPowers(a)
}

func sumPowers(a []*big.Int) *big.Int {
	var res *big.Int = big.NewInt(0)
	for _, n := range a {
		res = res.Add(res, n.Mul(n, n))
	}
	return res
}

func main() {
	n := getInput()
	res := calculate(n)
	fmt.Println(res.String())
}
