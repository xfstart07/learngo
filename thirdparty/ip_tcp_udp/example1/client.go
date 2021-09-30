package main

import (
	"fmt"
	"net"
)

func main() {
	sip := net.ParseIP("weixinote.dev")

	orgAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 9981}

	// 连接UDP连接
	conn, err := net.DialUDP("udp", orgAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 在终端读入一个字符，保持程序一直存在
	// b := make([]byte, 1)
	// os.Stdin.Read(b)

	conn.Write([]byte("hello"))

	// RemoteAddr 远程接受数据的地址
	fmt.Println("remote addr:", conn.RemoteAddr())

	data := make([]byte, 1024)
	n, err := conn.Read(data)

	fmt.Println("data:", string(data[:n]), " remote", conn.RemoteAddr())
}
