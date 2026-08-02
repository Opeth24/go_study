package main



import (
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8081")

	if err != nil {

		log.Println(err)

	}

	conn, err := listener.Accept()

	if err != nil {

		log.Println(err)

	}

	defer conn.Close()

	conn.Write([]byte("message"))

	time.Sleep(10)

	conn.Write([]byte("MesSaGe"))

	time.Sleep(10)

	conn.Write([]byte("MESSAGE"))

}