package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type (
	person struct {
		ID      int    `json:"ID"`
		Name    string `json:"Name"`
		Surname string `json:"Surname"`
	}
	student struct {
		Person person `json:"Person"`
		Class  int    `json:"Class"`
	}
)

var students []student

func readData() []student {
	var sData []student
	studentData, _ := os.Open("../../Files/StudentData.json")
	studentSlice, _ := ioutil.ReadAll(studentData)
	json.Unmarshal(studentSlice, &sData)
	return sData
}

func (st student) add(students *[]student) int {
	newID := (*students)[len(*students)-1].Person.ID
	*students = append(*students, st)
	(*students)[len(*students)-1].Person.ID = newID + 1
	return newID + 1
}

func (st student) replace(students *[]student) {
	for n, s := range *students {
		if st.Person.ID == s.Person.ID {
			(*students)[n] = st
			return
		}
	}
}

func delete(students *[]student, id int) {
	var changeID bool
	for n := range *students {
		if n != len(*students) {
			if changeID {
				(*students)[n].Person.ID--
			} else if (*students)[n].Person.ID == id {
				copy((*students)[n:], (*students)[n+1:])
				*students = (*students)[:len(*students)-1]
				if n != len(*students) {
					(*students)[n].Person.ID--
				}
				changeID = true
			}
		}
	}
}

func h1(w http.ResponseWriter, r *http.Request) {
	var st student

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")

	switch r.Method {
	case "POST":
		err := json.Unmarshal(body, &st)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		newID := st.add(&students)
		fmt.Fprintf(w, "Successfully added new student (ID - %d)", newID)
	case "PUT":
		err := json.Unmarshal(body, &st)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if st.Person.ID > students[len(students)-1].Person.ID ||
			st.Person.ID < students[0].Person.ID {
			w.WriteHeader(404)
			return
		}
		st.replace(&students)
		fmt.Fprintf(w, "Successfully replaced student (ID - %d)", st.Person.ID)
	case "PATCH":
		err := json.Unmarshal(body, &st)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if st.Person.ID > students[len(students)-1].Person.ID ||
			st.Person.ID < students[0].Person.ID {
			w.WriteHeader(404)
			return
		}
		st.replace(&students)
		fmt.Fprintf(w, "Successfully replaced student (ID - %d)", st.Person.ID)
	case "DELETE":
		delSt, err := strconv.Atoi(string(body))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if delSt > students[len(students)-1].Person.ID ||
			delSt < students[0].Person.ID {
			w.WriteHeader(404)
			return
		}
		delete(&students, delSt)

		fmt.Fprint(w, "Successfully deleted student")
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	var (
		js    []byte
		errJs error
	)

	parsedUrl := strings.Split(r.URL.Path, "/")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")
	w.Header().Set("Content-Type", "application/json")
	if parsedUrl[2] == "all" || parsedUrl[2] == "ALL" {
		js, errJs = json.Marshal(students)
	} else {
		index, err := strconv.Atoi(parsedUrl[2])
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if index <= 0 || index > len(students) {
			w.WriteHeader(404)
			return
		}
		js, errJs = json.Marshal(students[index-1])
	}

	if errJs != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(js)
}

func main() {
	students = readData()
	http.HandleFunc("/", h1)
	http.HandleFunc("/Students/", get)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
