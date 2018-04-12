package main

import (
	"bufio"
	"fmt"
	"net"
	_ "net"
	"strings"
)

func main() {
	fmt.Println("This is direct beginning of go program")
	listners, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		listners.Close()
		fmt.Println("Closing listning to new connections")
	}() // implicit call

	for {
		conn, err := listners.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
		}
		go handleNewConnection(conn)
	}
}

func handleNewConnection(connection net.Conn) {
	fmt.Println("Handling new connection request")
	defer connection.Close()
	bufRead := bufio.NewReader(connection)
	for {
		data, err := bufRead.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
		fmt.Fprintf(connection, strings.ToUpper(data))
		//_, err = bufio.NewWriter(connection).WriteString(strings.ToUpper(data))
		if err != nil {
			fmt.Println("Error writing back to the client")
		}
	}
}
