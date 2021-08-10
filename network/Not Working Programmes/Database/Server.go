package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

func init() {
	teacherData, _ := os.Open("../../Files/TeacherData.json")
	studentData, _ := os.Open("../../Files/StudentData.json")
	directorData, _ := os.Open("../../Files/DirectorData.json")

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
		listIDs()
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
		Person     person  `json:"Person"`
		Subject    string  `json:"Subject"`
		Salary     float64 `json:"Salary"`
		Classrooms []int   `json:"Classrooms"`
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

func (t teacher) create(tData *[]teacher) {
	newID := (*tData)[len(*tData)-1].Person.ID + 1
	*tData = append(*tData, t)
	(*tData)[len(*tData)-1].Person.ID = newID
}

func (s student) create(sData *[]student) {
	newID := (*sData)[len(*sData)-1].Person.ID + 1
	*sData = append(*sData, s)
	(*sData)[len(*sData)-1].Person.ID = newID
}

func (d director) create(dData *director) {
	*dData = d
}

func (t teacher) update(tData *[]teacher) {
	for n, c := range *tData {
		if c.Person.ID == t.Person.ID {
			(*tData)[n] = t
			return
		}
	}
	fmt.Println("ERROR, ID not found")
}

func (s student) update(sData *[]student) {
	for n, c := range *sData {
		if c.Person.ID == s.Person.ID {
			(*sData)[n] = s
			return
		}
	}
	fmt.Println("ERROR, ID not found")
}

func (d director) update(dData *director) {
	*dData = d
}

func (teacher) read(tData []teacher, id int) {
	for _, c := range tData {
		if c.Person.ID == id {
			fmt.Printf("id - %d\nname, surname - %s, %s\nsubject - %s\nsalary - %f\nclassrooms - %v\n\n", c.Person.ID, c.Person.Name, c.Person.Surname, c.Subject, c.Salary, c.Classrooms)
			return
		}
	}
}

func (student) read(sData []student, id int) {
	for _, c := range sData {
		if c.Person.ID == id {
			fmt.Printf("id - %d\nname, surname - %s, %s\nclass - %v\n\n", c.Person.ID, c.Person.Name, c.Person.Surname, c.Class)
			return
		}
	}
}

func (director) read(dData director) {
	fmt.Printf("id - %d\nname, surname - %s, %s\nroom - %s\nphone - %d\n\n", dData.Person.ID, dData.Person.Name, dData.Person.Surname, dData.Phone, dData.Room)
}

func (teacher) delete(tData *[]teacher, id int) {
	for n, c := range *tData {
		if c.Person.ID == id {
			copy((*tData)[n:], (*tData)[n+1:])
			*tData = (*tData)[:len(*tData)-1]
			return
		}
	}
	fmt.Println("ERROR, ID not found")
}

func (student) delete(sData *[]student, id int) {
	for n, c := range *sData {
		if c.Person.ID == id {
			copy((*sData)[n:], (*sData)[n+1:])
			*sData = (*sData)[:len(*sData)-1]
			return
		}
	}
	fmt.Println("ERROR, ID not found")
}

func (director) delete(dData *director) {
	dData = new(director)
}

func (teacher) listIDs(tData []teacher, conn *net.TCPConn) {
	for _, c := range sData {
		conn.Write([]byte(strconv.Itoa(c.Person.ID)))
	}
	conn.Write([]byte{'\x00'})
}

func (student) listIDs(sData []student, conn *net.TCPConn) {
	for _, c := range sData {
		conn.Write([]byte(strconv.Itoa(c.Person.ID)))
	}
	conn.Write([]byte{'\x00'})
}

func (director) listIDs(dData director, conn *net.TCPConn) {
	conn.Write([]byte(strconv.Itoa(dData.Person.ID)))
}

func commitActions(conn *net.TCPConn) {
	var (
		data [2048]byte
		acD  actionData
	)

	for {
		_, err := conn.Read(data[0:])
		if err != nil {
			fmt.Println("Error reading data")
			return
		}

		json.Unmarshal(data[0:], &acD)
		switch acD.Action {
		case "create_Teacher":
			acD.Teacher.create(&tData)
		case "create_Student":
			acD.Student.create(&sData)
		case "create_Director":
			acD.Director.create(&dData)
		case "update_Teacher":
			acD.Teacher.listIDs(tData, conn)
			acD.Teacher.update(&tData)
		case "update_Student":
			acD.Student.listIDs(sData, conn)
			acD.Student.update(&sData)
		case "update_Director":
			acD.Director.listIDs(dData, conn)
			acD.Director.update(&dData)
		case "read_Teacher":
			acD.Teacher.listIDs(tData, conn)
			acD.Teacher.read(tData, acD.Teacher.Person.ID)
		case "read_Student":
			acD.Student.listIDs(sData, conn)
			acD.Student.read(sData, acD.Student.Person.ID)
		case "read_Director":
			acD.Director.listIDs(dData, conn)
			acD.Director.read(dData)
		case "delete_Teacher":
			acD.Teacher.delete(&tData, acD.Teacher.Person.ID)
		case "delete_Student":
			acD.Student.delete(&sData, acD.Student.Person.ID)
		case "delete_Director":
			acD.Director.delete(&dData)
		}
	}
}

func startServer() *net.TCPConn {
	port := ":8080"
	protocol := "tcp"

	adr, err := net.ResolveTCPAddr(protocol, port)
	if err != nil {
		fmt.Println("Wrong addr")
		return nil
	}

	fmt.Println("Server working on addr " + adr.String())

	listener, err := net.ListenTCP(protocol, adr)
	if err != nil {
		fmt.Println("Error creating server")
		return nil
	}

	conn, _ := listener.AcceptTCP()
	return conn
}

func main() {
	conn := startServer()
	commitActions(conn)
}
