// ВНИМАНИЕ! Пример работает некорректно из-за слишком 
// позднего сбрасывания буфера вывода и закрытия файла.
// Корректная версия - в примере defer_0_good.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N:= 1200
	// Создаём файл numbers0.dat с числами от 1 до N
	file, err := os.Create("numbers0.dat")
	if err != nil {
		return
	}
	defer file.Close()
	buffCreate := bufio.NewWriter(file)
	defer buffCreate.Flush()
	for i := 1; i <= N; i++ {
		fmt.Fprintln(buffCreate, i)
	}

	// Считываем числа из файла numbers0.dat
	file, err = os.Open("numbers0.dat")
	if err != nil {
		return
	}
	defer file.Close()
	buffRead := bufio.NewReader(file)
	var (
		data []int
		c    int
	)
	for {
		if _, err = fmt.Fscanf(buffRead, "%d\n", &c); err != nil {
			break
		}
		data = append(data, c)
	}

	// Выводим их в обратном порядке в файл numbers0.res
	file, err = os.Create("numbers0.res")
	if err != nil {
		return
	}
	defer file.Close()
	buffWrite := bufio.NewWriter(file)
	defer buffWrite.Flush()
	for i := len(data) - 1; i >= 0; i-- {
		fmt.Fprintln(buffWrite, data[i])
	}
}
