package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var addr = flag.String("addr", ":7912", "udp server bing address")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	// 地址
	dstAddr, err := net.ResolveUDPAddr("udp", *addr)
	fmt.Println(dstAddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// 设置监听
	conn, err := net.ListenUDP("udp", dstAddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	recvBuff := make([]byte, 1024)
	for {
		log.Println("Ready to receive packets!")
		// 读取消息
		rn, rmAddr, err := conn.ReadFromUDP(recvBuff)
		if err != nil {
			log.Println("Error:", err)
			return
		}
		fmt.Printf("<<< Packet received from: %s, data: %s\n", rmAddr.String(), string(recvBuff[:rn]))

		fmt.Println(checkstrIPV4(string(recvBuff[:rn])))

	}
}

func checkstrIPV4(str string) bool {
	str = str[5:]
	fmt.Println("str ", str)

	ip := net.ParseIP(str)
	fmt.Println("ip ", ip)
	res := ip.To4()

	if res != nil {
		fmt.Println("是IP", str)
		return true
	}

	return false
}

func checkstrMAC(str string) bool {
	str = str[5:]
	fmt.Println("str ", str, "len", len(str))

	mac, err := net.ParseMAC(str)
	fmt.Println("mac ", mac)
	fmt.Println("mac error ", err)

	if err == nil && mac != nil {
		fmt.Println("是Mac地址", str)
		return true
	}

	return false
}
