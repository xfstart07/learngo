package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.UDPAddr{
		IP:   net.ParseIP("weixinote.dev"),
		Port: 9981,
	}
	// 监听一个IP 的 UDP
	listener, err := net.ListenUDP("udp", &ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印本地IP
	fmt.Println("Local: ", listener.LocalAddr())

	// buffer 数据
	data := make([]byte, 1024)
	for {
		// 读取UDP数据
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println("error during read: ", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])

		// 再写一个数据回去
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
