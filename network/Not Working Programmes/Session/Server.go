package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"time"
)

//var offset []byte
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

func init() {
	var err error
	fin, err = os.Open("Database.json")
	if err != nil {
		fmt.Println(err)
	}
}

func readDB(clCred credentials) (bool, error) {
	var (
		loginFound bool
		err        error
	)

	dec := json.NewDecoder(fin)
	dec.Token()
	for {
		errDec := dec.Decode(&userData)
		if errDec == io.EOF {
			break
		}

		if err != nil {
			return false, err
		}

		if userData.Credentials.Login == clCred.Login {
			loginFound = true
			break
		}
	}
	dec.Token()
	return loginFound, nil
}

func startSession(ip net.IP) {
	userSession.Id = 001
	userSession.UserId = 001
	userSession.IP = ip
	userSession.Open_Date = time.Now()
	dur, _ := time.ParseDuration("1m")
	userSession.Duration = dur
	userSession.Exp_Date = userSession.Open_Date.Add(dur)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")
	var clCred credentials
	fmt.Println("1")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "body")
		w.Write([]byte("Error reading body"))
		w.WriteHeader(500)
		return
	}
	fmt.Println("2")
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		w.Write([]byte("Not IP adress"))
		return
	}
	userIP := net.ParseIP(ip)
	fmt.Println("3")
	jserr := json.Unmarshal(body, &clCred)
	if jserr != nil {
		w.Write([]byte("Error unmarshaling"))
		w.WriteHeader(400)
		return
	}
	fmt.Println("4")

	loginFound, err := readDB(clCred)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("User - %v\n%x\n%s\nDatabase - %v\n%x\n%s\n", clCred.Password, clCred.Password, clCred.Password, userData.Credentials.Password, userData.Credentials.Password, userData.Credentials.Password)
	fmt.Println("5")
	if loginFound {
		fmt.Println("6")
		if userSession.Exp_Date.After(time.Now()) {
			startSession(userIP)
			w.Write([]byte("Session reopened"))
		}
		fmt.Println("7")
		if reflect.DeepEqual(clCred, userData.Credentials) {
			startSession(userIP)
			w.Write([]byte("Session successfully opened"))            //Если написать раньше редиректа - то выдаст ошибку на сервере http: superfluous response.WriteHeader call from main.login (Server.go:134)
			http.Redirect(w, r, "http://localhost:8080/welcome", 308) //500 sends automatically doesn't redirect
		} else {
			w.Write([]byte("Incorrect password, try again"))
		}
		fmt.Println("8")
	} else {
		fmt.Println("9")
		w.Write([]byte("User not found"))
	}
	fmt.Println("10")
}

func welcome(w http.ResponseWriter, r *http.Request) {
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

	fin, err := os.Open("Database.json") //Replace names in the database
	if err != nil {
		w.WriteHeader(500)
		return
	}

	dec := json.NewDecoder(fin)
	for {
		errDec := dec.Decode(&userData)
		if errDec == io.EOF {
			break
		}
		if errDec != nil {
			w.WriteHeader(500)
			return
		}
	}

	if string(body) == userData.Name {
		fmt.Fprintf(w, "Your name is %s", string(body))
	} else {
		fmt.Fprintf(w, "Your name is not %s", string(body))
	}
}

func main() {
	http.HandleFunc("/", login) //Проверять, есть ли тут активная сессия
	http.HandleFunc("/welcome", welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
