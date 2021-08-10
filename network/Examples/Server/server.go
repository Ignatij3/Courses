package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func (s Student) Info() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

type StudentList []Student

func (s StudentList) AddStudent(NewStudent []byte) error {
	var NStudent = Student{}
	err := json.Unmarshal(NewStudent, &NStudent)
	if err != nil {
		return err
	}

	s = append(s, NStudent)
	return nil
}

var Students StudentList

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")

	switch r.Method {
	case "GET":
		parsedUrl := strings.Split(r.URL.Path, "/")
		if parsedUrl[2] == "" {
			str := fmt.Sprintf(`{"Length": %d}`, len(Students))
			io.WriteString(w, str)
		} else {
			ind, _ := strconv.Atoi(parsedUrl[2])
			if len(Students) <= ind {
				w.WriteHeader(404)
				return
			}
			str, err := Students[ind].Info()
			if err != nil {
				w.WriteHeader(500)
				return
			}
			io.WriteString(w, str)
		}
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		err := Students.AddStudent(body)
		if err != nil {
			io.WriteString(w, "Not created due to "+err.Error())
			return
		}
		io.WriteString(w, "Created")
	}
}

func main() {
	Students = StudentList{{Name: "Henry", Age: 6}, {Name: "John", Age: 3}}
	http.HandleFunc("/Student/", GetStudents)
	http.ListenAndServe(":8080", nil)
}
