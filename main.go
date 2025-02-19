package main

import (
	"fmt"
	"net"
)

func main() {
	var ip string

	fmt.Print("give me a ip: ")
	fmt.Scan(&ip)

	conn, err := net.Dial("tcp", ip+":80")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message := "Hello, Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
	fmt.Println("Sent message to server:", message)
}
