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

func EncodingData() {
	data, err01 := os.Open("..\\Files\\students\\students.bin")
	if err01 != nil {
		fmt.Println(err01)
		return
	}
	defer data.Close()
	var Rdata bytes.Buffer
	_, err01 = Rdata.ReadFrom(data)
	if err01 != nil {
		fmt.Println(err01)
		return
	}
	var list []tStudent
	decoder := gob.NewDecoder(&Rdata)
	err01 = decoder.Decode(&list)
	if err01 != nil {
		fmt.Println(err01)
		return
	}
	var NameNum map[string]int
	var Parralel map[int]int
	var NLastName map[string]int
	NameNum = make(map[string]int)
	Parralel = make(map[int]int)
	NLastName =  make(map[string]int)
	Count := -1
	for _, names := range list {
		NameNum[names.FirstName]++
		Count++
	}
	for _, students := range list {
		Parralel[students.Group.Year]++
	}
	for _, surname := range list {
		NLastName[surname.LastName]++
	}
	var max int
	for _, d := range NLastName {
		if d > max {
			max = d
		}
	}
	for index, student := range list {
		fmt.Printf("%d) %s %s %d%c\n", index, student.FirstName, student.LastName, student.Group.Year, student.Group.Index)
	}
	for x := 0; x <= 25; x++ {
		fmt.Printf("\n")
	}
	fmt.Println("В школе учится", Count, "ученика")
	fmt.Println("Вот количество учеников каждой парралели", Parralel)
	fmt.Println("Самые частые фамилии встречается", max, "раз \nВот они:")
	for surname, d := range NLastName {
		if d == max {
			fmt.Println(surname)
		}
	}
}
func main () {
	EncodingData()
}
