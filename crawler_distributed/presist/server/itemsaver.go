// Author: xufei
// Date: 2018/12/18

package main

import (
	"flag"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/presist"
	"learngo/crawler_distributed/rpcsupport"

	"gopkg.in/olivere/elastic.v6"
)

var port = flag.String("port", "", "端口 :1234")

func main() {
	flag.Parse()

	err := serverRpc(*port, config.ElasticSearchDB)
	if err != nil {
		panic(err)
	}
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServerRpc(host, &presist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
