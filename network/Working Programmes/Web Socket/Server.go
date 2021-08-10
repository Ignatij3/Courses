package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	chatroom []*websocket.Conn
	waitBoth sync.WaitGroup
)

func addConnection(conn *websocket.Conn) {
	chatroom = append(chatroom, conn)
	fmt.Println("Connection established")
}

func chat(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  255,
		WriteBufferSize: 255,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Upgrading error: %v", err)
		return
	}

	addConnection(conn)
	waitBoth.Done()
	waitBoth.Wait()
	for {
		_, text, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Reading message error: %v", err)
			return
		}

		switch conn.UnderlyingConn().RemoteAddr().String() {
		case chatroom[0].UnderlyingConn().RemoteAddr().String():
			err := chatroom[1].WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				fmt.Printf("Writing message error: %v", err)
				return
			}
		case chatroom[1].UnderlyingConn().RemoteAddr().String():
			err := chatroom[0].WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				fmt.Printf("Writing message error: %v", err)
				return
			}
		}
	}
}

func main() {
	waitBoth.Add(2)
	http.HandleFunc("/chat", chat)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
