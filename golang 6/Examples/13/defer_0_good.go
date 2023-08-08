package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(n int, filename string) error {
	// Создаём файл filename с числами от 1 до n
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	buff := bufio.NewWriter(file)
	defer buff.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintln(buff, i)
	}
	return nil
}

func ReadFile(filename string) ([]int, error) {
	var a []int
	// Считываем числа из файла filename
	file, err := os.Open(filename)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()
	buff := bufio.NewReader(file)
	var c int
	for {
		if _, err = fmt.Fscanf(buff, "%d\n", &c); err != nil {
			break
		}
		a = append(a, c)
	}
	return a, nil
}

func WriteReverse(filename string, a []int) error {
	// Выводим слайс a в обратном порядке в файл filename
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	buff := bufio.NewWriter(file)
	defer buff.Flush()
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Fprintln(buff, a[i])
	}
	return nil
}

func main() {
	// Создаём файл numbers0.dat
	CreateFile(1200, "numbers0.dat")

	// Считываем числа из файла numbers0.dat
	data, _ := ReadFile("numbers0.dat")

	// Выводим их в обратном порядке в файл numbers0.res
	WriteReverse("numbers0.res", data)
}
