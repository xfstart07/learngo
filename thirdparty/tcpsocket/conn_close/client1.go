package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	conn.Close()
	log.Println("close ok")

	var buf = make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("read err", err)
	} else {
		log.Printf("read %d bytes, content %s \n", n, string(buf))
	}

	n, err = conn.Write(buf)
	if err != nil {
		log.Println("write err", err)
	} else {
		log.Printf("write %d bytes, content %s \n", n, string(buf))
	}

	time.Sleep(1000 * time.Second)
}
