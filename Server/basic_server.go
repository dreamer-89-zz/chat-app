package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type message struct {
	Message string `json:"message"`
}

func main1() {
	//http.Handle("/socket", websocket.Handler(Socket))
}

func Socket(ws *websocket.Conn) {
	fmt.Println("Reached body of the foo method")
	for {
		var m message
		if err := websocket.ReadJSON(ws, &m); err != nil {
			fmt.Println("Error retrieving websocket message")
		}
		fmt.Println("Reaceived message => ", m.Message)
		if err := websocket.WriteJSON(ws, "Response from server"); err != nil {
			fmt.Println("Error writing back to client")
			break
		}
	}
}
