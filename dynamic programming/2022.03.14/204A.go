package main

import (
	"fmt"
	"strconv"
)

const ASCII_ZERO = 48

func getData() (l, r uint64) {
	fmt.Scanf("%d %d", &l, &r)
	return
}

func main() {
	l, r := getData()
	res := calculate(l, r)
	fmt.Println(res)
}

func calculate(l, r uint64) uint64 {
	var res uint64
	larr := []byte(strconv.FormatUint(l, 10))
	rarr := []byte(strconv.FormatUint(r, 10))

	if larr[0] < larr[len(larr)-1] {
		l += (10 - (uint64(larr[len(larr)-1]) - ASCII_ZERO)) + (uint64(larr[0]) - ASCII_ZERO)
		larr = []byte(strconv.FormatUint(l, 10))

	} else if larr[0] > larr[len(larr)-1] {
		larr[len(larr)-1] = larr[0]
	}

	if rarr[0] > rarr[len(rarr)-1] {
		r -= (10 + uint64(rarr[len(rarr)-1]) - ASCII_ZERO) - (uint64(rarr[0]) - ASCII_ZERO)
		rarr = []byte(strconv.FormatUint(r, 10))

	} else if rarr[0] < rarr[len(rarr)-1] {
		rarr[len(rarr)-1] = rarr[0]
	}

	if len(larr)+1 < len(rarr) {
		for i := len(larr) + 1; i < len(rarr); i++ {
			res += 9 * exp(i-2)
		}
	}

	lcheck, _ := strconv.Atoi(string(larr))
	rcheck, _ := strconv.Atoi(string(rarr))
	if lcheck > rcheck {
		return 0
	}

	maxn := exp(len(larr)) - 1
	res += calcNum([]byte(strconv.FormatUint(uint64(maxn), 10))) - calcNum(larr) + 1

	return res + calcNum(rarr)
}

//Первый и последний элемент должны быть равны
func calcNum(a []byte) uint64 {
	if len(a) <= 2 {
		return uint64(a[0] - ASCII_ZERO)
	}

	border := uint64(a[0] - ASCII_ZERO)
	a = a[1 : len(a)-1]

	offset, _ := strconv.Atoi(string(a))
	return border*exp(len(a)) - (exp(len(a)) - 1 - uint64(offset))
}

func exp(n int) uint64 {
	var res uint64 = 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}
