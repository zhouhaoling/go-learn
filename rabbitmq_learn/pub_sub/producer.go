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

	//3. 创建交换机
	//参数：
	//1、name：交换机名称
	//2、kind:交换机类型
	//amqp.ExchangeDirect 定向
	//amqp.ExchangeFanout 扇形（广播），发送消息到每个队列
	//amqp.ExchangeTopic 通配符的方式
	//amqp.ExchangeHeaders 参数匹配
	//3、durable：是否持久化
	//4、autoDelete：自动删除
	//5、internal：内部使用 一般false
	//6、noWait bool,
	//7、args：参数
	exchangeName := "test_fanout"
	err = channel.ExchangeDeclare(exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)
	rl.ErrorHandle(err, "failed to declare a exchange")

	//创建队列
	queue1Name := "test_fanout_queue1"
	queue2Name := "test_fanout_queue2"
	_, err = channel.QueueDeclare(queue1Name, true, false, false, false, nil)
	rl.ErrorHandle(err, "failed to declare a queue1")
	_, err = channel.QueueDeclare(queue2Name, true, false, false, false, nil)
	rl.ErrorHandle(err, "failed to declare a queue2")

	//绑定队列和交换机，通过交换机名与队列名来实现绑定
	err = channel.QueueBind(queue1Name, "", exchangeName, false, nil)
	rl.ErrorHandle(err, "failed to bind exchange and queue")
	err = channel.QueueBind(queue2Name, "", exchangeName, false, nil)
	rl.ErrorHandle(err, "failed to bind exchange and queue")

	//发送消息，通过交换机来发送消息
	body := "日志信息：方法被调用，日志级别:info..."
	channel.Publish(
		exchangeName, //交换机名称，简单模式下使用默认的
		"",           //routing key 路由名称
		false,        //mandatory
		false,        //immediate
		amqp.Publishing{ //发送消息数据
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}
