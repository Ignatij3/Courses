package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	task00 = "../Files/task00.txt"
	task01 = "../Files/task01.txt"
	task02 = "../Files/task02.txt"
	task03 = "../Files/task03.txt"
	task04 = "../Files/task04.txt"
	task05 = "../Files/task05.txt"
	task06 = "../Files/task06.txt"
	task07 = "../Files/task07.txt"
	task08 = "../Files/task08.txt"
	task09 = "../Files/task09.txt"
	task10 = "../Files/task10.txt"
)

func isOk(a []float64) bool {
	for _, c := range a {
		if c != c || !math.IsInf(c, 0) {
			return false
		}
	}
	return true
}

func Calculate(slice [][]float64) ([]float64, bool) {
	var (
		multiplier, temp float64
		answers          []float64
	)

	for i := 0; i < len(slice); {
		answers = append(answers, 0.0)
	}

	for step := 0; step < len(slice); step++ {
		for row := 1 + step; row < len(slice); row++ {
			multiplier = slice[row][step]
			for col := step; col < len(slice[0]); col++ {
				//fmt.Printf("%v + %v(%v/%v) = ", slice[row][col], slice[step][col], -multiplier, slice[step][step])
				slice[row][col] += slice[step][col] * (-multiplier / slice[step][step])
				//fmt.Println(slice[row][col])
			}
			//fmt.Println()
		}
		if step == len(slice)-1 {
			break
		}
	}
	for row := len(slice) - 1; row >= 0; row-- {
		answers[row] = slice[row][row]
		for col := len(slice) - 1; col > row; col-- {
			temp = slice[row][col] * answers[col]
			slice[row][len(slice)] += -temp
		}
		answers[row] = slice[row][len(slice)] / answers[row]
	}

	return answers, isOk(answers)
}

func main() {
	var (
		a                 [][]float64
		c                 []float64
		row, pass, choose int
		filePath          string
	)

	fmt.Println("0) ", strings.Split(task00, "../Files/")[len(strings.Split(task00, "../Files/"))-1])
	fmt.Println("1) ", strings.Split(task01, "../Files/")[len(strings.Split(task01, "../Files/"))-1])
	fmt.Println("2) ", strings.Split(task02, "../Files/")[len(strings.Split(task02, "../Files/"))-1])
	fmt.Println("3) ", strings.Split(task03, "../Files/")[len(strings.Split(task03, "../Files/"))-1])
	fmt.Println("4) ", strings.Split(task04, "../Files/")[len(strings.Split(task04, "../Files/"))-1])
	fmt.Println("5) ", strings.Split(task05, "../Files/")[len(strings.Split(task05, "../Files/"))-1])
	fmt.Println("6) ", strings.Split(task06, "../Files/")[len(strings.Split(task06, "../Files/"))-1])
	fmt.Println("7) ", strings.Split(task07, "../Files/")[len(strings.Split(task07, "../Files/"))-1])
	fmt.Println("8) ", strings.Split(task08, "../Files/")[len(strings.Split(task08, "../Files/"))-1])
	fmt.Println("9) ", strings.Split(task09, "../Files/")[len(strings.Split(task09, "../Files/"))-1])
	fmt.Println("10) ", strings.Split(task10, "../Files/")[len(strings.Split(task10, "../Files/"))-1])

	fmt.Print("Выберите файл с уравнениями, который хотите решить (0 - 10): ")
	fmt.Scan(&choose)
	for choose < 0 && choose > 10 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз:")
		fmt.Scan(&choose)
	}
	switch choose {
	case 0:
		filePath = task00
	case 1:
		filePath = task01
	case 2:
		filePath = task02
	case 3:
		filePath = task03
	case 4:
		filePath = task04
	case 5:
		filePath = task05
	case 6:
		filePath = task06
	case 7:
		filePath = task07
	case 8:
		filePath = task08
	case 9:
		filePath = task09
	case 10:
		filePath = task10
	}
	fmt.Println()

	fin, _ := os.Open(filePath)
	defer fin.Close()

	for scanner := bufio.NewScanner(fin); scanner.Scan(); row++ {
		if pass > 0 {
			c = make([]float64, 0)
			for _, snum := range strings.Fields(scanner.Text()) {
				if x, err := strconv.ParseFloat(snum, 64); err == nil {
					c = append(c, x)
				}
			}
			fmt.Printf("Row #%2d: %v\n", row, c)
			a = append(a, c)
		} else {
			pass++
		}
	}
	fmt.Println()
	res, ok := Calculate(a)

	if ok {
		fmt.Println("Result-----------------------------------------------------------\n")
		for r, p := 0, 0; r < len(a); r++ {
			fmt.Printf("%d-е неизвестное - %v\n", p+1, res[r])
			p++
		}
		fmt.Println("\nResult-----------------------------------------------------------")
	}
}
