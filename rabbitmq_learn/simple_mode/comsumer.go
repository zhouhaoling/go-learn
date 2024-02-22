package main

import (
	rl "go-learn/rabbitmq_learn"
	"log"
)

func main() {
	conn, err := rl.RabbitMQConn()
	rl.ErrorHandle(err, "failed to connect to rabbitmq")
	defer conn.Close()

	channel, err := conn.Channel()
	rl.ErrorHandle(err, "failed to open a channel")
	defer channel.Close()

	//声明消息要发送的队列
	//如果没有一个名字叫hello的队列，则会创建该队列，如果有则不会创建
	q, err := channel.QueueDeclare(
		"simple:queue",
		false,
		false,
		false,
		false,
		nil,
	)
	rl.ErrorHandle(err, "failed to declare a queue")

	//定义消费者
	msgs, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	rl.ErrorHandle(err, "failed to register a consume")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
