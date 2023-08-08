package main

import (
	"bufio"
	"fmt"
	"os"
)

type matrix struct {
	rows, columns int
	data          [][]uint64
}

func (m *matrix) appendRow(row []uint64) {
	if m.columns == 0 {
		m.columns = len(row)
	}

	m.data = append(m.data, append([]uint64{}, row...))
	m.rows++
}

func newMatrixFrom(data [][]uint64) *matrix {
	matrix := &matrix{}
	matrix.rows = len(data)
	matrix.columns = len(data[0])
	matrix.data = make([][]uint64, matrix.rows)
	copy(matrix.data, data)
	return matrix
}

func (m *matrix) power(mat *matrix, exp uint64) {
	*m = *newIdentityMatrix(m.rows)
	if exp == 0 {
		return
	}

	for exp > 0 {
		if exp&1 == 1 {
			m.dotMod(m, mat)
		}
		mat.dotMod(mat, mat)
		exp >>= 1
	}
}

func (m *matrix) dotMod(lhs *matrix, rhs *matrix) {
	mat := newMatrix(lhs.rows, rhs.columns)
	var sum uint64
	for row := 0; row < lhs.rows; row++ {
		for column := 0; column < rhs.columns; column++ {
			sum = 0
			for i := 0; i < lhs.columns; i++ {
				sum = (lhs.data[row][i]*rhs.data[i][column] + sum) % MODULO
			}
			mat.data[row][column] = sum
		}
	}
	*m = *mat
}

func newMatrix(rows, columns int) *matrix {
	matrix := &matrix{}
	matrix.rows = rows
	matrix.columns = columns

	matrix.data = make([][]uint64, rows)
	for i := range matrix.data {
		matrix.data[i] = make([]uint64, columns)
	}

	return matrix
}

func newIdentityMatrix(size int) *matrix {
	matrix := newMatrix(size, size)
	for i := 0; i < size; i++ {
		matrix.data[i][i] = 1
	}
	return matrix
}

func newEmptyMatrix() *matrix {
	return &matrix{rows: 0, columns: 0, data: [][]uint64{}}
}

const MODULO = 1e9 + 7

func getData() ([]uint64, uint64) {
	var (
		distances        [100]uint64
		dist, max, depth uint64
		n                int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &depth)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &dist)
		distances[dist-1]++
		if dist > max {
			max = dist
		}
	}

	return distances[:max], depth
}

func main() { //Создать identity matrix, затем в начало добавить строку с правилом и снизу добавить для суммы ряд
	distances, depth := getData()
	res := calculate(distances, depth)
	fmt.Println(res)
}

func calculate(distances []uint64, depth uint64) uint64 {
	mat := makeMatrix(distances, depth)
	mat.power(newMatrixFrom(mat.data), depth)
	return (mat.data[mat.rows-1][0] + 1) % MODULO
}

func makeMatrix(distances []uint64, depth uint64) *matrix {
	distances = append(distances, 0)
	tailSetter := make([]uint64, len(distances))
	tailSetter[0] = 1

	mat := newEmptyMatrix()
	mat.appendRow(distances)

	for i := 0; i < len(distances)-2; {
		mat.appendRow(tailSetter)
		i++
		tailSetter[i] = 1
		tailSetter[i-1] = 0
	}

	distances[len(distances)-1] = 1
	mat.appendRow(distances)

	return mat
}
