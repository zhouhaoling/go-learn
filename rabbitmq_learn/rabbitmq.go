package rabbitmq_learn

import (
	"log"

	"github.com/streadway/amqp"
)

func RabbitMQConn() (conn *amqp.Connection, err error) {
	//连接rabbitmq
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	return
}

func ErrorHandle(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
