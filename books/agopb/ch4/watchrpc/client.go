// Author: xufei
// Date: 2019-08-22

package main

import (
	"fmt"
	"learngo/books/agopb/ch4/watchrpc/service"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "weixinote.dev:1234")
	if err != nil {
		log.Fatal(err)
	}

	doClientWork(client)
}

func doClientWork(client *rpc.Client) {
	go func() {
		fmt.Println("run watch...")
		var keyChanged string
		err := client.Call(service.ServiceName+".Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("watch key:", keyChanged)
	}()

	err := client.Call(service.ServiceName+".Set", [2]string{"name", "Leon"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	err = client.Call(service.ServiceName+".Set", [2]string{"name", "Leon Xu"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)
}
