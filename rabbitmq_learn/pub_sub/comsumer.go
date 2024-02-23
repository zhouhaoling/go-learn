package main

import (
	"fmt"
	rl "go-learn/rabbitmq_learn"
)

func main() {
	//连接rabbitmq服务器
	conn, err := rl.RabbitMQConn()
	rl.ErrorHandle(err, "failed to connect to rabbitmq")
	//关闭连接
	//defer conn.Close()

	//创建通道
	channel, err := conn.Channel()
	rl.ErrorHandle(err, "failed to open a channel ")
	//关闭通道
	//defer channel.Close()

	q, err := channel.QueueDeclare(
		"test_fanout_queue1",
		true,
		false,
		false,
		false,
		nil,
	)
	rl.ErrorHandle(err, "failed to declare a queue")

	// 4.将消息发布到声明的队列
	messages, err := channel.Consume( // 注册一个消费者（接收消息）
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range messages {
		fmt.Printf("Received a message: %s\n", d.Body)
		fmt.Println("将日志信息保存到数据库")
	}
}
