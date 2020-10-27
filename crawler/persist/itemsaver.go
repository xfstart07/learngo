// Author: xufei
// Date: 2018/12/12

package persist

import (
	"context"
	"learngo/crawler/engine"
	"log"

	"gopkg.in/olivere/elastic.v6"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return out, err
	}

	go func() {
		for item := range out {
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	resp, err := client.Index().
		Index(index).
		Type("zhenai").
		Id(item.ID).
		BodyJson(item.Payload).
		Do(context.Background())
	if err != nil {
		return err
	}

	log.Printf("%+v", resp)
	return err
}
