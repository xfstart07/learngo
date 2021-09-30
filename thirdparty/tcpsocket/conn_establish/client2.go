package main

import (
	"log"
	"net"
)

func main() {
	var sl []net.Conn
	for i := 1; i < 1000; i++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}
}

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}
