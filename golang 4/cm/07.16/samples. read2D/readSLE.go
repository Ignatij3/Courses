package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int

	// Чтение данных из файла 
	fin, _ := os.Open("system.dat")
	defer fin.Close()

	fmt.Fscanln(fin, &N)

	a := make([][]float64, N, N)
	c := make([]float64, N, N)
	scanner := bufio.NewScanner(fin)

	for row := 0;  row < N; row++ {
		a[row] = make([]float64, N, N)
		scanner.Scan();
		col := 0
		for _, xstr := range strings.Fields(scanner.Text()) {
			if x, err := strconv.ParseFloat(xstr, 64); err == nil {
				if col < N  {
					a[row][col] = x
				} else {
					c[row] = x
				}	
			}
			col++
		}
	}
}
