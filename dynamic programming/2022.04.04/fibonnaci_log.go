package main

import (
	"fmt"
	"matrix"
)

const MODULO = 1e9 + 7

func fibMatrixPower(n uint64) uint64 {
	result, _ := matrix.NewMatrixFrom([][]uint64{{0, 1}, {1, 1}})
	result.Power(result, n, MODULO)
	return result.At(0, 1)
}

func main() {
	fmt.Println(fibMatrixPower(369))
}
