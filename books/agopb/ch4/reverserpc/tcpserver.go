// Author: xufei
// Date: 2019-08-23

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	clientChan := make(chan *rpc.Client)
	go func() {
		doClientWork(clientChan)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		clientChan <- rpc.NewClient(conn)
	}
}

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer func() {
		e := client.Close()
		if e != nil {
			fmt.Println("关闭失败", e)
		}
		fmt.Println("关闭连接成功")
	}()

	fmt.Println("tcp 的 rpc 请求")
	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("获取成功", reply)
}
