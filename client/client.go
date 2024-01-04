package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	data := fmt.Sprintf("%s:%s", username, password)
	fmt.Fprintf(conn, data)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	serverResponse := string(buffer[:n])
	fmt.Println("Server response:", serverResponse)
}
