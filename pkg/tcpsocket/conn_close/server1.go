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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes,  error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
		}

		n, err = c.Write(buf)
		if err != nil {
			log.Println("write err", err)
		} else {
			log.Println("write succ", string(buf))
		}
	}
}
