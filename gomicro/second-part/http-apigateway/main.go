// Author: xufei
// Date: 2020/3/6

package main

import (
	"context"
	"encoding/json"
	"learngo/gomicro/second-part/proto/prime"
	"learngo/gomicro/second-part/proto/sum"
	"log"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/v2/web"
)

var (
	sumService   sum.SumService
	primeService prime.PrimeService
)

func main() {
	srv := web.NewService(
		web.Name("go.micro.learning.web.portal"),
		//web.Address(":8088"),
		web.StaticDir("html"),
	)

	_ = srv.Init()

	sumService = sum.NewSumService("go.micro.learning.srv.sum", srv.Options().Service.Client())
	primeService = prime.NewPrimeService("go.micro.learning.srv.prime", srv.Options().Service.Client())
	srv.HandleFunc("/portal/sum", SumHandle)
	srv.HandleFunc("/portal/prime", PrimeHandle)
	srv.HandleFunc("/portal/update", PatchHandle)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

func SumHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("sum handle")
	inputStr := r.URL.Query().Get("input")
	input, _ := strconv.ParseInt(inputStr, 10, 64)

	req := &sum.SumRequest{
		Input: input,
	}
	resp, err := sumService.GetSum(context.Background(), req)
	if err != nil {
		log.Printf("sum client request failed: %v", err)
	}

	_, _ = w.Write([]byte(strconv.Itoa(int(resp.Output))))
}

func PrimeHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("prime handle")

	inputStr := r.URL.Query().Get("input")
	input, _ := strconv.ParseInt(inputStr, 10, 64)

	req := &prime.PrimeRequest{
		Input: input,
	}
	resp, err := primeService.GetPrime(context.Background(), req)
	if err != nil {
		log.Printf("prime client request failed: %v", err)
	}

	ret, _ := json.Marshal(resp.Output)
	_, _ = w.Write(ret)
}

func PatchHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("patch", r.Method)
}
