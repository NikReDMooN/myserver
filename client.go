package main

import (
	"fmt"
	"net"
)

func doclient() {
	fmt.Println("client is connecting")
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("client connected")
	defer conn.Close()
	for {
		var message string
		fmt.Println("clinet is waiting for message")
		fmt.Scanln(&message)
		n, err := conn.Write([]byte(message))
		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}

		if message == "break" {
			fmt.Println("client closed")
			break
		}

	}
}
