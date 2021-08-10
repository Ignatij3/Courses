package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendCredentials() {
	var (
		cl http.Client
	)

	req, _ := http.NewRequest("POST", "http://localhost:8080/login", nil)
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
		fmt.Println(string(response))
		return
	}

	fmt.Println(string(response))
}

func sendData() {
	var (
		cl http.Client
	)

	req, _ := http.NewRequest("POST", "http://localhost:8080/login", nil)
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
	//Choose token, or password
	sendCredentials()
	sendData()
}
