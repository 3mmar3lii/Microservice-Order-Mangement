package rabbitmq

import (
	"context"
	"log"
)

// MessageHandler processes the raw message bytes. Return nil to Ack, or an error to Nack.
type MessageHandler func(ctx context.Context, body []byte) error

func (r *RabbitMq) Consume(ctx context.Context, cfg QueueConfig, consumerTag string, handler MessageHandler) error {
	q, err := r.channel.QueueDeclare(
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

	msgs, err := r.channel.ConsumeWithContext(
		ctx,
		q.Name,
		consumerTag,
		false,
		cfg.Exclusive, // exclusive
		false,         // noLocal (must be false in RabbitMQ)
		cfg.NoWait,    // noWait
		cfg.Args,      // args
	)
	if err != nil {
		return err
	}

	// Listen for messages in a background goroutine until ctx is canceled or channel closes
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-msgs:
				if !ok {
					return // RabbitMQ channel closed
				}

				if err := handler(ctx, d.Body); err != nil {
					log.Printf("error processing message from %s: %v", q.Name, err)
					// Nack(multiple: false, requeue: true) puts the message back in the queue to retry
					_ = d.Nack(false, true)
					continue
				}

				// Ack(multiple: false) tells RabbitMQ the message was processed safely
				_ = d.Ack(false)
			}
		}
	}()

	return nil
}
