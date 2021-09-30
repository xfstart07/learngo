package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	log.Println("接收值", msg.NSQDAddress, string(msg.Body))
	return nil
}

func main() {
	initConsumer("test", "test-ch", "weixinote.dev:4150")

	forever := make(chan int)
	<-forever

}

func initConsumer(topic, channel, address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	consume, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	consume.AddHandler(&ConsumerT{})
	if err := consume.ConnectToNSQD(address); err != nil {
		log.Fatalln(err)
	}
}
