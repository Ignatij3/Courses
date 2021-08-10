package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func listStudents(r *http.Response, index string) {
	var (
		st       []student
		singleSt student
		all      bool
		errJs    error
	)

	response, _ := ioutil.ReadAll(r.Body)
	if index == "ALL" || index == "all" {
		all = true
		errJs = json.Unmarshal(response, &st)
	} else {
		errJs = json.Unmarshal(response, &singleSt)
	}
	if errJs != nil {
		fmt.Println(string(response))
		return
	}

	if all {
		for n, s := range st {
			fmt.Printf("%d) (%d) %s %s (%d grade)\n", n+1, s.Person.ID, s.Person.Name, s.Person.Surname, s.Class)
		}
	} else {
		fmt.Printf("(%d) %s %s (%d grade)\n", singleSt.Person.ID, singleSt.Person.Name, singleSt.Person.Surname, singleSt.Class)
	}
}

func makeGetRequest(cl http.Client) {
	var (
		request *http.Request
		studNum string
	)

	fmt.Print("Enter student number (\"all\" to read all students): ")
	fmt.Scan(&studNum)

	request, _ = http.NewRequest("GET", "http://127.0.0.1:8080/Students/"+studNum, nil)
	r, err := cl.Do(request)
	if err != nil {
		fmt.Printf("Request error - %v\n", err)
	}
	listStudents(r, studNum)
}

func (st *student) getData(req string) {
	if req == "+id" {
		fmt.Print("Enter id (starts from 301): ")
		fmt.Scan(&st.Person.ID)
	}
	fmt.Print("Enter name: ")
	fmt.Scan(&st.Person.Name)
	fmt.Print("Enter surname: ")
	fmt.Scan(&st.Person.Surname)
	fmt.Print("Enter class: ")
	fmt.Scan(&st.Class)
}

func makePostRequest(cl http.Client) {
	var (
		request *http.Request
		st      student
	)

	st.getData("")
	stBody := student{
		Person: st.Person,
		Class:  st.Class,
	}

	marshalledBody, _ := json.Marshal(stBody)
	request, _ = http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer(marshalledBody))
	r, err := cl.Do(request)
	if err != nil {
		fmt.Printf("Request error - %v\n", err)
		return
	}
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}

func makePutRequest(cl http.Client) {
	var (
		request *http.Request
		st      student
	)

	st.getData("+id")
	stBody := student{
		Person: st.Person,
		Class:  st.Class,
	}
	marshalledBody, _ := json.Marshal(stBody)

	request, _ = http.NewRequest("PUT", "http://127.0.0.1:8080", bytes.NewBuffer(marshalledBody))
	r, err := cl.Do(request)
	if err != nil {
		fmt.Printf("Request error - %v\n", err)
		return
	}
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}

func makePatchRequest(cl http.Client) {
	var (
		request *http.Request
		st      student
	)

	st.getData("+id")
	stBody := student{
		Person: st.Person,
		Class:  st.Class,
	}
	marshalledBody, _ := json.Marshal(stBody)

	request, _ = http.NewRequest("PATCH", "http://127.0.0.1:8080", bytes.NewBuffer(marshalledBody))
	r, err := cl.Do(request)
	if err != nil {
		fmt.Printf("Request error - %v\n", err)
		return
	}
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}

func makeDeleteRequest(cl http.Client) {
	var (
		request *http.Request
		id      string
	)

	fmt.Print("Enter id you want to delete: ")
	fmt.Scan(&id)

	request, _ = http.NewRequest("DELETE", "http://127.0.0.1:8080", bytes.NewBuffer([]byte(id)))
	r, err := cl.Do(request)
	if err != nil {
		fmt.Printf("Request error - %v\n", err)
		return
	}
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}

func main() {
	var (
		cl          http.Client
		requestType string
	)

	for {
		fmt.Print("Enter request type (\"none\" to exit): ")
		fmt.Scan(&requestType)
		switch requestType {
		case "GET", "get":
			makeGetRequest(cl)
		case "POST", "post":
			makePostRequest(cl)
		case "PUT", "put":
			makePutRequest(cl)
		case "PATCH", "patch":
			makePatchRequest(cl)
		case "DELETE", "delete":
			makeDeleteRequest(cl)
		case "NONE", "none":
			break
		default:
			fmt.Println("Incorrect input")
		}
	}
}
