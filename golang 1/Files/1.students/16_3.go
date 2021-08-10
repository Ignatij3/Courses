package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type (
	tGroup struct {
		year int
		index rune
	}	
		
	tStudent struct  {
		firstName string
		lastName string
		group tGroup
	}	
)

func main() {

// Ввод данных

	fin, err := os.Open("students.csv")
	if err != nil {
		return
	}
	defer fin.Close()

	// Create a new Scanner for the input
	scanner := bufio.NewScanner(fin)
	var  (
		student tStudent
		data []tStudent
	)	

	// Считываем строки из файла students.csv, конвертируем 
	// их в tStudent и собираем в data []tStudent
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		student.firstName, student.lastName = line[1], line[0]
		gr:= []rune(line[2])
		student.group.index = gr[len(gr)-1]
		gr = gr[:len(gr)-1]
		student.group.year, _ = strconv.Atoi(string(gr))
		data = append(data, student)
	}
	
	// Составляем список классов и считаем количество учеников в каждом классе
	list := make(map[tGroup]int)
	for _, student := range data  {
		list[student.group]++	
	}
	// Выводим результаты
    for gr, amount := range list {
        fmt.Printf("%3d%c -> %d\n", gr.year, gr.index, amount)
    }
	
}

