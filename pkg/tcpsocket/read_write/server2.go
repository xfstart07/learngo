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
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read err", err)
			return
		}
		log.Println("read str = ", string(buf[:n]))
	}
}
