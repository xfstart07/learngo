// Author: xufei
// Date: 2018/12/12

package client

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler_distributed/rpcsupport"
	"log"
	"net/rpc"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		for item := range out {
			err := Save(client, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *rpc.Client, item engine.Item) error {
	var result string
	err := client.Call("ItemSaverService.Save", item, &result)
	if err != nil {
		return err
	}

	if result != "ok" {
		return fmt.Errorf("item save error")
	}

	return nil
}
