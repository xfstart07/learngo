package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	laddr = flag.String("laddr", "weixinote.dev:7913", "local server address")
	raddr = flag.String("raddr", "weixinote.dev:7912", "remote server address")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	// Resolving Address
	localAddr, err := net.ResolveUDPAddr("udp", *laddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", *raddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// 设置一个本地IP监听，返回一个连接
	conn, err := net.ListenUDP("udp", localAddr)
	// Exit if some error occured
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	// 写一个消息给服务器
	_, err = conn.WriteToUDP([]byte("hello client2"), remoteAddr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(">>> Packet sent to: ", *raddr)
	}
}
