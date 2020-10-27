// Author: xufei
// Date: 2019-08-23

package main

import (
	"fmt"
	"learngo/books/agopb/ch4/reverserpcv2/service"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
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

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			err := p.Register(service.NewHelloService(conn))
			if err != nil {
				log.Println("注册 rpc 失败", err)
				return
			}

			p.ServeConn(conn)
		}()
	}
}
