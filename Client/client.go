package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to local host socket
	connect, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text to send!")
	//writer := bufio.NewWriter(connect)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input from console")
		}
		//_, err = writer.WriteString(text)
		// Send to server socket
		fmt.Fprintf(connect, text)
		if err != nil {
			fmt.Println(err)
		}
		serverMessage, _ := bufio.NewReader(connect).ReadString('\n')
		fmt.Println(serverMessage)
	}
}
