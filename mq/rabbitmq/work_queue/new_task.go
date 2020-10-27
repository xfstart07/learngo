package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	// 创建连接
	// 打开一个通道
	// 声明一个队列
	// 发布消息

	conn, err := amqp.Dial("amqp://guest:guest@192.168.4.71:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"work_task",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyForm(os.Args)
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)

}

func bodyForm(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "hello"
	} else {
		return strings.Join(args[1:], " ")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
