// Author: xufei
// Date: 2020/5/10

package main

import (
	"learngo/gomicro/thiry-part/proto/learning"
	greeterH "learngo/gomicro/thiry-part/rpc/handler/greeter"
	learningH "learngo/gomicro/thiry-part/rpc/handler/learning"
	"log"

	"github.com/micro/go-micro/v2"
)

func main() {
	srv := micro.NewService(micro.Name("yw.dev.api.learning"))

	err := learning.RegisterGreeterHandler(srv.Server(), new(greeterH.Handler))
	if err != nil {
		log.Fatal(err)
	}

	err = learning.RegisterLearningHandler(srv.Server(), new(learningH.Handler))
	if err != nil {
		log.Fatal(err)
	}

	srv.Init()

	log.Fatal(srv.Run())
}
