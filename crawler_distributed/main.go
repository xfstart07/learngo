// Author: xufei
// Date: 2018/11/22

package main

import (
	"flag"
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/presist/client"
	"learngo/crawler_distributed/rpcsupport"
	client2 "learngo/crawler_distributed/worker/client"
	"log"
	"net/http"
	"net/rpc"
	"strings"

	_ "net/http/pprof"
)

var itemSaverPort = flag.String("item_saver", "", "item saver port")
var workerPort = flag.String("worker", "", "worker port(comma separate)")

func main() {
	flag.Parse()

	saver, err := client.ItemSaver(*itemSaverPort)
	if err != nil {
		panic(err)
	}

	pool := CreateClientPool(strings.Split(*workerPort, ","))
	worker, err := client2.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:       &scheduler.QueuedSchedule{},
		WorkerCount:     100,
		ItemSaver:       saver,
		WorkerProcessor: worker,
	}

	go func() {
		log.Println(http.ListenAndServe("weixinote.dev:8080", nil))
	}()

	e.Run(engine.Request{
		URL: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, "ParseCityList"),
	})
}

func CreateClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		newClient, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("new client err %v", err)
		} else {
			clients = append(clients, newClient)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
