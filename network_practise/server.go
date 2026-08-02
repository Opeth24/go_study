package networkpractise

import (
	"log"
	"net"
	"time"
)

func Server() {

	listener, err := net.Listen("tcp", "127.0.0.1:8089")

	if err != nil {

		log.Fatal("Ошибка запуска сервера: ", err)

	}
	defer listener.Close()

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