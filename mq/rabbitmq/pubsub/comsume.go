package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@weixinote.dev:5672/")
	failOnError(err, "连接MQ失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "创建通道失败")
	defer ch.Close()

	// 交换机声明
	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "交换机创建失败")

	// 队列声明
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "队列声明失败")

	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	failOnError(err, "绑定队列失败")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "注册消费者失败")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
