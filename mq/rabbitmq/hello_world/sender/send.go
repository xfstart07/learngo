package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// 创建连接
	conn, err := amqp.Dial("amqp://guest:guest@weixinote.dev:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 打开一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//声明一个队列，绑定到一个空的交换机
	q, err := ch.QueueDeclare(
		"message_task", // name 队列名
		false,          // durable   持久化
		false,          // delete when unused  //当空闲时删除
		false,          // exclusive  // 专用队列
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello"
	// 向交换机发布消息
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key // 路由名称
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
