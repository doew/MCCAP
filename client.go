package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Merry")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Received: ", message)

	if message == "Christmas!\n" {
		fmt.Fprintln(conn, "Chicken")
		message, _ := bufio.NewReader(conn).ReadString('C')
		fmt.Print("Received: ", message)
	}
}
