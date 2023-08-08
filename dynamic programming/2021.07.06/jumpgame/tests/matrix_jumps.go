package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WIN    byte = 'W'
	DRAW   byte = 'R'
	DEFEAT byte = 'D'
)

type point struct {
	row    uint16
	column uint16
}
type cell struct {
	state           byte
	accessible_from []point
}

func main() {
	matrix, starting_point := get_data()
	fmt.Println(matrix)
	cell_map := map_all_cells(matrix)
	print_output(cell_map, starting_point)
}

func get_data() ([][]int16, point) {
	var rows, columns uint16

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &rows, &columns)
	matrix := make([][]int16, rows)

	for i := range matrix {
		matrix[i] = make([]int16, columns)
		for j := 0; j < len(matrix[i])-1; j++ {
			fmt.Fscanf(reader, "%d ", &matrix[i][j])
		}
		fmt.Fscanf(reader, "%d\n", &matrix[i][len(matrix[i])-1])
	}

	var starting_point point
	fmt.Fscanf(reader, "%d %d\n", &starting_point.row, &starting_point.column)
	starting_point.row--
	starting_point.column--

	return matrix, starting_point
}

func map_all_cells(matrix [][]int16) [][]cell {
	rows, columns := int16(len(matrix)), int16(len(matrix[0]))

	cell_map := make([][]cell, rows)
	for i := range cell_map {
		cell_map[i] = make([]cell, columns)
	}

	move := int16(0)
	for row := range cell_map {
		for column := range cell_map[0] {
			move = matrix[row][column]
			if move > 0 {
				if int16(row)+move >= rows && int16(column)+move >= columns {
					cell_map[row][column].state = DEFEAT
				} else {
					if int16(row)+move < rows {
						cell_map[int16(row)+move][column].accessible_from = append(cell_map[int16(row)+move][column].accessible_from,
							point{row: uint16(row), column: uint16(column)})
					}
					if int16(column)+move < columns {
						cell_map[move][int16(column)+move].accessible_from = append(cell_map[row][int16(column)+move].accessible_from,
							point{row: uint16(row), column: uint16(column)})
					}
				}
			} else if move < 0 {
				if int16(row)+move < 0 && int16(column)+move < 0 {
					cell_map[row][column].state = DEFEAT
				} else {
					if int16(row)+move >= 0 {
						cell_map[int16(row)+move][column].accessible_from = append(cell_map[int16(row)+move][column].accessible_from,
							point{row: uint16(row), column: uint16(column)})
					}
					if int16(column)+move >= 0 {
						cell_map[row][int16(column)+move].accessible_from = append(cell_map[row][int16(column)+move].accessible_from,
							point{row: uint16(row), column: uint16(column)})
					}
				}
			} else {
				cell_map[row][column].state = DRAW
			}
		}
	}

	for row := range cell_map {
		for column := range cell_map[0] {
			if cell_map[row][column].state == 0 {
				cell_map[row][column].state = DRAW
			} else if cell_map[row][column].state == DEFEAT {
				for _, point := range cell_map[row][column].accessible_from {
					cell_map[point.row][point.column].state = WIN
				}
			}
		}
	}

	return cell_map
}

func print_output(cell_map [][]cell, starting_point point) {
	fmt.Printf("WIN: %v\nDRAW: %v\nDEFEAT: %v\n\n", WIN, DRAW, DEFEAT)
	fmt.Println(cell_map, starting_point)
	switch cell_map[starting_point.row][starting_point.column].state {
	case WIN:
		fmt.Println("Anton")
	case DRAW:
		fmt.Println("draw")
	case DEFEAT:
		fmt.Println("Yasha")
	}
}
