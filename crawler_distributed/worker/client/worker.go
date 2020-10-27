// Author: xufei
// Date: 2018/12/19

package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/worker"
	"log"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(request engine.Request) (engine.ParseResult, error) {
		log.Printf("start worker")

		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult
		client := <-clientChan
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		pResult := worker.DeserializeResult(sResult)

		return pResult, nil
	}, nil
}
