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

type (
	tRow   []float64
	tData2 []tRow
)

func Calculate(slice tData2) {
	var (
		multiplier, add, top float64
		answers              [30]float64
		originRow, originCol [30]int
		num                  int
	)

	for step := 0; step < len(slice); step++ { //Для пошагового продвижения вперёд
		for newRow := step; newRow < len(slice); newRow++ { //Ходит по линиям
			for newCol := step; newCol < len(slice[0])-1; newCol++ { //Ходит по столбцам
				if top < math.Abs(slice[newRow][newCol]) { //Проверяется, меньше ли top, чем число из слайса
					top = math.Abs(slice[newRow][newCol])           //Если меньше, то оно становится тем числом
					originRow[num], originCol[num] = newRow, newCol //И записываются координаты
				}
			}
		}

		if originRow[num] != step { //Если наибольшее число не находится на той же линии, что и угловое (на которое делится),
			for changeRow := step; changeRow < len(slice[0]); changeRow++ { // то меняю линии местами
				slice[step][changeRow], slice[originRow[num]][changeRow] = slice[originRow[num]][changeRow], slice[step][changeRow]
			} // ^^^^^^^ Тут прохожусь по линии (меняя индекс столбика) и меняю местами с "начальной" линией
		}
		if originCol[num] != step { //Если наибольшее число не находится в том же столбе, что и угловое (на которое делится),
			for changeCol := step; changeCol < len(slice); changeCol++ { // то меняю столбики местами
				slice[changeCol][step], slice[changeCol][originCol[num]] = slice[changeCol][originCol[num]], slice[changeCol][step]
			} // ^^^^^^^ Тут прохожусь по столбикам (меняя индекс линий) и меняю местами с "начальным" столбиком
		}

		for row := 1 + step; row < len(slice); row++ {
			multiplier = slice[row][step] //Выбирается число, которое делится на угловое число
			for col := step; col < len(slice[0]); col++ {
				//fmt.Printf("%v + %v(%v/%v) = ", slice[row][col], slice[step][col], -multiplier, slice[step][step])
				slice[row][col] += slice[step][col] * (-multiplier / slice[step][step]) //Это умножается на число на линии
				//fmt.Println(slice[row][col])
			}
			//fmt.Println()
		}
		top = 0 //Сбрасываю top
		num++   //Увеличиваю num, для того, что бы координаты не мешались
	}

	fmt.Println("-------------------------------------------------------------------\n")
	for r := 0; r < len(slice); r++ {
		for c := 0; c < len(slice[0]); c++ {
			fmt.Print(slice[r][c], " ") //Прохожусь по слайсу, что-бы показать результаты прямого хода
		}
		fmt.Println()
	}
	fmt.Println("\n-------------------------------------------------------------------\n")

	for row := len(slice) - 1; row >= 0; row-- {
		answers[row] = slice[row][row] //Ответам присваивается число, неизвестное которого надо найти
		for col := len(slice) - 1; col > row; col-- {
			add = slice[row][col] * answers[col] //Здесь в переменную загоняется результат произведения числа и вычисленной неизвестной для него
			slice[row][len(slice)] += -add       //И добавляется к тому числу, что стоит за = (Поэтому и -add)
		}
		answers[row] = slice[row][len(slice)] / answers[row] //В ту же ячейку идёт частное числа по ту сторону от = и самого числа
	} //(которое теперь отражает нужное неизвестное)

	fmt.Println("Result-----------------------------------------------------------\n")
	for r, p := 0, 1; r < len(slice); r++ {
		fmt.Printf("%d-е неизвестное - %v\n", p, answers[r]) //Выводится окончательный результат
		p++
	}
	fmt.Println("\nResult-----------------------------------------------------------")
}

func main() { //Preview
	var (
		a                 tData2
		c                 tRow
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
	fmt.Scan(&choose) //Тут выбирается, какой файл посмотреть
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

	fin, _ := os.Open(filePath) //Тут открывается файл
	defer fin.Close()

	for scanner := bufio.NewScanner(fin); scanner.Scan(); row++ {
		if pass > 0 {
			c = make([]float64, 0)
			for _, snum := range strings.Fields(scanner.Text()) {
				if x, err := strconv.ParseFloat(snum, 64); err == nil {
					c = append(c, x) //Тут он выбирает по одной цифре (в строчку)
				}
			}
			fmt.Printf("Row #%2d: %v\n", row, c)
			a = append(a, c) //Тут он построчно добавляет в слайс
		} else {
			pass++
		}
	}
	fmt.Println()
	Calculate(a) //Тут я передаю в функцию весь слайс и строчку (последнюю)
}
