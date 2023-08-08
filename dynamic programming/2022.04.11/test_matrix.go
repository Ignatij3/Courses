package main

import (
	"fmt"
	"matrix"
	"strconv"
)

func dotProduct() {
	fmtfloat := func(n float64) string { return strconv.FormatFloat(n, 'f', -1, 64) }

	a := &matrix.Matrix[float64]{} // PxC
	a.AppendRow([]float64{0.5, 13, 7, -8, 3})
	a.AppendRow([]float64{5, 1, 1, 3, 5})
	a.AppendRow([]float64{8, 7, 17, -0.7, -9})

	b, _ := matrix.NewMatrixFrom( // CxR
		[][]float64{
			{5, -3, -6},
			{5, -33, 6},
			{5, -43, 6},
			{5, -0.17, 6},
			{5, 1, -1},
		},
	)

	fmt.Printf("matrix A:\n%s\n\n", a.ToString(fmtfloat))
	fmt.Printf("matrix B:\n%s\n\n", b.ToString(fmtfloat))

	a.Dot(a, b, 0) // PxR
	fmt.Printf("result:\n%s\n\n", a.ToString(fmtfloat))
}

func fibonacci() { // 263308584
	fmtuint := func(n uint64) string { return fmt.Sprint(n) }

	fibmat, _ := matrix.NewMatrixFrom([][]uint64{{0, 1}, {1, 1}})
	fibmat.Power(fibmat, 369, 1e9+7)
	fmt.Printf("answer: %d\nmatrix:\n%s\n", fibmat.At(0, 1), fibmat.ToString(fmtuint))
}

func main() {
	dotProduct()
	fibonacci()
}
