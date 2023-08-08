package main

import (
	"fmt"
	"math"
	"strings"
)

const infinity = math.MaxInt64 //Чтобы при сложении не происходило переполнения

type matrix struct {
	rows, columns int
	data          [][]way
}

type way struct {
	length   uint64
	transfer int
}

func newMatrixFrom(data [][]way) *matrix {
	matrix := &matrix{}
	matrix.rows = len(data)
	matrix.columns = len(data[0])
	matrix.data = make([][]way, matrix.rows)
	copy(matrix.data, data)
	return matrix
}

func newMatrix(rows, columns int) *matrix {
	matrix := &matrix{}
	matrix.rows = rows
	matrix.columns = columns

	matrix.data = make([][]way, rows)
	for i := range matrix.data {
		matrix.data[i] = make([]way, columns)
	}

	return matrix
}

func getData() *matrix {
	mat := newMatrix(10, 10)
	mat.data[0] = convertToWay(0, 339, 351, 255, 264, infinity, 315, 99, 309, infinity)
	mat.data[1] = convertToWay(1, 0, 265, infinity, 311, 207, 243, 336, 271, 233)
	mat.data[2] = convertToWay(infinity, 212, 0, 315, 321, 67, 462, 126, 23, 45)
	mat.data[3] = convertToWay(infinity, 29, 159, 0, 71, infinity, 309, 254, infinity, 107)
	mat.data[4] = convertToWay(infinity, 396, infinity, 153, 0, 103, 24, 220, 149, 304)
	mat.data[5] = convertToWay(401, 110, 340, 44, 258, 0, 108, 257, 179, 261)
	mat.data[6] = convertToWay(330, 457, 470, 315, infinity, 73, 0, 31, 292, 262)
	mat.data[7] = convertToWay(infinity, 28, infinity, 130, infinity, 428, infinity, 0, 44, 157)
	mat.data[8] = convertToWay(239, 335, 260, 232, 403, 10, infinity, 376, 0, 20)
	mat.data[9] = convertToWay(infinity, infinity, 78, 72, 90, 32, 154, 112, 236, 0)
	return mat
}

func convertToWay(a ...uint64) []way {
	res := make([]way, len(a))
	for i := range a {
		res[i] = way{length: a[i], transfer: 0}
	}
	return res
}

func (m matrix) ToString() string {
	if m.data == nil || len(m.data) == 0 {
		return "[]"
	}

	var s strings.Builder
	for i := 0; i < m.rows; i++ {
		s.WriteByte('[')
		for j := 0; j < m.columns; j++ {
			if m.data[i][j].transfer != 0 {
				m.data[i][j].transfer++
			}
			s.WriteString(" " + fmt.Sprintf("%v", m.data[i][j]) + " ")
		}
		s.WriteByte(']')
		s.WriteByte('\n')
	}
	return s.String()[:s.Len()-1]
}

func main() {
	mat := getData()
	oldres := oldTransform(mat)
	res := transform(mat)
	fmt.Printf("Old:\n%v\n\nNew:\n%v\n", oldres.ToString(), res.ToString())
}

func oldTransform(mat *matrix) *matrix {
	res := &matrix{}
	temp := newMatrixFrom(mat.data)

	for transfer := 0; transfer < mat.rows; transfer++ {
		res = newMatrixPassingThrough(mat, transfer)

		for start := 0; start < mat.rows; start++ {
			for end := 0; end < mat.columns; end++ {
				if temp.data[start][end].length <= temp.data[start][transfer].length+temp.data[transfer][end].length {
					res.data[start][end].length = temp.data[start][end].length
					res.data[start][end].transfer = temp.data[start][end].transfer
				} else {
					res.data[start][end].length = temp.data[start][transfer].length + temp.data[transfer][end].length
					res.data[start][end].transfer = transfer
				}
			}
		}

		temp = newMatrixFrom(res.data)
	}

	return res
}

func transform(mat *matrix) *matrix {
	res := newMatrixFrom(mat.data)

	for transfer := 0; transfer < res.rows; transfer++ {
		for start := 0; start < res.rows; start++ {
			if start != transfer {
				for end := 0; end < res.columns; end++ {
					if end != transfer {
						if res.data[start][end].length > res.data[start][transfer].length+res.data[transfer][end].length {
							res.data[start][end].length = res.data[start][transfer].length + res.data[transfer][end].length
							res.data[start][end].transfer = transfer
						}
					}
				}
			}
		}
	}

	return res
}

//newMatrixPassingThrough empty matrix with duplicated n-th rows and columns of mat
func newMatrixPassingThrough(mat *matrix, n int) *matrix {
	newMat := newMatrix(mat.rows, mat.columns)
	for pos := 0; pos < mat.rows; pos++ {
		newMat.data[pos][n] = mat.data[pos][n]
		newMat.data[n][pos] = mat.data[n][pos]
	}
	return newMat
}
