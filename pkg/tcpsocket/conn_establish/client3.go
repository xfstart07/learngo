package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.DialTimeout("tcp", "104.236.176.96:80", 2*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Println("dial ok")
}
