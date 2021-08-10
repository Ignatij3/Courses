package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func receiveMessages(conn *websocket.Conn) {
	for {
		_, text, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Connection reading error: %v", err)
			return
		}
		fmt.Println(string(text))
	}
}

func startCommunication() {
	var message string

	addr := flag.String("addr", "localhost:8080", "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/chat"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	defer conn.Close()

	if err != nil {
		fmt.Printf("Dialing error: %v", err)
		return
	}
	fmt.Println("Connected")

	go receiveMessages(conn)
	for {
		fmt.Scan(&message)
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Printf("Writing message error: %v", err)
			return
		}
	}
}

func main() {
	startCommunication()
}
