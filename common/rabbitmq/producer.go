package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Producer(ctx context.Context, url string, cfg QueueConfig, body []byte) error {
	
	rabbitMQ := NewRabbitMq(url)
	if err := rabbitMQ.Connect(); err != nil {
		return err
	}
	// Make sure to close the connection/channel when done if opened per-publish
	defer rabbitMQ.Close()

	q, err := rabbitMQ.channel.QueueDeclare(
		cfg.Name,
		cfg.Durable,
		cfg.AutoDelete,
		cfg.Exclusive,
		cfg.NoWait,
		cfg.Args,
	)
	if err != nil {
		return err
	}

	return rabbitMQ.channel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}
