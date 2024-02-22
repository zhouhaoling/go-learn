package main

import (
	"fmt"
	rl "go-learn/rabbitmq_learn"
	"strconv"

	"github.com/streadway/amqp"
)

func Publish() {
	conn, err := rl.RabbitMQConn()
	rl.ErrorHandle(err, "failed to connect to rabbitmq")
	defer conn.Close()

	channel, err := conn.Channel()
	rl.ErrorHandle(err, "failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"queue", //队列名称
		true,    //是否持久化
		false,   //是否自动删除
		false,   //消息是否可以共享
		false,   //是否等待
		nil,     //其他参数
	)

	for i := 0; i < 10; i++ {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello world" + strconv.Itoa(i)),
		}
		if err := channel.Publish("", q.Name, false, false, message); err != nil {
			fmt.Println(err)
		}
	}

}
