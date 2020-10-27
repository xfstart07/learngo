// Author: xufei
// Date: 2018/12/19

package main

import (
	"learngo/frontend/api"
	"learngo/frontend/api/service"
	"log"
)

func main() {
	err := service.NewElasticClient()
	if err != nil {
		panic(err)
	}

	log.Fatal(api.NewApi(":8001"))
}
