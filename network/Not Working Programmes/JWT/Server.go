package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	userSession session
	fin         *os.File
	userData    user
)

type (
	user struct {
		Credentials credentials `json:"credentials"`
		Name        string
	}
	credentials struct {
		Login    string `json:"Login"`
		Password string `json:"Password"`
	}
	session struct {
		Id, UserId int
		IP         net.IP
		Open_Date  time.Time
		Duration   time.Duration
		Exp_Date   time.Time
	}
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")
	var (
		userData user
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

}

func main() {
	http.HandleFunc("/login", login) //Проверять, есть ли тут активная сессия
	//http.HandleFunc("/welcome", welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
