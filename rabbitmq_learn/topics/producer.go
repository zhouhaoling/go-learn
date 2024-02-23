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
	exchangeName := "test_topic"
	channel.ExchangeDeclare(exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)

	//4. 创建队列
	queue1Name := "test_topic_queue1"
	queue2Name := "test_topic_queue2"
	channel.QueueDeclare(queue1Name, true, false, false, false, nil)
	channel.QueueDeclare(queue2Name, true, false, false, false, nil)

	//5. 绑定队列和交换机
	//routing key 系统名称.日志级别
	//需求：所有error级别的日志村日数据库，所有order系统的日志存入数据库
	//队列1绑定
	channel.QueueBind(queue1Name, "#.error", exchangeName, false, nil)
	channel.QueueBind(queue1Name, "order.*", exchangeName, false, nil)
	//队列2绑定
	channel.QueueBind(queue2Name, "*.*", exchangeName, false, nil)

	//6. 发送消息
	body := "日志信息：delete方法被调用，日志级别:error..."
	err = channel.Publish(
		exchangeName, // exchange 交换机名称 简单模式下交换机会使用默认的""
		"goods.info", // routing key 路由名称
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{ //发送消息数据
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}
