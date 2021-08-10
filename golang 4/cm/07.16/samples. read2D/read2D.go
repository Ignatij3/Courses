package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type  (
	tRow []float64
	tData2 []tRow
)	

func main() {
	var (
		a	tData2
		c	tRow
		row int
	)

	// Чтение данных из файла 2D.dat
	fin, _ := os.Open("task00.txt")
	defer fin.Close()

	for scanner := bufio.NewScanner(fin); scanner.Scan(); row++ {
		c = make([]float64, 0)
		for _, snum := range strings.Fields(scanner.Text()) {
			if x, err := strconv.ParseFloat(snum, 64); err == nil {
				c = append(c, x)
			}
		}
		fmt.Printf("Row #%2d: %v\n", row, c)
		a = append(a, c)
	}
	fmt.Println(a)
}
