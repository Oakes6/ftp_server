package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	Users map[string]string
}

func ListenAndServe() {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8080})
	if err != nil {
		fmt.Printf("error listening to server: %v", err)
	}
	fmt.Println("Listening on port 8080...")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("error trying to accept connection: %v", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn *net.TCPConn) {
	// switch command {
	// case "ADD":
	// default:
	// 	fmt.Println("Command not found")
	// }
	fmt.Println("Received communication from client")
	input := bufio.NewScanner(conn)
	// _, err := conn.ReadFrom(conn)
	// if err != nil {
	// 	fmt.Printf("error reading from connection: %v", err)
	// }
	for input.Scan() {
		// grab verb and param
		input := input.Text()
		splitInput := strings.SplitN(input, " ", 2)
		verb, param := splitInput[0], splitInput[1]
		fmt.Printf("VERB: %s\tPARAM: %s\n", verb, param)
		ctx := context.Background()
		processCommand(ctx, verb, param)
	}

}

func processCommand(ctx context.Context, verb, param string) {
	switch verb {
	case "ADD":
		fmt.Println("Called the ADD verb")
	default:
		fmt.Println("Not a supported command")
	}
}

func main() {
	ListenAndServe()
}
