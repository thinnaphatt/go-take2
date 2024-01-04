package main

import (
	"fmt"
	"net"
)

const (
	username = "std1"
	password = "p@ssw0rd"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := string(buffer[:n])
	fmt.Println("Received:", clientData)

	if clientData == fmt.Sprintf("%s:%s", username, password) {
		conn.Write([]byte("Hello"))
	} else {
		conn.Write([]byte("Invalid credentials"))
	}
}

func main() {
	fmt.Println("Server is starting...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}
