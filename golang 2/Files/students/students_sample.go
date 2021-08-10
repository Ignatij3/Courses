package main

import (
	"fmt"
	"os"
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

	//  Проверка: открываем файл students.bin на чтение, ...
   	file, err := os.Open("..\\Files\\students.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// ... читаем из него данные в буфер buff2, ...
	var buff2 bytes.Buffer
	_, err = buff2.ReadFrom(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// ... и декодируем их в data2 []tStudent
	var data2 []tStudent
	dec := gob.NewDecoder(&buff2)
	if err := dec.Decode(&data2); err != nil {
		fmt.Println(err)
		return
	}

	// Пример обработки данных: составляем список классов
	// и считаем количество учеников в каждом классе
	list := make(map[tGroup]int)
	for _, student := range data2  {
		list[student.Group]++	
	}
	// Выводим результаты
    for gr, amount := range list {
        fmt.Printf("%3d%c -> %d\n", gr.Year, gr.Index, amount)
    }
    fmt.Printf("Всего в школе %d классов\n", len(list))
}

