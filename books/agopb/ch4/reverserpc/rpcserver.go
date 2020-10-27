// Author: xufei
// Date: 2019-08-23

package main

import (
	"fmt"
	"learngo/books/agopb/ch4/reverserpc/service"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	err := rpc.Register(new(service.HelloService))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := net.Dial("tcp", "weixinote.dev:1234")
		if err != nil {
			log.Println("拨号失败", err)
		}
		if conn == nil {
			fmt.Println("无连接")
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("拨号成功")

		rpc.ServeConn(conn)

		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("连接关闭")
	}
}
