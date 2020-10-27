// Author: xufei
// Date: 2018/12/18

package main

import (
	"flag"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"log"
)

var port = flag.String("port", "", "端口 :9000")

func main() {
	flag.Parse()

	log.Fatal(serverRpc(*port))
}

func serverRpc(host string) error {
	return rpcsupport.ServerRpc(host, &worker.CrawlService{})
}
