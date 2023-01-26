package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
)

type conn struct {
	conn   net.Conn
	path   string
	binary bool
}

func NewConnection(newConn net.Conn) *conn {
	return &conn{conn: newConn, path: "/", binary: true}
}

func (c *conn) printPath() {
	fmt.Fprint(c.conn, c.path)
}

func (c *conn) closeConn() {
	fmt.Fprint(c.conn, "221 Bye.\n")
	err := c.conn.Close()
	if err != nil {
		fmt.Printf("internal server error: %s", err)
	}
}

func (c *conn) handleConn() {

	input := bufio.NewScanner(c.conn)
	for input.Scan() {
		// grab verb and param
		input := input.Text()
		splitInput := strings.SplitN(input, " ", 2)
		var param string
		verb := splitInput[0]
		if len(splitInput) > 1 {
			param = splitInput[1]
		}
		fmt.Printf("VERB: %s\tPARAM: %s\n", verb, param)
		ctx := context.Background()
		c.processCommand(ctx, verb, param)
	}

}

func (c *conn) processCommand(ctx context.Context, verb, param string) {
	switch verb {
	case "USER":
		fmt.Fprintf(c.conn, "230 login successful.\n")
	case "PWD":
		c.printPath()
	case "QUIT":
		c.closeConn()
	default:
		fmt.Println("Not a supported command")
	}
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
		go NewConnection(conn).handleConn()
	}
}

func main() {
	ListenAndServe()
}
