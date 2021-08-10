package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func connect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  921600, //1280*720
		WriteBufferSize: 0,
	}

	conn, err := upgrader.Upgrade(w, r, nil) //http соединение становится websocket'ом
	if err != nil {
		fmt.Fprintf(w, "Upgrading error: %v", err)
		return
	}

	for { //В этом цикле сервер получает картинку и декодит её в файл
		img := getImage(conn)
		saveImage(img)
	}
}

func getImage(conn *websocket.Conn) image.Image {
	_, mssg, err := conn.ReadMessage() //Получает картинку в виде []byte
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(mssg)
	img, err1 := png.Decode(reader)
	if err1 != nil {
		log.Fatal(err1)
	}
	return img
}

func saveImage(img image.Image) {
	rand.Seed(time.Now().UnixNano())

	var b strings.Builder
	b.Grow(30)
	b.WriteString("userImages/") //Сохраняется в папку userImages
	b.WriteString("userImage")
	b.WriteString(strconv.Itoa(rand.Intn(1000000) + 100000))
	b.WriteString(".png")

	file, err := os.Create(b.String())
	defer file.Close()
	if err != nil { //Если папки "userImages" не существует
		makeDirectory()                   //Создаю папку
		file, err = os.Create(b.String()) //Создаю там картинку
	}

	png.Encode(file, img) //Переношу данные о картинке в файл
	fmt.Println("Image saved")
}

func makeDirectory() {
	err := os.Mkdir("UserImages", 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", connect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
