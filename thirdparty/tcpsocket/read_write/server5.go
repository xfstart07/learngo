package main

import (
	"log"
	"net"
	"time"
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
	time.Sleep(time.Second * 10)
	for {
		// read from the connection
		time.Sleep(5 * time.Second)
		var buf = make([]byte, 60000)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes,  error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
		}

		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
