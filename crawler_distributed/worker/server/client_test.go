// Author: xufei
// Date: 2018/12/18

package main

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"testing"
)

func TestClient(t *testing.T) {
	const host = ":1235"
	go serverRpc(host)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1532019062",
		Parser: worker.SerializedParser{
			Name: config.ProfileParser,
			Args: engine.ParseArgs{
				Name:   "柠檬",
				Gender: "女士",
			},
		},
	}

	result := worker.ParseResult{}
	err = client.Call(config.CrawlServiceRpc, &req, &result)
	if err != nil {
		panic(err)
	}

	t.Logf("%+v", result)
}
