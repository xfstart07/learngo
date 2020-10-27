package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error: ", err)
		return
	}
	defer conn.Close()

	data := make([]byte, 65536)
	var total int
	for {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Microsecond))
		n, err := conn.Write(data)
		if err != nil {
			total += n
			log.Printf("write %d bytes, err %s\n", n, err)
			break
		}
		total += n
		log.Printf("write %d bytes in total\n", total)
		// time.Sleep(1000 * time.Second)
	}
}
