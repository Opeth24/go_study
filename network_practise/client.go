package networkpractise

import (
	"fmt"
	"net"
	"strings"
)

func StartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println("Net error:", err)
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