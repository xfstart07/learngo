package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client2.go YOUR_CONTENT")
		return
	}

	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	time.Sleep(2 * time.Second)
	data := os.Args[1]
	conn.Write([]byte(data))
}
