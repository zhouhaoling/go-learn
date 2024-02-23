package main

import (
	rl "go-learn/rabbitmq_learn"

	"github.com/streadway/amqp"
)

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

	//创建交换机
	exchangeName := "test_direct"
	channel.ExchangeDeclare(exchangeName, amqp.ExchangeDirect, true, false, false, false, nil)

	//创建队列
	queue1Name := "test_direct_queue1"
	queue2Name := "test_direct_queue2"
	channel.QueueDeclare(queue1Name, true, false, false, false, nil)
	channel.QueueDeclare(queue2Name, true, false, false, false, nil)

	//5. 绑定队列和交换机
	//队列1绑定 test
	channel.QueueBind(queue1Name, "test", exchangeName, false, nil)
	//队列2绑定 info error warning
	channel.QueueBind(queue2Name, "info", exchangeName, false, nil)
	channel.QueueBind(queue2Name, "error", exchangeName, false, nil)
	channel.QueueBind(queue2Name, "warning", exchangeName, false, nil)

	//6. 发送消息
	body := "日志信息：delete方法被调用，日志级别:error..."
	err = channel.Publish(
		exchangeName, // exchange 交换机名称 简单模式下交换机会使用默认的""
		"error",      // routing key 路由名称, 该消息会发送到routing key为 error的队列
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{ //发送消息数据
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}
