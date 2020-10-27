package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

// 发给 NSQD
func main() {
	address := "weixinote.dev:4150"
	product := initProducter(address)

	for {
		publish(product, "ktv_info", "send")
		time.Sleep(5 * time.Second)
	}

}

func initProducter(address string) *nsq.Producer {
	product, err := nsq.NewProducer(address, nsq.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}
	return product
}

func publish(product *nsq.Producer, topic, message string) {
	if product != nil {
		log.Println("publish msg: ", message)
		if err := product.Publish(topic, []byte(message)); err != nil {
			log.Fatalln(err)
		}
	}
}
