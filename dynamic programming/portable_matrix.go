package main

import (
	"fmt"
)

type Matrix struct {
	rows, columns int
	data          [][]uint64
}

func newIdentityMatrix(size int) *Matrix {
	matrix := newMatrix(size, size)
	for i := 0; i < int(size); i++ {
		matrix.data[i][i] = 1
	}
	return matrix
}

func newMatrixFrom(data [][]uint64) *Matrix {
	matrix := &Matrix{}
	matrix.rows = len(data)
	matrix.columns = len(data[0])
	matrix.data = make([][]uint64, matrix.rows)
	copy(matrix.data, data)
	return matrix
}

func (m *Matrix) dot(lhs *Matrix, rhs *Matrix, modulo uint64) {
	mat := newMatrix(lhs.rows, rhs.columns)
	for row := 0; row < lhs.rows; row++ {
		for column := 0; column < rhs.columns; column++ {
			for i := 0; i < lhs.columns; i++ {
				mat.data[row][column] = (lhs.data[row][i]*rhs.data[i][column] + mat.data[row][column]) % modulo
			}
		}
	}
	*m = *mat
}

func (m *Matrix) power(base *Matrix, exp uint64, modulo uint64) {
	*m = *newIdentityMatrix(m.rows)
	for exp > 0 {
		if exp&1 == 1 {
			m.dot(m, base, modulo)
		}
		base.dot(base, base, modulo)
		exp >>= 1
	}
}

func newMatrix(rows, columns int) *Matrix {
	matrix := &Matrix{}
	matrix.rows = rows
	matrix.columns = columns

	matrix.data = make([][]uint64, rows)
	for i := range matrix.data {
		matrix.data[i] = make([]uint64, columns)
	}

	return matrix
}

func main() {
	matrix := newMatrixFrom([][]uint64{{0, 1}, {1, 1}})

	// newMatrixFrom(matrix.data) обязательно, иначе в функции matrix перезапишется
	matrix.power(newMatrixFrom(matrix.data), 369, 1e9+7)
	fmt.Println(matrix.data[0][1]) // 263308584
}
