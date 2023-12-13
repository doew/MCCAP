package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Received: ", message)
		if strings.TrimSpace(message) == "Merry" {
			conn.Write([]byte("Christmas!\n"))
		} else if strings.TrimSpace(message) == "Chicken" {
            // ASCII Art from https://textart.sh/topic/chicken
			conn.Write([]byte("\n              ████████        \n            ██      ▒▒██      \n          ██    ▒▒▒▒▒▒▓▓██    \n        ██  ░░▒▒▒▒▒▒▒▒▒▒▓▓██  \n        ██░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒██  \n      ██  ░░▒▒▒▒▒▒▒▒░░▒▒▒▒▓▓██\n      ██░░░░▒▒▒▒▒▒▒▒▒▒░░▒▒▓▓██\n      ██░░▒▒▒▒▒▒▒▒▒▒░░░░▒▒▓▓██\n        ██▒▒▒▒▒▒▒▒▒▒░░▒▒▒▒██  \n      ░░██░░▒▒▒▒▒▒▒▒▒▒▒▒▓▓██  \n        ░░██▒▒▒▒▒▒▒▒▒▒▓▓██    \n            ██▒▒▒▒▒▒▓▓██      \n            ██▒▒▒▒▒▒▓▓██      \n              ████████        \n              ██░░░░██        \n              ██░░░░██        \n              ██  ░░██        \n            ██    ░░░░██      \n          ██    ████░░░░██    \n            ████    ████      \nC")) // ここにチキンのASCIIアートを入れる
		}
	}
}
