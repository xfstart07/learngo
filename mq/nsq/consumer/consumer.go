// Author: xufei
// Date: 2019-10-15 16:32

package consumer

import (
	"encoding/json"
	"log"
	"time"

	"gopkg.in/olivere/elastic.v6"

	"github.com/nsqio/go-nsq"
)

type ConsumerHandler struct {
	q  *nsq.Consumer
	in chan record
	w  *work
}

func (h *ConsumerHandler) HandleMessage(message *nsq.Message) error {
	var rec record
	_ = json.Unmarshal(message.Body, &rec)
	h.in <- rec

	//log.Println("接收的消息", time.Now().UnixNano(), msg)

	return nil
}

func NewConsumer() {
	config := nsq.NewConfig()
	config.MaxBackoffDuration = 50 * time.Millisecond

	topicName := "publish"
	q, _ := nsq.NewConsumer(topicName, "ch", config)

	in := make(chan record, 10)
	h := &ConsumerHandler{
		q:  q,
		in: in,
		w:  newWorker(in),
	}
	q.AddHandler(h)

	err := q.ConnectToNSQD("weixinote.dev:4150")
	if err != nil {
		log.Println(err.Error())
	}
	<-q.StopChan
}

func newES() *elastic.Client {
	var opts []elastic.ClientOptionFunc

	opts = append(opts, elastic.SetURL("http://weixinote.dev:9200"))
	opts = append(opts, elastic.SetSniff(false))
	opts = append(opts, elastic.SetHealthcheckInterval(10*time.Second))

	client, err := elastic.NewClient(opts...)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
