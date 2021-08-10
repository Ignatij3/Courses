package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"bytes"
	"encoding/gob"
)

type  (
	tGroup struct {
		Year int
		Index rune
	}	
		
	tStudent struct  {
		FirstName string
		LastName string
		Group tGroup
	}	
)

func main() {

// Чтение данных в формате csv

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

// Конвертируем данные в формат struct tStudent

	// Считываем строки из файла students.csv, конвертируем 
	// их в tStudent и собираем в data []tStudent
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		student.FirstName, student.LastName = line[1], line[0]
		gr:= []rune(line[2])
		student.Group.Index = gr[len(gr)-1]
		gr = gr[:len(gr)-1]
		student.Group.Year, _ = strconv.Atoi(string(gr))
		data = append(data, student)
	}

// Записываем массив data []tStudent в файл students.bin сериализованно

	// gob-кодируем массив (слайс) данных data 
	var gobBuff bytes.Buffer
	enc := gob.NewEncoder(&gobBuff)
	err = enc.Encode(data) 
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Создаём файл students.bin ...
	fout, err := os.Create("students.bin")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    // ... и сбрасываем в него закодированные данные
    gobBuff.WriteTo(fout)
    if err != nil {
        fmt.Println(err)
        fout.Close()
        return
    }
    err = fout.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}
