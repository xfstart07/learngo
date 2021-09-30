package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	sip := net.ParseIP("192.168.4.122")

	srcAddr := &net.UDPAddr{
		IP:   sip,
		Port: 0,
	}
	dstAddr := &net.UDPAddr{
		IP:   sip,
		Port: 7912,
	}

	// DialUDP 建立发送数据的连接
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	// 写数据
	for {
		_, err = conn.Write([]byte("0014#00.05.29.C0.43.2A"))
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(">>> Packet sent to: ", *dstAddr)
		}

		time.Sleep(5 * time.Second)
	}
}
