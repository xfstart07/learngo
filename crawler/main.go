// Author: xufei
// Date: 2018/11/22

package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	saver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:       &scheduler.QueuedSchedule{},
		WorkerCount:     100,
		ItemSaver:       saver,
		WorkerProcessor: engine.Worker,
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
