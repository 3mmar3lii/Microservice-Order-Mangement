package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	url string
}

func NewRabbitMq(url string) *RabbitMq {
	return &RabbitMq{url:url}
}

func (r *RabbitMq) Connect() error {

	conn, err := amqp.Dial(r.url)
	r.failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := conn.Channel()
	r.failOnError(err, "Failed to open a channel")

	r.conn = conn
	r.channel = channel

	return nil
}



func (r *RabbitMq) failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		return 
	}
}

func (r *RabbitMq) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}