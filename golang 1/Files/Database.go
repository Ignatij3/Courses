package main
import (
		"os"
		"fmt"
		"strings"
		"strconv"
		"github.com/nsf/termbox-go"
		)
		
func exists (name string) bool {
	f, err:=os.Open("students.dat")
	if err==nil {
		f.Close()
		return true
	} else {
		return !os.IsNotExist(err)
	}
}

func main() {
	type (
		tGrade struct {
			GLetter string
			GNum int
		}
		tFullName struct {
			FName string
			LName string
		}
		tStudent struct {
			grade tGrade
			name tFullName
		}
	)
	var students [775]tStudent
	fin, _ := os.Open ("students.dat")
	Scanner := bufio.NewScanner(fin)
	i := 0
	for Scanner.Scan() {
		Info := strings.Split(Scanner.Text(), ' ')
		students[i].name.LName = Info[0]
		students[i].name.FName = Info[1]
		students[i].grade.GNum = strconv.Atoi(Info[2])
		students[i].grade.GLetter = Info[3]
		i++
		if i = len(students) {break}
	}
	fin.Close()
	for
	
	
	
	
	
}
