package main

import (
	"fmt"
	"net"
)

type (
	school interface {
		create()
		update()
		read()
		delete()
	}
	person struct {
		ID      int    `json:"ID"`
		Name    string `json:"Name"`
		Surname string `json:"Surname"`
	}
	teacher struct {
		Person     person  `json:"Person"` //ID starts with 2
		Subject    string  `json:"Subject"`
		Salary     float64 `json:"Salary"`
		Classrooms []int   `json:"Classrooms"`
	}
	student struct {
		Person person `json:"Person"` //ID starts with 3
		Class  int    `json:"Class"`
	}
	director struct {
		Person person `json:"Person"` //ID starts with 1
		Room   int    `json:"Room"`
		Phone  string `json:"Phone"`
	}
)

func (t *teacher) create() {
	var classroom int = 1
	fmt.Print("Enter Name:")
	fmt.Scan(&t.Person.Name)
	fmt.Print("Enter Surname:")
	fmt.Scan(&t.Person.Surname)
	fmt.Print("Enter Subject:")
	fmt.Scan(&t.Subject)
	fmt.Print("Enter Salary:")
	fmt.Scan(&t.Salary)
	for classroom != 0 {
		fmt.Print("Enter Classrooms (0 to exit):")
		fmt.Scan(&classroom)
		t.Classrooms = append(t.Classrooms, classroom)
	}
}

func (s *student) create() {
	fmt.Print("Enter Name:")
	fmt.Scan(&s.Person.Name)
	fmt.Print("Enter Surname:")
	fmt.Scan(&s.Person.Surname)
	fmt.Print("Enter Class:")
	fmt.Scan(&s.Class)
}

func (d *director) create() {
	fmt.Print("Enter Name:")
	fmt.Scan(&d.Person.Name)
	fmt.Print("Enter Surname:")
	fmt.Scan(&d.Person.Surname)
	fmt.Print("Enter Room:")
	fmt.Scan(&d.Room)
	fmt.Print("Enter Phone:")
	fmt.Scan(&d.Phone)
}

func (t teacher) update(conn *net.TCPConn) teacher {
	var (
		classroom int = 1
		id        [3]byte
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan((&t).Person.ID)
	fmt.Print("Enter Name:")
	fmt.Scan((&t).Person.Name)
	fmt.Print("Enter Surname:")
	fmt.Scan((&t).Person.Surname)
	fmt.Print("Enter Subject:")
	fmt.Scan((&t).Subject)
	fmt.Print("Enter Salary:")
	fmt.Scan((&t).Salary)
	for classroom != 0 {
		fmt.Print("Enter Classrooms (0 to exit):")
		fmt.Scan(&classroom)
		(&t).Classrooms = append((&t).Classrooms, classroom)
	}
	return t
}

func (s student) update(conn *net.TCPConn) student {
	var id [3]byte
	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan((&s).Person.ID)
	fmt.Print("Enter Name:")
	fmt.Scan((&s).Person.Name)
	fmt.Print("Enter Surname:")
	fmt.Scan((&s).Person.Surname)
	fmt.Print("Enter Class:")
	fmt.Scan((&s).Class)
	return s
}

func (d director) update(conn *net.TCPConn) director {
	var id [3]byte
	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan((&d).Person.ID)
	fmt.Print("Enter Name:")
	fmt.Scan((&d).Person.Name)
	fmt.Print("Enter Surname:")
	fmt.Scan((&d).Person.Surname)
	fmt.Print("Enter Room:")
	fmt.Scan((&d).Room)
	fmt.Print("Enter Phone:")
	fmt.Scan((&d).Phone)
	return d
}

func (t teacher) read(conn *net.TCPConn) int {
	var (
		id  [3]byte
		num int
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan(&num)
	return num
}

func (s student) read(conn *net.TCPConn) int {
	var (
		id  [3]byte
		num int
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan(&num)
	return num
}

func (d director) read(conn *net.TCPConn) int {
	var (
		id  [3]byte
		num int
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan(&num)
	return num
}

func (t teacher) delete(conn *net.TCPConn) int {
	var (
		id  [3]byte
		num int
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan(&num)
	return num
}

func (s student) delete(conn *net.TCPConn) int {
	var (
		id  [3]byte
		num int
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan(&num)
	return num
}

func (d director) delete(conn *net.TCPConn) int {
	var (
		id  [3]byte
		num int
	)

	for n := 1; id[0] != '\x00'; n++ {
		conn.Read(id[:])
		fmt.Printf("%d) %s\n", n, id)
	}

	fmt.Print("Enter ID:")
	fmt.Scan(&num)
	return num
}

func chooseDatabase() rune {
	var data int

	fmt.Println("1) Teachers")
	fmt.Println("2) Students")
	fmt.Println("3) Director")
	fmt.Print("What data you want to change? (choose number)")
	fmt.Scan(&data)

	for data < 1 || data > 3 {
		fmt.Print("Incorrect input, try again")
		fmt.Scan(&data)
	}

	switch data {
	case 1:
		return 'T'
	case 2:
		return 'S'
	case 3:
		return 'D'
	}
	return ' '
}

func chooseAction() rune {
	var action int

	fmt.Println("0) End connection")
	fmt.Println("1) Create")
	fmt.Println("2) Update")
	fmt.Println("3) Read")
	fmt.Println("4) Delete")
	fmt.Print("What you want to do? (choose number)")
	fmt.Scan(&action)

	for action < 0 || action > 4 {
		fmt.Print("Incorrect input, try again")
		fmt.Scan(&action)
	}

	switch action {
	case 1:
		return 'C'
	case 2:
		return 'U'
	case 3:
		return 'R'
	case 4:
		return 'D'
	}
	return 'E'
}

func sendActions() {
	var (
		data, action rune
		t            teacher
		s            student
		d            director
	)
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	for {
		action = chooseAction()
		data = chooseDatabase()

		switch action {
		case 'C':
			switch data {
			case 'T':
				t.create()
				fmt.Fprintf(conn, "%v", t)
			case 'S':
				s.create()
				fmt.Fprintf(conn, "%v", s)
			case 'D':
				d.create()
				fmt.Fprintf(conn, "%v", d)
			}
		case 'U':
			switch data {
			case 'T':
				t = t.update(conn)
				fmt.Fprintf(conn, "%v", t)
			case 'S':
				s = s.update(conn)
				fmt.Fprintf(conn, "%v", s)
			case 'D':
				d = d.update(conn)
				fmt.Fprintf(conn, "%v", d)
			}
		case 'R':
			switch data {
			case 'T':
				t = t.read(conn)
				fmt.Fprintf(conn, "%v", t)
			case 'S':
				s = s.read(conn)
				fmt.Fprintf(conn, "%v", s)
			case 'D':
				d = d.read(conn)
				fmt.Fprintf(conn, "%v", d)
			}
		case 'D':
			switch data {
			case 'T':
				t = t.delete(conn)
				fmt.Fprintf(conn, "%v", t)
			case 'S':
				s = s.delete(conn)
				fmt.Fprintf(conn, "%v", s)
			case 'D':
				d = d.delete(conn)
				fmt.Fprintf(conn, "%v", d)
			}
		case 'E':
			break
		}
	}
}

func main() {
	sendActions()
}
