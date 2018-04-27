package main

import (
	"flag"
	"fmt"
	"log"
	_ "net"
	"net/http"

	"golang.org/x/net/websocket"
)

type Message struct {
	Message string `json:"message"`
}

var url = flag.String("addr", "localhost:8080", "http service address")

func socket(ws *websocket.Conn) {
	for {
		var m Message

		if err := websocket.JSON.Receive(ws, &m); err != nil {
			fmt.Println(err)
			log.Println(err)
			break
		}

		log.Println("Received message : ", m)
		m2 := Message{Message: "Thanks for checking over"}
		if err := websocket.JSON.Send(ws, &m2); err != nil {
			fmt.Println(err)
			log.Println(err)
			break
		}
	}
}
func main() {
	// TODO ... Server should handle new connections and persist the same via websocket connections
	// 1. Anything typed here should be shown to clients attached to this server.
	// Future ... Multi-chat rooms, giving client option to connect with other clients. Peer to peer chat & group chats
	// 2. Server will just act as a hub connecting different clients.
	http.Handle("/socket", websocket.Handler(socket))
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Client requested foo!")
	})
	http.ListenAndServe(":9090", nil)
}
