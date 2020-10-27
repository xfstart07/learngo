package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.4.71:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"work_task", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 设置分派消息的策略，这里的配置：分派1个，预先分派设置为0，配置不是全局
	// 保证消息在分派到一个之后，只有等回复了 ack 之后才会再次分派。
	err = ch.Qos(1, 0, false)
	failOnError(err, "Failed to set Qos")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // 自动回复设置为 false
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false) // ask 手动回复， false 只发一次消息
		}
	}()

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
