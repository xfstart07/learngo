// Author: xufei
// Date: 2018/12/18

package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	log.Printf("listening %v...", host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
