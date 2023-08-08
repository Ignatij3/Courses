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

func newMatrixFrom(data [][]uint64) *matrix {
	matrix := &matrix{}
	matrix.rows = len(data)
	matrix.columns = len(data[0])
	matrix.data = make([][]uint64, matrix.rows)
	copy(matrix.data, data)
	return matrix
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

func allPoints(cities int) []int {
	res := make([]int, cities)
	for i := range res {
		res[i] = i
	}
	return res
}

func getData() {
	var cities int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &cities)

	mat := newMatrix(cities, cities)
	for i := range mat.data {
		for j := range mat.data[i] {
			fmt.Fscanf(reader, "%d", &mat.data[i][j])
		}
		fmt.Fscanf(reader, "\n")
	}

	var length, lastSum uint64
	var start, end int

	mat = transform(mat, allPoints(cities)...)
	lastSum = roadSum(mat)

	fmt.Fscanf(reader, "%d\n", &cities)
	for i := 0; i < cities; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &start, &end, &length)
		if mat.data[start-1][end-1] > length {
			mat.data[start-1][end-1] = length
			mat.data[end-1][start-1] = length
			mat = transform(mat, start-1, end-1)
			lastSum = roadSum(mat)
		}

		fmt.Println(lastSum)
	}
}

func transform(mat *matrix, transferPoints ...int) *matrix {
	res := newMatrixFrom(mat.data)

	for _, transfer := range transferPoints {
		for start := 0; start < res.rows; start++ {
			if start != transfer {
				for end := start + 1; end < res.columns; end++ {
					if end != transfer {
						if res.data[start][end] > res.data[start][transfer]+res.data[transfer][end] {
							res.data[start][end] = res.data[start][transfer] + res.data[transfer][end]
							res.data[end][start] = res.data[start][transfer] + res.data[transfer][end]
						}
					}
				}
			}
		}
	}

	return res
}

func roadSum(mat *matrix) uint64 {
	res := uint64(0)
	for i := range mat.data {
		for j := i + 1; j < mat.columns; j++ {
			res += mat.data[i][j]
		}
	}
	return res
}

func main() {
	getData()
}
