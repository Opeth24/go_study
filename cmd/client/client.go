package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Net error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error while listen accept: ", err)
	}
	defer conn.Close()


	message := make([]byte, 1024)
	for i := 0; i < 3; i++ {

		n, err := conn.Read(message)
		if err != nil {
			fmt.Println(err)
		}
		decodeStr := string(message[:n])
		fmt.Println(strings.ToUpper(decodeStr))
	}
}