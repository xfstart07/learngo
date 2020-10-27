// Author: xufei
// Date: 2019-08-22

package main

import (
	"io"
	"learngo/books/agopb/ch4/jsonrpc/service"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := service.RegisterHelloService(new(service.HelloService))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, req *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: req.Body,
			Writer:     w,
		}

		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			log.Println(err)
		}
	})

	log.Fatal(http.ListenAndServe(":1234", nil))
}

// client
// curl weixinote.dev:1234/jsonrpc -X POST --data '{"method":"jsonrpc/HelloService.Hello","params":["hello"],"id":0}'
