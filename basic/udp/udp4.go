package main

import (
"net"
"time"
	"fmt"
	"os"
)

func main() {
	service := ":6333"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError2(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError2(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError2(err error){
	if  err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}
