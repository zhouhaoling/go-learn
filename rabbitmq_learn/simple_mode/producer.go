package main

import (
	"encoding/json"
	rl "go-learn/rabbitmq_learn"

	"github.com/streadway/amqp"
)

type simpleDemo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func main() {
	//连接rabbitmq服务器
	conn, err := rl.RabbitMQConn()
	rl.ErrorHandle(err, "failed to connect to rabbitmq")
	//关闭连接
	defer conn.Close()

	//创建通道
	channel, err := conn.Channel()
	rl.ErrorHandle(err, "failed to open a channel ")
	//关闭通道
	defer channel.Close()

	//声明或者创建一个队列用来保存消息
	q, err := channel.QueueDeclare(
		//队列名称
		"simple:queue", //name
		false,          //durable
		false,          //delete when unused
		false,          // exclusive
		false,          //no-wait
		nil,            //arguments
	)
	rl.ErrorHandle(err, "failed to declare a queue")

	//消息
	data := simpleDemo{
		Name: "Tom",
		Addr: "Beijing",
	}
	dataBytes, err := json.Marshal(data)
	rl.ErrorHandle(err, "struct to json failed")

	//发布消息
	err = channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        dataBytes,
	})
	rl.ErrorHandle(err, "failed to publish a message")
}
