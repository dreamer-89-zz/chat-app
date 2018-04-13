package main

import (
	"bufio"
	"fmt"
	"net"
	_ "net"
	"strings"
	"flag"

	"github.com/gorilla/websocket"
)

var url = flag.String("addr", "localhost:8080", "http service address")
func main() {
	// TODO ... Server should handle new connections and persist the same via websocket connections
	// Anything typed here should be shown to clients attached to this server. 
	// Future ... Multi-chat rooms, giving client option to connect with other clients. Peer to peer chat & group chats
	// Server will just act as a hub connecting different clients.
}
