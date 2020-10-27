package main

import (
	"log"
	"net"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Println("dail ok")
}
