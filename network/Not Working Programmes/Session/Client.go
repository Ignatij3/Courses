package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var adress string = "http://localhost:8080/"

type credentials struct {
	Login    string
	Password string
}

func sendCredentials() {
	var (
		cred credentials
		cl   http.Client
		data string
	)

	fmt.Print("Enter login: ")
	fmt.Scan(&data)
	cred.Login = data

	fmt.Print("Enter password: ")
	fmt.Scan(&data)

	hsh := sha512.New()
	hsh.Write([]byte(data))
	cred.Password = fmt.Sprintf("%x", hsh.Sum(nil))

	marshalledBody, errjs := json.Marshal(cred)
	if errjs != nil {
		fmt.Println(errjs)
		return
	}

	//fmt.Printf("%v\n%x\n%s\n", marshalledBody, marshalledBody, marshalledBody)

	req, _ := http.NewRequest("POST", adress, bytes.NewBuffer(marshalledBody))
	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 && resp.StatusCode != 308 {
		fmt.Println(resp.StatusCode)
		fmt.Println(string(response))
		return
	}

	if string(response) != "http://localhost:8080/welcome" {
		fmt.Println(string(response))
	} else {
		adress = string(response)
	}
}

func sendData() {
	var (
		cl   http.Client
		name string
	)

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	req, _ := http.NewRequest("POST", adress, bytes.NewBuffer([]byte(name)))
	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return
	}

	fmt.Println(string(response))
}

func main() {
	for {
		if adress == "http://localhost:8080/" {
			sendCredentials()
		} else {
			sendData()
		}
	}
}
