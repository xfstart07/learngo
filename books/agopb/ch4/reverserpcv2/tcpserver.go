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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
	defer client.Close()

	var reply string

	err := client.Call("HelloService.Login", "user:password", &reply)
	if err != nil {
		log.Println("rpc call err", err)
		return
	}
	fmt.Println("登录", reply)

	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("获取成功", reply)
}
