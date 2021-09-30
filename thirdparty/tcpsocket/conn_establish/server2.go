package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln("error listen:", err)
	}
	defer listen.Close()
	log.Println("listen ok")

	var i int
	for {
		time.Sleep(10 * time.Second)
		if _, err := listen.Accept(); err != nil {
			log.Println(err)
			break
		}
		i++
		log.Printf("%d: accept a new connection", i)
	}
}
