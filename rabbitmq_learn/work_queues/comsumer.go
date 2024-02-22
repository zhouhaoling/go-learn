package main

import (
	"fmt"
	rl "go-learn/rabbitmq_learn"
)

func Consume(name string) {
	conn, err := rl.RabbitMQConn()
	rl.ErrorHandle(err, "failed to connect to rabbitmq")
	defer conn.Close()

	channel, err := conn.Channel()
	rl.ErrorHandle(err, "failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare("queue", true, false, false, false, nil)
	rl.ErrorHandle(err, "failed to declare a queue")

	consume, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	rl.ErrorHandle(err, "failed to register a consume")

	for message := range consume {
		fmt.Printf("name:%s   body:%s\n", name, string(message.Body))
	}
}
