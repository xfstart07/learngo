package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@weixinote.dev:5672/")
	failOnError(err, "连接MQ失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "创建通道失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "创建交换机失败")

	body := bodyForm(os.Args)
	err = ch.Publish(
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "发布消息失败")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func bodyForm(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "hello"
	} else {
		return strings.Join(args[1:], " ")
	}
}
