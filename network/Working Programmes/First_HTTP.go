package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type student struct {
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	GNumber int    `json:"GNumber"`
	GLetter string `json:"GLetter"`
}

func createData() [100]student {
	var s [100]student

	for n := range s {
		s[n].Name = "Name" + strconv.Itoa(n+1)
		s[n].Surname = "Surname" + strconv.Itoa(n+1)
		s[n].GNumber = rand.Intn(12) + 1
		switch rand.Intn(5) {
		case 0:
			s[n].GLetter = "a"
		case 1:
			s[n].GLetter = "b"
		case 2:
			s[n].GLetter = "c"
		case 3:
			s[n].GLetter = "d"
		case 4:
			s[n].GLetter = "e"
		}
	}
	return s
}

func main() {
	students := createData()

	h1 := func(w http.ResponseWriter, r *http.Request) {
		strNum := strings.SplitN(r.URL.Path, "/", 2)
		n, err := strconv.Atoi(strNum[1])
		if err != nil || n <= 0 || n > 100 {
			http.Error(w, "404 NOT FOUND", 404)
			return
		}
		fmt.Fprintf(w, "%s %s (%d%s)\n", students[n-1].Name, students[n-1].Surname, students[n-1].GNumber, string(students[n-1].GLetter))
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
