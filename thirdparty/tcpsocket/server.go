package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	log.Println("Listen 8888...")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// read
		// write
	}
}
