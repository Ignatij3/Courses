package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	teacherData, _ := os.Open("../Files/TeacherData.json")
	studentData, _ := os.Open("../Files/StudentData.json")
	directorData, _ := os.Open("../Files/DirectorData.json")

	teacherSlice, _ := ioutil.ReadAll(teacherData)
	studentSlice, _ := ioutil.ReadAll(studentData)
	directorSlice, _ := ioutil.ReadAll(directorData)

	json.Unmarshal(teacherSlice, &tData)
	json.Unmarshal(studentSlice, &sData)
	json.Unmarshal(directorSlice, &dData)
}

var (
	tData []teacher
	sData []student
	dData director
)

type (
	school interface {
		create()
		update()
		read()
		delete()
	}
	actionData struct {
		Action   string   `json:"Action"`
		Teacher  teacher  `json:"Teacher"`
		Student  student  `json:"Student"`
		Director director `json:"Director"`
	}
	person struct {
		ID      int    `json:"ID"`
		Name    string `json:"Name"`
		Surname string `json:"Surname"`
	}
	teacher struct {
		Person    person  `json:"Person"`
		Subject   string  `json:"Subject"`
		Salary    float64 `json:"Salary"`
		Classroom int     `json:"Classroom"`
	}
	student struct {
		Person person `json:"Person"`
		Class  int    `json:"Class"`
	}
	director struct {
		Person person `json:"Person"`
		Room   int    `json:"Room"`
		Phone  string `json:"Phone"`
	}
)

func (teacher) create(tData *[]teacher, acD actionData) {
	newID := (*tData)[len(*tData)-1].ID + 1
	/*d = append(*d, a.PData)
	(*d)[len(*d)-1].ID = newID*/
}

func (student) create() {
	newID := 0
	/*for range *d {
		newID++
	}*/
	/*d = append(*d, a.PData)
	(*d)[len(*d)-1].ID = newID*/
}

func (director) create() {
	newID := 0
	/*for range *d {
		newID++
	}*/
	/*d = append(*d, a.PData)
	(*d)[len(*d)-1].ID = newID*/
}

/*
func () update() {
	for n, c := range *d {
		if c.ID == a.PData.ID {
			(*d)[n] = a.PData
			return
		}
	}
}

func () read() {
	for _, c := range *d {
		fmt.Printf("id - %d\nname, surname, personal code - %s, %s (%s)\nsubject - %s\nsalary - %f\nclassrooms - %v\n\n", c.ID, c.Person.Name, c.Person.Surname, , c.Subject, c.Salary, c.Classrooms)
	}
}

func () delete() {
	for n, c := range *d {
		if c.ID == a.PData.ID {
			copy((*d)[n:], (*d)[n+1:])
			*d = (*d)[:len(*d)-1]
			return
		}
	}
}*/

func commitActions(actions []actionData) tData {
	var teachers tData
	for _, c := range actions {
		switch c.Action {
		case "create":
			teachers.create(c)
		case "update":
			teachers.update(c)
		case "read":
			teachers.read(c)
		case "delete":
			teachers.delete(c)
		}
	}
	return teachers
}

func main() {
	var (
		aData    []actionData
		singData actionData
	)

	file, _ := os.Open("../Files/ActionData.json")
	byteSlice, _ := ioutil.ReadAll(file)

	dec := json.NewDecoder(strings.NewReader(string(byteSlice)))
	dec.Token()
	for dec.More() {
		dec.Decode(&singData)
		aData = append(aData, singData)
	}
	dec.Token()

	teachers := commitActions(aData)
	fmt.Println("teachers", teachers)
}
